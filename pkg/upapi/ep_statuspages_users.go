package upapi

import (
	"context"
)

type StatusPageUser struct {
	PK        int64  `json:"pk"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsActive  bool   `json:"is_active,omitempty"`
}

func (s StatusPageUser) PrimaryKey() PrimaryKey {
	return PrimaryKey(s.PK)
}

type StatusPageUserListResponse struct {
	Count   int64            `json:"count,omitempty"`
	Results []StatusPageUser `json:"results,omitempty"`
}

func (r StatusPageUserListResponse) List() []StatusPageUser {
	return r.Results
}

func (r StatusPageUserListResponse) CountItems() int64 {
	return r.Count
}

type StatusPageUserResponse StatusPageUser

func (r StatusPageUserResponse) Item() StatusPageUser {
	return StatusPageUser(r)
}

type StatusPageUserCreateUpdateResponse struct {
	Results StatusPageUser `json:"results,omitempty"`
}

func (r StatusPageUserCreateUpdateResponse) Item() StatusPageUser {
	return r.Results
}

type StatusPageUserListOptions struct {
	Page     int64  `url:"page,omitempty"`
	PageSize int64  `url:"page_size,omitempty"`
	Search   string `url:"search,omitempty"`
	Ordering string `url:"ordering,omitempty"`
}

type StatusPageUserEndpoint interface {
	Create(context.Context, StatusPageUser) (*StatusPageUser, error)
	Update(context.Context, PrimaryKeyable, StatusPageUser) (*StatusPageUser, error)
	List(context.Context, StatusPageUserListOptions) (*ListResult[StatusPageUser], error)
	Get(context.Context, PrimaryKeyable) (*StatusPageUser, error)
	Delete(context.Context, PrimaryKeyable) error
}

type statusPageUserEndpointImpl struct {
	EndpointLister[StatusPageUserListResponse, StatusPageUser, StatusPageUserListOptions]
	EndpointCreator[StatusPageUser, StatusPageUserCreateUpdateResponse, StatusPageUser]
	EndpointUpdater[StatusPageUser, StatusPageUserCreateUpdateResponse, StatusPageUser]
	EndpointGetter[StatusPageUserResponse, StatusPageUser]
	EndpointDeleter
}
