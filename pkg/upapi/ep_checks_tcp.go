package upapi

import (
	"context"

	"github.com/shopspring/decimal"
)

type CheckTCP struct {
	Name                   string          `json:"name,omitempty"`
	ContactGroups          *[]string       `json:"contact_groups,omitempty"`
	Locations              []string        `json:"locations,omitempty"`
	Tags                   []string        `json:"tags,omitempty"`
	IsPaused               bool            `json:"is_paused"`
	Interval               int64           `json:"msp_interval,omitempty"`
	Address                string          `json:"msp_address"`
	Port                   int64           `json:"msp_port,omitempty"`
	SendString             string          `json:"msp_send_string,omitempty"`
	ExpectString           string          `json:"msp_expect_string,omitempty"`
	Sensitivity            int64           `json:"msp_sensitivity,omitempty"`
	NumRetries             int64           `json:"msp_num_retries,omitempty"`
	UseIpVersion           string          `json:"msp_use_ip_version,omitempty"`
	UptimeSLA              decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string          `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool            `json:"msp_include_in_global_metrics"`
	Encryption             string          `json:"msp_encryption"`
}

type checksEndpointTCPImpl struct {
	EndpointCreator[CheckTCP, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckTCP, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointTCPImpl) CreateTCP(ctx context.Context, check CheckTCP) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointTCPImpl) UpdateTCP(ctx context.Context, pk PrimaryKeyable, check CheckTCP) (*Check, error) {
	return c.Update(ctx, pk, check)
}
