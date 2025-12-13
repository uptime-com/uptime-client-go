package upapi

import "context"

type CheckPageSpeed struct {
	Name          string               `json:"name,omitempty"`
	ContactGroups *[]string            `json:"contact_groups,omitempty"`
	Locations     []string             `json:"locations,omitempty"`
	Tags          []string             `json:"tags,omitempty"`
	IsPaused      bool                 `json:"is_paused"`
	Address       string               `json:"msp_address"`
	Interval      int64                `json:"msp_interval,omitempty"`
	Username      string               `json:"msp_username,omitempty"`
	Password      string               `json:"msp_password,omitempty"`
	Headers       string               `json:"msp_headers,omitempty"`
	Script        string               `json:"msp_script,omitempty"`
	NumRetries    int64                `json:"msp_num_retries,omitempty"`
	Notes         string               `json:"msp_notes,omitempty"`
	Config        CheckPageSpeedConfig `json:"pagespeedconfig,omitempty"`
}

type checksEndpointPageSpeedImpl struct {
	EndpointCreator[CheckPageSpeed, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckPageSpeed, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointPageSpeedImpl) CreatePageSpeed(ctx context.Context, check CheckPageSpeed) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointPageSpeedImpl) UpdatePageSpeed(ctx context.Context, pk PrimaryKeyable, check CheckPageSpeed) (*Check, error) {
	return c.Update(ctx, pk, check)
}
