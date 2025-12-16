package upapi

import (
	"context"

	"github.com/shopspring/decimal"
)

type CheckRUM struct {
	Name                   string          `json:"name,omitempty"`
	ContactGroups          *[]string       `json:"contact_groups,omitempty"`
	Locations              []string        `json:"locations,omitempty"`
	Tags                   []string        `json:"tags,omitempty"`
	IsPaused               bool            `json:"is_paused"`
	Address                string          `json:"msp_address"`
	Threshold              int64           `json:"msp_threshold,omitempty"`
	UptimeSLA              decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	Notes                  string          `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool            `json:"msp_include_in_global_metrics"`
}

type checksEndpointRUMImpl struct {
	EndpointCreator[CheckRUM, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckRUM, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointRUMImpl) CreateRUM(ctx context.Context, check CheckRUM) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointRUMImpl) UpdateRUM(ctx context.Context, pk PrimaryKeyable, check CheckRUM) (*Check, error) {
	return c.Update(ctx, pk, check)
}
