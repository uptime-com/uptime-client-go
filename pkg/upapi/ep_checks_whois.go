package upapi

import (
	"context"

	"github.com/shopspring/decimal"
)

type CheckWHOIS struct {
	Name          string          `json:"name,omitempty"`
	ContactGroups *[]string       `json:"contact_groups,omitempty"`
	Locations     []string        `json:"locations,omitempty"`
	Tags          []string        `json:"tags,omitempty"`
	IsPaused      bool            `json:"is_paused"`
	Address       string          `json:"msp_address"`
	ExpectString  string          `json:"msp_expect_string,omitempty"`
	Threshold     int64           `json:"msp_threshold,omitempty"`
	NumRetries    int64           `json:"msp_num_retries,omitempty"`
	UptimeSLA     decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	Notes         string          `json:"msp_notes,omitempty"`
}

type checksEndpointWHOISImpl struct {
	EndpointCreator[CheckWHOIS, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckWHOIS, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointWHOISImpl) CreateWHOIS(ctx context.Context, check CheckWHOIS) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointWHOISImpl) UpdateWHOIS(ctx context.Context, pk PrimaryKeyable, check CheckWHOIS) (*Check, error) {
	return c.Update(ctx, pk, check)
}
