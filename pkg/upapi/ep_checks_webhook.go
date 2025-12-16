package upapi

import (
	"context"

	"github.com/shopspring/decimal"
)

type CheckWebhook struct {
	Name                   string          `json:"name,omitempty"`
	ContactGroups          *[]string       `json:"contact_groups,omitempty"`
	Locations              []string        `json:"locations,omitempty"`
	Tags                   []string        `json:"tags,omitempty"`
	IsPaused               bool            `json:"is_paused"`
	UptimeSLA              decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string          `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool            `json:"msp_include_in_global_metrics"`
	WebhookUrl             string          `json:"webhook_url,omitempty"`
}

type checksEndpointWebhookImpl struct {
	EndpointCreator[CheckWebhook, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckWebhook, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointWebhookImpl) CreateWebhook(ctx context.Context, check CheckWebhook) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointWebhookImpl) UpdateWebhook(ctx context.Context, pk PrimaryKeyable, check CheckWebhook) (*Check, error) {
	return c.Update(ctx, pk, check)
}
