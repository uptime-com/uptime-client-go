package upapi

import (
	"context"

	"github.com/shopspring/decimal"
)

type CheckIMAP struct {
	Name                   string          `json:"name,omitempty"`
	ContactGroups          *[]string       `json:"contact_groups,omitempty"`
	Locations              []string        `json:"locations,omitempty"`
	Tags                   []string        `json:"tags,omitempty"`
	IsPaused               bool            `json:"is_paused"`
	Interval               int64           `json:"msp_interval,omitempty"`
	Address                string          `json:"msp_address"`
	Port                   int64           `json:"msp_port,omitempty"`
	ExpectString           string          `json:"msp_expect_string,omitempty"`
	Encryption             string          `json:"msp_encryption,omitempty"`
	Sensitivity            int64           `json:"msp_sensitivity,omitempty"`
	NumRetries             int64           `json:"msp_num_retries,omitempty"`
	UseIPVersion           string          `json:"msp_use_ip_version,omitempty"`
	UptimeSLA              decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string          `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool            `json:"msp_include_in_global_metrics"`
}

type checksEndpointIMAPImpl struct {
	EndpointCreator[CheckIMAP, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckIMAP, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointIMAPImpl) CreateIMAP(ctx context.Context, check CheckIMAP) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointIMAPImpl) UpdateIMAP(ctx context.Context, pk PrimaryKeyable, check CheckIMAP) (*Check, error) {
	return c.Update(ctx, pk, check)
}
