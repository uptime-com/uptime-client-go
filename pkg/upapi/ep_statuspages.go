package upapi

import "context"

type StatusPage struct {
	PK                        int64  `json:"pk,omitempty"`
	URL                       string `json:"url,omitempty"`
	Name                      string `json:"name"`
	VisibilityLevel           string `json:"visibility_level"`
	Description               string `json:"description"`
	PageType                  string `json:"page_type"`
	Slug                      string `json:"slug"`
	CNAME                     string `json:"cname"`
	AllowSubscriptions        bool   `json:"allow_subscriptions"`
	AllowSearchIndexing       bool   `json:"allow_search_indexing"`
	AllowDrillDown            bool   `json:"allow_drill_down"`
	AuthUsername              string `json:"auth_username"`
	AuthPassword              string `json:"auth_password"`
	ShowStatusTab             bool   `json:"show_status_tab"`
	ShowActiveIncidents       bool   `json:"show_active_incidents"`
	ShowComponentResponseTime bool   `json:"show_component_response_time"`
	ShowHistoryTab            bool   `json:"show_history_tab"`
	DefaultHistoryDateRange   int64  `json:"default_history_date_range"`
	UptimeCalculationType     string `json:"uptime_calculation_type"`
	ShowHistorySnake          bool   `json:"show_history_snake"`
	ShowComponentHistory      bool   `json:"show_component_history"`
	ShowSummaryMetrics        bool   `json:"show_summary_metrics"`
	ShowPastIncidents         bool   `json:"show_past_incidents"`
	AllowPdfReport            bool   `json:"allow_pdf_report"`
	GoogleAnalyticsCode       string `json:"google_analytics_code"`
	ContactEmail              string `json:"contact_email"`
	EmailFrom                 string `json:"email_from"`
	EmailReplyTo              string `json:"email_reply_to"`
	CustomHeaderHtml          string `json:"custom_header_html"`
	CustomFooterHtml          string `json:"custom_footer_html"`
	CustomCss                 string `json:"custom_css"`
	CompanyWebsiteUrl         string `json:"company_website_url"`
	Timezone                  string `json:"timezone"`
}

func (s StatusPage) PrimaryKey() PrimaryKey {
	return PrimaryKey(s.PK)
}

type StatusPageListOptions struct {
	Page            int64  `url:"page,omitempty"`
	PageSize        int64  `url:"page_size,omitempty"`
	Search          string `url:"search,omitempty"`
	Ordering        string `url:"ordering,omitempty"`
	VisibilityLevel string `url:"visibility_level,omitempty"`
}

type StatusPageListResponse struct {
	Count   int64        `json:"count,omitempty"`
	Results []StatusPage `json:"results,omitempty"`
}

func (r StatusPageListResponse) List() []StatusPage {
	return r.Results
}

type StatusPageResponse StatusPage

func (r StatusPageResponse) Item() StatusPage {
	return StatusPage(r)
}

type StatusPageCreateUpdateResponse struct {
	Results StatusPage `json:"results,omitempty"`
}

func (r StatusPageCreateUpdateResponse) Item() StatusPage {
	return r.Results
}

type StatusPagesEndpoint interface {
	List(context.Context, StatusPageListOptions) ([]StatusPage, error)
	Create(context.Context, StatusPage) (*StatusPage, error)
	Update(context.Context, PrimaryKeyable, StatusPage) (*StatusPage, error)
	Get(context.Context, PrimaryKeyable) (*StatusPage, error)
	Delete(context.Context, PrimaryKeyable) error
}

func NewStatusPagesEndpoint(cbd CBD) StatusPagesEndpoint {
	const endpoint = "statuspages"
	return &statusPagesEndpointImpl{
		EndpointLister:  NewEndpointLister[StatusPageListResponse, StatusPage, StatusPageListOptions](cbd, endpoint),
		EndpointCreator: NewEndpointCreator[StatusPage, StatusPageCreateUpdateResponse, StatusPage](cbd, endpoint),
		EndpointUpdater: NewEndpointUpdater[StatusPage, StatusPageCreateUpdateResponse, StatusPage](cbd, endpoint),
		EndpointGetter:  NewEndpointGetter[StatusPageResponse, StatusPage](cbd, endpoint),
		EndpointDeleter: NewEndpointDeleter(cbd, endpoint),
	}
}

type statusPagesEndpointImpl struct {
	EndpointLister[StatusPageListResponse, StatusPage, StatusPageListOptions]
	EndpointCreator[StatusPage, StatusPageCreateUpdateResponse, StatusPage]
	EndpointUpdater[StatusPage, StatusPageCreateUpdateResponse, StatusPage]
	EndpointGetter[StatusPageResponse, StatusPage]
	EndpointDeleter
}
