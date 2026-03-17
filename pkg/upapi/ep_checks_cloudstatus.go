package upapi

import "context"

type CheckCloudStatusConfig struct {
	ServiceName string `json:"service_name,omitempty"`
}

type CheckCloudStatus struct {
	Name              string                 `json:"name,omitempty"`
	ContactGroups     *[]string              `json:"contact_groups,omitempty"`
	Locations         []string               `json:"locations,omitempty"`
	Tags              []string               `json:"tags,omitempty"`
	IsPaused          *bool                  `json:"is_paused,omitempty"`
	CloudStatusConfig CheckCloudStatusConfig `json:"cloudstatusconfig,omitempty"`
}

type checksEndpointCloudStatusImpl struct {
	EndpointCreator[CheckCloudStatus, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckCloudStatus, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointCloudStatusImpl) CreateCloudStatus(ctx context.Context, check CheckCloudStatus) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointCloudStatusImpl) UpdateCloudStatus(ctx context.Context, pk PrimaryKeyable, check CheckCloudStatus) (*Check, error) {
	return c.Update(ctx, pk, check)
}
