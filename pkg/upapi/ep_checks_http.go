package upapi

import (
	"context"

	"github.com/shopspring/decimal"
)

type CheckHTTP struct {
	Name                   string          `json:"name,omitempty"`
	ContactGroups          *[]string       `json:"contact_groups,omitempty"`
	Locations              []string        `json:"locations,omitempty"`
	Tags                   []string        `json:"tags,omitempty"`
	IsPaused               bool            `json:"is_paused"`
	Interval               int64           `json:"msp_interval,omitempty"`
	Address                string          `json:"msp_address"`
	Port                   int64           `json:"msp_port,omitempty"`
	Username               string          `json:"msp_username,omitempty"`
	Password               string          `json:"msp_password,omitempty"`
	Proxy                  string          `json:"msp_proxy,omitempty"`
	StatusCode             string          `json:"msp_status_code"`
	SendString             string          `json:"msp_send_string,omitempty"`
	ExpectString           string          `json:"msp_expect_string,omitempty"`
	ExpectStringType       string          `json:"msp_expect_string_type,omitempty"`
	Encryption             string          `json:"msp_encryption"`
	Threshold              int64           `json:"msp_threshold,omitempty"`
	Headers                string          `json:"msp_headers,omitempty"`
	Version                int64           `json:"msp_version,omitempty"`
	Sensitivity            int64           `json:"msp_sensitivity,omitempty"`
	NumRetries             int64           `json:"msp_num_retries,omitempty"`
	UseIPVersion           string          `json:"msp_use_ip_version,omitempty"`
	UptimeSLA              decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string          `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool            `json:"msp_include_in_global_metrics"`
}

type checksEndpointHTTPImpl struct {
	EndpointCreator[CheckHTTP, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckHTTP, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointHTTPImpl) CreateHTTP(ctx context.Context, check CheckHTTP) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointHTTPImpl) UpdateHTTP(ctx context.Context, pk PrimaryKeyable, check CheckHTTP) (*Check, error) {
	return c.Update(ctx, pk, check)
}
