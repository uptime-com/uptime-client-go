package uptime

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

// OutageService handles communication with the outage related
// methods of the Uptime.com API.
//
// Uptime.com API docs: https://uptime.com/api/v1/docs/#!/outages/
type OutageService service

// Outage represents an outage reported by Uptime.com.
type Outage struct {
	PK                         int64     `json:"pk,omitempty"`
	URL                        string    `json:"url,omitempty"`
	CreatedAt                  time.Time `json:"created_at,omitempty"`
	ResolvedAt                 time.Time `json:"resolved_at,omitempty"`
	DurationSecs               int64     `json:"duration_secs,omitempty"`
	IgnoreAlertURL             string    `json:"ignore_alert_url,omitempty"`
	CheckPK                    int64     `json:"check_pk,omitempty"`
	CheckURL                   string    `json:"check_url,omitempty"`
	CheckAddresss              string    `json:"check_address,omitempty"`
	CheckName                  string    `json:"check_name,omitempty"`
	CheckMonitoringServiceType string    `json:"check_monitoring_service_type,omitempty"`
	StateIsUp                  bool      `json:"state_is_up,omitempty"`
	Ignored                    bool      `json:"ignored,omitempty"`
	NumLocationsDown           int       `json:"num_locations_down,omitempty"`
	AllAlerts                  *[]Alert  `json:"all_alerts,omitempty"`
}

// Alert represents an alert generated during an outage.
type Alert struct {
	PK                   int64      `json:"pk,omitempty"`
	URL                  string     `json:"url,omitempty"`
	CreatedAt            *time.Time `json:"created_at,omitempty"`
	MonitoringServerName string     `json:"monitoring_server_name,omitempty"`
	MonitoringServerIPv4 *net.IP    `json:"monitoring_server_ipv4,omitempty"`
	MonitoringServerIPv6 *net.IP    `json:"monitoring_server_ipv6,omitempty"`
	Location             string     `json:"location,omitempty"`
	Output               string     `json:"output,omitempty"`
}

// OutageListResponse represents a page of Outage results returned by
// the Uptime.com API.
type OutageListResponse struct {
	Count    int       `json:"count,omitempty"`
	Next     string    `json:"next,omitempty"`
	Previous string    `json:"previous,omitempty"`
	Results  []*Outage `json:"results,omitempty"`
}

// OutageListOptions specifies the optional parameters to the OutagesService.List
// and OutagesService.ListByServiceType methods.
type OutageListOptions struct {
	Page                       int    `url:"page,omitempty"`
	PageSize                   int    `url:"page_size,omitempty"`
	Search                     string `url:"search,omitempty"`
	Ordering                   string `url:"ordering,omitempty"`
	CheckMonitoringServiceType string `url:"check_monitoring_service_type,omitempty"`
}

// List all outages on the account.
func (s *OutageService) List(ctx context.Context, opt *OutageListOptions) ([]*Outage, *http.Response, error) {
	u := "outages"
	return s.listOutages(ctx, u, opt)
}

func (s *OutageService) listOutages(ctx context.Context, url string, opt *OutageListOptions) ([]*Outage, *http.Response, error) {
	u, err := addOptions(url, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var outages OutageListResponse
	resp, err := s.client.Do(ctx, req, &outages)
	if err != nil {
		return nil, resp, err
	}

	return outages.Results, resp, nil
}

// Get a single outage.
func (s *OutageService) Get(ctx context.Context, pk string) (*Outage, *http.Response, error) {
	u := fmt.Sprintf("outages/%v", pk)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var o *Outage
	resp, err := s.client.Do(ctx, req, &o)
	if err != nil {
		return nil, resp, err
	}
	return o, resp, nil
}
