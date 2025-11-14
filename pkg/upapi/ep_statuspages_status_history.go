package upapi

import (
	"context"
	"fmt"
)

// StatusPageStatusHistory represents a historical status entry
type StatusPageStatusHistory struct {
	PK          int64  `json:"pk,omitempty"`
	Status      string `json:"status,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
	ComponentPK *int64 `json:"component_pk,omitempty"`
}

func (s StatusPageStatusHistory) PrimaryKey() PrimaryKey {
	return PrimaryKey(s.PK)
}

// StatusPageStatusHistoryListOptions provides filtering options for status history
type StatusPageStatusHistoryListOptions struct {
	Page        int64  `url:"page,omitempty"`
	PageSize    int64  `url:"page_size,omitempty"`
	Search      string `url:"search,omitempty"`
	Ordering    string `url:"ordering,omitempty"`
	Status      string `url:"status,omitempty"`
	ComponentPK int64  `url:"component_pk,omitempty"`
	DateFrom    string `url:"date_from,omitempty"`
	DateTo      string `url:"date_to,omitempty"`
}

// StatusPageStatusHistoryListResponse wraps the list response
type StatusPageStatusHistoryListResponse struct {
	Count   int64                      `json:"count,omitempty"`
	Results []StatusPageStatusHistory `json:"results,omitempty"`
}

func (r StatusPageStatusHistoryListResponse) List() []StatusPageStatusHistory {
	return r.Results
}

// StatusPageStatusHistoryResponse wraps a single status history item
type StatusPageStatusHistoryResponse StatusPageStatusHistory

func (r StatusPageStatusHistoryResponse) Item() StatusPageStatusHistory {
	return StatusPageStatusHistory(r)
}

// StatusPageStatusHistoryEndpoint provides access to status page status history
type StatusPageStatusHistoryEndpoint interface {
	List(context.Context, StatusPageStatusHistoryListOptions) ([]StatusPageStatusHistory, error)
	Get(context.Context, PrimaryKeyable) (*StatusPageStatusHistory, error)
}

type statusPageStatusHistoryEndpointImpl struct {
	EndpointLister[StatusPageStatusHistoryListResponse, StatusPageStatusHistory, StatusPageStatusHistoryListOptions]
	EndpointGetter[StatusPageStatusHistoryResponse, StatusPageStatusHistory]
}

// NewStatusPageStatusHistoryEndpoint creates a new status page status history endpoint
func NewStatusPageStatusHistoryEndpoint(cbd CBD, statusPagePK PrimaryKeyable) StatusPageStatusHistoryEndpoint {
	endpoint := fmt.Sprintf("statuspages/%d/history", statusPagePK.PrimaryKey())
	return &statusPageStatusHistoryEndpointImpl{
		EndpointLister: NewEndpointLister[StatusPageStatusHistoryListResponse, StatusPageStatusHistory, StatusPageStatusHistoryListOptions](cbd, endpoint),
		EndpointGetter: NewEndpointGetter[StatusPageStatusHistoryResponse, StatusPageStatusHistory](cbd, endpoint),
	}
}
