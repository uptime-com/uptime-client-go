# go-uptime - a Go client library for Uptime.com

## Supported resources:
* Checks
* Tags
* Outages

## Installation
`go get -u github.com/uptime-com/rest-api-clients/golang/uptime`

## Documentation
To view godocs locally, run `godoc`. Open http://localhost:6060 in a web browser and navigate to the go-uptime package under Third party.

The [Uptime.com API Docs](https://uptime.com/api/v1/docs/) may also be a useful reference.

## Usage Examples
### Instantiate a Client
```go
import (
    "context"
    uptime "github.com/uptime-com/rest-api-clients/golang/uptime"
)

clientConfig := &uptime.Config {
    Token:            "my-api-token",
    RateMilliseconds: 2000,
}

client, err := uptime.NewClient(clientConfig)
```

### Create a Check
```go
c := &uptime.Check {
    CheckType:     "HTTP",
    Address:       "https://uptime.com",
    Interval:      1,
    Threshold:     15,
    Locations:     []string{"US-East", "GBR", "AUT"},
    ContactGroups: []string{"Default"},
    Tags:          []string{"testing"},
}

ctx := context.Background()
newCheck, response, err := client.Checks.Create(ctx, c)
```

### Get a list of Outages
```go
options := &uptime.OutageListOptions {
    Ordering: "-created_at",
}

ctx := context.Background()
outages, resp, err := client.Outages.List(ctx, options)
```

### Delete a Tag
```go
tagId = 12345

ctx := context.Background()
resp, err := client.Tags.Delete(ctx, tagId)
```

## Credits
go-uptime was originally created by [Kyle Gentle](https://github.com/kylegentle), with support from Elias Laham and the Dev Team at Uptime.com.
