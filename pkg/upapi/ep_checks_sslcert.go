package upapi

import (
	"context"

	"github.com/shopspring/decimal"
)

type CheckSSLCert struct {
	Name          string             `json:"name,omitempty"`
	ContactGroups *[]string          `json:"contact_groups,omitempty"`
	Locations     []string           `json:"locations,omitempty"`
	Tags          []string           `json:"tags,omitempty"`
	IsPaused      bool               `json:"is_paused"`
	Protocol      string             `json:"msp_protocol,omitempty"`
	Address       string             `json:"msp_address"`
	Port          int64              `json:"msp_port,omitempty"`
	Threshold     int64              `json:"msp_threshold"`
	NumRetries    int64              `json:"msp_num_retries,omitempty"`
	UptimeSLA     decimal.Decimal    `json:"msp_uptime_sla,omitempty"`
	Notes         string             `json:"msp_notes,omitempty"`
	SSLConfig     CheckSSLCertConfig `json:"sslconfig,omitempty"`
}

type checksEndpointSSLCertImpl struct {
	EndpointCreator[CheckSSLCert, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckSSLCert, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointSSLCertImpl) CreateSSLCert(ctx context.Context, check CheckSSLCert) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointSSLCertImpl) UpdateSSLCert(ctx context.Context, pk PrimaryKeyable, check CheckSSLCert) (*Check, error) {
	return c.Update(ctx, pk, check)
}
