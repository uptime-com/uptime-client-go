package upapi

import (
	"context"

	"github.com/shopspring/decimal"
)

type CheckSMTP struct {
	Name                   string          `json:"name,omitempty"`
	ContactGroups          *[]string       `json:"contact_groups,omitempty"`
	Locations              []string        `json:"locations,omitempty"`
	Tags                   []string        `json:"tags,omitempty"`
	IsPaused               bool            `json:"is_paused"`
	Interval               int64           `json:"msp_interval"`
	Address                string          `json:"msp_address"`
	Port                   int64           `json:"msp_port,omitempty"`
	Username               string          `json:"msp_username,omitempty"`
	Password               string          `json:"msp_password,omitempty"`
	ExpectString           string          `json:"msp_expect_string,omitempty"`
	Encryption             string          `json:"msp_encryption,omitempty"`
	Sensitivity            int64           `json:"msp_sensitivity,omitempty"`
	NumRetries             int64           `json:"msp_num_retries,omitempty"`
	UseIpVersion           string          `json:"msp_use_ip_version,omitempty"`
	UptimeSLA              decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string          `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool            `json:"msp_include_in_global_metrics"`
}

type checksEndpointSMTPImpl struct {
	EndpointCreator[CheckSMTP, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckSMTP, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointSMTPImpl) CreateSMTP(ctx context.Context, check CheckSMTP) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointSMTPImpl) UpdateSMTP(ctx context.Context, pk PrimaryKeyable, check CheckSMTP) (*Check, error) {
	return c.Update(ctx, pk, check)
}
