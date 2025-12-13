package upapi

import (
	"context"

	"github.com/shopspring/decimal"
)

type CheckDNS struct {
	Name                   string          `json:"name,omitempty"`
	ContactGroups          *[]string       `json:"contact_groups,omitempty"`
	Locations              []string        `json:"locations,omitempty"`
	Tags                   []string        `json:"tags,omitempty"`
	IsPaused               bool            `json:"is_paused"`
	Interval               int64           `json:"msp_interval,omitempty"`
	Address                string          `json:"msp_address"`
	DNSServer              string          `json:"msp_dns_server,omitempty"`
	DNSRecordType          string          `json:"msp_dns_record_type,omitempty"`
	ExpectString           string          `json:"msp_expect_string,omitempty"`
	Threshold              int64           `json:"msp_threshold,omitempty"`
	Sensitivity            int64           `json:"msp_sensitivity,omitempty"`
	NumRetries             int64           `json:"msp_num_retries,omitempty"`
	UptimeSLA              decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string          `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool            `json:"msp_include_in_global_metrics"`
}

type checksEndpointDNSImpl struct {
	EndpointCreator[CheckDNS, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckDNS, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointDNSImpl) CreateDNS(ctx context.Context, check CheckDNS) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointDNSImpl) UpdateDNS(ctx context.Context, pk PrimaryKeyable, check CheckDNS) (*Check, error) {
	return c.Update(ctx, pk, check)
}
