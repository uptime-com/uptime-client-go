package upapi

import "context"

// AccountUsageItem represents a single usage/limit entry for the account plan.
type AccountUsageItem struct {
	Name    string `json:"name"`
	Current int64  `json:"current"`
	Max     int64  `json:"max"`
}

type AccountUsageListResponse []AccountUsageItem

type AccountUsageListOptions struct{}

func (r AccountUsageListResponse) List() []AccountUsageItem {
	return r
}

func (r AccountUsageListResponse) CountItems() int64 {
	return int64(len(r))
}

// AccountUsageEndpoint provides access to account usage and plan limits.
type AccountUsageEndpoint interface {
	List(ctx context.Context) (*ListResult[AccountUsageItem], error)
}

func NewAccountUsageEndpoint(cbd CBD) AccountUsageEndpoint {
	return &accountUsageEndpointImpl{
		EndpointLister: NewEndpointLister[AccountUsageListResponse, AccountUsageItem, AccountUsageListOptions](cbd, "auth/account-usage"),
	}
}

type accountUsageEndpointImpl struct {
	EndpointLister[AccountUsageListResponse, AccountUsageItem, AccountUsageListOptions]
}

func (e accountUsageEndpointImpl) List(ctx context.Context) (*ListResult[AccountUsageItem], error) {
	return e.EndpointLister.List(ctx, AccountUsageListOptions{})
}
