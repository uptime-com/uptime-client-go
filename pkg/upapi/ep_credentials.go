package upapi

import (
	"context"
)

type CredentialSecret struct {
	Certificate string `json:"certificate,omitempty"`
	Key         string `json:"key,omitempty"`
	Password    string `json:"password,omitempty"`
	Passphrase  string `json:"passphrase,omitempty"`
	Secret      string `json:"secret,omitempty"`
}

type Credential struct {
	PK                   int64            `json:"id,omitempty"`
	DisplayName          string           `json:"display_name,omitempty"`
	Description          string           `json:"description,omitempty"`
	CredentialType       string           `json:"credential_type"`
	Hint                 string           `json:"hint,omitempty"`
	Username             string           `json:"username,omitempty"`
	Version              string           `json:"version,omitempty"`
	UsedSecretProperties []string         `json:"used_secret_properties,omitempty"`
	CreatedBy            int64            `json:"created_by,omitempty"`
	Secret               CredentialSecret `json:"secret"`
}

func (c Credential) PrimaryKey() PrimaryKey {
	return PrimaryKey(c.PK)
}

type CredentialListOptions struct {
	Page              int64  `url:"page,omitempty"`
	PageSize          int64  `url:"page_size,omitempty"`
	Search            string `url:"search,omitempty"`
	Ordering          string `url:"ordering,omitempty"`
	HasOnCallSchedule bool   `url:"has_on_call_schedule,omitempty"`
}

type CredentialListResponse struct {
	Count   int64        `json:"count,omitempty"`
	Results []Credential `json:"results,omitempty"`
}

func (r CredentialListResponse) List() []Credential {
	return r.Results
}

type CredentialResponse Credential

func (r CredentialResponse) Item() Credential {
	return Credential(r)
}

type CredentialCreateUpdateResponse struct {
	Results Credential `json:"results,omitempty"`
}

func (r CredentialCreateUpdateResponse) Item() Credential {
	return r.Results
}

type CredentialEndpoint interface {
	List(context.Context, CredentialListOptions) ([]Credential, error)
	Create(context.Context, Credential) (*Credential, error)
	Update(context.Context, PrimaryKeyable, Credential) (*Credential, error)
	Get(context.Context, PrimaryKeyable) (*Credential, error)
	Delete(context.Context, PrimaryKeyable) error
}

func NewCredentialsEndpoint(cbd CBD) CredentialEndpoint {
	const endpoint = "credentials"
	return &credentialsEndpointImpl{
		EndpointLister:  NewEndpointLister[CredentialListResponse, Credential, CredentialListOptions](cbd, endpoint),
		EndpointCreator: NewEndpointCreator[Credential, CredentialCreateUpdateResponse](cbd, endpoint),
		EndpointUpdater: NewEndpointUpdater[Credential, CredentialCreateUpdateResponse](cbd, endpoint),
		EndpointGetter:  NewEndpointGetter[CredentialResponse](cbd, endpoint),
		EndpointDeleter: NewEndpointDeleter(cbd, endpoint),
	}
}

type credentialsEndpointImpl struct {
	EndpointLister[CredentialListResponse, Credential, CredentialListOptions]
	EndpointCreator[Credential, CredentialCreateUpdateResponse, Credential]
	EndpointUpdater[Credential, CredentialCreateUpdateResponse, Credential]
	EndpointGetter[CredentialResponse, Credential]
	EndpointDeleter
}
