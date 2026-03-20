# AGENTS.md

Guide for AI agents and human developers working on this codebase.

## Project overview

Go client library and CLI tool (`upctl`) for the [Uptime.com API](https://uptime.com/api/v1/docs/).

- **Module**: `github.com/uptime-com/uptime-client-go/v2`
- **Go version**: 1.18+ (uses generics)
- **Development model**: trunk-based (commits go directly to `main`)

## Repository structure

```
pkg/upapi/           # Public API client library
  api.go             # API interface, client factory, endpoint registration
  endpoint.go        # Generic CRUD endpoint implementations (Getter, Lister, Creator, Updater, Deleter)
  cbd.go             # Core building dependencies: HTTP client, request builder, response decoder
  options.go         # Client options (auth, base URL, rate limiting, retry, tracing)
  error.go           # Structured API error type
  ep_*.go            # Individual endpoint implementations (one file per resource)

internal/upctl/      # CLI tool implementation (not importable)
  cli.go             # Root command, global flags, API initialization
  cli_*.go           # Subcommands (one file per resource)
  output.go          # JSON/spew output formatting
  utils.go           # Flag binding via reflection, PK parsing

cmd/upctl/           # CLI entrypoint
  main.go            # main() function
```

## Architecture

### CBD pattern (Client, Builder, Decoder)

The library is built on three core interfaces composed into `CBD`:

```go
type CBD interface {
Doer            // Do(*http.Request) (*http.Response, error)
RequestBuilder  // BuildRequest(ctx, method, endpoint, args, data) (*http.Request, error)
ResponseDecoder // DecodeResponse(*http.Response, data) error
}
```

Client options use a decorator pattern, each wrapping the CBD to add behavior:
`WithToken`, `WithBearerToken`, `WithBaseURL`, `WithUserAgent`, `WithRateLimit`, `WithRetry`, `WithTrace`,
`WithSubaccount`.

### Generic endpoint helpers

Five generic types handle standard CRUD operations:

| Type                                       | HTTP method | URL pattern        |
|--------------------------------------------|-------------|--------------------|
| `EndpointGetter[Response, Item]`           | GET         | `{endpoint}/{pk}/` |
| `EndpointLister[Response, Item, Options]`  | GET         | `{endpoint}/`      |
| `EndpointCreator[Request, Response, Item]` | POST        | `{endpoint}/`      |
| `EndpointUpdater[Request, Response, Item]` | PATCH       | `{endpoint}/{pk}/` |
| `EndpointDeleter`                          | DELETE      | `{endpoint}/{pk}/` |

Response types must implement `Itemable[T]` (for single items) or `Listable[T]` (for lists).

### Custom actions

Endpoints with non-CRUD operations (e.g., `Alerts.RootCause`, `Users.Deactivate`) make direct CBD calls:

```go
req, err := e.cbd.BuildRequest(ctx, "POST", path, nil, nil)
rs, err := e.cbd.Do(req)
defer rs.Body.Close()
// check status, decode response
```

Always append a trailing slash to paths (the API requires it).

### Nested sub-endpoints

Status pages demonstrate the pattern for sub-resources:

```go
api.StatusPages().Components(statusPagePK).List(ctx, opts)
```

The parent endpoint returns a sub-endpoint scoped to a specific parent PK.

## Adding a new endpoint

### API layer (`pkg/upapi/`)

1. Create `ep_yourresource.go`
2. Define the main type with JSON tags and `PrimaryKey()` method
3. Define response wrappers implementing `Itemable`/`Listable` interfaces
4. Define `ListOptions` struct with `url` tags
5. Define the endpoint interface
6. Implement using generic helpers (embed them in your impl struct)
7. Register in `api.go`: add to `API` interface, `apiImpl` struct, `New()` factory, and accessor method
8. Write tests using `httptest.NewServer`

### CLI layer (`internal/upctl/`)

1. Create `cli_yourresource.go`
2. Define parent command with `Use`, `Aliases`, `Short`
3. Add subcommands: `list`, `get`, `create`, `update`, `delete` as appropriate
4. Use `Bind()` to map flags struct fields to CLI flags (supports: string, int32, int64, float64, bool, []string, *[]
   string, nested structs)
5. Register via `init()` functions

### Conventions

- JSON tags: use `omitempty` for optional fields (enables PATCH partial updates)
- Pointer fields: use `*bool`, `*int64` when zero-value vs absent distinction matters
- URL query params: use `url` tags on options structs (parsed by `go-querystring`)
- Endpoint paths: lowercase, URL-safe (e.g., `"contacts"`, `"auth/account-usage"`)
- List defaults: `Page: 1, PageSize: 100, Ordering: "pk"`
- Pass the bound flags struct directly to API calls (do not construct a new struct with hardcoded values)

## Testing

```bash
go test ./...
```

Tests use `httptest.NewServer` to mock API responses. No external services required.

The `Bind()` utility has its own tests in `internal/upctl/utils_test.go`.

## Release process

### Versioning: EffVer (Intended Effort Versioning)

This project follows [EffVer](https://effver.org), which versions by the effort consumers need to adopt a release, not
by the technical nature of the change.

Given version `MACRO.MESO.MICRO`:

- **Micro** bump (e.g., v2.9.0 -> v2.9.1): adopting this release should require minimal effort. Bug fixes, new
  endpoints, additive features. Existing code continues to work without changes.
- **Meso** bump (e.g., v2.9.1 -> v2.10.0): adopting this release may require some effort. Small breaking changes to
  specific types or method signatures. Callers of affected APIs need updates, but the migration path is straightforward.
- **Macro** bump (e.g., v2.10.0 -> v3.0.0): adopting this release requires significant effort. Architectural changes,
  module path changes, widespread API redesign. Note: a Go major version bump also requires changing the module import
  path.

The key distinction from SemVer: EffVer acknowledges that "every bug has users." Rather than pretending bug fixes are
always safe, it communicates the expected downstream impact honestly. A behavioral change that technically fixes a bug
but breaks existing workflows warrants a meso bump, not a micro one.

### How to release

Releases are automated via GoReleaser running in the `release.yaml` GitHub Actions workflow, triggered by `v*` tags.

To release:

```bash
# Check the current latest tag
git tag --sort=-v:refname | head -5

# Tag the release (on main)
git tag v2.10.0

# Push the tag
git push origin v2.10.0
```

This triggers the workflow which:

1. Builds `upctl` binaries for darwin/linux/windows (amd64, arm64)
2. Builds and pushes Docker images to `uptimecom/upctl`
3. Creates a GitHub release with archives and checksums
4. Signs checksums with GPG

Pre-release versions (e.g., `v2.10.0-rc1`) are automatically detected and marked as pre-release on GitHub.

### Release artifacts

- **Binaries**: tar.gz archives per platform (`upctl_{version}_{os}_{arch}.tar.gz`)
- **Docker**: `uptimecom/upctl:latest`, `uptimecom/upctl:{version}`, `uptimecom/upctl:{major}`
- **Checksums**: SHA-256, GPG-signed
