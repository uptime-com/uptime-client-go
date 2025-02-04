package upapi

import "context"

type Dashboard struct {
	PK                         int64    `json:"pk,omitempty"`
	ServicesSelected           []string `json:"services_selected"`
	ServicesTags               []string `json:"services_tags"`
	Ordering                   int64    `json:"ordering"`
	Name                       string   `json:"name"`
	IsPinned                   bool     `json:"is_pinned"`
	MetricsShowSection         bool     `json:"metrics_show_section"`
	MetricsForAllChecks        bool     `json:"metrics_for_all_checks"`
	ServicesShowSection        bool     `json:"services_show_section"`
	ServicesNumToShow          int64    `json:"services_num_to_show"`
	ServicesIncludeUp          bool     `json:"services_include_up"`
	ServicesIncludeDown        bool     `json:"services_include_down"`
	ServicesIncludePaused      bool     `json:"services_include_paused"`
	ServicesIncludeMaintenance bool     `json:"services_include_maintenance"`
	ServicesPrimarySort        string   `json:"services_primary_sort"`
	ServicesSecondarySort      string   `json:"services_secondary_sort"`
	ServicesShowUptime         bool     `json:"services_show_uptime"`
	ServicesShowResponseTime   bool     `json:"services_show_response_time"`
	AlertsShowSection          bool     `json:"alerts_show_section"`
	AlertsForAllChecks         bool     `json:"alerts_for_all_checks"`
	AlertsIncludeIgnored       bool     `json:"alerts_include_ignored"`
	AlertsincludeResolved      bool     `json:"alerts_include_resolved"`
	AlertsnumToShow            int64    `json:"alerts_num_to_show"`
}

func (d Dashboard) PrimaryKey() PrimaryKey {
	return PrimaryKey(d.PK)
}

type DashboardListOptions struct {
	Page     int64  `url:"page,omitempty"`
	PageSize int64  `url:"page_size,omitempty"`
	Search   string `url:"search,omitempty"`
	Ordering string `url:"ordering,omitempty"`
}

type DashboardListResponse struct {
	Count   int64       `json:"count,omitempty"`
	Results []Dashboard `json:"results,omitempty"`
}

func (r DashboardListResponse) List() []Dashboard {
	return r.Results
}

type DashboardResponse Dashboard

func (d DashboardResponse) Item() Dashboard {
	return Dashboard(d)
}

type DashboardCreateUpdateResponse struct {
	Results Dashboard `json:"results,omitempty"`
}

func (r DashboardCreateUpdateResponse) Item() Dashboard {
	return r.Results
}

type DashboardsEndpoint interface {
	List(context.Context, DashboardListOptions) ([]Dashboard, error)
	Create(context.Context, Dashboard) (*Dashboard, error)
	Update(context.Context, PrimaryKeyable, Dashboard) (*Dashboard, error)
	Get(context.Context, PrimaryKeyable) (*Dashboard, error)
	Delete(context.Context, PrimaryKeyable) error
}

func NewDashboardsEndpoint(cbd CBD) DashboardsEndpoint {
	const endpoint = "dashboards"
	return &dashboardsEndpointImpl{
		EndpointLister:  NewEndpointLister[DashboardListResponse, Dashboard, DashboardListOptions](cbd, endpoint),
		EndpointCreator: NewEndpointCreator[Dashboard, DashboardCreateUpdateResponse, Dashboard](cbd, endpoint),
		EndpointUpdater: NewEndpointUpdater[Dashboard, DashboardCreateUpdateResponse, Dashboard](cbd, endpoint),
		EndpointGetter:  NewEndpointGetter[DashboardResponse, Dashboard](cbd, endpoint),
		EndpointDeleter: NewEndpointDeleter(cbd, endpoint),
	}
}

type dashboardsEndpointImpl struct {
	EndpointLister[DashboardListResponse, Dashboard, DashboardListOptions]
	EndpointCreator[Dashboard, DashboardCreateUpdateResponse, Dashboard]
	EndpointUpdater[Dashboard, DashboardCreateUpdateResponse, Dashboard]
	EndpointGetter[DashboardResponse, Dashboard]
	EndpointDeleter
}
