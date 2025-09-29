package upapi

import (
	"context"
)

type StatusPageSubsDomainBlockList struct {
	PK     int64  `json:"id"`
	Domain string `json:"domain"`
}

func (s StatusPageSubsDomainBlockList) PrimaryKey() PrimaryKey {
	return PrimaryKey(s.PK)
}

type StatusPageSubsDomainBlockListListResponse struct {
	Count   int64                           `json:"count,omitempty"`
	Results []StatusPageSubsDomainBlockList `json:"results,omitempty"`
}

func (r StatusPageSubsDomainBlockListListResponse) List() []StatusPageSubsDomainBlockList {
	return r.Results
}

type StatusPageSubsDomainBlockListResponse StatusPageSubsDomainBlockList

func (r StatusPageSubsDomainBlockListResponse) Item() StatusPageSubsDomainBlockList {
	return StatusPageSubsDomainBlockList(r)
}

type StatusPageSubsDomainBlockListCreateUpdateResponse struct {
	Results StatusPageSubsDomainBlockList `json:"results,omitempty"`
}

func (r StatusPageSubsDomainBlockListCreateUpdateResponse) Item() StatusPageSubsDomainBlockList {
	return r.Results
}

type StatusPageSubsDomainBlockListListOptions struct {
	Page     int64  `url:"page,omitempty"`
	PageSize int64  `url:"page_size,omitempty"`
	Search   string `url:"search,omitempty"`
	Ordering string `url:"ordering,omitempty"`
}

type StatusPageSubsDomainBlockListEndpoint interface {
	Create(context.Context, StatusPageSubsDomainBlockList) (*StatusPageSubsDomainBlockList, error)
	List(context.Context, StatusPageSubsDomainBlockListListOptions) ([]StatusPageSubsDomainBlockList, error)
	Update(context.Context, PrimaryKeyable, StatusPageSubsDomainBlockList) (*StatusPageSubsDomainBlockList, error)
	Get(context.Context, PrimaryKeyable) (*StatusPageSubsDomainBlockList, error)
	Delete(context.Context, PrimaryKeyable) error
}

type statusPageSubsDomainBlockListEndpointImpl struct {
	EndpointLister[StatusPageSubsDomainBlockListListResponse, StatusPageSubsDomainBlockList, StatusPageSubsDomainBlockListListOptions]
	EndpointCreator[StatusPageSubsDomainBlockList, StatusPageSubsDomainBlockListCreateUpdateResponse, StatusPageSubsDomainBlockList]
	EndpointUpdater[StatusPageSubsDomainBlockList, StatusPageSubsDomainBlockListCreateUpdateResponse, StatusPageSubsDomainBlockList]
	EndpointGetter[StatusPageSubsDomainBlockListResponse, StatusPageSubsDomainBlockList]
	EndpointDeleter
}
