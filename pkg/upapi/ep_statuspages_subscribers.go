package upapi

import (
	"context"
)

type StatusPageSubscriber struct {
	PK                 int64  `json:"id"`
	Target             string `json:"target"`
	Type               string `json:"type"`
	ForceValidationSMS bool   `json:"force_validation_sms"`
}

func (s StatusPageSubscriber) PrimaryKey() PrimaryKey {
	return PrimaryKey(s.PK)
}

type StatusPageSubscriberListResponse struct {
	Count   int64                  `json:"count,omitempty"`
	Results []StatusPageSubscriber `json:"results,omitempty"`
}

func (r StatusPageSubscriberListResponse) List() []StatusPageSubscriber {
	return r.Results
}

type StatusPageSubscriberResponse StatusPageSubscriber

func (r StatusPageSubscriberResponse) Item() StatusPageSubscriber {
	return StatusPageSubscriber(r)
}

type StatusPageSubscriberCreateResponse struct {
	Results StatusPageSubscriber `json:"results,omitempty"`
}

func (r StatusPageSubscriberCreateResponse) Item() StatusPageSubscriber {
	return r.Results
}

type StatusPageSubscriberListOptions struct {
	Page     int64  `url:"page,omitempty"`
	PageSize int64  `url:"page_size,omitempty"`
	Search   string `url:"search,omitempty"`
	Ordering string `url:"ordering,omitempty"`
}

type StatusPageSubscriberEndpoint interface {
	Create(context.Context, StatusPageSubscriber) (*StatusPageSubscriber, error)
	List(context.Context, StatusPageSubscriberListOptions) ([]StatusPageSubscriber, error)
	Get(context.Context, PrimaryKeyable) (*StatusPageSubscriber, error)
	Delete(context.Context, PrimaryKeyable) error
}

type statusPageSubscriberEndpointImpl struct {
	EndpointLister[StatusPageSubscriberListResponse, StatusPageSubscriber, StatusPageSubscriberListOptions]
	EndpointCreator[StatusPageSubscriber, StatusPageSubscriberCreateResponse, StatusPageSubscriber]
	EndpointGetter[StatusPageSubscriberResponse, StatusPageSubscriber]
	EndpointDeleter
}
