package upapi

import (
	"context"

	"github.com/shopspring/decimal"
)

type CheckSSH struct {
	Name                   string          `json:"name,omitempty"`
	ContactGroups          *[]string       `json:"contact_groups,omitempty"`
	Locations              []string        `json:"locations,omitempty"`
	Tags                   []string        `json:"tags,omitempty"`
	IsPaused               bool            `json:"is_paused"`
	Interval               int64           `json:"msp_interval,omitempty"`
	Address                string          `json:"msp_address"`
	Port                   int64           `json:"msp_port,omitempty"`
	Sensitivity            int64           `json:"msp_sensitivity,omitempty"`
	NumRetries             int64           `json:"msp_num_retries,omitempty"`
	UseIpVersion           string          `json:"msp_use_ip_version,omitempty"`
	UptimeSLA              decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string          `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool            `json:"msp_include_in_global_metrics"`
}

type checksEndpointSSHImpl struct {
	EndpointCreator[CheckSSH, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckSSH, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointSSHImpl) CreateSSH(ctx context.Context, check CheckSSH) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointSSHImpl) UpdateSSH(ctx context.Context, pk PrimaryKeyable, check CheckSSH) (*Check, error) {
	return c.Update(ctx, pk, check)
}
