package upapi

import (
	"context"
)

type SLAReportGroup struct {
	ID            int      `json:"id,omitempty"`
	Name          string   `json:"name,omitempty"`
	GroupServices []string `json:"group_services,omitempty"`
}

func (s SLAReportGroup) PrimaryKey() PrimaryKey {
	return PrimaryKey(s.ID)
}

type SLAReportGroupListResponse struct {
	Count   int64            `json:"count,omitempty"`
	Results []SLAReportGroup `json:"results,omitempty"`
}

func (r SLAReportGroupListResponse) List() []SLAReportGroup {
	return r.Results
}

type SLAReportGroupResponse SLAReportGroup

func (r SLAReportGroupResponse) Item() SLAReportGroup {
	return SLAReportGroup(r)
}

type SLAReportGroupCreateResponse struct {
	Results SLAReportGroup `json:"results,omitempty"`
}

func (r SLAReportGroupCreateResponse) Item() SLAReportGroup {
	return r.Results
}

type SLAReportGroupCreateUpdateResponse struct {
	Results SLAReportGroup `json:"results,omitempty"`
}

func (r SLAReportGroupCreateUpdateResponse) Item() SLAReportGroup {
	return r.Results
}

type SLAReportGroupListOptions struct {
	Page     int64  `url:"page,omitempty"`
	PageSize int64  `url:"page_size,omitempty"`
	Search   string `url:"search,omitempty"`
	Ordering string `url:"ordering,omitempty"`
}

type SLAReportsGroupsEndpoint interface {
	Create(context.Context, SLAReportGroup) (*SLAReportGroup, error)
	List(context.Context, SLAReportGroupListOptions) ([]SLAReportGroup, error)
	Get(context.Context, PrimaryKeyable) (*SLAReportGroup, error)
	Delete(context.Context, PrimaryKeyable) error
}

type slaReportsGroupsEndpointImpl struct {
	EndpointLister[SLAReportGroupListResponse, SLAReportGroup, SLAReportGroupListOptions]
	EndpointCreator[SLAReportGroup, SLAReportGroupCreateResponse, SLAReportGroup]
	EndpointGetter[SLAReportGroupResponse, SLAReportGroup]
	EndpointDeleter
}
