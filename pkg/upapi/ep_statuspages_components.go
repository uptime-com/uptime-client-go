package upapi

import (
	"context"
)

type StatusPageComponent struct {
	PK             int64  `json:"pk"`
	URL            string `json:"url,omitempty"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	IsGroup        bool   `json:"is_group,omitempty"`
	GroupID        *int64 `json:"group_id"`
	ServiceID      *int64 `json:"service_id"`
	Status         string `json:"status,omitempty"`
	AutoStatusDown string `json:"auto_status_down,omitempty"`
	AutoStatusUp   string `json:"auto_status_up,omitempty"`
}

func (s StatusPageComponent) PrimaryKey() PrimaryKey {
	return PrimaryKey(s.PK)
}

type StatusPageComponentListResponse struct {
	Count   int64                 `json:"count,omitempty"`
	Results []StatusPageComponent `json:"results,omitempty"`
}

func (r StatusPageComponentListResponse) List() []StatusPageComponent {
	return r.Results
}

func (r StatusPageComponentListResponse) CountItems() int64 {
	return r.Count
}

type StatusPageComponentResponse StatusPageComponent

func (r StatusPageComponentResponse) Item() StatusPageComponent {
	return StatusPageComponent(r)
}

type StatusPageComponentCreateUpdateResponse struct {
	Results StatusPageComponent `json:"results,omitempty"`
}

func (r StatusPageComponentCreateUpdateResponse) Item() StatusPageComponent {
	return r.Results
}

type StatusPageComponentListOptions struct {
	Page     int64  `url:"page,omitempty"`
	PageSize int64  `url:"page_size,omitempty"`
	Search   string `url:"search,omitempty"`
	Ordering string `url:"ordering,omitempty"`
}

type StatusPageComponentEndpoint interface {
	Create(context.Context, StatusPageComponent) (*StatusPageComponent, error)
	Update(context.Context, PrimaryKeyable, StatusPageComponent) (*StatusPageComponent, error)
	List(context.Context, StatusPageComponentListOptions) (*ListResult[StatusPageComponent], error)
	Get(context.Context, PrimaryKeyable) (*StatusPageComponent, error)
	Delete(context.Context, PrimaryKeyable) error
}

type statusPageComponentEndpointImpl struct {
	EndpointLister[StatusPageComponentListResponse, StatusPageComponent, StatusPageComponentListOptions]
	EndpointCreator[StatusPageComponent, StatusPageComponentCreateUpdateResponse, StatusPageComponent]
	EndpointUpdater[StatusPageComponent, StatusPageComponentCreateUpdateResponse, StatusPageComponent]
	EndpointGetter[StatusPageComponentResponse, StatusPageComponent]
	EndpointDeleter
}
