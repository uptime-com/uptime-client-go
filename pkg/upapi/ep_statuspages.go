package upapi

import (
	"context"
	"fmt"
)

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
	MaxVisibleComponentDays   int64  `json:"max_visible_component_days,omitempty"`
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

	AllowSubscriptionsEmail   bool   `json:"allow_subscriptions_email"`
	AllowSubscriptionsRss     bool   `json:"allow_subscriptions_rss"`
	AllowSubscriptionsSlack   bool   `json:"allow_subscriptions_slack"`
	AllowSubscriptionsSms     bool   `json:"allow_subscriptions_sms"`
	AllowSubscriptionsWebhook bool   `json:"allow_subscriptions_webhook"`
	HideEmptyTabsHistory      bool   `json:"hide_empty_tabs_history"`
	Theme                     string `json:"theme"`
	CustomHeaderBgColorHex    string `json:"custom_header_bg_color_hex"`
	CustomHeaderTextColorHex  string `json:"custom_header_text_color_hex"`
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
	Components(PrimaryKeyable) StatusPageComponentEndpoint
	Incidents(PrimaryKeyable) StatusPageIncidentEndpoint
	Metrics(PrimaryKeyable) StatusPageMetricEndpoint
	Subscribers(PrimaryKeyable) StatusPageSubscriberEndpoint
	SubscriptionDomainAllowList(PrimaryKeyable) StatusPageSubsDomainAllowListEndpoint
	SubscriptionDomainBlockList(PrimaryKeyable) StatusPageSubsDomainBlockListEndpoint
	Users(PrimaryKeyable) StatusPageUserEndpoint
	CurrentStatus(PrimaryKeyable) StatusPageCurrentStatusEndpoint
	StatusHistory(PrimaryKeyable) StatusPageStatusHistoryEndpoint
}

func NewStatusPagesEndpoint(cbd CBD) StatusPagesEndpoint {
	const endpoint = "statuspages"
	return &statusPagesEndpointImpl{
		cbd:             cbd,
		endpoint:        endpoint,
		EndpointLister:  NewEndpointLister[StatusPageListResponse, StatusPage, StatusPageListOptions](cbd, endpoint),
		EndpointCreator: NewEndpointCreator[StatusPage, StatusPageCreateUpdateResponse, StatusPage](cbd, endpoint),
		EndpointUpdater: NewEndpointUpdater[StatusPage, StatusPageCreateUpdateResponse, StatusPage](cbd, endpoint),
		EndpointGetter:  NewEndpointGetter[StatusPageResponse, StatusPage](cbd, endpoint),
		EndpointDeleter: NewEndpointDeleter(cbd, endpoint),
	}
}

type statusPagesEndpointImpl struct {
	cbd      CBD
	endpoint string
	EndpointLister[StatusPageListResponse, StatusPage, StatusPageListOptions]
	EndpointCreator[StatusPage, StatusPageCreateUpdateResponse, StatusPage]
	EndpointUpdater[StatusPage, StatusPageCreateUpdateResponse, StatusPage]
	EndpointGetter[StatusPageResponse, StatusPage]
	EndpointDeleter
}

func (c *statusPagesEndpointImpl) Components(pk PrimaryKeyable) StatusPageComponentEndpoint {
	endpoint := fmt.Sprintf("%s/%d/components", c.endpoint, pk.PrimaryKey())
	return &statusPageComponentEndpointImpl{
		EndpointLister:  NewEndpointLister[StatusPageComponentListResponse, StatusPageComponent, StatusPageComponentListOptions](c.cbd, endpoint),
		EndpointCreator: NewEndpointCreator[StatusPageComponent, StatusPageComponentCreateUpdateResponse](c.cbd, endpoint),
		EndpointUpdater: NewEndpointUpdater[StatusPageComponent, StatusPageComponentCreateUpdateResponse](c.cbd, endpoint),
		EndpointGetter:  NewEndpointGetter[StatusPageComponentResponse](c.cbd, endpoint),
		EndpointDeleter: NewEndpointDeleter(c.cbd, endpoint),
	}
}

func (c *statusPagesEndpointImpl) Incidents(pk PrimaryKeyable) StatusPageIncidentEndpoint {
	endpoint := fmt.Sprintf("%s/%d/incidents", c.endpoint, pk.PrimaryKey())
	return &statusPageIncidentEndpointImpl{
		EndpointLister:  NewEndpointLister[StatusPageIncidentListResponse, StatusPageIncident, StatusPageIncidentListOptions](c.cbd, endpoint),
		EndpointCreator: NewEndpointCreator[StatusPageIncident, StatusPageIncidentCreateUpdateResponse](c.cbd, endpoint),
		EndpointUpdater: NewEndpointUpdater[StatusPageIncident, StatusPageIncidentCreateUpdateResponse](c.cbd, endpoint),
		EndpointGetter:  NewEndpointGetter[StatusPageIncidentResponse](c.cbd, endpoint),
		EndpointDeleter: NewEndpointDeleter(c.cbd, endpoint),
	}
}

func (c *statusPagesEndpointImpl) Metrics(pk PrimaryKeyable) StatusPageMetricEndpoint {
	endpoint := fmt.Sprintf("%s/%d/metrics", c.endpoint, pk.PrimaryKey())
	return &statusPageMetricEndpointImpl{
		EndpointLister:  NewEndpointLister[StatusPageMetricListResponse, StatusPageMetric, StatusPageMetricListOptions](c.cbd, endpoint),
		EndpointCreator: NewEndpointCreator[StatusPageMetric, StatusPageMetricCreateUpdateResponse](c.cbd, endpoint),
		EndpointUpdater: NewEndpointUpdater[StatusPageMetric, StatusPageMetricCreateUpdateResponse](c.cbd, endpoint),
		EndpointGetter:  NewEndpointGetter[StatusPageMetricResponse](c.cbd, endpoint),
		EndpointDeleter: NewEndpointDeleter(c.cbd, endpoint),
	}
}

func (c *statusPagesEndpointImpl) Subscribers(pk PrimaryKeyable) StatusPageSubscriberEndpoint {
	endpoint := fmt.Sprintf("%s/%d/subscribers", c.endpoint, pk.PrimaryKey())
	return &statusPageSubscriberEndpointImpl{
		EndpointLister:  NewEndpointLister[StatusPageSubscriberListResponse, StatusPageSubscriber, StatusPageSubscriberListOptions](c.cbd, endpoint),
		EndpointCreator: NewEndpointCreator[StatusPageSubscriber, StatusPageSubscriberCreateResponse](c.cbd, endpoint),
		EndpointGetter:  NewEndpointGetter[StatusPageSubscriberResponse](c.cbd, endpoint),
		EndpointDeleter: NewEndpointDeleter(c.cbd, endpoint),
	}
}

func (c *statusPagesEndpointImpl) SubscriptionDomainAllowList(pk PrimaryKeyable) StatusPageSubsDomainAllowListEndpoint {
	endpoint := fmt.Sprintf("%s/%d/subscription-domain-allow-list", c.endpoint, pk.PrimaryKey())
	return &statusPageSubsDomainAllowListEndpointImpl{
		EndpointLister:  NewEndpointLister[StatusPageSubsDomainAllowListListResponse, StatusPageSubsDomainAllowList, StatusPageSubsDomainAllowListListOptions](c.cbd, endpoint),
		EndpointCreator: NewEndpointCreator[StatusPageSubsDomainAllowList, StatusPageSubsDomainAllowListCreateUpdateResponse](c.cbd, endpoint),
		EndpointUpdater: NewEndpointUpdater[StatusPageSubsDomainAllowList, StatusPageSubsDomainAllowListCreateUpdateResponse](c.cbd, endpoint),
		EndpointGetter:  NewEndpointGetter[StatusPageSubsDomainAllowListResponse](c.cbd, endpoint),
		EndpointDeleter: NewEndpointDeleter(c.cbd, endpoint),
	}
}

func (c *statusPagesEndpointImpl) SubscriptionDomainBlockList(pk PrimaryKeyable) StatusPageSubsDomainBlockListEndpoint {
	endpoint := fmt.Sprintf("%s/%d/subscription-domain-block-list", c.endpoint, pk.PrimaryKey())
	return &statusPageSubsDomainBlockListEndpointImpl{
		EndpointLister:  NewEndpointLister[StatusPageSubsDomainBlockListListResponse, StatusPageSubsDomainBlockList, StatusPageSubsDomainBlockListListOptions](c.cbd, endpoint),
		EndpointCreator: NewEndpointCreator[StatusPageSubsDomainBlockList, StatusPageSubsDomainBlockListCreateUpdateResponse](c.cbd, endpoint),
		EndpointUpdater: NewEndpointUpdater[StatusPageSubsDomainBlockList, StatusPageSubsDomainBlockListCreateUpdateResponse](c.cbd, endpoint),
		EndpointGetter:  NewEndpointGetter[StatusPageSubsDomainBlockListResponse](c.cbd, endpoint),
		EndpointDeleter: NewEndpointDeleter(c.cbd, endpoint),
	}
}

func (c *statusPagesEndpointImpl) Users(pk PrimaryKeyable) StatusPageUserEndpoint {
	endpoint := fmt.Sprintf("%s/%d/users", c.endpoint, pk.PrimaryKey())
	return &statusPageUserEndpointImpl{
		EndpointLister:  NewEndpointLister[StatusPageUserListResponse, StatusPageUser, StatusPageUserListOptions](c.cbd, endpoint),
		EndpointCreator: NewEndpointCreator[StatusPageUser, StatusPageUserCreateUpdateResponse](c.cbd, endpoint),
		EndpointUpdater: NewEndpointUpdater[StatusPageUser, StatusPageUserCreateUpdateResponse](c.cbd, endpoint),
		EndpointGetter:  NewEndpointGetter[StatusPageUserResponse](c.cbd, endpoint),
		EndpointDeleter: NewEndpointDeleter(c.cbd, endpoint),
	}
}

func (c *statusPagesEndpointImpl) CurrentStatus(pk PrimaryKeyable) StatusPageCurrentStatusEndpoint {
	return NewStatusPageCurrentStatusEndpoint(c.cbd, pk)
}

func (c *statusPagesEndpointImpl) StatusHistory(pk PrimaryKeyable) StatusPageStatusHistoryEndpoint {
	return NewStatusPageStatusHistoryEndpoint(c.cbd, pk)
}
