package upapi

import "context"

// CloudStatusGroupListItem is the read-only item shape returned by the
// /checks/cloudstatus-groups/ lookup endpoint. It is intentionally separate
// from CloudStatusGroup, which has a custom MarshalJSON that serializes as a
// bare integer for the check-config write path - re-marshaling a list item
// through that type would silently drop the Name.
type CloudStatusGroupListItem struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type CloudStatusGroupListOptions struct {
	Page     int64  `url:"page,omitempty"`
	PageSize int64  `url:"page_size,omitempty"`
	Search   string `url:"search,omitempty"`
	Ordering string `url:"ordering,omitempty"`
}

type CloudStatusGroupListResponse struct {
	Count   int64                      `json:"count,omitempty"`
	Results []CloudStatusGroupListItem `json:"results,omitempty"`
}

func (r CloudStatusGroupListResponse) List() []CloudStatusGroupListItem {
	return r.Results
}

func (r CloudStatusGroupListResponse) CountItems() int64 {
	return r.Count
}

type checksEndpointCloudStatusGroupsImpl struct {
	EndpointLister[CloudStatusGroupListResponse, CloudStatusGroupListItem, CloudStatusGroupListOptions]
}

func (c checksEndpointCloudStatusGroupsImpl) ListCloudStatusGroups(
	ctx context.Context, opts CloudStatusGroupListOptions,
) (*ListResult[CloudStatusGroupListItem], error) {
	return c.List(ctx, opts)
}
