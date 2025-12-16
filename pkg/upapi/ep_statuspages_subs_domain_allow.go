package upapi

import (
	"context"
)

type StatusPageSubsDomainAllowList struct {
	PK     int64  `json:"id"`
	Domain string `json:"domain"`
}

func (s StatusPageSubsDomainAllowList) PrimaryKey() PrimaryKey {
	return PrimaryKey(s.PK)
}

type StatusPageSubsDomainAllowListListResponse struct {
	Count   int64                           `json:"count,omitempty"`
	Results []StatusPageSubsDomainAllowList `json:"results,omitempty"`
}

func (r StatusPageSubsDomainAllowListListResponse) List() []StatusPageSubsDomainAllowList {
	return r.Results
}

func (r StatusPageSubsDomainAllowListListResponse) CountItems() int64 {
	return r.Count
}

type StatusPageSubsDomainAllowListResponse StatusPageSubsDomainAllowList

func (r StatusPageSubsDomainAllowListResponse) Item() StatusPageSubsDomainAllowList {
	return StatusPageSubsDomainAllowList(r)
}

type StatusPageSubsDomainAllowListCreateUpdateResponse struct {
	Results StatusPageSubsDomainAllowList `json:"results,omitempty"`
}

func (r StatusPageSubsDomainAllowListCreateUpdateResponse) Item() StatusPageSubsDomainAllowList {
	return r.Results
}

type StatusPageSubsDomainAllowListListOptions struct {
	Page     int64  `url:"page,omitempty"`
	PageSize int64  `url:"page_size,omitempty"`
	Search   string `url:"search,omitempty"`
	Ordering string `url:"ordering,omitempty"`
}

type StatusPageSubsDomainAllowListEndpoint interface {
	Create(context.Context, StatusPageSubsDomainAllowList) (*StatusPageSubsDomainAllowList, error)
	List(context.Context, StatusPageSubsDomainAllowListListOptions) (*ListResult[StatusPageSubsDomainAllowList], error)
	Update(context.Context, PrimaryKeyable, StatusPageSubsDomainAllowList) (*StatusPageSubsDomainAllowList, error)
	Get(context.Context, PrimaryKeyable) (*StatusPageSubsDomainAllowList, error)
	Delete(context.Context, PrimaryKeyable) error
}

type statusPageSubsDomainAllowListEndpointImpl struct {
	EndpointLister[StatusPageSubsDomainAllowListListResponse, StatusPageSubsDomainAllowList, StatusPageSubsDomainAllowListListOptions]
	EndpointCreator[StatusPageSubsDomainAllowList, StatusPageSubsDomainAllowListCreateUpdateResponse, StatusPageSubsDomainAllowList]
	EndpointUpdater[StatusPageSubsDomainAllowList, StatusPageSubsDomainAllowListCreateUpdateResponse, StatusPageSubsDomainAllowList]
	EndpointGetter[StatusPageSubsDomainAllowListResponse, StatusPageSubsDomainAllowList]
	EndpointDeleter
}
