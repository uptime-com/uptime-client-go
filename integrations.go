package uptime

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type IntegrationService service

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

type IntegrationListResponse struct {
	Count    int            `json:"count,omitempty"`
	Next     string         `json:"next,omitempty"`
	Previous string         `json:"previous,omitempty"`
	Results  []*Integration `json:"results,omitempty"`
}

type IntegrationListOptions struct {
	Page     int    `url:"page,omitempty"`
	PageSize int    `url:"page_size,omitempty"`
	Search   string `url:"search,omitempty"`
	Ordering string `url:"ordering,omitempty"`
	Module   string `url:"module,omitempty"`
}

type IntegrationResponse struct {
	Messages map[string]interface{} `json:"messages,omitempty"`
	Results  Integration            `json:"results,omitempty"`
}

func (s *IntegrationService) List(ctx context.Context, opt *IntegrationListOptions) ([]*Integration, *http.Response, error) {
	u := "integrations"
	clResp, resp, err := s.listIntegrations(ctx, u, opt)
	return clResp.Results, resp, err
}

func (s *IntegrationService) ListAll(ctx context.Context, opt *IntegrationListOptions) ([]*Integration, error) {
	u := "integrations"
	opt.Page = 1

	result := []*Integration{}

	clResp, _, err := s.listIntegrations(ctx, u, opt)
	if err != nil {
		return nil, err
	}
	result = append(result, clResp.Results...)

	for clResp.Next != "" {
		opt.Page++
		clResp, _, err = s.listIntegrations(ctx, u, opt)
		if err != nil {
			return nil, err
		}
		result = append(result, clResp.Results...)
	}

	return result, err
}

func (s *IntegrationService) listIntegrations(ctx context.Context, url string, opt *IntegrationListOptions) (*IntegrationListResponse, *http.Response, error) {
	u, err := addOptions(url, opt)
	if err != nil {
		return nil, nil, err
	}

	req, _ := s.client.NewRequest("GET", u, nil)

	var integrations IntegrationListResponse
	resp, err := s.client.Do(ctx, req, &integrations)
	if err != nil {
		return nil, nil, err
	}
	return &integrations, resp, nil
}

// Create a new integration in Uptime.com based on the provided Integration.
func (s *IntegrationService) Create(ctx context.Context, integration *Integration) (*Integration, *http.Response, error) {
	suffix := strings.ToLower(strings.Replace(integration.Module, "_", "-", -1))
	u := fmt.Sprintf("integrations/add-%v", suffix)

	req, err := s.client.NewRequest("POST", u, integration)
	if err != nil {
		return nil, nil, err
	}

	cr := &IntegrationResponse{}
	resp, err := s.client.Do(ctx, req, cr)
	if err != nil {
		return nil, resp, err
	}

	return &cr.Results, resp, nil
}

func (s *IntegrationService) Get(ctx context.Context, pk int) (*Integration, *http.Response, error) {
	u := fmt.Sprintf("integrations/%v", pk)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	c := &Integration{}
	resp, err := s.client.Do(ctx, req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, nil
}

// Update an integration.
func (s *IntegrationService) Update(ctx context.Context, integration *Integration) (*Integration, *http.Response, error) {
	u := fmt.Sprintf("integrations/%v", integration.PK)
	if integration.PK == 0 {
		return nil, nil, fmt.Errorf("Error updating integration with empty PK")
	}

	req, err := s.client.NewRequest("PATCH", u, integration)
	if err != nil {
		return nil, nil, err
	}

	cr := &IntegrationResponse{}
	resp, err := s.client.Do(ctx, req, cr)
	if err != nil {
		return nil, resp, err
	}

	return &cr.Results, resp, nil
}

// Delete an integration.
func (s *IntegrationService) Delete(ctx context.Context, pk int) (*http.Response, error) {
	u := fmt.Sprintf("integrations/%v", pk)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
