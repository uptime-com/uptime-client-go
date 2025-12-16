package upapi

import (
	"context"

	"github.com/shopspring/decimal"
)

type CheckHeartbeat struct {
	Name                   string          `json:"name,omitempty"`
	ContactGroups          *[]string       `json:"contact_groups,omitempty"`
	Tags                   []string        `json:"tags,omitempty"`
	IsPaused               bool            `json:"is_paused"`
	Interval               int64           `json:"msp_interval,omitempty"`
	UptimeSLA              decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string          `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool            `json:"msp_include_in_global_metrics"`
	HeartbeatURL           string          `json:"heartbeat_url,omitempty"`
}

type checksEndpointHeartbeatImpl struct {
	EndpointCreator[CheckHeartbeat, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckHeartbeat, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointHeartbeatImpl) CreateHeartbeat(ctx context.Context, check CheckHeartbeat) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointHeartbeatImpl) UpdateHeartbeat(ctx context.Context, pk PrimaryKeyable, check CheckHeartbeat) (*Check, error) {
	return c.Update(ctx, pk, check)
}
