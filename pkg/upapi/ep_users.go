package upapi

import (
	"context"
	"fmt"
)

type UserAccount struct {
	Name               string `json:"name,omitempty"`
	Timezone           string `json:"timezone,omitempty"`
	DataRegion         string `json:"data_region,omitempty"`
	FreeTrialExpiresAt string `json:"free_trial_expires_at,omitempty"`
}

type User struct {
	PK                  int64        `json:"pk,omitempty"`
	URL                 string       `json:"url,omitempty"`
	FirstName           string       `json:"first_name,omitempty"`
	LastName            string       `json:"last_name,omitempty"`
	Email               string       `json:"email,omitempty"`
	Password            string       `json:"password,omitempty"`
	IsActive            bool         `json:"is_active,omitempty"`
	IsPrimary           bool         `json:"is_primary,omitempty"`
	AccessLevel         string       `json:"access_level,omitempty"`
	IsAPIEnabled        bool         `json:"is_api_enabled,omitempty"`
	NotifyPaidInvoices  bool         `json:"notify_paid_invoices,omitempty"`
	AssignedSubaccounts []string     `json:"assigned_subaccounts,omitempty"`
	RequireTwoFactor    string       `json:"require_two_factor,omitempty"`
	MustTwoFactor       bool         `json:"must_two_factor,omitempty"`
	Timezone            string       `json:"timezone,omitempty"`
	Account             *UserAccount `json:"account,omitempty"`
}

func (u User) PrimaryKey() PrimaryKey {
	return PrimaryKey(u.PK)
}

type UserListResponse struct {
	Count    int64  `json:"count,omitempty"`
	Next     string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
	Results  []User `json:"results,omitempty"`
}

func (r UserListResponse) List() []User {
	return r.Results
}

type UserListOptions struct {
	Page        int64  `url:"page,omitempty"`
	PageSize    int64  `url:"page_size,omitempty"`
	Search      string `url:"search,omitempty"`
	Ordering    string `url:"ordering,omitempty"`
	Email       string `url:"email,omitempty"`
	AccessLevel string `url:"access_level,omitempty"`
	Subaccount  string `url:"subaccount,omitempty"`
}

type UserItemResponse struct {
	User `json:",inline"`
}

func (u UserItemResponse) Item() User {
	return u.User
}

type UserCreateUpdateResponse struct {
	Results User `json:"results"`
}

func (u UserCreateUpdateResponse) Item() User {
	return u.Results
}

type UserCreateRequest struct {
	FirstName           string   `json:"first_name,omitempty" flag:"first-name"`
	LastName            string   `json:"last_name,omitempty" flag:"last-name"`
	Email               string   `json:"email" flag:"email"`
	Password            string   `json:"password" flag:"password"`
	AccessLevel         string   `json:"access_level,omitempty" flag:"access-level"`
	IsAPIEnabled        bool     `json:"is_api_enabled,omitempty" flag:"api-enabled"`
	NotifyPaidInvoices  bool     `json:"notify_paid_invoices,omitempty" flag:"notify-paid-invoices"`
	AssignedSubaccounts []string `json:"assigned_subaccounts,omitempty" flag:"assigned-subaccounts"`
	RequireTwoFactor    string   `json:"require_two_factor,omitempty" flag:"require-two-factor"`
}

type UserUpdateRequest struct {
	FirstName           string   `json:"first_name,omitempty" flag:"first-name"`
	LastName            string   `json:"last_name,omitempty" flag:"last-name"`
	Email               string   `json:"email,omitempty" flag:"email"`
	Password            string   `json:"password,omitempty" flag:"password"`
	AccessLevel         string   `json:"access_level,omitempty" flag:"access-level"`
	IsAPIEnabled        *bool    `json:"is_api_enabled,omitempty" flag:"api-enabled"`
	NotifyPaidInvoices  *bool    `json:"notify_paid_invoices,omitempty" flag:"notify-paid-invoices"`
	AssignedSubaccounts []string `json:"assigned_subaccounts,omitempty" flag:"assigned-subaccounts"`
	RequireTwoFactor    string   `json:"require_two_factor,omitempty" flag:"require-two-factor"`
}

type UsersEndpoint interface {
	List(context.Context, UserListOptions) ([]User, error)
	Create(context.Context, UserCreateRequest) (*User, error)
	Get(context.Context, PrimaryKeyable) (*User, error)
	Update(context.Context, PrimaryKeyable, UserUpdateRequest) (*User, error)
	Delete(context.Context, PrimaryKeyable) error
	Deactivate(context.Context, PrimaryKeyable) (*User, error)
	Reactivate(context.Context, PrimaryKeyable) (*User, error)
}

func NewUsersEndpoint(cbd CBD) UsersEndpoint {
	const endpoint = "users"
	return &usersEndpointImpl{
		CBD:             cbd,
		endpoint:        endpoint,
		EndpointLister:  NewEndpointLister[UserListResponse, User, UserListOptions](cbd, endpoint),
		EndpointGetter:  NewEndpointGetter[UserItemResponse, User](cbd, endpoint),
		EndpointCreator: NewEndpointCreator[UserCreateRequest, UserCreateUpdateResponse, User](cbd, endpoint),
		EndpointUpdater: NewEndpointUpdater[UserUpdateRequest, UserCreateUpdateResponse, User](cbd, endpoint),
		EndpointDeleter: NewEndpointDeleter(cbd, endpoint),
	}
}

type usersEndpointImpl struct {
	CBD
	endpoint string
	EndpointLister[UserListResponse, User, UserListOptions]
	EndpointGetter[UserItemResponse, User]
	EndpointCreator[UserCreateRequest, UserCreateUpdateResponse, User]
	EndpointUpdater[UserUpdateRequest, UserCreateUpdateResponse, User]
	EndpointDeleter
}

func (e *usersEndpointImpl) Deactivate(ctx context.Context, pk PrimaryKeyable) (*User, error) {
	path := fmt.Sprintf("%s/%d/deactivate/", e.endpoint, pk.PrimaryKey())
	req, err := e.BuildRequest(ctx, "POST", path, nil, nil)
	if err != nil {
		return nil, err
	}
	rs, err := e.Do(req)
	if err != nil {
		return nil, err
	}
	defer rs.Body.Close()
	if rs.StatusCode != 200 {
		return nil, ErrorFromResponse(rs)
	}
	var resp UserCreateUpdateResponse
	if err := e.DecodeResponse(rs, &resp); err != nil {
		return nil, err
	}
	user := resp.Item()
	return &user, nil
}

func (e *usersEndpointImpl) Reactivate(ctx context.Context, pk PrimaryKeyable) (*User, error) {
	path := fmt.Sprintf("%s/%d/reactivate/", e.endpoint, pk.PrimaryKey())
	req, err := e.BuildRequest(ctx, "POST", path, nil, nil)
	if err != nil {
		return nil, err
	}
	rs, err := e.Do(req)
	if err != nil {
		return nil, err
	}
	defer rs.Body.Close()
	if rs.StatusCode != 200 {
		return nil, ErrorFromResponse(rs)
	}
	var resp UserCreateUpdateResponse
	if err := e.DecodeResponse(rs, &resp); err != nil {
		return nil, err
	}
	user := resp.Item()
	return &user, nil
}
