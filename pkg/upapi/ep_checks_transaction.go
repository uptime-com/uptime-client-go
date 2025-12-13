package upapi

import (
	"context"

	"github.com/shopspring/decimal"
)

type CheckTransaction struct {
	Name                   string          `json:"name,omitempty"`
	ContactGroups          *[]string       `json:"contact_groups,omitempty"`
	Locations              []string        `json:"locations,omitempty"`
	Tags                   []string        `json:"tags,omitempty"`
	IsPaused               bool            `json:"is_paused"`
	Interval               int64           `json:"msp_interval,omitempty"`
	Threshold              int64           `json:"msp_threshold,omitempty"`
	Script                 string          `json:"msp_script,omitempty"`
	Sensitivity            int64           `json:"msp_sensitivity,omitempty"`
	NumRetries             int64           `json:"msp_num_retries,omitempty"`
	UptimeSLA              decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string          `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool            `json:"msp_include_in_global_metrics"`
}

type checksEndpointTransactionImpl struct {
	EndpointCreator[CheckTransaction, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckTransaction, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointTransactionImpl) CreateTransaction(ctx context.Context, check CheckTransaction) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointTransactionImpl) UpdateTransaction(ctx context.Context, pk PrimaryKeyable, check CheckTransaction) (*Check, error) {
	return c.Update(ctx, pk, check)
}
