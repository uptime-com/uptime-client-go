package upapi

import "context"

type CheckLocationListResponse struct {
	Locations []string `json:"locations"`
}

func (r CheckLocationListResponse) List() []string {
	return r.Locations
}

func (r CheckLocationListResponse) CountItems() int64 {
	return int64(len(r.Locations))
}

type CheckLocationListOptions struct{}

type checksEndpointLocationsImpl struct {
	EndpointLister[CheckLocationListResponse, string, CheckLocationListOptions]
}

func (c checksEndpointLocationsImpl) ListLocations(ctx context.Context) (*ListResult[string], error) {
	return c.List(ctx, CheckLocationListOptions{})
}
