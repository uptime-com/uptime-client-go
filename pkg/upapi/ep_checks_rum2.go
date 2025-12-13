package upapi

import (
	"context"

	"github.com/shopspring/decimal"
)

type CheckRUM2 struct {
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

type checksEndpointRUM2Impl struct {
	EndpointCreator[CheckRUM2, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckRUM2, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointRUM2Impl) CreateRUM2(ctx context.Context, check CheckRUM2) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointRUM2Impl) UpdateRUM2(ctx context.Context, pk PrimaryKeyable, check CheckRUM2) (*Check, error) {
	return c.Update(ctx, pk, check)
}
