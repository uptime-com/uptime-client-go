package uptime

import (
	"context"
	"fmt"

	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/domainr/whois"
	whoisparser "github.com/likexian/whois-parser"
)

type CheckService service

// Check represents a check in Uptime.com.
type Check struct {
	PK        int    `json:"pk,omitempty"`
	CheckType string `json:"check_type,omitempty"`
	URL       string `json:"url,omitempty"`
	Name      string `json:"name,omitempty"`

	Address   string `json:"msp_address,omitempty"`
	Port      int    `json:"msp_port,omitempty"`
	IPVersion string `json:"msp_use_ip_version,omitempty"`

	Interval    int      `json:"msp_interval,omitempty"`
	Locations   []string `json:"locations,omitempty"`
	Sensitivity int      `json:"msp_sensitivity,omitempty"`
	Threshold   int      `json:"msp_threshold,omitempty"`

	Headers      string `json:"msp_headers,omitempty"`
	Username     string `json:"msp_username,omitempty"`
	Password     string `json:"msp_password,omitempty"`
	SendString   string `json:"msp_send_string,omitempty"`
	ExpectString string `json:"msp_expect_string,omitempty"`

	ContactGroups []string `json:"contact_groups,omitempty"`
	Tags          []string `json:"tags,omitempty"`
	Escalations   []string `json:"escalations,omitempty"`

	Notes                  string `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool   `json:"msp_include_in_global_metrics,omitempty"`

	// For DNS checks
	DNSServer     string `json:"msp_dns_server,omitempty"`
	DNSRecordType string `json:"msp_dns_record_type,omitempty"`

	// For IMAP, POP checks
	Encryption string `json:"msp_encrytion,omitempty"`

	// For Transaction checks
	Script string `json:"msp_script,omitempty"`

	// For SSL checks
	Protocol string `json:"msp_protocol,omitempty"`
}

type CheckListResponse struct {
	Count    int      `json:"count,omitempty"`
	Next     string   `json:"next,omitempty"`
	Previous string   `json:"previous,omitempty"`
	Results  []*Check `json:"results,omitempty"`
}

// CheckListOptions specifies the optional parameters to the CheckService.List method.
type CheckListOptions struct {
	Page                  int      `url:"page,omitempty"`
	PageSize              int      `url:"page_size,omitempty"`
	Search                string   `url:"search,omitempty"`
	Ordering              string   `url:"ordering,omitempty"`
	MonitoringServiceType string   `url:"monitoring_service_type,omitempty"`
	IsPaused              bool     `url:"is_paused,omitempty"`
	StateIsUp             bool     `url:"state_is_up,omitempty"`
	Tag                   []string `url:"tag,omitempty"`
}

type CheckResponse struct {
	Messages map[string]interface{} `json:"messages,omitempty"`
	Results  Check                  `json:"results,omitempty"`
}

func (s *CheckService) List(ctx context.Context, opt *CheckListOptions) ([]*Check, *http.Response, error) {
	u := "checks"
	clResp, resp, err := s.listChecks(ctx, u, opt)
	return clResp.Results, resp, err
}

func (s *CheckService) ListAll(ctx context.Context, opt *CheckListOptions) ([]*Check, error) {
	u := "checks"
	opt.Page = 1

	result := []*Check{}

	clResp, _, err := s.listChecks(ctx, u, opt)
	if err != nil {
		return nil, err
	}
	result = append(result, clResp.Results...)

	for clResp.Next != "" {
		opt.Page++
		clResp, _, err = s.listChecks(ctx, u, opt)
		if err != nil {
			return nil, err
		}
		result = append(result, clResp.Results...)
	}

	return result, err
}

func (s *CheckService) listChecks(ctx context.Context, url string, opt *CheckListOptions) (*CheckListResponse, *http.Response, error) {
	u, err := addOptions(url, opt)
	if err != nil {
		return nil, nil, err
	}

	req, _ := s.client.NewRequest("GET", u, nil)

	var checks CheckListResponse
	resp, err := s.client.Do(ctx, req, &checks)
	if err != nil {
		return nil, nil, err
	}
	return &checks, resp, nil
}

// Create a new check in Uptime.com based on the provided Check.
func (s *CheckService) Create(ctx context.Context, check *Check) (*Check, *http.Response, error) {
	suffix := strings.ToLower(strings.Replace(check.CheckType, "_", "-", -1))
	u := fmt.Sprintf("checks/add-%v", suffix)

	// Get a Whois expected string if we don't have one.
	if check.CheckType == "WHOIS" && check.ExpectString == "" {
		s, err := uptimeWhoisCheckExpectString(check.Address)
		if err != nil {
			return nil, nil, err
		}
		check.ExpectString = s
	}

	req, err := s.client.NewRequest("POST", u, check)
	if err != nil {
		return nil, nil, err
	}

	cr := &CheckResponse{}
	resp, err := s.client.Do(ctx, req, cr)
	if err != nil {
		return nil, resp, err
	}

	return &cr.Results, resp, nil
}

func (s *CheckService) Get(ctx context.Context, pk int) (*Check, *http.Response, error) {
	u := fmt.Sprintf("checks/%v", pk)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	c := &Check{}
	resp, err := s.client.Do(ctx, req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, nil
}

// Update a check.
func (s *CheckService) Update(ctx context.Context, check *Check) (*Check, *http.Response, error) {
	u := fmt.Sprintf("checks/%v", check.PK)
	if check.PK == 0 {
		return nil, nil, fmt.Errorf("Error updating check with empty PK")
	}

	req, err := s.client.NewRequest("PATCH", u, check)
	if err != nil {
		return nil, nil, err
	}

	cr := &CheckResponse{}
	resp, err := s.client.Do(ctx, req, cr)
	if err != nil {
		return nil, resp, err
	}

	return &cr.Results, resp, nil
}

// Delete a check.
func (s *CheckService) Delete(ctx context.Context, pk int) (*http.Response, error) {
	u := fmt.Sprintf("checks/%v", pk)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

// CheckStatsOptions specifies the parameters to the CheckService.Stats method.
type CheckStatsOptions struct {
	StartDate              string
	EndDate                string
	Location               string
	LocationsResponseTimes bool
	IncludeAlerts          bool
	Download               bool
	PDF                    bool
}

// CheckStatsResponse represents the API's response to a Stats query.
type CheckStatsResponse struct {
	StartDate  string           `json:"start_date"`
	EndDate    string           `json:"end_date"`
	Statistics []*CheckStats    `json:"statistics"`
	Totals     CheckStatsTotals `json:"totals"`
}

// CheckStats represents the check statistics for a given day.
type CheckStats struct {
	Date         string `json:"date"`
	Outages      int    `json:"outages"`
	DowntimeSecs int    `json:"downtime_secs"`
}

// CheckStatsTotals represents the 'totals' section of check statistics in Uptime.com.
type CheckStatsTotals struct {
	Outages      int   `json:"outages,omitempty"`
	DowntimeSecs int64 `json:"downtime_secs,omitempty"`
}

// Stats gets statistics on the specified check.
func (s *CheckService) Stats(ctx context.Context, pk int, opt *CheckStatsOptions) (*CheckStatsResponse, *http.Response, error) {
	u := fmt.Sprintf("checks/%v/stats/?start_date=%s&end_date=%s&location=%s&locations_response_times=%t&include_alerts=%t&download=%t&pdf=%t",
		pk,
		opt.StartDate,
		opt.EndDate,
		url.QueryEscape(opt.Location),
		opt.LocationsResponseTimes,
		opt.IncludeAlerts,
		opt.Download,
		opt.PDF)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	c := &CheckStatsResponse{}
	resp, err := s.client.Do(ctx, req, c)
	if err != nil {
		return nil, resp, err
	}
	return c, resp, nil
}

// When creating a new Whois check through the UI, Uptime.com automatically generates
// a custom-formatted string describing elements to verify in a Whois response. This
// function replicates that behavior.
func uptimeWhoisCheckExpectString(domain string) (string, error) {
	req, err := whois.NewRequest(domain)
	if err != nil {
		return "", fmt.Errorf("Error constructing request for %v: %v", domain, err)
	}

	resp, err := whois.DefaultClient.Fetch(req)
	if err != nil {
		return "", fmt.Errorf("Error getting whois for %v: %v", domain, err)
	}

	result, err := whoisparser.Parse(resp.String())
	if err != nil {
		return "", fmt.Errorf("Error parsing raw Whois from %s: %v", domain, err)
	}

	expiry := uptimeWhoisTimeStr(result.Domain.ExpirationDate)
	nameServers := result.Domain.NameServers
	name := strings.ToLower(result.Registrar.Name)
	whoisInfo := fmt.Sprintf("expires: %s\nnameservers: %s\nregistrar: %s", expiry, nameServers, name)

	return whoisInfo, nil
}

func uptimeWhoisTimeStr(ts string) string {
	t, _ := time.Parse(time.RFC3339, ts)
	return t.Format("2006-01-02")
}
