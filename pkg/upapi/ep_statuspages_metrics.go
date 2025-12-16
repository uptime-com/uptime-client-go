package upapi

import (
	"context"
)

type StatusPageMetric struct {
	PK        int64  `json:"pk"`
	URL       string `json:"url,omitempty"`
	Name      string `json:"name"`
	ServiceID int64  `json:"service_id"`
	IsVisible bool   `json:"is_visible"`
}

func (s StatusPageMetric) PrimaryKey() PrimaryKey {
	return PrimaryKey(s.PK)
}

type StatusPageMetricListResponse struct {
	Count   int64              `json:"count,omitempty"`
	Results []StatusPageMetric `json:"results,omitempty"`
}

func (r StatusPageMetricListResponse) List() []StatusPageMetric {
	return r.Results
}

func (r StatusPageMetricListResponse) CountItems() int64 {
	return r.Count
}

type StatusPageMetricResponse StatusPageMetric

func (r StatusPageMetricResponse) Item() StatusPageMetric {
	return StatusPageMetric(r)
}

type StatusPageMetricCreateUpdateResponse struct {
	Results StatusPageMetric `json:"results,omitempty"`
}

func (r StatusPageMetricCreateUpdateResponse) Item() StatusPageMetric {
	return r.Results
}

type StatusPageMetricListOptions struct {
	Page     int64  `url:"page,omitempty"`
	PageSize int64  `url:"page_size,omitempty"`
	Search   string `url:"search,omitempty"`
	Ordering string `url:"ordering,omitempty"`
}

type StatusPageMetricEndpoint interface {
	Create(context.Context, StatusPageMetric) (*StatusPageMetric, error)
	Update(context.Context, PrimaryKeyable, StatusPageMetric) (*StatusPageMetric, error)
	List(context.Context, StatusPageMetricListOptions) (*ListResult[StatusPageMetric], error)
	Get(context.Context, PrimaryKeyable) (*StatusPageMetric, error)
	Delete(context.Context, PrimaryKeyable) error
}

type statusPageMetricEndpointImpl struct {
	EndpointLister[StatusPageMetricListResponse, StatusPageMetric, StatusPageMetricListOptions]
	EndpointCreator[StatusPageMetric, StatusPageMetricCreateUpdateResponse, StatusPageMetric]
	EndpointUpdater[StatusPageMetric, StatusPageMetricCreateUpdateResponse, StatusPageMetric]
	EndpointGetter[StatusPageMetricResponse, StatusPageMetric]
	EndpointDeleter
}
