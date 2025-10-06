package upapi

import (
	"context"
	"time"
)

type ServiceVariableCredential struct {
	ID                   int64    `json:"id,omitempty"`
	CredentialType       string   `json:"credential_type,omitempty"`
	DisplayName          string   `json:"display_name,omitempty"`
	Description          string   `json:"description,omitempty"`
	Hint                 string   `json:"hint,omitempty"`
	Username             string   `json:"username,omitempty"`
	Version              string   `json:"version,omitempty"`
	UsedSecretProperties []string `json:"used_secret_properties,omitempty"`
	CreatedBy            int64    `json:"created_by,omitempty"`
}

type ServiceVariable struct {
	ID           int64                      `json:"id,omitempty"`
	CredentialID int64                      `json:"credential_id,omitempty"`
	Credential   *ServiceVariableCredential `json:"credential,omitempty"`
	PropertyName string                     `json:"property_name,omitempty"`
	VariableName string                     `json:"variable_name,omitempty"`
	DeletedAt    *time.Time                 `json:"deleted_at,omitempty"`
	Account      string                     `json:"account,omitempty"`
	Service      string                     `json:"service,omitempty"`
}

func (s ServiceVariable) PrimaryKey() PrimaryKey {
	return PrimaryKey(s.ID)
}

type ServiceVariableListResponse struct {
	Count    int64             `json:"count,omitempty"`
	Next     string            `json:"next,omitempty"`
	Previous string            `json:"previous,omitempty"`
	Results  []ServiceVariable `json:"results,omitempty"`
}

func (r ServiceVariableListResponse) List() []ServiceVariable {
	return r.Results
}

type ServiceVariableListOptions struct {
	Page     int64  `url:"page,omitempty"`
	PageSize int64  `url:"page_size,omitempty"`
	Search   string `url:"search,omitempty"`
	Ordering string `url:"ordering,omitempty"`
}

type ServiceVariableItemResponse struct {
	Results ServiceVariable `json:"results,omitempty"`
}

func (s ServiceVariableItemResponse) Item() ServiceVariable {
	return s.Results
}

type ServiceVariableCreateUpdateResponse struct {
	Results ServiceVariable `json:"results,omitempty"`
}

func (s ServiceVariableCreateUpdateResponse) Item() ServiceVariable {
	return s.Results
}

type ServiceVariableCreateRequest struct {
	CredentialID int64  `json:"credential_id" flag:"credential-id"`
	PropertyName string `json:"property_name" flag:"property-name"`
	ServiceID    int64  `json:"service_id" flag:"service-id"`
	VariableName string `json:"variable_name" flag:"variable-name"`
}

type ServiceVariableUpdateRequest struct {
	CredentialID int64  `json:"credential_id,omitempty" flag:"credential-id"`
	PropertyName string `json:"property_name,omitempty" flag:"property-name"`
	ServiceID    int64  `json:"service_id,omitempty" flag:"service-id"`
	VariableName string `json:"variable_name,omitempty" flag:"variable-name"`
}

type ServiceVariablesEndpoint interface {
	List(context.Context, ServiceVariableListOptions) ([]ServiceVariable, error)
	Create(context.Context, ServiceVariableCreateRequest) (*ServiceVariable, error)
	Get(context.Context, PrimaryKeyable) (*ServiceVariable, error)
	Update(context.Context, PrimaryKeyable, ServiceVariableUpdateRequest) (*ServiceVariable, error)
	Delete(context.Context, PrimaryKeyable) error
}

func NewServiceVariablesEndpoint(cbd CBD) ServiceVariablesEndpoint {
	const endpoint = "servicevariables"
	return &serviceVariablesEndpointImpl{
		EndpointLister:  NewEndpointLister[ServiceVariableListResponse, ServiceVariable, ServiceVariableListOptions](cbd, endpoint),
		EndpointGetter:  NewEndpointGetter[ServiceVariableItemResponse, ServiceVariable](cbd, endpoint),
		EndpointCreator: NewEndpointCreator[ServiceVariableCreateRequest, ServiceVariableCreateUpdateResponse, ServiceVariable](cbd, endpoint),
		EndpointUpdater: NewEndpointUpdater[ServiceVariableUpdateRequest, ServiceVariableCreateUpdateResponse, ServiceVariable](cbd, endpoint),
		EndpointDeleter: NewEndpointDeleter(cbd, endpoint),
	}
}

type serviceVariablesEndpointImpl struct {
	EndpointLister[ServiceVariableListResponse, ServiceVariable, ServiceVariableListOptions]
	EndpointGetter[ServiceVariableItemResponse, ServiceVariable]
	EndpointCreator[ServiceVariableCreateRequest, ServiceVariableCreateUpdateResponse, ServiceVariable]
	EndpointUpdater[ServiceVariableUpdateRequest, ServiceVariableCreateUpdateResponse, ServiceVariable]
	EndpointDeleter
}
