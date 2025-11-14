package upapi

import (
	"context"
	"fmt"
	"net"
	"time"
)

// AlertItem represents an alert/incident in the Uptime.com monitoring system.
// Alerts are generated when a check detects an issue from a monitoring location.
type AlertItem struct {
	PK                         int64      `json:"pk,omitempty"`
	URL                        string     `json:"url,omitempty"`
	CreatedAt                  *time.Time `json:"created_at,omitempty"`
	ResolvedAt                 *time.Time `json:"resolved_at,omitempty"`
	MonitoringServerName       string     `json:"monitoring_server_name,omitempty"`
	MonitoringServerIPv4       *net.IP    `json:"monitoring_server_ipv4,omitempty"`
	MonitoringServerIPv6       *net.IP    `json:"monitoring_server_ipv6,omitempty"`
	Location                   string     `json:"location,omitempty"`
	Output                     string     `json:"output,omitempty"`
	StateIsUp                  bool       `json:"state_is_up,omitempty"`
	Ignored                    bool       `json:"ignored,omitempty"`
	CheckPK                    int64      `json:"check_pk,omitempty"`
	CheckURL                   string     `json:"check_url,omitempty"`
	CheckAddress               string     `json:"check_address,omitempty"`
	CheckName                  string     `json:"check_name,omitempty"`
	CheckMonitoringServiceType string     `json:"check_monitoring_service_type,omitempty"`
}

func (a AlertItem) PrimaryKey() PrimaryKey {
	return PrimaryKey(a.PK)
}

// AlertListResponse represents a paginated list of alerts.
type AlertListResponse struct {
	Count    int64       `json:"count,omitempty"`
	Next     string      `json:"next,omitempty"`
	Previous string      `json:"previous,omitempty"`
	Results  []AlertItem `json:"results,omitempty"`
}

func (r AlertListResponse) List() []AlertItem {
	return r.Results
}

// AlertListOptions specifies the optional parameters for listing alerts.
type AlertListOptions struct {
	Page                       int64  `url:"page,omitempty"`
	PageSize                   int64  `url:"page_size,omitempty"`
	Search                     string `url:"search,omitempty"`
	Ordering                   string `url:"ordering,omitempty"`
	StateIsUp                  *bool  `url:"state_is_up,omitempty"`
	CheckPK                    int64  `url:"check_pk,omitempty"`
	CheckMonitoringServiceType string `url:"check_monitoring_service_type,omitempty"`
	CheckTag                   string `url:"check_tag,omitempty"`
	StartDate                  string `url:"start_date,omitempty"`
	EndDate                    string `url:"end_date,omitempty"`
}

// AlertResponse represents a single alert response.
type AlertResponse AlertItem

func (r AlertResponse) Item() AlertItem {
	return AlertItem(r)
}

// AlertRootCause represents root cause analysis data for an alert.
type AlertRootCause struct {
	PK                   int64      `json:"pk,omitempty"`
	URL                  string     `json:"url,omitempty"`
	CreatedAt            *time.Time `json:"created_at,omitempty"`
	MonitoringServerName string     `json:"monitoring_server_name,omitempty"`
	MonitoringServerIPv4 *net.IP    `json:"monitoring_server_ipv4,omitempty"`
	MonitoringServerIPv6 *net.IP    `json:"monitoring_server_ipv6,omitempty"`
	Location             string     `json:"location,omitempty"`
	Output               string     `json:"output,omitempty"`
	RootCauseData        string     `json:"root_cause_data,omitempty"`
}

// AlertsEndpoint defines the interface for interacting with alert resources.
type AlertsEndpoint interface {
	List(context.Context, AlertListOptions) ([]AlertItem, error)
	Get(context.Context, PrimaryKeyable) (*AlertItem, error)
	RootCause(context.Context, PrimaryKeyable) (*AlertRootCause, error)
	Ignore(context.Context, PrimaryKeyable) (*AlertItem, error)
}

// NewAlertsEndpoint creates a new alerts endpoint.
func NewAlertsEndpoint(cbd CBD) AlertsEndpoint {
	const endpoint = "alerts"
	return &alertsEndpointImpl{
		cbd:            cbd,
		endpoint:       endpoint,
		EndpointLister: NewEndpointLister[AlertListResponse, AlertItem, AlertListOptions](cbd, endpoint),
		EndpointGetter: NewEndpointGetter[AlertResponse, AlertItem](cbd, endpoint),
	}
}

type alertsEndpointImpl struct {
	cbd      CBD
	endpoint string
	EndpointLister[AlertListResponse, AlertItem, AlertListOptions]
	EndpointGetter[AlertResponse, AlertItem]
}

// RootCause retrieves the root cause analysis data for a specific alert.
func (e *alertsEndpointImpl) RootCause(ctx context.Context, pk PrimaryKeyable) (*AlertRootCause, error) {
	path := fmt.Sprintf("%s/alert/%d/root-cause/", e.endpoint, pk.PrimaryKey())
	req, err := e.cbd.BuildRequest(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, err
	}
	rs, err := e.cbd.Do(req)
	if err != nil {
		return nil, err
	}
	defer rs.Body.Close()
	if rs.StatusCode != 200 {
		return nil, ErrorFromResponse(rs)
	}
	var resp AlertRootCause
	if err := e.cbd.DecodeResponse(rs, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Ignore toggles the ignore state of an alert.
func (e *alertsEndpointImpl) Ignore(ctx context.Context, pk PrimaryKeyable) (*AlertItem, error) {
	path := fmt.Sprintf("%s/%d/ignore/", e.endpoint, pk.PrimaryKey())
	req, err := e.cbd.BuildRequest(ctx, "POST", path, nil, nil)
	if err != nil {
		return nil, err
	}
	rs, err := e.cbd.Do(req)
	if err != nil {
		return nil, err
	}
	defer rs.Body.Close()
	if rs.StatusCode != 200 {
		return nil, ErrorFromResponse(rs)
	}
	var resp AlertResponse
	if err := e.cbd.DecodeResponse(rs, &resp); err != nil {
		return nil, err
	}
	alert := resp.Item()
	return &alert, nil
}
