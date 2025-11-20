package upapi

import (
	"context"
	"net"
	"time"
)

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
	CheckAddress               string    `json:"check_address,omitempty"`
	CheckName                  string    `json:"check_name,omitempty"`
	CheckMonitoringServiceType string    `json:"check_monitoring_service_type,omitempty"`
	StateIsUp                  bool      `json:"state_is_up,omitempty"`
	Ignored                    bool      `json:"ignored,omitempty"`
	NumLocationsDown           int64     `json:"num_locations_down,omitempty"`
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
	Count    int64    `json:"count,omitempty"`
	Next     string   `json:"next,omitempty"`
	Previous string   `json:"previous,omitempty"`
	Results  []Outage `json:"results,omitempty"`
}

func (o OutageListResponse) List() []Outage {
	return o.Results
}

type OutageResponse Outage

func (o OutageResponse) Item() Outage {
	return Outage(o)
}

// OutageListOptions specifies the optional parameters to the OutagesService.List
// and OutagesService.ListByServiceType methods.
type OutageListOptions struct {
	Page                       int64  `url:"page,omitempty"`
	PageSize                   int64  `url:"page_size,omitempty"`
	Search                     string `url:"search,omitempty"`
	Ordering                   string `url:"ordering,omitempty"`
	CheckMonitoringServiceType string `url:"check_monitoring_service_type,omitempty"`
}

type OutagesEndpoint interface {
	List(context.Context, OutageListOptions) ([]Outage, error)
	Get(context.Context, PrimaryKeyable) (*Outage, error)
}

func NewOutagesEndpoint(cbd CBD) OutagesEndpoint {
	const endpoint = "outages"
	return &outagesEndpointImpl{
		EndpointLister: NewEndpointLister[OutageListResponse, Outage, OutageListOptions](cbd, endpoint),
		EndpointGetter: NewEndpointGetter[OutageResponse, Outage](cbd, endpoint),
	}
}

type outagesEndpointImpl struct {
	EndpointLister[OutageListResponse, Outage, OutageListOptions]
	EndpointGetter[OutageResponse, Outage]
}
