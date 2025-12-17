# uptime-client-go

A Go client library for Uptime.com

## Breaking Changes

### v2.6.0

The `List()` methods on all endpoints now return `*ListResult[Item]` instead of `[]Item` to expose pagination metadata (total count) from API responses.

**Migration:**
```go
// Before (v2.5.x and earlier)
checks, err := api.Checks().List(ctx, upapi.CheckListOptions{})
if err != nil {
    return err
}
for _, check := range checks {
    fmt.Println(check.Name)
}

// After (v2.6.0+)
result, err := api.Checks().List(ctx, upapi.CheckListOptions{})
if err != nil {
    return err
}
for _, check := range result.Items {
    fmt.Println(check.Name)
}

// Access total count for pagination
fmt.Printf("Showing %d of %d checks\n", len(result.Items), result.TotalCount)
```

### v2.5.0

The `ContactGroups` field type has changed from `[]string` to `*[]string` across all check types to properly support PATCH requests.

**Migration:**
```go
// Before (v2.4.x and earlier)
check := upapi.CheckHTTP{
    ContactGroups: []string{"Default"},
}

// After (v2.5.0+)
check := upapi.CheckHTTP{
    ContactGroups: &[]string{"Default"},
}

// To explicitly set empty contact groups (clears the field)
check := upapi.CheckHTTP{
    ContactGroups: &[]string{},
}

// To omit the field (useful for PATCH - won't update the field)
check := upapi.CheckHTTP{
    ContactGroups: nil,
}
```

## Supported resources:

* Checks
* Dashboards
* Tags
* Outages
* Integrations
* Probe servers
* Contact groups (partial)

## Installation

### Command line tool (`upctl`)

Downdload the latest release from the [releases page](./releases) or install from sources:

```bash
go install github.com/uptime-com/uptime-client-go/v2/cmd/upctl@latest
```

#### Authentication

Obtain API token from [Uptime.com](https://uptime.com/api/tokens) and set it as an environment variable:

```bash
export UPCTL_TOKEN=your-api-token
```

### Library

```bash
go get -u github.com/uptime-com/uptime-client-go/v2@latest
```

## Run Tests

```bash
go test -v ./uptime
```

## Documentation

To view godocs locally, run `godoc`. Open http://localhost:6060 in a web browser and navigate to the go-uptime package
under Third party.

The [Uptime.com API Docs](https://uptime.com/api/v1/docs/) may also be a useful reference.

## Usage Examples

Please see the [examples](./examples) directory for usage examples.

## Credits

Contributions are welcome! Please feel free to fork and submit a pull request with any improvements.

Original version was created by [Kyle Gentle](https://github.com/kylegentle), with support from Elias Laham and the
Dev Team at Uptime.com.

### Contributors

See [contributors page](./graphs/contributors).
