package upapi

import (
	"context"
	"fmt"
	"net/http"
)

// Subaccount represents a subaccount in Uptime.com.
type Subaccount struct {
	PK   int64  `json:"pk,omitempty"`
	Name string `json:"name" flag:"name"`
	URL  string `json:"url,omitempty"`
}

func (s Subaccount) PrimaryKey() PrimaryKey {
	return PrimaryKey(s.PK)
}

// SubaccountItemResponse wraps a single subaccount response.
type SubaccountItemResponse struct {
	Subaccount `json:",inline"`
}

func (s SubaccountItemResponse) Item() Subaccount {
	return s.Subaccount
}

// SubaccountCreateRequest represents the request body for creating a subaccount.
type SubaccountCreateRequest struct {
	Name string `json:"name" flag:"name"`
}

// SubaccountUpdateRequest represents the request body for updating a subaccount.
type SubaccountUpdateRequest struct {
	Name string `json:"name" flag:"name"`
}

// SubaccountPacks represents the request/response for transferring packs.
// Positive num transfers from main account to subaccount.
// Negative num transfers from subaccount to main account.
type SubaccountPacks struct {
	Num int64 `json:"num" flag:"num"`
}

// SubaccountsEndpoint provides access to subaccount operations.
type SubaccountsEndpoint interface {
	List(context.Context) ([]Subaccount, error)
	Create(context.Context, SubaccountCreateRequest) (*Subaccount, error)
	Get(context.Context, PrimaryKeyable) (*Subaccount, error)
	Update(context.Context, PrimaryKeyable, SubaccountUpdateRequest) (*Subaccount, error)
	TransferPacks(context.Context, PrimaryKeyable, SubaccountPacks) error
}

// NewSubaccountsEndpoint creates a new subaccounts endpoint.
func NewSubaccountsEndpoint(cbd CBD) SubaccountsEndpoint {
	const endpoint = "auth/subaccounts"
	return &subaccountsEndpointImpl{
		CBD:             cbd,
		endpoint:        endpoint,
		EndpointGetter:  NewEndpointGetter[SubaccountItemResponse, Subaccount](cbd, endpoint),
		EndpointCreator: NewEndpointCreator[SubaccountCreateRequest, SubaccountItemResponse, Subaccount](cbd, endpoint),
		EndpointUpdater: NewEndpointUpdater[SubaccountUpdateRequest, SubaccountItemResponse, Subaccount](cbd, endpoint),
	}
}

type subaccountsEndpointImpl struct {
	CBD
	endpoint string
	EndpointGetter[SubaccountItemResponse, Subaccount]
	EndpointCreator[SubaccountCreateRequest, SubaccountItemResponse, Subaccount]
	EndpointUpdater[SubaccountUpdateRequest, SubaccountItemResponse, Subaccount]
}

// List returns all subaccounts. Note: This endpoint returns a simple array, not a paginated response.
func (e *subaccountsEndpointImpl) List(ctx context.Context) ([]Subaccount, error) {
	path := e.endpoint + "/"
	req, err := e.BuildRequest(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, err
	}
	rs, err := e.Do(req)
	if err != nil {
		return nil, err
	}
	defer rs.Body.Close()
	if rs.StatusCode != http.StatusOK {
		return nil, ErrorFromResponse(rs)
	}
	var subaccounts []Subaccount
	if err := e.DecodeResponse(rs, &subaccounts); err != nil {
		return nil, err
	}
	return subaccounts, nil
}

// TransferPacks transfers packs between main account and subaccount.
// Positive num transfers from main to subaccount, negative transfers from subaccount to main.
func (e *subaccountsEndpointImpl) TransferPacks(ctx context.Context, pk PrimaryKeyable, packs SubaccountPacks) error {
	path := fmt.Sprintf("%s/%d/allocation/", e.endpoint, pk.PrimaryKey())
	req, err := e.BuildRequest(ctx, "POST", path, packs, nil)
	if err != nil {
		return err
	}
	rs, err := e.Do(req)
	if err != nil {
		return err
	}
	defer rs.Body.Close()
	if rs.StatusCode != http.StatusOK && rs.StatusCode != http.StatusCreated {
		return ErrorFromResponse(rs)
	}
	return nil
}
