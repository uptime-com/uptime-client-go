package upapi

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/shopspring/decimal"
)

type CheckSSLCertConfig struct {
	Protocol         string `json:"ssl_cert_protocol,omitempty" flag:"sslcert.protocol"`
	CRL              bool   `json:"ssl_cert_crl" flag:"sslcert.crl"`
	FirstElementOnly bool   `json:"ssl_cert_first_element_only" flag:"sslcert.first-element-only"`
	Match            string `json:"ssl_cert_match,omitempty" flag:"sslcert.match"`
	Issuer           string `json:"ssl_cert_issuer,omitempty" flag:"sslcert.issuer"`
	MinVersion       string `json:"ssl_cert_minimum_ssl_tls_version,omitempty" flag:"sslcert.min-version"`
	Fingerprint      string `json:"ssl_cert_fingerprint,omitempty" flag:"sslcert.fingerprint"`
	SelfSigned       bool   `json:"ssl_cert_selfsigned,omitempty" flag:"sslcert.self-signed"`
	URL              string `json:"ssl_cert_file,omitempty" flag:"sslcert.url"`
}

type CheckPageSpeedConfig struct {
	EmulatedDevice       string `json:"emulated_device,omitempty"`
	ConnectionThrottling string `json:"connection_throttling,omitempty"`
	ExcludeURLs          string `json:"exclude_urls,omitempty"`
	UptimeGradeThreshold string `json:"uptime_grade_threshold,omitempty"`
}

type CheckGroupConfig struct {
	CheckServices               []string `json:"group_check_services,omitempty"`
	CheckTags                   []string `json:"group_check_tags,omitempty"`
	CheckDownCondition          string   `json:"group_check_down_condition,omitempty"`
	UptimePercentCalculation    string   `json:"group_uptime_percent_calculation,omitempty"`
	ResponseTimeCalculationMode string   `json:"group_response_time_calculation_mode,omitempty"`
	ResponseTimeCheckType       string   `json:"group_response_time_check_type,omitempty"`
	ResponseTimeSingleCheck     string   `json:"group_response_time_single_check,omitempty"`
}

// Check represents a check in Uptime.com.
type Check struct {
	PK                     int64           `json:"pk,omitempty"`
	URL                    string          `json:"url,omitempty"`
	StatsURL               string          `json:"stats_url,omitempty"`
	AlertsURL              string          `json:"alerts_url,omitempty"`
	Name                   string          `json:"name,omitempty"`
	CachedResponseTime     float64         `json:"cached_response_time,omitempty"`
	ContactGroups          []string        `json:"contact_groups"`
	CreatedAt              time.Time       `json:"created_at,omitempty"`
	ModifiedAt             time.Time       `json:"modified_at,omitempty"`
	Locations              []string        `json:"locations,omitempty"`
	Tags                   []string        `json:"tags,omitempty"`
	CheckType              string          `json:"check_type,omitempty"`
	Escalations            string          `json:"-,omitempty"` // TODO
	MonitoringServiceType  string          `json:"monitoring_service_type,omitempty"`
	IsPaused               bool            `json:"is_paused,omitempty"`
	IsUnderMaintenance     bool            `json:"is_under_maintenance,omitempty"`
	StateIsUp              bool            `json:"state_is_up,omitempty"`
	StateChangedAt         time.Time       `json:"state_changed_at,omitempty"`
	HeartbeatURL           string          `json:"heartbeat_url,omitempty"`
	WebhookURL             string          `json:"webhook_url,omitempty"`
	Protocol               string          `json:"msp_protocol,omitempty"`
	Interval               int64           `json:"msp_interval,omitempty"`
	Address                string          `json:"msp_address"`
	Port                   int64           `json:"msp_port,omitempty"`
	Username               string          `json:"msp_username,omitempty"`
	Password               string          `json:"msp_password,omitempty"`
	Proxy                  string          `json:"msp_proxy,omitempty"`
	DNSServer              string          `json:"msp_dns_server,omitempty"`
	DNSRecordType          string          `json:"msp_dns_record_type,omitempty"`
	StatusCode             string          `json:"msp_status_code,omitempty"`
	SendString             string          `json:"msp_send_string,omitempty"`
	ExpectString           string          `json:"msp_expect_string,omitempty"`
	ExpectStringType       string          `json:"msp_expect_string_type,omitempty"`
	Encryption             string          `json:"msp_encryption,omitempty"`
	Threshold              int64           `json:"msp_threshold,omitempty"`
	Headers                string          `json:"msp_headers,omitempty"`
	Script                 string          `json:"msp_script,omitempty"`
	Version                int64           `json:"msp_version,omitempty"`
	Sensitivity            int64           `json:"msp_sensitivity,omitempty"`
	NumRetries             int64           `json:"msp_num_retries,omitempty"`
	UseIPVersion           string          `json:"msp_use_ip_version,omitempty"`
	UptimeSLA              decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string          `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool            `json:"msp_include_in_global_metrics,omitempty"`

	Maintenance *CheckMaintenance `json:"maintenance,omitempty"`

	SSLConfig       *CheckSSLCertConfig   `json:"sslconfig,omitempty"`
	PageSpeedConfig *CheckPageSpeedConfig `json:"pagespeedconfig,omitempty"`
	GroupConfig     *CheckGroupConfig     `json:"groupcheckconfig,omitempty"`
}

func (c Check) PrimaryKey() PrimaryKey {
	return PrimaryKey(c.PK)
}

type CheckGetResponse Check

func (c CheckGetResponse) Item() Check {
	return Check(c)
}

type CheckListResponse struct {
	Count   int64   `json:"count,omitempty"`
	Results []Check `json:"results,omitempty"`
}

func (r CheckListResponse) List() []Check {
	return r.Results
}

// CheckListOptions specifies the optional parameters to the CheckService.List method.
type CheckListOptions struct {
	Page                  int64    `url:"page,omitempty"`
	PageSize              int64    `url:"page_size,omitempty"`
	Search                string   `url:"search,omitempty"`
	Ordering              string   `url:"ordering,omitempty"`
	MonitoringServiceType string   `url:"monitoring_service_type,omitempty"`
	IsPaused              bool     `url:"is_paused"`
	StateIsUp             bool     `url:"state_is_up,omitempty"`
	Tag                   []string `url:"tag,omitempty"`
}

type CheckCreateUpdateResponse struct {
	Messages map[string]interface{} `json:"messages,omitempty"`
	Results  Check                  `json:"results,omitempty"`
}

func (r CheckCreateUpdateResponse) Item() Check {
	return r.Results
}

// CheckStatsOptions specifies the parameters to /api/v1/checks/{pk}/stats/ endpoint
type CheckStatsOptions struct {
	StartDate              string `url:"start_date,omitempty"`
	EndDate                string `url:"end_date,omitempty"`
	Location               string `url:"location,omitempty"`
	LocationsResponseTimes bool   `url:"locations_response_times,omitempty"`
	IncludeAlerts          bool   `url:"include_alerts,omitempty"`
	Download               bool   `url:"download,omitempty"`
	PDF                    bool   `url:"pdf,omitempty"`
}

// CheckStatsResponse represents the API response to a Stats query
type CheckStatsResponse struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Totals    struct {
		Outages      int64 `json:"outages,omitempty"`
		DowntimeSecs int64 `json:"downtime_secs,omitempty"`
	} `json:"totals"`
	Statistics []CheckStats `json:"statistics"`
}

func (c CheckStatsResponse) List() []CheckStats {
	return c.Statistics
}

type CheckStats struct {
	Date                   string   `json:"date"`
	Outages                int64    `json:"outages"`
	DowntimeSecs           int64    `json:"downtime_secs"`
	Uptime                 *float64 `json:"uptime,omitempty"`
	ResponseTime           *float64 `json:"response_time,omitempty"`
	ResponseTimeDatapoints [][]any  `json:"response_time_datapoints,omitempty"`
}

type ChecksEndpoint interface {
	List(context.Context, CheckListOptions) ([]Check, error)
	Get(context.Context, PrimaryKeyable) (*Check, error)
	Delete(context.Context, PrimaryKeyable) error
	Stats(context.Context, PrimaryKeyable, CheckStatsOptions) ([]CheckStats, error)

	CreateAPI(context.Context, CheckAPI) (*Check, error)
	UpdateAPI(context.Context, PrimaryKeyable, CheckAPI) (*Check, error)

	CreateBlacklist(context.Context, CheckBlacklist) (*Check, error)
	UpdateBlacklist(context.Context, PrimaryKeyable, CheckBlacklist) (*Check, error)

	CreateDNS(context.Context, CheckDNS) (*Check, error)
	UpdateDNS(context.Context, PrimaryKeyable, CheckDNS) (*Check, error)

	CreateGroup(context.Context, CheckGroup) (*Check, error)
	UpdateGroup(context.Context, PrimaryKeyable, CheckGroup) (*Check, error)

	CreateHeartbeat(context.Context, CheckHeartbeat) (*Check, error)
	UpdateHeartbeat(context.Context, PrimaryKeyable, CheckHeartbeat) (*Check, error)

	CreateHTTP(context.Context, CheckHTTP) (*Check, error)
	UpdateHTTP(context.Context, PrimaryKeyable, CheckHTTP) (*Check, error)

	CreateICMP(context.Context, CheckICMP) (*Check, error)
	UpdateICMP(context.Context, PrimaryKeyable, CheckICMP) (*Check, error)

	CreateIMAP(context.Context, CheckIMAP) (*Check, error)
	UpdateIMAP(context.Context, PrimaryKeyable, CheckIMAP) (*Check, error)

	CreateMalware(context.Context, CheckMalware) (*Check, error)
	UpdateMalware(context.Context, PrimaryKeyable, CheckMalware) (*Check, error)

	CreateNTP(context.Context, CheckNTP) (*Check, error)
	UpdateNTP(context.Context, PrimaryKeyable, CheckNTP) (*Check, error)

	CreatePOP(context.Context, CheckPOP) (*Check, error)
	UpdatePOP(context.Context, PrimaryKeyable, CheckPOP) (*Check, error)

	CreateRUM(context.Context, CheckRUM) (*Check, error)
	UpdateRUM(context.Context, PrimaryKeyable, CheckRUM) (*Check, error)

	CreateRUM2(context.Context, CheckRUM2) (*Check, error)
	UpdateRUM2(context.Context, PrimaryKeyable, CheckRUM2) (*Check, error)

	CreateSMTP(context.Context, CheckSMTP) (*Check, error)
	UpdateSMTP(context.Context, PrimaryKeyable, CheckSMTP) (*Check, error)

	CreateSSH(context.Context, CheckSSH) (*Check, error)
	UpdateSSH(context.Context, PrimaryKeyable, CheckSSH) (*Check, error)

	CreateSSLCert(context.Context, CheckSSLCert) (*Check, error)
	UpdateSSLCert(context.Context, PrimaryKeyable, CheckSSLCert) (*Check, error)

	CreateTCP(context.Context, CheckTCP) (*Check, error)
	UpdateTCP(context.Context, PrimaryKeyable, CheckTCP) (*Check, error)

	CreateTransaction(context.Context, CheckTransaction) (*Check, error)
	UpdateTransaction(context.Context, PrimaryKeyable, CheckTransaction) (*Check, error)

	CreateUDP(context.Context, CheckUDP) (*Check, error)
	UpdateUDP(context.Context, PrimaryKeyable, CheckUDP) (*Check, error)

	CreateWebhook(context.Context, CheckWebhook) (*Check, error)
	UpdateWebhook(context.Context, PrimaryKeyable, CheckWebhook) (*Check, error)

	CreateWHOIS(context.Context, CheckWHOIS) (*Check, error)
	UpdateWHOIS(context.Context, PrimaryKeyable, CheckWHOIS) (*Check, error)

	CreatePageSpeed(context.Context, CheckPageSpeed) (*Check, error)
	UpdatePageSpeed(context.Context, PrimaryKeyable, CheckPageSpeed) (*Check, error)

	UpdateMaintenance(context.Context, PrimaryKeyable, CheckMaintenance) (*Check, error)
}

func NewChecksEndpoint(cbd CBD) ChecksEndpoint {
	endpoint := "checks"
	return &checksEndpointImpl{
		checksEndpointAPIImpl: checksEndpointAPIImpl{
			EndpointCreator: NewEndpointCreator[CheckAPI, CheckCreateUpdateResponse, Check](cbd, endpoint+"/add-api"),
			EndpointUpdater: NewEndpointUpdater[CheckAPI, CheckCreateUpdateResponse, Check](cbd, endpoint),
		},
		checksEndpointBlacklistImpl: checksEndpointBlacklistImpl{
			EndpointCreator: NewEndpointCreator[CheckBlacklist, CheckCreateUpdateResponse, Check](cbd, endpoint+"/add-blacklist"),
			EndpointUpdater: NewEndpointUpdater[CheckBlacklist, CheckCreateUpdateResponse, Check](cbd, endpoint),
		},
		checksEndpointDNSImpl: checksEndpointDNSImpl{
			EndpointCreator: NewEndpointCreator[CheckDNS, CheckCreateUpdateResponse, Check](cbd, endpoint+"/add-dns"),
			EndpointUpdater: NewEndpointUpdater[CheckDNS, CheckCreateUpdateResponse, Check](cbd, endpoint),
		},
		checksEndpointGroupImpl: checksEndpointGroupImpl{
			EndpointCreator: NewEndpointCreator[CheckGroup, CheckCreateUpdateResponse, Check](cbd, endpoint+"/add-group"),
			EndpointUpdater: NewEndpointUpdater[CheckGroup, CheckCreateUpdateResponse, Check](cbd, endpoint),
		},
		checksEndpointHeartbeatImpl: checksEndpointHeartbeatImpl{
			EndpointCreator: NewEndpointCreator[CheckHeartbeat, CheckCreateUpdateResponse, Check](cbd, endpoint+"/add-heartbeat"),
			EndpointUpdater: NewEndpointUpdater[CheckHeartbeat, CheckCreateUpdateResponse, Check](cbd, endpoint),
		},
		checksEndpointHTTPImpl: checksEndpointHTTPImpl{
			EndpointCreator: NewEndpointCreator[CheckHTTP, CheckCreateUpdateResponse, Check](cbd, endpoint+"/add-http"),
			EndpointUpdater: NewEndpointUpdater[CheckHTTP, CheckCreateUpdateResponse, Check](cbd, endpoint),
		},
		checksEndpointICMPImpl: checksEndpointICMPImpl{
			EndpointCreator: NewEndpointCreator[CheckICMP, CheckCreateUpdateResponse, Check](cbd, endpoint+"/add-icmp"),
			EndpointUpdater: NewEndpointUpdater[CheckICMP, CheckCreateUpdateResponse, Check](cbd, endpoint),
		},
		checksEndpointIMAPImpl: checksEndpointIMAPImpl{
			EndpointCreator: NewEndpointCreator[CheckIMAP, CheckCreateUpdateResponse, Check](cbd, endpoint+"/add-imap"),
			EndpointUpdater: NewEndpointUpdater[CheckIMAP, CheckCreateUpdateResponse, Check](cbd, endpoint),
		},
		checksEndpointMalwareImpl: checksEndpointMalwareImpl{
			EndpointCreator: NewEndpointCreator[CheckMalware, CheckCreateUpdateResponse, Check](cbd, endpoint+"/add-malware"),
			EndpointUpdater: NewEndpointUpdater[CheckMalware, CheckCreateUpdateResponse, Check](cbd, endpoint),
		},
		checksEndpointNTPImpl: checksEndpointNTPImpl{
			EndpointCreator: NewEndpointCreator[CheckNTP, CheckCreateUpdateResponse, Check](cbd, endpoint+"/add-ntp"),
			EndpointUpdater: NewEndpointUpdater[CheckNTP, CheckCreateUpdateResponse, Check](cbd, endpoint),
		},
		checksEndpointPOPImpl: checksEndpointPOPImpl{
			EndpointCreator: NewEndpointCreator[CheckPOP, CheckCreateUpdateResponse, Check](cbd, endpoint+"/add-pop"),
			EndpointUpdater: NewEndpointUpdater[CheckPOP, CheckCreateUpdateResponse, Check](cbd, endpoint),
		},
		checksEndpointRUMImpl: checksEndpointRUMImpl{
			EndpointCreator: NewEndpointCreator[CheckRUM, CheckCreateUpdateResponse, Check](cbd, endpoint+"/add-rum"),
			EndpointUpdater: NewEndpointUpdater[CheckRUM, CheckCreateUpdateResponse, Check](cbd, endpoint),
		},
		checksEndpointRUM2Impl: checksEndpointRUM2Impl{
			EndpointCreator: NewEndpointCreator[CheckRUM2, CheckCreateUpdateResponse, Check](cbd, endpoint+"/add-rum2"),
			EndpointUpdater: NewEndpointUpdater[CheckRUM2, CheckCreateUpdateResponse, Check](cbd, endpoint),
		},
		checksEndpointSMTPImpl: checksEndpointSMTPImpl{
			EndpointCreator: NewEndpointCreator[CheckSMTP, CheckCreateUpdateResponse, Check](cbd, endpoint+"/add-smtp"),
			EndpointUpdater: NewEndpointUpdater[CheckSMTP, CheckCreateUpdateResponse, Check](cbd, endpoint),
		},
		checksEndpointSSHImpl: checksEndpointSSHImpl{
			EndpointCreator: NewEndpointCreator[CheckSSH, CheckCreateUpdateResponse, Check](cbd, endpoint+"/add-ssh"),
			EndpointUpdater: NewEndpointUpdater[CheckSSH, CheckCreateUpdateResponse, Check](cbd, endpoint),
		},
		checksEndpointSSLCertImpl: checksEndpointSSLCertImpl{
			EndpointCreator: NewEndpointCreator[CheckSSLCert, CheckCreateUpdateResponse, Check](cbd, endpoint+"/add-ssl-cert"),
			EndpointUpdater: NewEndpointUpdater[CheckSSLCert, CheckCreateUpdateResponse, Check](cbd, endpoint),
		},
		checksEndpointTCPImpl: checksEndpointTCPImpl{
			EndpointCreator: NewEndpointCreator[CheckTCP, CheckCreateUpdateResponse, Check](cbd, endpoint+"/add-tcp"),
			EndpointUpdater: NewEndpointUpdater[CheckTCP, CheckCreateUpdateResponse, Check](cbd, endpoint),
		},
		checksEndpointTransactionImpl: checksEndpointTransactionImpl{
			EndpointCreator: NewEndpointCreator[CheckTransaction, CheckCreateUpdateResponse, Check](cbd, endpoint+"/add-transaction"),
			EndpointUpdater: NewEndpointUpdater[CheckTransaction, CheckCreateUpdateResponse, Check](cbd, endpoint),
		},
		checksEndpointUDPImpl: checksEndpointUDPImpl{
			EndpointCreator: NewEndpointCreator[CheckUDP, CheckCreateUpdateResponse, Check](cbd, endpoint+"/add-udp"),
			EndpointUpdater: NewEndpointUpdater[CheckUDP, CheckCreateUpdateResponse, Check](cbd, endpoint),
		},
		checksEndpointWebhookImpl: checksEndpointWebhookImpl{
			EndpointCreator: NewEndpointCreator[CheckWebhook, CheckCreateUpdateResponse, Check](cbd, endpoint+"/add-webhook"),
			EndpointUpdater: NewEndpointUpdater[CheckWebhook, CheckCreateUpdateResponse, Check](cbd, endpoint),
		},
		checksEndpointWHOISImpl: checksEndpointWHOISImpl{
			EndpointCreator: NewEndpointCreator[CheckWHOIS, CheckCreateUpdateResponse, Check](cbd, endpoint+"/add-whois"),
			EndpointUpdater: NewEndpointUpdater[CheckWHOIS, CheckCreateUpdateResponse, Check](cbd, endpoint),
		},
		checksStatsEndpointImpl: checksStatsEndpointImpl{
			endpoint: NewEndpointLister[CheckStatsResponse, CheckStats, CheckStatsOptions](&checksNestedEndpointCBD{CBD: cbd}, endpoint+"/%d/stats"),
		},
		checksEndpointPageSpeedImpl: checksEndpointPageSpeedImpl{
			EndpointCreator: NewEndpointCreator[CheckPageSpeed, CheckCreateUpdateResponse, Check](cbd, endpoint+"/add-pagespeed"),
			EndpointUpdater: NewEndpointUpdater[CheckPageSpeed, CheckCreateUpdateResponse, Check](cbd, endpoint),
		},
		checksEndpointMaintenanceImpl: checksEndpointMaintenanceImpl{
			EndpointUpdater: NewEndpointUpdater[CheckMaintenance, CheckCreateUpdateResponse, Check](
				&checksNestedEndpointCBD{CBD: cbd, EndpointSuffix: "maintenance/"}, endpoint,
			),
		},
		EndpointLister:  NewEndpointLister[CheckListResponse, Check, CheckListOptions](cbd, endpoint),
		EndpointGetter:  NewEndpointGetter[CheckGetResponse, Check](cbd, endpoint),
		EndpointDeleter: NewEndpointDeleter(cbd, endpoint),
	}
}

type checksEndpointImpl struct {
	checksEndpointAPIImpl
	checksEndpointBlacklistImpl
	checksEndpointDNSImpl
	checksEndpointGroupImpl
	checksEndpointHeartbeatImpl
	checksEndpointHTTPImpl
	checksEndpointICMPImpl
	checksEndpointIMAPImpl
	checksEndpointMalwareImpl
	checksEndpointNTPImpl
	checksEndpointPOPImpl
	checksEndpointRUMImpl
	checksEndpointRUM2Impl
	checksEndpointSMTPImpl
	checksEndpointSSHImpl
	checksEndpointSSLCertImpl
	checksEndpointTCPImpl
	checksEndpointTransactionImpl
	checksEndpointUDPImpl
	checksEndpointWebhookImpl
	checksEndpointWHOISImpl
	checksStatsEndpointImpl
	checksEndpointPageSpeedImpl
	checksEndpointMaintenanceImpl
	EndpointLister[CheckListResponse, Check, CheckListOptions]
	EndpointGetter[CheckCreateUpdateResponse, Check]
	EndpointUpdater[Check, CheckCreateUpdateResponse, Check]
	EndpointDeleter
}

type checksPKCtxKey struct{}

type checksStatsEndpointImpl struct {
	endpoint EndpointLister[CheckStatsResponse, CheckStats, CheckStatsOptions]
}

func (c *checksEndpointImpl) Stats(ctx context.Context, pk PrimaryKeyable, opts CheckStatsOptions) ([]CheckStats, error) {
	ctx = context.WithValue(ctx, checksPKCtxKey{}, pk)
	return c.endpoint.List(ctx, opts)
}

type checksNestedEndpointCBD struct {
	CBD
	EndpointSuffix string
}

func (c checksNestedEndpointCBD) BuildRequest(ctx context.Context, method string, endpoint string, args any, data any) (*http.Request, error) {
	if c.EndpointSuffix != "" {
		endpoint += c.EndpointSuffix
	}
	pk, ok := ctx.Value(checksPKCtxKey{}).(PrimaryKey)
	if ok {
		endpoint = fmt.Sprintf(endpoint, pk.PrimaryKey())
	}
	return c.CBD.BuildRequest(ctx, method, endpoint, args, data)
}

type CheckAPI struct {
	Name                   string          `json:"name,omitempty"`
	ContactGroups          []string        `json:"contact_groups,omitempty"`
	Locations              []string        `json:"locations,omitempty"`
	Tags                   []string        `json:"tags,omitempty"`
	IsPaused               bool            `json:"is_paused"`
	Interval               int64           `json:"msp_interval,omitempty"`
	Threshold              int64           `json:"msp_threshold,omitempty"`
	Script                 string          `json:"msp_script,omitempty"`
	Sensitivity            int64           `json:"msp_sensitivity,omitempty"`
	NumRetries             int64           `json:"msp_num_retries,omitempty"`
	UseIPVersion           string          `json:"msp_use_ip_version,omitempty"`
	UptimeSLA              decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string          `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool            `json:"msp_include_in_global_metrics"`
}

type checksEndpointAPIImpl struct {
	EndpointCreator[CheckAPI, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckAPI, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointAPIImpl) CreateAPI(ctx context.Context, check CheckAPI) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointAPIImpl) UpdateAPI(ctx context.Context, pk PrimaryKeyable, check CheckAPI) (*Check, error) {
	return c.Update(ctx, pk, check)
}

type CheckBlacklist struct {
	Name          string          `json:"name,omitempty"`
	ContactGroups []string        `json:"contact_groups,omitempty"`
	Locations     []string        `json:"locations,omitempty"`
	Tags          []string        `json:"tags,omitempty"`
	IsPaused      bool            `json:"is_paused"`
	Address       string          `json:"msp_address"`
	NumRetries    int64           `json:"msp_num_retries,omitempty"`
	UptimeSLA     decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	Notes         string          `json:"msp_notes,omitempty"`
}

type checksEndpointBlacklistImpl struct {
	EndpointCreator[CheckBlacklist, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckBlacklist, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointBlacklistImpl) CreateBlacklist(ctx context.Context, check CheckBlacklist) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointBlacklistImpl) UpdateBlacklist(ctx context.Context, pk PrimaryKeyable, check CheckBlacklist) (*Check, error) {
	return c.Update(ctx, pk, check)
}

type CheckDNS struct {
	Name                   string          `json:"name,omitempty"`
	ContactGroups          []string        `json:"contact_groups,omitempty"`
	Locations              []string        `json:"locations,omitempty"`
	Tags                   []string        `json:"tags,omitempty"`
	IsPaused               bool            `json:"is_paused"`
	Interval               int64           `json:"msp_interval,omitempty"`
	Address                string          `json:"msp_address"`
	DNSServer              string          `json:"msp_dns_server,omitempty"`
	DNSRecordType          string          `json:"msp_dns_record_type,omitempty"`
	ExpectString           string          `json:"msp_expect_string,omitempty"`
	Threshold              int64           `json:"msp_threshold,omitempty"`
	Sensitivity            int64           `json:"msp_sensitivity,omitempty"`
	NumRetries             int64           `json:"msp_num_retries,omitempty"`
	UptimeSLA              decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string          `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool            `json:"msp_include_in_global_metrics"`
}

type checksEndpointDNSImpl struct {
	EndpointCreator[CheckDNS, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckDNS, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointDNSImpl) CreateDNS(ctx context.Context, check CheckDNS) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointDNSImpl) UpdateDNS(ctx context.Context, pk PrimaryKeyable, check CheckDNS) (*Check, error) {
	return c.Update(ctx, pk, check)
}

type CheckGroup struct {
	Name                   string           `json:"name,omitempty"`
	ContactGroups          []string         `json:"contact_groups"`
	Locations              []string         `json:"locations,omitempty"`
	Tags                   []string         `json:"tags,omitempty"`
	IsPaused               bool             `json:"is_paused"`
	UptimeSLA              decimal.Decimal  `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        decimal.Decimal  `json:"msp_response_time_sla,omitempty"`
	Notes                  string           `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool             `json:"msp_include_in_global_metrics"`
	Config                 CheckGroupConfig `json:"groupcheckconfig,omitempty"`
}

type checksEndpointGroupImpl struct {
	EndpointCreator[CheckGroup, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckGroup, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointGroupImpl) CreateGroup(ctx context.Context, check CheckGroup) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointGroupImpl) UpdateGroup(ctx context.Context, pk PrimaryKeyable, check CheckGroup) (*Check, error) {
	return c.Update(ctx, pk, check)
}

type CheckHeartbeat struct {
	Name                   string          `json:"name,omitempty"`
	ContactGroups          []string        `json:"contact_groups,omitempty"`
	Tags                   []string        `json:"tags,omitempty"`
	IsPaused               bool            `json:"is_paused"`
	Interval               int64           `json:"msp_interval,omitempty"`
	UptimeSLA              decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string          `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool            `json:"msp_include_in_global_metrics"`
	HeartbeatURL           string          `json:"heartbeat_url,omitempty"`
}

type checksEndpointHeartbeatImpl struct {
	EndpointCreator[CheckHeartbeat, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckHeartbeat, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointHeartbeatImpl) CreateHeartbeat(ctx context.Context, check CheckHeartbeat) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointHeartbeatImpl) UpdateHeartbeat(ctx context.Context, pk PrimaryKeyable, check CheckHeartbeat) (*Check, error) {
	return c.Update(ctx, pk, check)
}

type CheckHTTP struct {
	Name                   string          `json:"name,omitempty"`
	ContactGroups          []string        `json:"contact_groups,omitempty"`
	Locations              []string        `json:"locations,omitempty"`
	Tags                   []string        `json:"tags,omitempty"`
	IsPaused               bool            `json:"is_paused"`
	Interval               int64           `json:"msp_interval,omitempty"`
	Address                string          `json:"msp_address"`
	Port                   int64           `json:"msp_port,omitempty"`
	Username               string          `json:"msp_username,omitempty"`
	Password               string          `json:"msp_password,omitempty"`
	Proxy                  string          `json:"msp_proxy,omitempty"`
	StatusCode             string          `json:"msp_status_code"`
	SendString             string          `json:"msp_send_string,omitempty"`
	ExpectString           string          `json:"msp_expect_string,omitempty"`
	ExpectStringType       string          `json:"msp_expect_string_type,omitempty"`
	Encryption             string          `json:"msp_encryption"`
	Threshold              int64           `json:"msp_threshold,omitempty"`
	Headers                string          `json:"msp_headers,omitempty"`
	Version                int64           `json:"msp_version,omitempty"`
	Sensitivity            int64           `json:"msp_sensitivity,omitempty"`
	NumRetries             int64           `json:"msp_num_retries,omitempty"`
	UseIPVersion           string          `json:"msp_use_ip_version,omitempty"`
	UptimeSLA              decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string          `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool            `json:"msp_include_in_global_metrics"`
}

type checksEndpointHTTPImpl struct {
	EndpointCreator[CheckHTTP, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckHTTP, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointHTTPImpl) CreateHTTP(ctx context.Context, check CheckHTTP) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointHTTPImpl) UpdateHTTP(ctx context.Context, pk PrimaryKeyable, check CheckHTTP) (*Check, error) {
	return c.Update(ctx, pk, check)
}

type CheckICMP struct {
	Name                   string          `json:"name,omitempty"`
	ContactGroups          []string        `json:"contact_groups,omitempty"`
	Locations              []string        `json:"locations,omitempty"`
	Tags                   []string        `json:"tags,omitempty"`
	IsPaused               bool            `json:"is_paused"`
	Interval               int64           `json:"msp_interval,omitempty"`
	Address                string          `json:"msp_address"`
	Sensitivity            int64           `json:"msp_sensitivity,omitempty"`
	NumRetries             int64           `json:"msp_num_retries,omitempty"`
	UseIPVersion           string          `json:"msp_use_ip_version,omitempty"`
	UptimeSLA              decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string          `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool            `json:"msp_include_in_global_metrics"`
}

type checksEndpointICMPImpl struct {
	EndpointCreator[CheckICMP, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckICMP, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointICMPImpl) CreateICMP(ctx context.Context, check CheckICMP) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointICMPImpl) UpdateICMP(ctx context.Context, pk PrimaryKeyable, check CheckICMP) (*Check, error) {
	return c.Update(ctx, pk, check)
}

type CheckIMAP struct {
	Name                   string          `json:"name,omitempty"`
	ContactGroups          []string        `json:"contact_groups,omitempty"`
	Locations              []string        `json:"locations,omitempty"`
	Tags                   []string        `json:"tags,omitempty"`
	IsPaused               bool            `json:"is_paused"`
	Interval               int64           `json:"msp_interval,omitempty"`
	Address                string          `json:"msp_address"`
	Port                   int64           `json:"msp_port,omitempty"`
	ExpectString           string          `json:"msp_expect_string,omitempty"`
	Encryption             string          `json:"msp_encryption,omitempty"`
	Sensitivity            int64           `json:"msp_sensitivity,omitempty"`
	NumRetries             int64           `json:"msp_num_retries,omitempty"`
	UseIPVersion           string          `json:"msp_use_ip_version,omitempty"`
	UptimeSLA              decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string          `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool            `json:"msp_include_in_global_metrics"`
}

type checksEndpointIMAPImpl struct {
	EndpointCreator[CheckIMAP, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckIMAP, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointIMAPImpl) CreateIMAP(ctx context.Context, check CheckIMAP) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointIMAPImpl) UpdateIMAP(ctx context.Context, pk PrimaryKeyable, check CheckIMAP) (*Check, error) {
	return c.Update(ctx, pk, check)
}

type CheckMalware struct {
	Name          string          `json:"name,omitempty"`
	ContactGroups []string        `json:"contact_groups,omitempty"`
	Locations     []string        `json:"locations,omitempty"`
	Tags          []string        `json:"tags,omitempty"`
	IsPaused      bool            `json:"is_paused"`
	Address       string          `json:"msp_address"`
	NumRetries    int64           `json:"msp_num_retries,omitempty"`
	UptimeSLA     decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	Notes         string          `json:"msp_notes,omitempty"`
}

type checksEndpointMalwareImpl struct {
	EndpointCreator[CheckMalware, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckMalware, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointMalwareImpl) CreateMalware(ctx context.Context, check CheckMalware) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointMalwareImpl) UpdateMalware(ctx context.Context, pk PrimaryKeyable, check CheckMalware) (*Check, error) {
	return c.Update(ctx, pk, check)
}

type CheckNTP struct {
	Name                   string          `json:"name,omitempty"`
	ContactGroups          []string        `json:"contact_groups,omitempty"`
	Locations              []string        `json:"locations,omitempty"`
	Tags                   []string        `json:"tags,omitempty"`
	IsPaused               bool            `json:"is_paused"`
	Interval               int64           `json:"msp_interval"`
	Address                string          `json:"msp_address"`
	Port                   int64           `json:"msp_port,omitempty"`
	Threshold              int64           `json:"msp_threshold,omitempty"`
	Sensitivity            int64           `json:"msp_sensitivity,omitempty"`
	NumRetries             int64           `json:"msp_num_retries,omitempty"`
	UseIPVersion           string          `json:"msp_use_ip_version,omitempty"`
	UptimeSLA              decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string          `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool            `json:"msp_include_in_global_metrics"`
}

type checksEndpointNTPImpl struct {
	EndpointCreator[CheckNTP, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckNTP, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointNTPImpl) CreateNTP(ctx context.Context, check CheckNTP) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointNTPImpl) UpdateNTP(ctx context.Context, pk PrimaryKeyable, check CheckNTP) (*Check, error) {
	return c.Update(ctx, pk, check)
}

type CheckPOP struct {
	Name                   string          `json:"name,omitempty"`
	ContactGroups          []string        `json:"contact_groups,omitempty"`
	Locations              []string        `json:"locations,omitempty"`
	Tags                   []string        `json:"tags,omitempty"`
	IsPaused               bool            `json:"is_paused"`
	Interval               int64           `json:"msp_interval,omitempty"`
	Address                string          `json:"msp_address"`
	Port                   int64           `json:"msp_port,omitempty"`
	ExpectString           string          `json:"msp_expect_string,omitempty"`
	Encryption             string          `json:"msp_encryption,omitempty"`
	Sensitivity            int64           `json:"msp_sensitivity,omitempty"`
	NumRetries             int64           `json:"msp_num_retries,omitempty"`
	UseIPVersion           string          `json:"msp_use_ip_version,omitempty"`
	UptimeSLA              decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string          `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool            `json:"msp_include_in_global_metrics"`
}

type checksEndpointPOPImpl struct {
	EndpointCreator[CheckPOP, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckPOP, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointPOPImpl) CreatePOP(ctx context.Context, check CheckPOP) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointPOPImpl) UpdatePOP(ctx context.Context, pk PrimaryKeyable, check CheckPOP) (*Check, error) {
	return c.Update(ctx, pk, check)
}

type CheckRUM struct {
	Name                   string          `json:"name,omitempty"`
	ContactGroups          []string        `json:"contact_groups,omitempty"`
	Locations              []string        `json:"locations,omitempty"`
	Tags                   []string        `json:"tags,omitempty"`
	IsPaused               bool            `json:"is_paused"`
	Address                string          `json:"msp_address"`
	Threshold              int64           `json:"msp_threshold,omitempty"`
	UptimeSLA              decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	Notes                  string          `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool            `json:"msp_include_in_global_metrics"`
}

type checksEndpointRUMImpl struct {
	EndpointCreator[CheckRUM, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckRUM, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointRUMImpl) CreateRUM(ctx context.Context, check CheckRUM) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointRUMImpl) UpdateRUM(ctx context.Context, pk PrimaryKeyable, check CheckRUM) (*Check, error) {
	return c.Update(ctx, pk, check)
}

type CheckRUM2 struct {
	Name                   string          `json:"name,omitempty"`
	ContactGroups          []string        `json:"contact_groups,omitempty"`
	Locations              []string        `json:"locations,omitempty"`
	Tags                   []string        `json:"tags,omitempty"`
	IsPaused               bool            `json:"is_paused"`
	Address                string          `json:"msp_address"`
	Threshold              int64           `json:"msp_threshold,omitempty"`
	UptimeSLA              decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	Notes                  string          `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool            `json:"msp_include_in_global_metrics"`
}

type checksEndpointRUM2Impl struct {
	EndpointCreator[CheckRUM2, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckRUM2, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointRUM2Impl) CreateRUM2(ctx context.Context, check CheckRUM2) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointRUM2Impl) UpdateRUM2(ctx context.Context, pk PrimaryKeyable, check CheckRUM2) (*Check, error) {
	return c.Update(ctx, pk, check)
}

type CheckSMTP struct {
	Name                   string          `json:"name,omitempty"`
	ContactGroups          []string        `json:"contact_groups,omitempty"`
	Locations              []string        `json:"locations,omitempty"`
	Tags                   []string        `json:"tags,omitempty"`
	IsPaused               bool            `json:"is_paused"`
	Interval               int64           `json:"msp_interval"`
	Address                string          `json:"msp_address"`
	Port                   int64           `json:"msp_port,omitempty"`
	Username               string          `json:"msp_username,omitempty"`
	Password               string          `json:"msp_password,omitempty"`
	ExpectString           string          `json:"msp_expect_string,omitempty"`
	Encryption             string          `json:"msp_encryption,omitempty"`
	Sensitivity            int64           `json:"msp_sensitivity,omitempty"`
	NumRetries             int64           `json:"msp_num_retries,omitempty"`
	UseIpVersion           string          `json:"msp_use_ip_version,omitempty"`
	UptimeSLA              decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string          `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool            `json:"msp_include_in_global_metrics"`
}

type checksEndpointSMTPImpl struct {
	EndpointCreator[CheckSMTP, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckSMTP, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointSMTPImpl) CreateSMTP(ctx context.Context, check CheckSMTP) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointSMTPImpl) UpdateSMTP(ctx context.Context, pk PrimaryKeyable, check CheckSMTP) (*Check, error) {
	return c.Update(ctx, pk, check)
}

type CheckSSH struct {
	Name                   string          `json:"name,omitempty"`
	ContactGroups          []string        `json:"contact_groups,omitempty"`
	Locations              []string        `json:"locations,omitempty"`
	Tags                   []string        `json:"tags,omitempty"`
	IsPaused               bool            `json:"is_paused"`
	Interval               int64           `json:"msp_interval,omitempty"`
	Address                string          `json:"msp_address"`
	Port                   int64           `json:"msp_port,omitempty"`
	Sensitivity            int64           `json:"msp_sensitivity,omitempty"`
	NumRetries             int64           `json:"msp_num_retries,omitempty"`
	UseIpVersion           string          `json:"msp_use_ip_version,omitempty"`
	UptimeSLA              decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string          `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool            `json:"msp_include_in_global_metrics"`
}

type checksEndpointSSHImpl struct {
	EndpointCreator[CheckSSH, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckSSH, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointSSHImpl) CreateSSH(ctx context.Context, check CheckSSH) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointSSHImpl) UpdateSSH(ctx context.Context, pk PrimaryKeyable, check CheckSSH) (*Check, error) {
	return c.Update(ctx, pk, check)
}

type CheckSSLCert struct {
	Name          string             `json:"name,omitempty"`
	ContactGroups []string           `json:"contact_groups,omitempty"`
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

type CheckTCP struct {
	Name                   string          `json:"name,omitempty"`
	ContactGroups          []string        `json:"contact_groups,omitempty"`
	Locations              []string        `json:"locations,omitempty"`
	Tags                   []string        `json:"tags,omitempty"`
	IsPaused               bool            `json:"is_paused"`
	Interval               int64           `json:"msp_interval,omitempty"`
	Address                string          `json:"msp_address"`
	Port                   int64           `json:"msp_port,omitempty"`
	SendString             string          `json:"msp_send_string,omitempty"`
	ExpectString           string          `json:"msp_expect_string,omitempty"`
	Sensitivity            int64           `json:"msp_sensitivity,omitempty"`
	NumRetries             int64           `json:"msp_num_retries,omitempty"`
	UseIpVersion           string          `json:"msp_use_ip_version,omitempty"`
	UptimeSLA              decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string          `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool            `json:"msp_include_in_global_metrics"`
	Encryption             string          `json:"msp_encryption"`
}

type checksEndpointTCPImpl struct {
	EndpointCreator[CheckTCP, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckTCP, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointTCPImpl) CreateTCP(ctx context.Context, check CheckTCP) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointTCPImpl) UpdateTCP(ctx context.Context, pk PrimaryKeyable, check CheckTCP) (*Check, error) {
	return c.Update(ctx, pk, check)
}

type CheckTransaction struct {
	Name                   string          `json:"name,omitempty"`
	ContactGroups          []string        `json:"contact_groups,omitempty"`
	Locations              []string        `json:"locations,omitempty"`
	Tags                   []string        `json:"tags,omitempty"`
	IsPaused               bool            `json:"is_paused"`
	Interval               int64           `json:"msp_interval,omitempty"`
	Threshold              int64           `json:"msp_threshold,omitempty"`
	Script                 string          `json:"msp_script,omitempty"`
	Sensitivity            int64           `json:"msp_sensitivity,omitempty"`
	NumRetries             int64           `json:"msp_num_retries,omitempty"`
	UptimeSLA              decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string          `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool            `json:"msp_include_in_global_metrics"`
}

type checksEndpointTransactionImpl struct {
	EndpointCreator[CheckTransaction, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckTransaction, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointTransactionImpl) CreateTransaction(ctx context.Context, check CheckTransaction) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointTransactionImpl) UpdateTransaction(ctx context.Context, pk PrimaryKeyable, check CheckTransaction) (*Check, error) {
	return c.Update(ctx, pk, check)
}

type CheckUDP struct {
	Name                   string          `json:"name,omitempty"`
	ContactGroups          []string        `json:"contact_groups,omitempty"`
	Locations              []string        `json:"locations,omitempty"`
	Tags                   []string        `json:"tags,omitempty"`
	IsPaused               bool            `json:"is_paused"`
	Interval               int64           `json:"msp_interval,omitempty"`
	Address                string          `json:"msp_address"`
	Port                   int64           `json:"msp_port,omitempty"`
	SendString             string          `json:"msp_send_string,omitempty"`
	ExpectString           string          `json:"msp_expect_string,omitempty"`
	Sensitivity            int64           `json:"msp_sensitivity,omitempty"`
	NumRetries             int64           `json:"msp_num_retries,omitempty"`
	UseIpVersion           string          `json:"msp_use_ip_version,omitempty"`
	UptimeSLA              decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string          `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool            `json:"msp_include_in_global_metrics"`
}

type checksEndpointUDPImpl struct {
	EndpointCreator[CheckUDP, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckUDP, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointUDPImpl) CreateUDP(ctx context.Context, check CheckUDP) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointUDPImpl) UpdateUDP(ctx context.Context, pk PrimaryKeyable, check CheckUDP) (*Check, error) {
	return c.Update(ctx, pk, check)
}

type CheckWebhook struct {
	Name                   string          `json:"name,omitempty"`
	ContactGroups          []string        `json:"contact_groups,omitempty"`
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

type CheckWHOIS struct {
	Name          string          `json:"name,omitempty"`
	ContactGroups []string        `json:"contact_groups,omitempty"`
	Locations     []string        `json:"locations,omitempty"`
	Tags          []string        `json:"tags,omitempty"`
	IsPaused      bool            `json:"is_paused"`
	Address       string          `json:"msp_address"`
	ExpectString  string          `json:"msp_expect_string,omitempty"`
	Threshold     int64           `json:"msp_threshold,omitempty"`
	NumRetries    int64           `json:"msp_num_retries,omitempty"`
	UptimeSLA     decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	Notes         string          `json:"msp_notes,omitempty"`
}

type checksEndpointWHOISImpl struct {
	EndpointCreator[CheckWHOIS, CheckCreateUpdateResponse, Check]
	EndpointUpdater[CheckWHOIS, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointWHOISImpl) CreateWHOIS(ctx context.Context, check CheckWHOIS) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointWHOISImpl) UpdateWHOIS(ctx context.Context, pk PrimaryKeyable, check CheckWHOIS) (*Check, error) {
	return c.Update(ctx, pk, check)
}

type CheckPageSpeed struct {
	Name          string               `json:"name,omitempty"`
	ContactGroups []string             `json:"contact_groups,omitempty"`
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

type CheckMaintenanceSchedule struct {
	Type          string `json:"type"`
	FromTime      string `json:"from_time,omitempty"`
	ToTime        string `json:"to_time,omitempty"`
	Monthday      int    `json:"monthday,omitempty"`
	MonthdayFrom  int    `json:"monthday_from,omitempty"`
	MonthdayTo    int    `json:"monthday_to,omitempty"`
	OnceStartDate string `json:"once_start_date,omitempty"`
	OnceEndDate   string `json:"once_end_date,omitempty"`
	Weekdays      []int  `json:"weekdays,omitempty"`
}

type CheckMaintenance struct {
	State    string                     `json:"state,omitempty"`
	Schedule []CheckMaintenanceSchedule `json:"schedule,omitempty"`
}

type checksEndpointMaintenanceImpl struct {
	EndpointUpdater[CheckMaintenance, CheckCreateUpdateResponse, Check]
}

func (c checksEndpointMaintenanceImpl) UpdateMaintenance(ctx context.Context, pk PrimaryKeyable, maintenance CheckMaintenance) (*Check, error) {
	return c.Update(ctx, pk, maintenance)
}
