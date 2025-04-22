package upapi

import (
	"context"
)

type IncidentUpdate struct {
	PK            int64  `json:"id,omitempty"`
	URL           string `json:"url,omitempty"`
	Description   string `json:"description"`
	IncidentState string `json:"incident_state"`
	UpdatedAt     string `json:"updated_at"`
}

type IncidentAffectedComponent struct {
	PK int64 `json:"id"`
}

type IncidentAffectedComponentEntity struct {
	PK        int64                     `json:"id"`
	Status    string                    `json:"status"`
	Component IncidentAffectedComponent `json:"component"`
}

type StatusPageIncident struct {
	PK   int64  `json:"pk"`
	URL  string `json:"url,omitempty"`
	Name string `json:"name"`

	IncludeInGlobalMetrics bool `json:"include_in_global_metrics,omitempty"`

	Updates            []IncidentUpdate                  `json:"updates,omitempty"`
	AffectedComponents []IncidentAffectedComponentEntity `json:"affected_components,omitempty"`

	IncidentType                     string `json:"incident_type,omitempty"`
	StartsAt                         string `json:"starts_at"`
	EndsAt                           string `json:"ends_at,omitempty"`
	UpdateComponentStatus            bool   `json:"update_component_status,omitempty"`
	NotifySubscribers                bool   `json:"notify_subscribers,omitempty"`
	SendMaintenanceStartNotification bool   `json:"send_maintenance_start_notification,omitempty"`

	IsGroup        bool   `json:"is_group,omitempty"`
	GroupID        int64  `json:"group_id,omitempty"`
	ServiceID      int64  `json:"service_id,omitempty"`
	Status         string `json:"status,omitempty"`
	AutoStatusDowt string `json:"auto_status_down,omitempty"`
	AutoStatusUp   string `json:"auto_status_up,omitempty"`
}

func (s StatusPageIncident) PrimaryKey() PrimaryKey {
	return PrimaryKey(s.PK)
}

type StatusPageIncidentListResponse struct {
	Count   int64                `json:"count,omitempty"`
	Results []StatusPageIncident `json:"results,omitempty"`
}

func (r StatusPageIncidentListResponse) List() []StatusPageIncident {
	return r.Results
}

type StatusPageIncidentResponse StatusPageIncident

func (r StatusPageIncidentResponse) Item() StatusPageIncident {
	return StatusPageIncident(r)
}

type StatusPageIncidentCreateUpdateResponse struct {
	Results StatusPageIncident `json:"results,omitempty"`
}

func (r StatusPageIncidentCreateUpdateResponse) Item() StatusPageIncident {
	return r.Results
}

type StatusPageIncidentListOptions struct {
	Page     int64  `url:"page,omitempty"`
	PageSize int64  `url:"page_size,omitempty"`
	Search   string `url:"search,omitempty"`
	Ordering string `url:"ordering,omitempty"`
}

type StatusPageIncidentEndpoint interface {
	Create(context.Context, StatusPageIncident) (*StatusPageIncident, error)
	Update(context.Context, PrimaryKeyable, StatusPageIncident) (*StatusPageIncident, error)
	List(context.Context, StatusPageIncidentListOptions) ([]StatusPageIncident, error)
	Get(context.Context, PrimaryKeyable) (*StatusPageIncident, error)
	Delete(context.Context, PrimaryKeyable) error
}

type statusPageIncidentEndpointImpl struct {
	EndpointLister[StatusPageIncidentListResponse, StatusPageIncident, StatusPageIncidentListOptions]
	EndpointCreator[StatusPageIncident, StatusPageIncidentCreateUpdateResponse, StatusPageIncident]
	EndpointUpdater[StatusPageIncident, StatusPageIncidentCreateUpdateResponse, StatusPageIncident]
	EndpointGetter[StatusPageIncidentResponse, StatusPageIncident]
	EndpointDeleter
}
