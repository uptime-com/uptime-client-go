package upapi

import (
	"context"

	"github.com/shopspring/decimal"
)

type CheckGroup struct {
	Name                   string           `json:"name,omitempty"`
	ContactGroups          *[]string        `json:"contact_groups,omitempty"`
	Locations              []string         `json:"locations,omitempty"`
	Tags                   []string         `json:"tags,omitempty"`
	IsPaused               bool             `json:"is_paused"`
	UptimeSLA              decimal.Decimal  `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        decimal.Decimal  `json:"msp_response_time_sla,omitempty"`
	Notes                  string           `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool             `json:"msp_include_in_global_metrics"`
	Config                 CheckGroupConfig `json:"groupcheckconfig,omitempty"`
}

type checksEndpointGroupImpl struct {
	EndpointCreator[CheckGroup, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckGroup, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointGroupImpl) CreateGroup(ctx context.Context, check CheckGroup) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointGroupImpl) UpdateGroup(ctx context.Context, pk PrimaryKeyable, check CheckGroup) (*Check, error) {
	return c.Update(ctx, pk, check)
}
