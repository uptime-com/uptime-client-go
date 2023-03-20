package upapi

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

// Integration represents an integration in Uptime.com.
type Integration struct {
	PK   int    `json:"pk,omitempty"`
	URL  string `json:"url,omitempty"`
	Name string `json:"name,omitempty"`

	Module        string   `json:"module,omitempty"`
	ContactGroups []string `json:"contact_groups,omitempty"`
	APIEndpoint   string   `json:"api_endpoint,omitempty"`
	APIKey        string   `json:"api_key,omitempty"`
	Teams         string   `json:"teams,omitempty"`
	Tags          string   `json:"tags,omitempty"`
	AutoResolve   bool     `json:"autoresolve,omitempty"`
}

func (i Integration) PrimaryKey() int {
	return i.PK
}

type IntegrationListResponse struct {
	Count    int           `json:"count,omitempty"`
	Next     string        `json:"next,omitempty"`
	Previous string        `json:"previous,omitempty"`
	Results  []Integration `json:"results,omitempty"`
}

func (r IntegrationListResponse) List() []Integration {
	return r.Results
}

type IntegrationListOptions struct {
	Page     int    `url:"page,omitempty"`
	PageSize int    `url:"page_size,omitempty"`
	Search   string `url:"search,omitempty"`
	Ordering string `url:"ordering,omitempty"`
	Module   string `url:"module,omitempty"`
}

type IntegrationResponse struct {
	Results Integration `json:"results,omitempty"`
}

func (r IntegrationResponse) Item() *Integration {
	return &r.Results
}

type IntegrationsEndpoint interface {
	List(context.Context, IntegrationListOptions) ([]Integration, error)
	Get(context.Context, PrimaryKey) (*Integration, error)
	Create(context.Context, Integration) (*Integration, error)
	Update(context.Context, Integration) (*Integration, error)
	Delete(context.Context, PrimaryKey) error
}

func NewIntegrationsEndpoint(cbd CBD) IntegrationsEndpoint {
	const endpoint = "integrations"
	return &integrationsEndpointImpl{
		EndpointLister:  NewEndpointLister[IntegrationListResponse, Integration, IntegrationListOptions](cbd, endpoint),
		EndpointGetter:  NewEndpointGetter[PrimaryKey, IntegrationResponse, Integration](cbd, endpoint),
		EndpointCreator: NewEndpointCreator[Integration, IntegrationResponse, Integration](&integrationsCreateCBD{cbd}, endpoint),
		EndpointUpdater: NewEndpointUpdater[Integration, IntegrationResponse, Integration](cbd, endpoint),
		EndpointDeleter: NewEndpointDeleter[PrimaryKey](cbd, endpoint),
	}
}

type integrationsEndpointImpl struct {
	EndpointLister[IntegrationListResponse, Integration, IntegrationListOptions]
	EndpointGetter[PrimaryKey, IntegrationResponse, Integration]
	EndpointCreator[Integration, IntegrationResponse, Integration]
	EndpointUpdater[Integration, IntegrationResponse, Integration]
	EndpointDeleter[PrimaryKey]
}

type integrationsCreateCBD struct {
	CBD
}

func (c *integrationsCreateCBD) BuildRequest(ctx context.Context, method string, endpoint string, args any, body any) (*http.Request, error) {
	if method != http.MethodPost {
		panic("only POST requests allowed here; this is always a programming error")
	}
	data, ok := body.(Integration)
	if !ok {
		panic("only uptime.Integration objects allowed here; this is always a programming error")
	}
	endpoint = fmt.Sprintf("%sadd-%s/", endpoint, strings.ToLower(data.Module))
	return c.CBD.BuildRequest(ctx, method, endpoint, args, body)
}
