package upapi

import (
	"context"

	"github.com/shopspring/decimal"
)

type CheckNTP struct {
	Name                   string          `json:"name,omitempty"`
	ContactGroups          *[]string       `json:"contact_groups,omitempty"`
	Locations              []string        `json:"locations,omitempty"`
	Tags                   []string        `json:"tags,omitempty"`
	IsPaused               bool            `json:"is_paused"`
	Interval               int64           `json:"msp_interval"`
	Address                string          `json:"msp_address"`
	Port                   int64           `json:"msp_port,omitempty"`
	Threshold              int64           `json:"msp_threshold,omitempty"`
	Sensitivity            int64           `json:"msp_sensitivity,omitempty"`
	NumRetries             int64           `json:"msp_num_retries,omitempty"`
	UseIPVersion           string          `json:"msp_use_ip_version,omitempty"`
	UptimeSLA              decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string          `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool            `json:"msp_include_in_global_metrics"`
}

type checksEndpointNTPImpl struct {
	EndpointCreator[CheckNTP, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckNTP, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointNTPImpl) CreateNTP(ctx context.Context, check CheckNTP) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointNTPImpl) UpdateNTP(ctx context.Context, pk PrimaryKeyable, check CheckNTP) (*Check, error) {
	return c.Update(ctx, pk, check)
}
