package upapi

import "context"

type CloudStatusService struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Title    string `json:"title,omitempty"`
	SubTitle string `json:"sub_title,omitempty"`
	GroupID  int64  `json:"group_id"`
	Group    string `json:"group,omitempty"`
}

type CloudStatusServiceListOptions struct {
	Page     int64  `url:"page,omitempty"`
	PageSize int64  `url:"page_size,omitempty"`
	Group    string `url:"group,omitempty"`
	Search   string `url:"search,omitempty"`
	Ordering string `url:"ordering,omitempty"`
}

type CloudStatusServiceListResponse struct {
	Count   int64                `json:"count,omitempty"`
	Results []CloudStatusService `json:"results,omitempty"`
}

func (r CloudStatusServiceListResponse) List() []CloudStatusService {
	return r.Results
}

func (r CloudStatusServiceListResponse) CountItems() int64 {
	return r.Count
}

type checksEndpointCloudStatusServicesImpl struct {
	EndpointLister[CloudStatusServiceListResponse, CloudStatusService, CloudStatusServiceListOptions]
}

func (c checksEndpointCloudStatusServicesImpl) ListCloudStatusServices(
	ctx context.Context, opts CloudStatusServiceListOptions,
) (*ListResult[CloudStatusService], error) {
	return c.List(ctx, opts)
}
