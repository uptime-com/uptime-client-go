package upapi

import (
	"bytes"
	"context"
	"encoding/json"
)

// CloudStatusGroup references a cloud status group by ID. On write
// (POST/PUT) it serializes as the group's integer ID, matching the legacy
// request format. On read (GET) the server returns a nested
// `{"id": ..., "name": ...}` object, so UnmarshalJSON accepts either
// shape - bare integer or nested object.
type CloudStatusGroup struct {
	ID   int64  `json:"id"`
	Name string `json:"name,omitempty"`
}

func (g CloudStatusGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.ID)
}

func (g *CloudStatusGroup) UnmarshalJSON(data []byte) error {
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) == 0 || string(trimmed) == "null" {
		return nil
	}
	if trimmed[0] == '{' {
		type raw CloudStatusGroup
		var r raw
		if err := json.Unmarshal(data, &r); err != nil {
			return err
		}
		*g = CloudStatusGroup(r)
		return nil
	}
	var id int64
	if err := json.Unmarshal(data, &id); err != nil {
		return err
	}
	g.ID = id
	return nil
}

// CheckCloudStatusConfig describes the cloudstatusconfig payload accepted by
// the /checks/add-cloudstatus and /checks/{pk} endpoints.
//
// There are two ways to configure what is monitored:
//
//  1. Legacy: set ServiceName to a single component name. Deprecated server-side
//     in favour of the group + monitoring_type model below.
//  2. Group-based: set Group to the cloud status group ID and MonitoringType to
//     either "ALL" (every service in the group) or "SPECIFIC" (only entries
//     listed in Services and/or ServiceTitles).
//
// Use the GET /api/v1/checks/cloudstatus-groups/ and
// /api/v1/checks/cloudstatus-services/ endpoints to discover valid IDs.
type CheckCloudStatusConfig struct {
	// NotifyOnlyOnDown opts out of maintenance notifications.
	NotifyOnlyOnDown bool `json:"notify_only_on_down,omitempty"`

	// ServiceName is the legacy (deprecated) cloud status component name or ID.
	ServiceName string `json:"service_name,omitempty"`

	// Group is the cloud status group to monitor. On write only the numeric
	// ID is sent; on read the server also reports the group name.
	Group *CloudStatusGroup `json:"group,omitempty"`

	// MonitoringType selects how Group is monitored: "ALL" or "SPECIFIC".
	MonitoringType string `json:"monitoring_type,omitempty"`

	// Services is the list of specific service IDs to monitor when
	// MonitoringType is "SPECIFIC".
	Services []int64 `json:"services,omitempty"`

	// ServiceTitles auto-monitors current and future services whose title
	// matches any entry in this list when MonitoringType is "SPECIFIC".
	ServiceTitles []string `json:"service_titles,omitempty"`
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
