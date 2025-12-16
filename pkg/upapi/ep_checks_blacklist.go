package upapi

import (
	"context"

	"github.com/shopspring/decimal"
)

type CheckBlacklist struct {
	Name          string          `json:"name,omitempty"`
	ContactGroups *[]string       `json:"contact_groups,omitempty"`
	Locations     []string        `json:"locations,omitempty"`
	Tags          []string        `json:"tags,omitempty"`
	IsPaused      bool            `json:"is_paused"`
	Address       string          `json:"msp_address"`
	NumRetries    int64           `json:"msp_num_retries,omitempty"`
	UptimeSLA     decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	Notes         string          `json:"msp_notes,omitempty"`
}

type checksEndpointBlacklistImpl struct {
	EndpointCreator[CheckBlacklist, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckBlacklist, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointBlacklistImpl) CreateBlacklist(ctx context.Context, check CheckBlacklist) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointBlacklistImpl) UpdateBlacklist(ctx context.Context, pk PrimaryKeyable, check CheckBlacklist) (*Check, error) {
	return c.Update(ctx, pk, check)
}
