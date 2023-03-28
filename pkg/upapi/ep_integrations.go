package upapi

import (
	"context"
)

// Integration represents an integration in Uptime.com.
type Integration struct {
	PK            int      `json:"pk,omitempty"`
	URL           string   `json:"url,omitempty"`
	Name          string   `json:"name,omitempty"`
	Module        string   `json:"module,omitempty"`
	ContactGroups []string `json:"contact_groups,omitempty"`
	IsErrored     bool     `json:"is_errored,omitempty"`
	LastError     string   `json:"last_error,omitempty"`
}

func (i Integration) PrimaryKey() PrimaryKey {
	return PrimaryKey(i.PK)
}

type IntegrationListResponse struct {
	Count    int           `json:"count,omitempty"`
	Next     string        `json:"next,omitempty"`
	Previous string        `json:"previous,omitempty"`
	Results  []Integration `json:"results,omitempty"`
}

func (r IntegrationListResponse) List() []Integration {
	return r.Results
}

type IntegrationListOptions struct {
	Page     int    `url:"page,omitempty"`
	PageSize int    `url:"page_size,omitempty"`
	Search   string `url:"search,omitempty"`
	Ordering string `url:"ordering,omitempty"`
	Module   string `url:"module,omitempty"`
}

type IntegrationResponse struct {
	Results Integration `json:"results,omitempty"`
}

func (r IntegrationResponse) Item() Integration {
	return r.Results
}

type IntegrationsEndpoint interface {
	List(context.Context, IntegrationListOptions) ([]Integration, error)
	Get(context.Context, PrimaryKeyable) (*Integration, error)
	Delete(context.Context, PrimaryKeyable) error
	CreateCachet(context.Context, IntegrationCachet) (*Integration, error)
	UpdateCachet(context.Context, PrimaryKeyable, IntegrationCachet) (*Integration, error)
	CreateDatadog(context.Context, IntegrationDatadog) (*Integration, error)
	UpdateDatadog(context.Context, PrimaryKeyable, IntegrationDatadog) (*Integration, error)
	CreateGeckoboard(context.Context, IntegrationGeckoboard) (*Integration, error)
	UpdateGeckoboard(context.Context, PrimaryKeyable, IntegrationGeckoboard) (*Integration, error)
	CreateJiraServicedesk(context.Context, IntegrationJiraServicedesk) (*Integration, error)
	UpdateJiraServiceDesk(context.Context, PrimaryKeyable, IntegrationJiraServicedesk) (*Integration, error)
	CreateKlipfolio(context.Context, IntegrationKlipfolio) (*Integration, error)
	UpdateKlipfolio(context.Context, PrimaryKeyable, IntegrationKlipfolio) (*Integration, error)
	CreateLibrato(context.Context, IntegrationLibrato) (*Integration, error)
	UpdateLibrato(context.Context, PrimaryKeyable, IntegrationLibrato) (*Integration, error)
	CreateMicrosoftTeams(context.Context, IntegrationMicrosoftTeams) (*Integration, error)
	UpdateMicrosoftTeams(context.Context, PrimaryKeyable, IntegrationMicrosoftTeams) (*Integration, error)
	CreateOpsgenie(context.Context, IntegrationOpsgenie) (*Integration, error)
	UpdateOpsgenie(context.Context, PrimaryKeyable, IntegrationOpsgenie) (*Integration, error)
	CreatePagerduty(context.Context, IntegrationPagerduty) (*Integration, error)
	UpdatePagerduty(context.Context, PrimaryKeyable, IntegrationPagerduty) (*Integration, error)
	CreatePushbullet(context.Context, IntegrationPushbullet) (*Integration, error)
	UpdatePushbullet(context.Context, PrimaryKeyable, IntegrationPushbullet) (*Integration, error)
	CreatePushover(context.Context, IntegrationPushover) (*Integration, error)
	UpdatePushover(context.Context, PrimaryKeyable, IntegrationPushover) (*Integration, error)
	CreateSlack(context.Context, IntegrationSlack) (*Integration, error)
	UpdateSlack(context.Context, PrimaryKeyable, IntegrationSlack) (*Integration, error)
	CreateStatus(context.Context, IntegrationStatus) (*Integration, error)
	UpdateStatus(context.Context, PrimaryKeyable, IntegrationStatus) (*Integration, error)
	CreateStatuspage(context.Context, IntegrationStatuspage) (*Integration, error)
	UpdateStatuspage(context.Context, PrimaryKeyable, IntegrationStatuspage) (*Integration, error)
	CreateTwitter(context.Context, IntegrationTwitter) (*Integration, error)
	UpdateTwitter(context.Context, PrimaryKeyable, IntegrationTwitter) (*Integration, error)
	CreateVictorops(context.Context, IntegrationVictorops) (*Integration, error)
	UpdateVictorops(context.Context, PrimaryKeyable, IntegrationVictorops) (*Integration, error)
	CreateWavefront(context.Context, IntegrationWavefront) (*Integration, error)
	UpdateWavefront(context.Context, PrimaryKeyable, IntegrationWavefront) (*Integration, error)
	CreateWebhook(context.Context, IntegrationWebhook) (*Integration, error)
	UpdateWebhook(context.Context, PrimaryKeyable, IntegrationWebhook) (*Integration, error)
	CreateZapier(context.Context, IntegrationZapier) (*Integration, error)
	UpdateZapier(context.Context, PrimaryKeyable, IntegrationZapier) (*Integration, error)
}

func NewIntegrationsEndpoint(cbd CBD) IntegrationsEndpoint {
	const endpoint = "integrations"
	return &integrationsEndpointImpl{
		integrationsEndpointCachetImpl: integrationsEndpointCachetImpl{
			EndpointCreator: NewEndpointCreator[IntegrationCachet, IntegrationResponse, Integration](cbd, endpoint+"/add-cachet"),
			EndpointUpdater: NewEndpointUpdater[IntegrationCachet, IntegrationResponse, Integration](cbd, endpoint),
		},
		integrationsEndpointDatadogImpl: integrationsEndpointDatadogImpl{
			EndpointCreator: NewEndpointCreator[IntegrationDatadog, IntegrationResponse, Integration](cbd, endpoint+"/add-datadog"),
			EndpointUpdater: NewEndpointUpdater[IntegrationDatadog, IntegrationResponse, Integration](cbd, endpoint),
		},
		integrationsEndpointGeckoboardImpl: integrationsEndpointGeckoboardImpl{
			EndpointCreator: NewEndpointCreator[IntegrationGeckoboard, IntegrationResponse, Integration](cbd, endpoint+"/add-geckoboard"),
			EndpointUpdater: NewEndpointUpdater[IntegrationGeckoboard, IntegrationResponse, Integration](cbd, endpoint),
		},
		integrationsEndpointJiraServicedeskImpl: integrationsEndpointJiraServicedeskImpl{
			EndpointCreator: NewEndpointCreator[IntegrationJiraServicedesk, IntegrationResponse, Integration](cbd, endpoint+"/add-jiraservicedesk"),
			EndpointUpdater: NewEndpointUpdater[IntegrationJiraServicedesk, IntegrationResponse, Integration](cbd, endpoint),
		},
		integrationsEndpointKlipfolioImpl: integrationsEndpointKlipfolioImpl{
			EndpointCreator: NewEndpointCreator[IntegrationKlipfolio, IntegrationResponse, Integration](cbd, endpoint+"/add-klipfolio"),
			EndpointUpdater: NewEndpointUpdater[IntegrationKlipfolio, IntegrationResponse, Integration](cbd, endpoint),
		},
		integrationsEndpointLibratoImpl: integrationsEndpointLibratoImpl{
			EndpointCreator: NewEndpointCreator[IntegrationLibrato, IntegrationResponse, Integration](cbd, endpoint+"/add-librato"),
			EndpointUpdater: NewEndpointUpdater[IntegrationLibrato, IntegrationResponse, Integration](cbd, endpoint),
		},
		integrationsEndpointMicrosoftTeamsImpl: integrationsEndpointMicrosoftTeamsImpl{
			EndpointCreator: NewEndpointCreator[IntegrationMicrosoftTeams, IntegrationResponse, Integration](cbd, endpoint+"/add-microsoft-teams"),
			EndpointUpdater: NewEndpointUpdater[IntegrationMicrosoftTeams, IntegrationResponse, Integration](cbd, endpoint),
		},
		integrationsEndpointOpsgenieImpl: integrationsEndpointOpsgenieImpl{
			EndpointCreator: NewEndpointCreator[IntegrationOpsgenie, IntegrationResponse, Integration](cbd, endpoint+"/add-opsgenie"),
			EndpointUpdater: NewEndpointUpdater[IntegrationOpsgenie, IntegrationResponse, Integration](cbd, endpoint),
		},
		integrationsEndpointPagerdutyImpl: integrationsEndpointPagerdutyImpl{
			EndpointCreator: NewEndpointCreator[IntegrationPagerduty, IntegrationResponse, Integration](cbd, endpoint+"/add-pagerduty"),
			EndpointUpdater: NewEndpointUpdater[IntegrationPagerduty, IntegrationResponse, Integration](cbd, endpoint),
		},
		integrationsEndpointPushbulletImpl: integrationsEndpointPushbulletImpl{
			EndpointCreator: NewEndpointCreator[IntegrationPushbullet, IntegrationResponse, Integration](cbd, endpoint+"/add-pushbullet"),
			EndpointUpdater: NewEndpointUpdater[IntegrationPushbullet, IntegrationResponse, Integration](cbd, endpoint),
		},
		integrationsEndpointPushoverImpl: integrationsEndpointPushoverImpl{
			EndpointCreator: NewEndpointCreator[IntegrationPushover, IntegrationResponse, Integration](cbd, endpoint+"/add-pushover"),
			EndpointUpdater: NewEndpointUpdater[IntegrationPushover, IntegrationResponse, Integration](cbd, endpoint),
		},
		integrationsEndpointSlackImpl: integrationsEndpointSlackImpl{
			EndpointCreator: NewEndpointCreator[IntegrationSlack, IntegrationResponse, Integration](cbd, endpoint+"/add-slack"),
			EndpointUpdater: NewEndpointUpdater[IntegrationSlack, IntegrationResponse, Integration](cbd, endpoint),
		},
		integrationsEndpointStatusImpl: integrationsEndpointStatusImpl{
			EndpointCreator: NewEndpointCreator[IntegrationStatus, IntegrationResponse, Integration](cbd, endpoint+"/add-status"),
			EndpointUpdater: NewEndpointUpdater[IntegrationStatus, IntegrationResponse, Integration](cbd, endpoint),
		},
		integrationsEndpointStatuspageImpl: integrationsEndpointStatuspageImpl{
			EndpointCreator: NewEndpointCreator[IntegrationStatuspage, IntegrationResponse, Integration](cbd, endpoint+"/add-statuspage"),
			EndpointUpdater: NewEndpointUpdater[IntegrationStatuspage, IntegrationResponse, Integration](cbd, endpoint),
		},
		integrationsEndpointTwitterImpl: integrationsEndpointTwitterImpl{
			EndpointCreator: NewEndpointCreator[IntegrationTwitter, IntegrationResponse, Integration](cbd, endpoint+"/add-twitter"),
			EndpointUpdater: NewEndpointUpdater[IntegrationTwitter, IntegrationResponse, Integration](cbd, endpoint),
		},
		integrationsEndpointVictoropsImpl: integrationsEndpointVictoropsImpl{
			EndpointCreator: NewEndpointCreator[IntegrationVictorops, IntegrationResponse, Integration](cbd, endpoint+"/add-victorops"),
			EndpointUpdater: NewEndpointUpdater[IntegrationVictorops, IntegrationResponse, Integration](cbd, endpoint),
		},
		integrationsEndpointWavefrontImpl: integrationsEndpointWavefrontImpl{
			EndpointCreator: NewEndpointCreator[IntegrationWavefront, IntegrationResponse, Integration](cbd, endpoint+"/add-wavefront"),
			EndpointUpdater: NewEndpointUpdater[IntegrationWavefront, IntegrationResponse, Integration](cbd, endpoint),
		},
		integrationsEndpointWebhookImpl: integrationsEndpointWebhookImpl{
			EndpointCreator: NewEndpointCreator[IntegrationWebhook, IntegrationResponse, Integration](cbd, endpoint+"/add-webhook"),
			EndpointUpdater: NewEndpointUpdater[IntegrationWebhook, IntegrationResponse, Integration](cbd, endpoint),
		},
		integrationsEndpointZapierImpl: integrationsEndpointZapierImpl{
			EndpointCreator: NewEndpointCreator[IntegrationZapier, IntegrationResponse, Integration](cbd, endpoint+"/add-zapier"),
			EndpointUpdater: NewEndpointUpdater[IntegrationZapier, IntegrationResponse, Integration](cbd, endpoint),
		},
		EndpointLister:  NewEndpointLister[IntegrationListResponse, Integration, IntegrationListOptions](cbd, endpoint),
		EndpointGetter:  NewEndpointGetter[IntegrationResponse, Integration](cbd, endpoint),
		EndpointDeleter: NewEndpointDeleter(cbd, endpoint),
	}
}

type integrationsEndpointImpl struct {
	integrationsEndpointCachetImpl
	integrationsEndpointDatadogImpl
	integrationsEndpointGeckoboardImpl
	integrationsEndpointJiraServicedeskImpl
	integrationsEndpointKlipfolioImpl
	integrationsEndpointLibratoImpl
	integrationsEndpointMicrosoftTeamsImpl
	integrationsEndpointOpsgenieImpl
	integrationsEndpointPagerdutyImpl
	integrationsEndpointPushbulletImpl
	integrationsEndpointPushoverImpl
	integrationsEndpointSlackImpl
	integrationsEndpointStatusImpl
	integrationsEndpointStatuspageImpl
	integrationsEndpointTwitterImpl
	integrationsEndpointVictoropsImpl
	integrationsEndpointWavefrontImpl
	integrationsEndpointWebhookImpl
	integrationsEndpointZapierImpl
	EndpointLister[IntegrationListResponse, Integration, IntegrationListOptions]
	EndpointGetter[IntegrationResponse, Integration]
	EndpointDeleter
}

type IntegrationCachet struct {
	PK            int      `json:"pk,omitempty"`
	CachetURL     string   `json:"url,omitempty"`
	Name          string   `json:"name,omitempty"`
	ContactGroups []string `json:"contact_groups,omitempty"`
	Token         string   `json:"token,omitempty"`
	Component     string   `json:"component,omitempty"`
	Metric        string   `json:"metric,omitempty"`
}

func (i IntegrationCachet) PrimaryKey() PrimaryKey {
	return PrimaryKey(i.PK)
}

type integrationsEndpointCachetImpl struct {
	EndpointCreator[IntegrationCachet, IntegrationResponse, Integration]
	EndpointUpdater[IntegrationCachet, IntegrationResponse, Integration]
}

func (i integrationsEndpointCachetImpl) CreateCachet(ctx context.Context, data IntegrationCachet) (*Integration, error) {
	return i.EndpointCreator.Create(ctx, data)
}

func (i integrationsEndpointCachetImpl) UpdateCachet(ctx context.Context, pk PrimaryKeyable, data IntegrationCachet) (*Integration, error) {
	return i.EndpointUpdater.Update(ctx, pk, data)
}

type IntegrationDatadog struct {
	Name          string   `json:"name,omitempty"`
	ContactGroups []string `json:"contact_groups,omitempty"`
	APIKey        string   `json:"api_key,omitempty"`
	APPKey        string   `json:"app_key,omitempty"`
	Region        string   `json:"region,omitempty"`
}

type integrationsEndpointDatadogImpl struct {
	EndpointCreator[IntegrationDatadog, IntegrationResponse, Integration]
	EndpointUpdater[IntegrationDatadog, IntegrationResponse, Integration]
}

func (i integrationsEndpointDatadogImpl) CreateDatadog(ctx context.Context, data IntegrationDatadog) (*Integration, error) {
	return i.EndpointCreator.Create(ctx, data)
}

func (i integrationsEndpointDatadogImpl) UpdateDatadog(ctx context.Context, pk PrimaryKeyable, data IntegrationDatadog) (*Integration, error) {
	return i.EndpointUpdater.Update(ctx, pk, data)
}

type IntegrationGeckoboard struct {
	Name          string   `json:"name,omitempty"`
	ContactGroups []string `json:"contact_groups,omitempty"`
	APIKey        string   `json:"api_key,omitempty"`
	DatasetName   string   `json:"dataset_name,omitempty"`
}

type integrationsEndpointGeckoboardImpl struct {
	EndpointCreator[IntegrationGeckoboard, IntegrationResponse, Integration]
	EndpointUpdater[IntegrationGeckoboard, IntegrationResponse, Integration]
}

func (i integrationsEndpointGeckoboardImpl) CreateGeckoboard(ctx context.Context, data IntegrationGeckoboard) (*Integration, error) {
	return i.EndpointCreator.Create(ctx, data)
}

func (i integrationsEndpointGeckoboardImpl) UpdateGeckoboard(ctx context.Context, pk PrimaryKeyable, data IntegrationGeckoboard) (*Integration, error) {
	return i.EndpointUpdater.Update(ctx, pk, data)
}

type IntegrationJiraServicedesk struct {
	Name                     string   `json:"name,omitempty"`
	ContactGroups            []string `json:"contact_groups,omitempty"`
	APIEmail                 string   `json:"api_email,omitempty"`
	APIToken                 string   `json:"api_token,omitempty"`
	JiraSubdomain            string   `json:"jira_subdomain,omitempty"`
	ProjectKey               string   `json:"project_key,omitempty"`
	Labels                   string   `json:"labels,omitempty"`
	CustomFieldIdAccountName int      `json:"custom_field_id_account_name,omitempty"`
	CustomFieldIdCheckName   int      `json:"custom_field_id_check_name,omitempty"`
	CustomFieldIdCheckUrl    int      `json:"custom_field_id_check_url,omitempty"`
	CustomFieldsJson         string   `json:"custom_fields_json,omitempty"`
}

type integrationsEndpointJiraServicedeskImpl struct {
	EndpointCreator[IntegrationJiraServicedesk, IntegrationResponse, Integration]
	EndpointUpdater[IntegrationJiraServicedesk, IntegrationResponse, Integration]
}

func (i integrationsEndpointJiraServicedeskImpl) CreateJiraServicedesk(ctx context.Context, data IntegrationJiraServicedesk) (*Integration, error) {
	return i.EndpointCreator.Create(ctx, data)
}

func (i integrationsEndpointJiraServicedeskImpl) UpdateJiraServiceDesk(ctx context.Context, pk PrimaryKeyable, data IntegrationJiraServicedesk) (*Integration, error) {
	return i.EndpointUpdater.Update(ctx, pk, data)
}

type IntegrationKlipfolio struct {
	Name           string   `json:"name,omitempty"`
	ContactGroups  []string `json:"contact_groups,omitempty"`
	APIKey         string   `json:"api_key,omitempty"`
	DataSourceName string   `json:"data_source_name,omitempty"`
}

type integrationsEndpointKlipfolioImpl struct {
	EndpointCreator[IntegrationKlipfolio, IntegrationResponse, Integration]
	EndpointUpdater[IntegrationKlipfolio, IntegrationResponse, Integration]
}

func (i integrationsEndpointKlipfolioImpl) CreateKlipfolio(ctx context.Context, data IntegrationKlipfolio) (*Integration, error) {
	return i.EndpointCreator.Create(ctx, data)
}

func (i integrationsEndpointKlipfolioImpl) UpdateKlipfolio(ctx context.Context, pk PrimaryKeyable, data IntegrationKlipfolio) (*Integration, error) {
	return i.EndpointUpdater.Update(ctx, pk, data)
}

type IntegrationLibrato struct {
	Name          string   `json:"name,omitempty"`
	ContactGroups []string `json:"contact_groups,omitempty"`
	Email         string   `json:"email,omitempty"`
	APIToken      string   `json:"api_token,omitempty"`
	MetricName    string   `json:"metric_name,omitempty"`
}

type integrationsEndpointLibratoImpl struct {
	EndpointCreator[IntegrationLibrato, IntegrationResponse, Integration]
	EndpointUpdater[IntegrationLibrato, IntegrationResponse, Integration]
}

func (i integrationsEndpointLibratoImpl) CreateLibrato(ctx context.Context, data IntegrationLibrato) (*Integration, error) {
	return i.EndpointCreator.Create(ctx, data)
}

func (i integrationsEndpointLibratoImpl) UpdateLibrato(ctx context.Context, pk PrimaryKeyable, data IntegrationLibrato) (*Integration, error) {
	return i.EndpointUpdater.Update(ctx, pk, data)
}

type IntegrationMicrosoftTeams struct {
	Name          string   `json:"name,omitempty"`
	ContactGroups []string `json:"contact_groups,omitempty"`
	WebhookUrl    string   `json:"webhook_url,omitempty"`
}

type integrationsEndpointMicrosoftTeamsImpl struct {
	EndpointCreator[IntegrationMicrosoftTeams, IntegrationResponse, Integration]
	EndpointUpdater[IntegrationMicrosoftTeams, IntegrationResponse, Integration]
}

func (i integrationsEndpointMicrosoftTeamsImpl) CreateMicrosoftTeams(ctx context.Context, data IntegrationMicrosoftTeams) (*Integration, error) {
	return i.EndpointCreator.Create(ctx, data)
}

func (i integrationsEndpointMicrosoftTeamsImpl) UpdateMicrosoftTeams(ctx context.Context, pk PrimaryKeyable, data IntegrationMicrosoftTeams) (*Integration, error) {
	return i.EndpointUpdater.Update(ctx, pk, data)
}

type IntegrationOpsgenie struct {
	Name          string   `json:"name,omitempty"`
	ContactGroups []string `json:"contact_groups,omitempty"`
	APIEndpoint   string   `json:"api_endpoint,omitempty"`
	APIKey        string   `json:"api_key,omitempty"`
	Teams         string   `json:"teams,omitempty"`
	Tags          string   `json:"tags,omitempty"`
	Autoresolve   bool     `json:"autoresolve"`
}

type integrationsEndpointOpsgenieImpl struct {
	EndpointCreator[IntegrationOpsgenie, IntegrationResponse, Integration]
	EndpointUpdater[IntegrationOpsgenie, IntegrationResponse, Integration]
}

func (i integrationsEndpointOpsgenieImpl) CreateOpsgenie(ctx context.Context, data IntegrationOpsgenie) (*Integration, error) {
	return i.EndpointCreator.Create(ctx, data)
}

func (i integrationsEndpointOpsgenieImpl) UpdateOpsgenie(ctx context.Context, pk PrimaryKeyable, data IntegrationOpsgenie) (*Integration, error) {
	return i.EndpointUpdater.Update(ctx, pk, data)
}

type IntegrationPagerduty struct {
	Name          string   `json:"name,omitempty"`
	ContactGroups []string `json:"contact_groups,omitempty"`
	ServiceKey    string   `json:"service_key,omitempty"`
	Autoresolve   bool     `json:"autoresolve"`
}

type integrationsEndpointPagerdutyImpl struct {
	EndpointCreator[IntegrationPagerduty, IntegrationResponse, Integration]
	EndpointUpdater[IntegrationPagerduty, IntegrationResponse, Integration]
}

func (i integrationsEndpointPagerdutyImpl) CreatePagerduty(ctx context.Context, data IntegrationPagerduty) (*Integration, error) {
	return i.EndpointCreator.Create(ctx, data)
}

func (i integrationsEndpointPagerdutyImpl) UpdatePagerduty(ctx context.Context, pk PrimaryKeyable, data IntegrationPagerduty) (*Integration, error) {
	return i.EndpointUpdater.Update(ctx, pk, data)
}

type IntegrationPushbullet struct {
	Name          string   `json:"name,omitempty"`
	ContactGroups []string `json:"contact_groups,omitempty"`
	Email         string   `json:"email,omitempty"`
}

type integrationsEndpointPushbulletImpl struct {
	EndpointCreator[IntegrationPushbullet, IntegrationResponse, Integration]
	EndpointUpdater[IntegrationPushbullet, IntegrationResponse, Integration]
}

func (i integrationsEndpointPushbulletImpl) CreatePushbullet(ctx context.Context, data IntegrationPushbullet) (*Integration, error) {
	return i.EndpointCreator.Create(ctx, data)
}

func (i integrationsEndpointPushbulletImpl) UpdatePushbullet(ctx context.Context, pk PrimaryKeyable, data IntegrationPushbullet) (*Integration, error) {
	return i.EndpointUpdater.Update(ctx, pk, data)
}

type IntegrationPushover struct {
	Name          string   `json:"name,omitempty"`
	ContactGroups []string `json:"contact_groups,omitempty"`
	User          string   `json:"user,omitempty"`
	Priority      int      `json:"priority,omitempty"`
}

type integrationsEndpointPushoverImpl struct {
	EndpointCreator[IntegrationPushover, IntegrationResponse, Integration]
	EndpointUpdater[IntegrationPushover, IntegrationResponse, Integration]
}

func (i integrationsEndpointPushoverImpl) CreatePushover(ctx context.Context, data IntegrationPushover) (*Integration, error) {
	return i.EndpointCreator.Create(ctx, data)
}

func (i integrationsEndpointPushoverImpl) UpdatePushover(ctx context.Context, pk PrimaryKeyable, data IntegrationPushover) (*Integration, error) {
	return i.EndpointUpdater.Update(ctx, pk, data)
}

type IntegrationSlack struct {
	Name          string   `json:"name,omitempty"`
	ContactGroups []string `json:"contact_groups,omitempty"`
	WebhookURL    string   `json:"webhook_url,omitempty"`
	Channel       string   `json:"channel,omitempty"`
}

type integrationsEndpointSlackImpl struct {
	EndpointCreator[IntegrationSlack, IntegrationResponse, Integration]
	EndpointUpdater[IntegrationSlack, IntegrationResponse, Integration]
}

func (i integrationsEndpointSlackImpl) CreateSlack(ctx context.Context, data IntegrationSlack) (*Integration, error) {
	return i.EndpointCreator.Create(ctx, data)
}

func (i integrationsEndpointSlackImpl) UpdateSlack(ctx context.Context, pk PrimaryKeyable, data IntegrationSlack) (*Integration, error) {
	return i.EndpointUpdater.Update(ctx, pk, data)
}

type IntegrationStatus struct {
	Name          string   `json:"name,omitempty"`
	ContactGroups []string `json:"contact_groups,omitempty"`
	StatuspageID  string   `json:"statuspage_id,omitempty"`
	APIID         string   `json:"api_id,omitempty"`
	APIKey        string   `json:"api_key,omitempty"`
	Component     string   `json:"component,omitempty"`
	Container     string   `json:"container,omitempty"`
	Metric        string   `json:"metric,omitempty"`
}

type integrationsEndpointStatusImpl struct {
	EndpointCreator[IntegrationStatus, IntegrationResponse, Integration]
	EndpointUpdater[IntegrationStatus, IntegrationResponse, Integration]
}

func (i integrationsEndpointStatusImpl) CreateStatus(ctx context.Context, data IntegrationStatus) (*Integration, error) {
	return i.EndpointCreator.Create(ctx, data)
}

func (i integrationsEndpointStatusImpl) UpdateStatus(ctx context.Context, pk PrimaryKeyable, data IntegrationStatus) (*Integration, error) {
	return i.EndpointUpdater.Update(ctx, pk, data)
}

type IntegrationStatuspage struct {
	Name          string   `json:"name,omitempty"`
	ContactGroups []string `json:"contact_groups,omitempty"`
	APIKey        string   `json:"api_key,omitempty"`
	Page          string   `json:"page,omitempty"`
	Component     string   `json:"component,omitempty"`
	Metric        string   `json:"metric,omitempty"`
}

type integrationsEndpointStatuspageImpl struct {
	EndpointCreator[IntegrationStatuspage, IntegrationResponse, Integration]
	EndpointUpdater[IntegrationStatuspage, IntegrationResponse, Integration]
}

func (i integrationsEndpointStatuspageImpl) CreateStatuspage(ctx context.Context, data IntegrationStatuspage) (*Integration, error) {
	return i.EndpointCreator.Create(ctx, data)
}

func (i integrationsEndpointStatuspageImpl) UpdateStatuspage(ctx context.Context, pk PrimaryKeyable, data IntegrationStatuspage) (*Integration, error) {
	return i.EndpointUpdater.Update(ctx, pk, data)
}

type IntegrationTwitter struct {
	Name             string   `json:"name,omitempty"`
	ContactGroups    []string `json:"contact_groups,omitempty"`
	OauthToken       string   `json:"oauth_token,omitempty"`
	OauthTokenSecret string   `json:"oauth_token_secret,omitempty"`
}

type integrationsEndpointTwitterImpl struct {
	EndpointCreator[IntegrationTwitter, IntegrationResponse, Integration]
	EndpointUpdater[IntegrationTwitter, IntegrationResponse, Integration]
}

func (i integrationsEndpointTwitterImpl) CreateTwitter(ctx context.Context, data IntegrationTwitter) (*Integration, error) {
	return i.EndpointCreator.Create(ctx, data)
}

func (i integrationsEndpointTwitterImpl) UpdateTwitter(ctx context.Context, pk PrimaryKeyable, data IntegrationTwitter) (*Integration, error) {
	return i.EndpointUpdater.Update(ctx, pk, data)
}

type IntegrationVictorops struct {
	Name          string   `json:"name,omitempty"`
	ContactGroups []string `json:"contact_groups,omitempty"`
	ServiceKey    string   `json:"service_key,omitempty"`
	RoutingKey    string   `json:"routing_key,omitempty"`
}

type integrationsEndpointVictoropsImpl struct {
	EndpointCreator[IntegrationVictorops, IntegrationResponse, Integration]
	EndpointUpdater[IntegrationVictorops, IntegrationResponse, Integration]
}

func (i integrationsEndpointVictoropsImpl) CreateVictorops(ctx context.Context, data IntegrationVictorops) (*Integration, error) {
	return i.EndpointCreator.Create(ctx, data)
}

func (i integrationsEndpointVictoropsImpl) UpdateVictorops(ctx context.Context, pk PrimaryKeyable, data IntegrationVictorops) (*Integration, error) {
	return i.EndpointUpdater.Update(ctx, pk, data)
}

type IntegrationWavefront struct {
	Name          string   `json:"name,omitempty"`
	ContactGroups []string `json:"contact_groups,omitempty"`
	WavefrontUrl  string   `json:"wavefront_url,omitempty"`
	APIToken      string   `json:"api_token,omitempty"`
}

type integrationsEndpointWavefrontImpl struct {
	EndpointCreator[IntegrationWavefront, IntegrationResponse, Integration]
	EndpointUpdater[IntegrationWavefront, IntegrationResponse, Integration]
}

func (i integrationsEndpointWavefrontImpl) CreateWavefront(ctx context.Context, data IntegrationWavefront) (*Integration, error) {
	return i.EndpointCreator.Create(ctx, data)
}

func (i integrationsEndpointWavefrontImpl) UpdateWavefront(ctx context.Context, pk PrimaryKeyable, data IntegrationWavefront) (*Integration, error) {
	return i.EndpointUpdater.Update(ctx, pk, data)
}

type IntegrationWebhook struct {
	Name             string   `json:"name,omitempty"`
	ContactGroups    []string `json:"contact_groups,omitempty"`
	PostbackUrl      string   `json:"postback_url,omitempty"`
	Headers          string   `json:"headers,omitempty"`
	UseLegacyPayload bool     `json:"use_legacy_payload"`
}

type integrationsEndpointWebhookImpl struct {
	EndpointCreator[IntegrationWebhook, IntegrationResponse, Integration]
	EndpointUpdater[IntegrationWebhook, IntegrationResponse, Integration]
}

func (i integrationsEndpointWebhookImpl) CreateWebhook(ctx context.Context, data IntegrationWebhook) (*Integration, error) {
	return i.EndpointCreator.Create(ctx, data)
}

func (i integrationsEndpointWebhookImpl) UpdateWebhook(ctx context.Context, pk PrimaryKeyable, data IntegrationWebhook) (*Integration, error) {
	return i.EndpointUpdater.Update(ctx, pk, data)
}

type IntegrationZapier struct {
	Name          string   `json:"name,omitempty"`
	ContactGroups []string `json:"contact_groups,omitempty"`
	WebhookUrl    string   `json:"webhook_url,omitempty"`
}

type integrationsEndpointZapierImpl struct {
	EndpointCreator[IntegrationZapier, IntegrationResponse, Integration]
	EndpointUpdater[IntegrationZapier, IntegrationResponse, Integration]
}

func (i integrationsEndpointZapierImpl) CreateZapier(ctx context.Context, data IntegrationZapier) (*Integration, error) {
	return i.EndpointCreator.Create(ctx, data)
}

func (i integrationsEndpointZapierImpl) UpdateZapier(ctx context.Context, pk PrimaryKeyable, data IntegrationZapier) (*Integration, error) {
	return i.EndpointUpdater.Update(ctx, pk, data)
}
