package upapi

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/shopspring/decimal"
)

// Check represents a check in Uptime.com.
type Check struct {
	PK                     int              `json:"pk,omitempty"`
	URL                    string           `json:"url,omitempty"`
	StatsURL               string           `json:"stats_url,omitempty"`
	AlertsURL              string           `json:"alerts_url,omitempty"`
	Name                   string           `json:"name,omitempty"`
	CachedResponseTime     float32          `json:"cached_response_time,omitempty"`
	ContactGroups          []string         `json:"contact_groups"`
	CreatedAt              time.Time        `json:"created_at,omitempty"`
	ModifiedAt             time.Time        `json:"modified_at,omitempty"`
	Locations              []string         `json:"locations,omitempty"`
	Tags                   []string         `json:"tags,omitempty"`
	CheckType              string           `json:"check_type,omitempty"`
	Escalations            string           `json:"-,omitempty"` // TODO
	Maintenance            string           `json:"-,omitempty"` // TODO
	MonitoringServiceType  string           `json:"monitoring_service_type,omitempty"`
	IsPaused               bool             `json:"is_paused,omitempty"`
	IsUnderMaintenance     bool             `json:"is_under_maintenance,omitempty"`
	StateIsUp              bool             `json:"state_is_up,omitempty"`
	StateChangedAt         time.Time        `json:"state_changed_at,omitempty"`
	Protocol               string           `json:"msp_protocol,omitempty"`
	Interval               int              `json:"msp_interval,omitempty"`
	Address                string           `json:"msp_address,omitempty"`
	Port                   int              `json:"msp_port,omitempty"`
	Username               string           `json:"msp_username,omitempty"`
	Password               string           `json:"msp_password,omitempty"`
	Proxy                  string           `json:"msp_proxy,omitempty"`
	DNSServer              string           `json:"msp_dns_server,omitempty"`
	DNSRecordType          string           `json:"msp_dns_record_type,omitempty"`
	StatusCode             string           `json:"msp_status_code,omitempty"`
	SendString             string           `json:"msp_send_string,omitempty"`
	ExpectString           string           `json:"msp_expect_string,omitempty"`
	ExpectStringType       string           `json:"msp_expect_string_type,omitempty"`
	Encryption             string           `json:"msp_encryption,omitempty"`
	Threshold              int              `json:"msp_threshold,omitempty"`
	Headers                string           `json:"msp_headers,omitempty"`
	Script                 string           `json:"msp_script,omitempty"`
	Version                int              `json:"msp_version,omitempty"`
	Sensitivity            int              `json:"msp_sensitivity,omitempty"`
	NumRetries             int              `json:"msp_num_retries,omitempty"`
	UseIPVersion           string           `json:"msp_use_ip_version,omitempty"`
	UptimeSLA              *decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        *decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string           `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool             `json:"msp_include_in_global_metrics,omitempty"`
}

func (c Check) PrimaryKey() int {
	return c.PK
}

type CheckListResponse struct {
	Count   int     `json:"count,omitempty"`
	Results []Check `json:"results,omitempty"`
}

func (r CheckListResponse) List() []Check {
	return r.Results
}

// CheckListOptions specifies the optional parameters to the CheckService.List method.
type CheckListOptions struct {
	Page                  int      `url:"page,omitempty"`
	PageSize              int      `url:"page_size,omitempty"`
	Search                string   `url:"search,omitempty"`
	Ordering              string   `url:"ordering,omitempty"`
	MonitoringServiceType string   `url:"monitoring_service_type,omitempty"`
	IsPaused              bool     `url:"is_paused,omitempty"`
	StateIsUp             bool     `url:"state_is_up,omitempty"`
	Tag                   []string `url:"tag,omitempty"`
}

type CheckResponse struct {
	Messages map[string]interface{} `json:"messages,omitempty"`
	Results  Check                  `json:"results,omitempty"`
}

func (r CheckResponse) Item() *Check {
	return &r.Results
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
		Outages      int   `json:"outages,omitempty"`
		DowntimeSecs int64 `json:"downtime_secs,omitempty"`
	} `json:"totals"`
	Statistics []CheckStats `json:"statistics"`
}

func (c CheckStatsResponse) List() []CheckStats {
	return c.Statistics
}

type CheckStats struct {
	Date                   string   `json:"date"`
	Outages                int      `json:"outages"`
	DowntimeSecs           int      `json:"downtime_secs"`
	Uptime                 *float64 `json:"uptime,omitempty"`
	ResponseTime           *float64 `json:"response_time,omitempty"`
	ResponseTimeDatapoints [][]any  `json:"response_time_datapoints,omitempty"`
}

type ChecksEndpoint interface {
	List(context.Context, CheckListOptions) ([]Check, error)
	Get(context.Context, PrimaryKey) (*Check, error)
	Delete(context.Context, PrimaryKey) error
	Stats(context.Context, PrimaryKey, CheckStatsOptions) ([]CheckStats, error)

	CreateAPI(ctx context.Context, check CheckAPI) (*Check, error)
	UpdateAPI(ctx context.Context, check CheckAPI) (*Check, error)

	CreateBlacklist(ctx context.Context, check CheckBlacklist) (*Check, error)
	UpdateBlacklist(ctx context.Context, check CheckBlacklist) (*Check, error)

	CreateDNS(ctx context.Context, check CheckDNS) (*Check, error)
	UpdateDNS(ctx context.Context, check CheckDNS) (*Check, error)

	CreateGroup(ctx context.Context, check CheckGroup) (*Check, error)
	UpdateGroup(ctx context.Context, check CheckGroup) (*Check, error)

	CreateHeartbeat(ctx context.Context, check CheckHeartbeat) (*Check, error)
	UpdateHeartbeat(ctx context.Context, check CheckHeartbeat) (*Check, error)

	CreateHTTP(ctx context.Context, check CheckHTTP) (*Check, error)
	UpdateHTTP(ctx context.Context, check CheckHTTP) (*Check, error)

	CreateICMP(ctx context.Context, check CheckICMP) (*Check, error)
	UpdateICMP(ctx context.Context, check CheckICMP) (*Check, error)

	CreateIMAP(ctx context.Context, check CheckIMAP) (*Check, error)
	UpdateIMAP(ctx context.Context, check CheckIMAP) (*Check, error)

	CreateMalware(ctx context.Context, check CheckMalware) (*Check, error)
	UpdateMalware(ctx context.Context, check CheckMalware) (*Check, error)

	CreateNTP(ctx context.Context, check CheckNTP) (*Check, error)
	UpdateNTP(ctx context.Context, check CheckNTP) (*Check, error)

	CreatePOP(ctx context.Context, check CheckPOP) (*Check, error)
	UpdatePOP(ctx context.Context, check CheckPOP) (*Check, error)

	CreateRUM(ctx context.Context, check CheckRUM) (*Check, error)
	UpdateRUM(ctx context.Context, check CheckRUM) (*Check, error)

	CreateRUM2(ctx context.Context, check CheckRUM2) (*Check, error)
	UpdateRUM2(ctx context.Context, check CheckRUM2) (*Check, error)

	CreateSMTP(ctx context.Context, check CheckSMTP) (*Check, error)
	UpdateSMTP(ctx context.Context, check CheckSMTP) (*Check, error)

	CreateSSH(ctx context.Context, check CheckSSH) (*Check, error)
	UpdateSSH(ctx context.Context, check CheckSSH) (*Check, error)

	CreateSSLCert(ctx context.Context, check CheckSSLCert) (*Check, error)
	UpdateSSLCert(ctx context.Context, check CheckSSLCert) (*Check, error)

	CreateTCP(ctx context.Context, check CheckTCP) (*Check, error)
	UpdateTCP(ctx context.Context, check CheckTCP) (*Check, error)

	CreateTransaction(ctx context.Context, check CheckTransaction) (*Check, error)
	UpdateTransaction(ctx context.Context, check CheckTransaction) (*Check, error)

	CreateUDP(ctx context.Context, check CheckUDP) (*Check, error)
	UpdateUDP(ctx context.Context, check CheckUDP) (*Check, error)

	CreateWebhook(ctx context.Context, check CheckWebhook) (*Check, error)
	UpdateWebhook(ctx context.Context, check CheckWebhook) (*Check, error)

	CreateWHOIS(ctx context.Context, check CheckWHOIS) (*Check, error)
	UpdateWHOIS(ctx context.Context, check CheckWHOIS) (*Check, error)
}

func NewChecksEndpoint(cbd CBD) ChecksEndpoint {
	endpoint := "checks"
	return &checksEndpointImpl{
		checksEndpointAPIImpl: checksEndpointAPIImpl{
			EndpointCreator: NewEndpointCreator[CheckAPI, CheckResponse, Check](cbd, endpoint+"/add-api"),
			EndpointUpdater: NewEndpointUpdater[CheckAPI, CheckResponse, Check](cbd, endpoint),
		},
		checksEndpointBlacklistImpl: checksEndpointBlacklistImpl{
			EndpointCreator: NewEndpointCreator[CheckBlacklist, CheckResponse, Check](cbd, endpoint+"/add-blacklist"),
			EndpointUpdater: NewEndpointUpdater[CheckBlacklist, CheckResponse, Check](cbd, endpoint),
		},
		checksEndpointDNSImpl: checksEndpointDNSImpl{
			EndpointCreator: NewEndpointCreator[CheckDNS, CheckResponse, Check](cbd, endpoint+"/add-dns"),
			EndpointUpdater: NewEndpointUpdater[CheckDNS, CheckResponse, Check](cbd, endpoint),
		},
		checksEndpointGroupImpl: checksEndpointGroupImpl{
			EndpointCreator: NewEndpointCreator[CheckGroup, CheckResponse, Check](cbd, endpoint+"/add-group"),
			EndpointUpdater: NewEndpointUpdater[CheckGroup, CheckResponse, Check](cbd, endpoint),
		},
		checksEndpointHeartbeatImpl: checksEndpointHeartbeatImpl{
			EndpointCreator: NewEndpointCreator[CheckHeartbeat, CheckResponse, Check](cbd, endpoint+"/add-heartbeat"),
			EndpointUpdater: NewEndpointUpdater[CheckHeartbeat, CheckResponse, Check](cbd, endpoint),
		},
		checksEndpointHTTPImpl: checksEndpointHTTPImpl{
			EndpointCreator: NewEndpointCreator[CheckHTTP, CheckResponse, Check](cbd, endpoint+"/add-http"),
			EndpointUpdater: NewEndpointUpdater[CheckHTTP, CheckResponse, Check](cbd, endpoint),
		},
		checksEndpointICMPImpl: checksEndpointICMPImpl{
			EndpointCreator: NewEndpointCreator[CheckICMP, CheckResponse, Check](cbd, endpoint+"/add-icmp"),
			EndpointUpdater: NewEndpointUpdater[CheckICMP, CheckResponse, Check](cbd, endpoint),
		},
		checksEndpointIMAPImpl: checksEndpointIMAPImpl{
			EndpointCreator: NewEndpointCreator[CheckIMAP, CheckResponse, Check](cbd, endpoint+"/add-imap"),
			EndpointUpdater: NewEndpointUpdater[CheckIMAP, CheckResponse, Check](cbd, endpoint),
		},
		checksEndpointMalwareImpl: checksEndpointMalwareImpl{
			EndpointCreator: NewEndpointCreator[CheckMalware, CheckResponse, Check](cbd, endpoint+"/add-malware"),
			EndpointUpdater: NewEndpointUpdater[CheckMalware, CheckResponse, Check](cbd, endpoint),
		},
		checksEndpointNTPImpl: checksEndpointNTPImpl{
			EndpointCreator: NewEndpointCreator[CheckNTP, CheckResponse, Check](cbd, endpoint+"/add-ntp"),
			EndpointUpdater: NewEndpointUpdater[CheckNTP, CheckResponse, Check](cbd, endpoint),
		},
		checksEndpointPOPImpl: checksEndpointPOPImpl{
			EndpointCreator: NewEndpointCreator[CheckPOP, CheckResponse, Check](cbd, endpoint+"/add-pop"),
			EndpointUpdater: NewEndpointUpdater[CheckPOP, CheckResponse, Check](cbd, endpoint),
		},
		checksEndpointRUMImpl: checksEndpointRUMImpl{
			EndpointCreator: NewEndpointCreator[CheckRUM, CheckResponse, Check](cbd, endpoint+"/add-rum"),
			EndpointUpdater: NewEndpointUpdater[CheckRUM, CheckResponse, Check](cbd, endpoint),
		},
		checksEndpointRUM2Impl: checksEndpointRUM2Impl{
			EndpointCreator: NewEndpointCreator[CheckRUM2, CheckResponse, Check](cbd, endpoint+"/add-rum2"),
			EndpointUpdater: NewEndpointUpdater[CheckRUM2, CheckResponse, Check](cbd, endpoint),
		},
		checksEndpointSMTPImpl: checksEndpointSMTPImpl{
			EndpointCreator: NewEndpointCreator[CheckSMTP, CheckResponse, Check](cbd, endpoint+"/add-smtp"),
			EndpointUpdater: NewEndpointUpdater[CheckSMTP, CheckResponse, Check](cbd, endpoint),
		},
		checksEndpointSSHImpl: checksEndpointSSHImpl{
			EndpointCreator: NewEndpointCreator[CheckSSH, CheckResponse, Check](cbd, endpoint+"/add-ssh"),
			EndpointUpdater: NewEndpointUpdater[CheckSSH, CheckResponse, Check](cbd, endpoint),
		},
		checksEndpointSSLCertImpl: checksEndpointSSLCertImpl{
			EndpointCreator: NewEndpointCreator[CheckSSLCert, CheckResponse, Check](cbd, endpoint+"/add-ssl-cert"),
			EndpointUpdater: NewEndpointUpdater[CheckSSLCert, CheckResponse, Check](cbd, endpoint),
		},
		checksEndpointTCPImpl: checksEndpointTCPImpl{
			EndpointCreator: NewEndpointCreator[CheckTCP, CheckResponse, Check](cbd, endpoint+"/add-tcp"),
			EndpointUpdater: NewEndpointUpdater[CheckTCP, CheckResponse, Check](cbd, endpoint),
		},
		checksEndpointUDPImpl: checksEndpointUDPImpl{
			EndpointCreator: NewEndpointCreator[CheckUDP, CheckResponse, Check](cbd, endpoint+"/add-udp"),
			EndpointUpdater: NewEndpointUpdater[CheckUDP, CheckResponse, Check](cbd, endpoint),
		},
		checksEndpointWebhookImpl: checksEndpointWebhookImpl{
			EndpointCreator: NewEndpointCreator[CheckWebhook, CheckResponse, Check](cbd, endpoint+"/add-webhook"),
			EndpointUpdater: NewEndpointUpdater[CheckWebhook, CheckResponse, Check](cbd, endpoint),
		},
		checksEndpointWHOISImpl: checksEndpointWHOISImpl{
			EndpointCreator: NewEndpointCreator[CheckWHOIS, CheckResponse, Check](cbd, endpoint+"/add-whois"),
			EndpointUpdater: NewEndpointUpdater[CheckWHOIS, CheckResponse, Check](cbd, endpoint),
		},
		checksStatsEndpointImpl: checksStatsEndpointImpl{
			endpoint: NewEndpointLister[CheckStatsResponse, CheckStats, CheckStatsOptions](&checksStatsEndpointCBD{cbd}, endpoint+"/%d/stats"),
		},
		EndpointLister:  NewEndpointLister[CheckListResponse, Check, CheckListOptions](cbd, endpoint),
		EndpointGetter:  NewEndpointGetter[PrimaryKey, CheckResponse, Check](cbd, endpoint),
		EndpointDeleter: NewEndpointDeleter[PrimaryKey](cbd, endpoint),
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
	EndpointLister[CheckListResponse, Check, CheckListOptions]
	EndpointGetter[PrimaryKey, CheckResponse, Check]
	EndpointUpdater[Check, CheckResponse, Check]
	EndpointDeleter[PrimaryKey]
}

type checksStatsCtxKey struct{}

type checksStatsEndpointImpl struct {
	endpoint EndpointLister[CheckStatsResponse, CheckStats, CheckStatsOptions]
}

func (c *checksEndpointImpl) Stats(ctx context.Context, pk PrimaryKey, opts CheckStatsOptions) ([]CheckStats, error) {
	ctx = context.WithValue(ctx, checksStatsCtxKey{}, pk)
	return c.endpoint.List(ctx, opts)
}

type checksStatsEndpointCBD struct {
	CBD
}

func (c checksStatsEndpointCBD) BuildRequest(ctx context.Context, method string, endpoint string, args any, data any) (*http.Request, error) {
	pk := ctx.Value(checksStatsCtxKey{}).(PrimaryKey)
	endpoint = fmt.Sprintf(endpoint, pk.PrimaryKey())
	return c.CBD.BuildRequest(ctx, method, endpoint, args, data)
}

type CheckAPI struct {
	PK                     int              `json:"pk,omitempty"`
	Name                   string           `json:"name,omitempty"`
	ContactGroups          []string         `json:"contact_groups,omitempty"`
	Locations              []string         `json:"locations,omitempty"`
	Tags                   []string         `json:"tags,omitempty"`
	IsPaused               bool             `json:"is_paused,omitempty"`
	Interval               int              `json:"msp_interval,omitempty"`
	Threshold              int              `json:"msp_threshold,omitempty"`
	Script                 string           `json:"msp_script,omitempty"`
	Sensitivity            int              `json:"msp_sensitivity,omitempty"`
	NumRetries             int              `json:"msp_num_retries,omitempty"`
	UseIPVersion           string           `json:"msp_use_ip_version,omitempty"`
	UptimeSLA              *decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        *decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string           `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool             `json:"msp_include_in_global_metrics,omitempty"`
}

func (c CheckAPI) PrimaryKey() int {
	return c.PK
}

type checksEndpointAPIImpl struct {
	EndpointCreator[CheckAPI, CheckResponse, Check]
	EndpointUpdater[CheckAPI, CheckResponse, Check]
}

func (c checksEndpointAPIImpl) CreateAPI(ctx context.Context, check CheckAPI) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointAPIImpl) UpdateAPI(ctx context.Context, check CheckAPI) (*Check, error) {
	return c.Update(ctx, check)
}

type CheckBlacklist struct {
	PK                    int              `json:"pk,omitempty"`
	Name                  string           `json:"name,omitempty"`
	ContactGroups         []string         `json:"contact_groups,omitempty"`
	Locations             []string         `json:"locations,omitempty"`
	Tags                  []string         `json:"tags,omitempty"`
	MonitoringServiceType string           `json:"monitoring_service_type,omitempty"`
	IsPaused              bool             `json:"is_paused,omitempty"`
	Address               string           `json:"msp_address,omitempty"`
	NumRetries            int              `json:"msp_num_retries,omitempty"`
	UptimeSLA             *decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	Notes                 string           `json:"msp_notes,omitempty"`
}

func (c CheckBlacklist) PrimaryKey() int {
	return c.PK
}

type checksEndpointBlacklistImpl struct {
	EndpointCreator[CheckBlacklist, CheckResponse, Check]
	EndpointUpdater[CheckBlacklist, CheckResponse, Check]
}

func (c checksEndpointBlacklistImpl) CreateBlacklist(ctx context.Context, check CheckBlacklist) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointBlacklistImpl) UpdateBlacklist(ctx context.Context, check CheckBlacklist) (*Check, error) {
	return c.Update(ctx, check)
}

type CheckDNS struct {
	PK                     int              `json:"pk,omitempty"`
	Name                   string           `json:"name,omitempty"`
	ContactGroups          []string         `json:"contact_groups,omitempty"`
	Locations              []string         `json:"locations,omitempty"`
	Tags                   []string         `json:"tags,omitempty"`
	IsPaused               bool             `json:"is_paused,omitempty"`
	Interval               int              `json:"msp_interval,omitempty"`
	Address                string           `json:"msp_address,omitempty"`
	DnsServer              string           `json:"msp_dns_server,omitempty"`
	DnsRecordType          string           `json:"msp_dns_record_type,omitempty"`
	ExpectString           string           `json:"msp_expect_string,omitempty"`
	Threshold              int              `json:"msp_threshold,omitempty"`
	Sensitivity            int              `json:"msp_sensitivity,omitempty"`
	NumRetries             int              `json:"msp_num_retries,omitempty"`
	UptimeSLA              *decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        *decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string           `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool             `json:"msp_include_in_global_metrics,omitempty"`
}

func (c CheckDNS) PrimaryKey() int {
	return c.PK
}

type checksEndpointDNSImpl struct {
	EndpointCreator[CheckDNS, CheckResponse, Check]
	EndpointUpdater[CheckDNS, CheckResponse, Check]
}

func (c checksEndpointDNSImpl) CreateDNS(ctx context.Context, check CheckDNS) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointDNSImpl) UpdateDNS(ctx context.Context, check CheckDNS) (*Check, error) {
	return c.Update(ctx, check)
}

type CheckGroup struct {
	PK                     int              `json:"pk,omitempty"`
	Name                   string           `json:"name,omitempty"`
	ContactGroups          []string         `json:"contact_groups,omitempty"`
	Locations              []string         `json:"locations,omitempty"`
	Tags                   []string         `json:"tags,omitempty"`
	IsPaused               bool             `json:"is_paused,omitempty"`
	UptimeSLA              *decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        *decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string           `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool             `json:"msp_include_in_global_metrics,omitempty"`
	Config                 struct {
		CheckServices               []string `json:"group_check_services,omitempty"`
		CheckTags                   []string `json:"group_check_tags,omitempty"`
		CheckDownCondition          string   `json:"group_check_down_condition,omitempty"`
		UptimePercentCalculation    string   `json:"group_uptime_percent_calculation,omitempty"`
		ResponseTimeCalculationMode string   `json:"group_response_time_calculation_mode,omitempty"`
		ResponseTimeCheckType       string   `json:"group_response_time_check_type,omitempty"`
		ResponseTimeSingleCheck     string   `json:"group_response_time_single_check,omitempty"`
	} `json:"groupcheckconfig,omitempty"`
}

func (c CheckGroup) PrimaryKey() int {
	return c.PK
}

type checksEndpointGroupImpl struct {
	EndpointCreator[CheckGroup, CheckResponse, Check]
	EndpointUpdater[CheckGroup, CheckResponse, Check]
}

func (c checksEndpointGroupImpl) CreateGroup(ctx context.Context, check CheckGroup) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointGroupImpl) UpdateGroup(ctx context.Context, check CheckGroup) (*Check, error) {
	return c.Update(ctx, check)
}

type CheckHeartbeat struct {
	PK                     int              `json:"pk,omitempty"`
	Name                   string           `json:"name,omitempty"`
	ContactGroups          []string         `json:"contact_groups,omitempty"`
	Locations              []string         `json:"locations,omitempty"`
	Tags                   []string         `json:"tags,omitempty"`
	CheckType              string           `json:"check_type,omitempty"`
	IsPaused               bool             `json:"is_paused,omitempty"`
	Interval               int              `json:"msp_interval,omitempty"`
	UptimeSLA              *decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        *decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string           `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool             `json:"msp_include_in_global_metrics,omitempty"`
	HeartbeatURL           string           `json:"heartbeat_url,omitempty"`
}

func (c CheckHeartbeat) PrimaryKey() int {
	return c.PK
}

type checksEndpointHeartbeatImpl struct {
	EndpointCreator[CheckHeartbeat, CheckResponse, Check]
	EndpointUpdater[CheckHeartbeat, CheckResponse, Check]
}

func (c checksEndpointHeartbeatImpl) CreateHeartbeat(ctx context.Context, check CheckHeartbeat) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointHeartbeatImpl) UpdateHeartbeat(ctx context.Context, check CheckHeartbeat) (*Check, error) {
	return c.Update(ctx, check)
}

type CheckHTTP struct {
	PK                     int              `json:"pk,omitempty"`
	Name                   string           `json:"name,omitempty"`
	ContactGroups          []string         `json:"contact_groups,omitempty"`
	Locations              []string         `json:"locations,omitempty"`
	Tags                   []string         `json:"tags,omitempty"`
	IsPaused               bool             `json:"is_paused,omitempty"`
	Interval               int              `json:"msp_interval,omitempty"`
	Address                string           `json:"msp_address,omitempty"`
	Port                   int              `json:"msp_port,omitempty"`
	Username               string           `json:"msp_username,omitempty"`
	Password               string           `json:"msp_password,omitempty"`
	Proxy                  string           `json:"msp_proxy,omitempty"`
	StatusCode             string           `json:"msp_status_code,omitempty"`
	SendString             string           `json:"msp_send_string,omitempty"`
	ExpectString           string           `json:"msp_expect_string,omitempty"`
	ExpectStringType       string           `json:"msp_expect_string_type,omitempty"`
	Encryption             string           `json:"msp_encryption,omitempty"`
	Threshold              int              `json:"msp_threshold,omitempty"`
	Headers                string           `json:"msp_headers,omitempty"`
	Version                int              `json:"msp_version,omitempty"`
	Sensitivity            int              `json:"msp_sensitivity,omitempty"`
	NumRetries             int              `json:"msp_num_retries,omitempty"`
	UseIPVersion           string           `json:"msp_use_ip_version,omitempty"`
	UptimeSLA              *decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        *decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string           `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool             `json:"msp_include_in_global_metrics,omitempty"`
}

func (c CheckHTTP) PrimaryKey() int {
	return c.PK
}

type checksEndpointHTTPImpl struct {
	EndpointCreator[CheckHTTP, CheckResponse, Check]
	EndpointUpdater[CheckHTTP, CheckResponse, Check]
}

func (c checksEndpointHTTPImpl) CreateHTTP(ctx context.Context, check CheckHTTP) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointHTTPImpl) UpdateHTTP(ctx context.Context, check CheckHTTP) (*Check, error) {
	return c.Update(ctx, check)
}

type CheckICMP struct {
	PK                     int              `json:"pk,omitempty"`
	Name                   string           `json:"name,omitempty"`
	ContactGroups          []string         `json:"contact_groups,omitempty"`
	Locations              []string         `json:"locations,omitempty"`
	Tags                   []string         `json:"tags,omitempty"`
	IsPaused               bool             `json:"is_paused,omitempty"`
	Interval               int              `json:"msp_interval,omitempty"`
	Address                string           `json:"msp_address,omitempty"`
	Sensitivity            int              `json:"msp_sensitivity,omitempty"`
	NumRetries             int              `json:"msp_num_retries,omitempty"`
	UseIPVersion           string           `json:"msp_use_ip_version,omitempty"`
	UptimeSLA              *decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        *decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string           `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool             `json:"msp_include_in_global_metrics,omitempty"`
}

func (c CheckICMP) PrimaryKey() int {
	return c.PK
}

type checksEndpointICMPImpl struct {
	EndpointCreator[CheckICMP, CheckResponse, Check]
	EndpointUpdater[CheckICMP, CheckResponse, Check]
}

func (c checksEndpointICMPImpl) CreateICMP(ctx context.Context, check CheckICMP) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointICMPImpl) UpdateICMP(ctx context.Context, check CheckICMP) (*Check, error) {
	return c.Update(ctx, check)
}

type CheckIMAP struct {
	PK                     int              `json:"pk,omitempty"`
	Name                   string           `json:"name,omitempty"`
	ContactGroups          []string         `json:"contact_groups,omitempty"`
	Locations              []string         `json:"locations,omitempty"`
	Tags                   []string         `json:"tags,omitempty"`
	IsPaused               bool             `json:"is_paused,omitempty"`
	Interval               int              `json:"msp_interval,omitempty"`
	Address                string           `json:"msp_address,omitempty"`
	Port                   int              `json:"msp_port,omitempty"`
	ExpectString           string           `json:"msp_expect_string,omitempty"`
	Encryption             string           `json:"msp_encryption,omitempty"`
	Sensitivity            int              `json:"msp_sensitivity,omitempty"`
	NumRetries             int              `json:"msp_num_retries,omitempty"`
	UseIPVersion           string           `json:"msp_use_ip_version,omitempty"`
	UptimeSLA              *decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        *decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string           `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool             `json:"msp_include_in_global_metrics,omitempty"`
}

func (c CheckIMAP) PrimaryKey() int {
	return c.PK
}

type checksEndpointIMAPImpl struct {
	EndpointCreator[CheckIMAP, CheckResponse, Check]
	EndpointUpdater[CheckIMAP, CheckResponse, Check]
}

func (c checksEndpointIMAPImpl) CreateIMAP(ctx context.Context, check CheckIMAP) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointIMAPImpl) UpdateIMAP(ctx context.Context, check CheckIMAP) (*Check, error) {
	return c.Update(ctx, check)
}

type CheckMalware struct {
	PK            int              `json:"pk,omitempty"`
	Name          string           `json:"name,omitempty"`
	ContactGroups []string         `json:"contact_groups,omitempty"`
	Locations     []string         `json:"locations,omitempty"`
	Tags          []string         `json:"tags,omitempty"`
	CheckType     string           `json:"check_type,omitempty"`
	IsPaused      bool             `json:"is_paused,omitempty"`
	Address       string           `json:"msp_address,omitempty"`
	NumRetries    int              `json:"msp_num_retries,omitempty"`
	UptimeSLA     *decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	Notes         string           `json:"msp_notes,omitempty"`
}

func (c CheckMalware) PrimaryKey() int {
	return c.PK
}

type checksEndpointMalwareImpl struct {
	EndpointCreator[CheckMalware, CheckResponse, Check]
	EndpointUpdater[CheckMalware, CheckResponse, Check]
}

func (c checksEndpointMalwareImpl) CreateMalware(ctx context.Context, check CheckMalware) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointMalwareImpl) UpdateMalware(ctx context.Context, check CheckMalware) (*Check, error) {
	return c.Update(ctx, check)
}

type CheckNTP struct {
	PK                     int              `json:"pk,omitempty"`
	Name                   string           `json:"name,omitempty"`
	ContactGroups          []string         `json:"contact_groups,omitempty"`
	Locations              []string         `json:"locations,omitempty"`
	Tags                   []string         `json:"tags,omitempty"`
	IsPaused               bool             `json:"is_paused,omitempty"`
	Interval               int              `json:"msp_interval"`
	Address                string           `json:"msp_address,omitempty"`
	Port                   int              `json:"msp_port,omitempty"`
	Threshold              int              `json:"msp_threshold,omitempty"`
	Sensitivity            int              `json:"msp_sensitivity,omitempty"`
	NumRetries             int              `json:"msp_num_retries,omitempty"`
	UseIPVersion           string           `json:"msp_use_ip_version,omitempty"`
	UptimeSLA              *decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        *decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string           `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool             `json:"msp_include_in_global_metrics,omitempty"`
}

func (c CheckNTP) PrimaryKey() int {
	return c.PK
}

type checksEndpointNTPImpl struct {
	EndpointCreator[CheckNTP, CheckResponse, Check]
	EndpointUpdater[CheckNTP, CheckResponse, Check]
}

func (c checksEndpointNTPImpl) CreateNTP(ctx context.Context, check CheckNTP) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointNTPImpl) UpdateNTP(ctx context.Context, check CheckNTP) (*Check, error) {
	return c.Update(ctx, check)
}

type CheckPOP struct {
	PK                     int              `json:"pk,omitempty"`
	Name                   string           `json:"name,omitempty"`
	ContactGroups          []string         `json:"contact_groups,omitempty"`
	Locations              []string         `json:"locations,omitempty"`
	Tags                   []string         `json:"tags,omitempty"`
	IsPaused               bool             `json:"is_paused,omitempty"`
	Interval               int              `json:"msp_interval,omitempty"`
	Address                string           `json:"msp_address,omitempty"`
	Port                   int              `json:"msp_port,omitempty"`
	ExpectString           string           `json:"msp_expect_string,omitempty"`
	Encryption             string           `json:"msp_encryption,omitempty"`
	Sensitivity            int              `json:"msp_sensitivity,omitempty"`
	NumRetries             int              `json:"msp_num_retries,omitempty"`
	UseIPVersion           string           `json:"msp_use_ip_version,omitempty"`
	UptimeSLA              *decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        *decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string           `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool             `json:"msp_include_in_global_metrics,omitempty"`
}

func (c CheckPOP) PrimaryKey() int {
	return c.PK
}

type checksEndpointPOPImpl struct {
	EndpointCreator[CheckPOP, CheckResponse, Check]
	EndpointUpdater[CheckPOP, CheckResponse, Check]
}

func (c checksEndpointPOPImpl) CreatePOP(ctx context.Context, check CheckPOP) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointPOPImpl) UpdatePOP(ctx context.Context, check CheckPOP) (*Check, error) {
	return c.Update(ctx, check)
}

type CheckRUM struct {
	PK                     int              `json:"pk,omitempty"`
	Name                   string           `json:"name,omitempty"`
	ContactGroups          []string         `json:"contact_groups,omitempty"`
	Locations              []string         `json:"locations,omitempty"`
	Tags                   []string         `json:"tags,omitempty"`
	IsPaused               bool             `json:"is_paused,omitempty"`
	Address                string           `json:"msp_address,omitempty"`
	Threshold              int              `json:"msp_threshold,omitempty"`
	UptimeSLA              *decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	Notes                  string           `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool             `json:"msp_include_in_global_metrics,omitempty"`
}

func (c CheckRUM) PrimaryKey() int {
	return c.PK
}

type checksEndpointRUMImpl struct {
	EndpointCreator[CheckRUM, CheckResponse, Check]
	EndpointUpdater[CheckRUM, CheckResponse, Check]
}

func (c checksEndpointRUMImpl) CreateRUM(ctx context.Context, check CheckRUM) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointRUMImpl) UpdateRUM(ctx context.Context, check CheckRUM) (*Check, error) {
	return c.Update(ctx, check)
}

type CheckRUM2 struct {
	PK                     int              `json:"pk,omitempty"`
	Name                   string           `json:"name,omitempty"`
	ContactGroups          []string         `json:"contact_groups,omitempty"`
	Locations              []string         `json:"locations,omitempty"`
	Tags                   []string         `json:"tags,omitempty"`
	IsPaused               bool             `json:"is_paused,omitempty"`
	Address                string           `json:"msp_address,omitempty"`
	UptimeSLA              *decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	Notes                  string           `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool             `json:"msp_include_in_global_metrics,omitempty"`
}

func (c CheckRUM2) PrimaryKey() int {
	return c.PK
}

type checksEndpointRUM2Impl struct {
	EndpointCreator[CheckRUM2, CheckResponse, Check]
	EndpointUpdater[CheckRUM2, CheckResponse, Check]
}

func (c checksEndpointRUM2Impl) CreateRUM2(ctx context.Context, check CheckRUM2) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointRUM2Impl) UpdateRUM2(ctx context.Context, check CheckRUM2) (*Check, error) {
	return c.Update(ctx, check)
}

type CheckSMTP struct {
	PK                     int              `json:"pk,omitempty"`
	Name                   string           `json:"name,omitempty"`
	ContactGroups          []string         `json:"contact_groups,omitempty"`
	Locations              []string         `json:"locations,omitempty"`
	Tags                   []string         `json:"tags,omitempty"`
	IsPaused               bool             `json:"is_paused,omitempty"`
	Interval               int              `json:"msp_interval"`
	Address                string           `json:"msp_address"`
	Port                   int              `json:"msp_port,omitempty"`
	Username               string           `json:"msp_username,omitempty"`
	Password               string           `json:"msp_password,omitempty"`
	ExpectString           string           `json:"msp_expect_string,omitempty"`
	Encryption             string           `json:"msp_encryption,omitempty"`
	Sensitivity            int              `json:"msp_sensitivity,omitempty"`
	NumRetries             int              `json:"msp_num_retries,omitempty"`
	UseIpVersion           string           `json:"msp_use_ip_version,omitempty"`
	UptimeSLA              *decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        *decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string           `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool             `json:"msp_include_in_global_metrics,omitempty"`
}

func (c CheckSMTP) PrimaryKey() int {
	return c.PK
}

type checksEndpointSMTPImpl struct {
	EndpointCreator[CheckSMTP, CheckResponse, Check]
	EndpointUpdater[CheckSMTP, CheckResponse, Check]
}

func (c checksEndpointSMTPImpl) CreateSMTP(ctx context.Context, check CheckSMTP) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointSMTPImpl) UpdateSMTP(ctx context.Context, check CheckSMTP) (*Check, error) {
	return c.Update(ctx, check)
}

type CheckSSH struct {
	PK                     int              `json:"pk,omitempty"`
	Name                   string           `json:"name,omitempty"`
	ContactGroups          []string         `json:"contact_groups,omitempty"`
	Locations              []string         `json:"locations,omitempty"`
	Tags                   []string         `json:"tags,omitempty"`
	IsPaused               bool             `json:"is_paused,omitempty"`
	Interval               int              `json:"msp_interval,omitempty"`
	Address                string           `json:"msp_address,omitempty"`
	Port                   int              `json:"msp_port,omitempty"`
	Sensitivity            int              `json:"msp_sensitivity,omitempty"`
	NumRetries             int              `json:"msp_num_retries,omitempty"`
	UseIpVersion           string           `json:"msp_use_ip_version,omitempty"`
	UptimeSLA              *decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        *decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string           `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool             `json:"msp_include_in_global_metrics,omitempty"`
}

func (c CheckSSH) PrimaryKey() int {
	return c.PK
}

type checksEndpointSSHImpl struct {
	EndpointCreator[CheckSSH, CheckResponse, Check]
	EndpointUpdater[CheckSSH, CheckResponse, Check]
}

func (c checksEndpointSSHImpl) CreateSSH(ctx context.Context, check CheckSSH) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointSSHImpl) UpdateSSH(ctx context.Context, check CheckSSH) (*Check, error) {
	return c.Update(ctx, check)
}

type CheckSSLCert struct {
	PK            int              `json:"pk,omitempty"`
	Name          string           `json:"name,omitempty"`
	ContactGroups []string         `json:"contact_groups,omitempty"`
	Locations     []string         `json:"locations,omitempty"`
	Tags          []string         `json:"tags,omitempty"`
	IsPaused      bool             `json:"is_paused,omitempty"`
	Protocol      string           `json:"msp_protocol,omitempty"`
	Address       string           `json:"msp_address"`
	Port          int              `json:"msp_port,omitempty"`
	Threshold     int              `json:"msp_threshold"`
	NumRetries    int              `json:"msp_num_retries,omitempty"`
	UptimeSLA     *decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	Notes         string           `json:"msp_notes,omitempty"`
	Config        struct {
		SSLCertProtocol         string `json:"ssl_cert_protocol,omitempty"`
		SSLCertCRL              bool   `json:"ssl_cert_crl,omitempty"`
		SSLCertFirstElementOnly bool   `json:"ssl_cert_first_element_only,omitempty"`
		SSLCertMatch            string `json:"ssl_cert_match,omitempty"`
		SSLCertIssuer           string `json:"ssl_cert_issuer,omitempty"`
		SSLCertMinVersion       string `json:"ssl_cert_minimum_ssl_tls_version,omitempty"`
		SSLCertFingerprint      string `json:"ssl_cert_fingerprint,omitempty"`
		SSLCertSelfSigned       bool   `json:"ssl_cert_selfsigned,omitempty"`
		SSLCertFile             string `json:"ssl_cert_file,omitempty"`
	} `json:"sslconfig,omitempty"`
}

func (c CheckSSLCert) PrimaryKey() int {
	return c.PK
}

type checksEndpointSSLCertImpl struct {
	EndpointCreator[CheckSSLCert, CheckResponse, Check]
	EndpointUpdater[CheckSSLCert, CheckResponse, Check]
}

func (c checksEndpointSSLCertImpl) CreateSSLCert(ctx context.Context, check CheckSSLCert) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointSSLCertImpl) UpdateSSLCert(ctx context.Context, check CheckSSLCert) (*Check, error) {
	return c.Update(ctx, check)
}

type CheckTCP struct {
	PK                     int              `json:"pk,omitempty"`
	Name                   string           `json:"name,omitempty"`
	ContactGroups          []string         `json:"contact_groups,omitempty"`
	Locations              []string         `json:"locations,omitempty"`
	Tags                   []string         `json:"tags,omitempty"`
	IsPaused               bool             `json:"is_paused,omitempty"`
	Interval               int              `json:"msp_interval,omitempty"`
	Address                string           `json:"msp_address,omitempty"`
	Port                   int              `json:"msp_port,omitempty"`
	SendString             string           `json:"msp_send_string,omitempty"`
	ExpectString           string           `json:"msp_expect_string,omitempty"`
	Sensitivity            int              `json:"msp_sensitivity,omitempty"`
	NumRetries             int              `json:"msp_num_retries,omitempty"`
	UseIpVersion           string           `json:"msp_use_ip_version,omitempty"`
	UptimeSLA              *decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        *decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string           `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool             `json:"msp_include_in_global_metrics,omitempty"`
}

func (c CheckTCP) PrimaryKey() int {
	return c.PK
}

type checksEndpointTCPImpl struct {
	EndpointCreator[CheckTCP, CheckResponse, Check]
	EndpointUpdater[CheckTCP, CheckResponse, Check]
}

func (c checksEndpointTCPImpl) CreateTCP(ctx context.Context, check CheckTCP) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointTCPImpl) UpdateTCP(ctx context.Context, check CheckTCP) (*Check, error) {
	return c.Update(ctx, check)
}

type CheckTransaction struct {
	PK                     int              `json:"pk,omitempty"`
	Name                   string           `json:"name,omitempty"`
	ContactGroups          []string         `json:"contact_groups,omitempty"`
	Locations              []string         `json:"locations,omitempty"`
	Tags                   []string         `json:"tags,omitempty"`
	IsPaused               bool             `json:"is_paused,omitempty"`
	Interval               int              `json:"msp_interval,omitempty"`
	Threshold              int              `json:"msp_threshold,omitempty"`
	Script                 string           `json:"msp_script,omitempty"`
	Sensitivity            int              `json:"msp_sensitivity,omitempty"`
	NumRetries             int              `json:"msp_num_retries,omitempty"`
	UptimeSLA              *decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        *decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string           `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool             `json:"msp_include_in_global_metrics,omitempty"`
}

func (c CheckTransaction) PrimaryKey() int {
	return c.PK
}

type checksEndpointTransactionImpl struct {
	EndpointCreator[CheckTransaction, CheckResponse, Check]
	EndpointUpdater[CheckTransaction, CheckResponse, Check]
}

func (c checksEndpointTransactionImpl) CreateTransaction(ctx context.Context, check CheckTransaction) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointTransactionImpl) UpdateTransaction(ctx context.Context, check CheckTransaction) (*Check, error) {
	return c.Update(ctx, check)
}

type CheckUDP struct {
	PK                     int              `json:"pk,omitempty"`
	Name                   string           `json:"name,omitempty"`
	ContactGroups          []string         `json:"contact_groups,omitempty"`
	Locations              []string         `json:"locations,omitempty"`
	Tags                   []string         `json:"tags,omitempty"`
	IsPaused               bool             `json:"is_paused,omitempty"`
	Interval               int              `json:"msp_interval,omitempty"`
	Address                string           `json:"msp_address,omitempty"`
	Port                   int              `json:"msp_port,omitempty"`
	SendString             string           `json:"msp_send_string,omitempty"`
	ExpectString           string           `json:"msp_expect_string,omitempty"`
	Sensitivity            int              `json:"msp_sensitivity,omitempty"`
	NumRetries             int              `json:"msp_num_retries,omitempty"`
	UseIpVersion           string           `json:"msp_use_ip_version,omitempty"`
	UptimeSLA              *decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        *decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string           `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool             `json:"msp_include_in_global_metrics,omitempty"`
}

func (c CheckUDP) PrimaryKey() int {
	return c.PK
}

type checksEndpointUDPImpl struct {
	EndpointCreator[CheckUDP, CheckResponse, Check]
	EndpointUpdater[CheckUDP, CheckResponse, Check]
}

func (c checksEndpointUDPImpl) CreateUDP(ctx context.Context, check CheckUDP) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointUDPImpl) UpdateUDP(ctx context.Context, check CheckUDP) (*Check, error) {
	return c.Update(ctx, check)
}

type CheckWebhook struct {
	PK                     int              `json:"pk,omitempty"`
	Name                   string           `json:"name,omitempty"`
	ContactGroups          []string         `json:"contact_groups,omitempty"`
	Locations              []string         `json:"locations,omitempty"`
	Tags                   []string         `json:"tags,omitempty"`
	IsPaused               bool             `json:"is_paused,omitempty"`
	UptimeSLA              *decimal.Decimal `json:"msp_uptime_sla,omitempty"`
	ResponseTimeSLA        *decimal.Decimal `json:"msp_response_time_sla,omitempty"`
	Notes                  string           `json:"msp_notes,omitempty"`
	IncludeInGlobalMetrics bool             `json:"msp_include_in_global_metrics,omitempty"`
	WebhookUrl             string           `json:"webhook_url,omitempty"`
}

func (c CheckWebhook) PrimaryKey() int {
	return c.PK
}

type checksEndpointWebhookImpl struct {
	EndpointCreator[CheckWebhook, CheckResponse, Check]
	EndpointUpdater[CheckWebhook, CheckResponse, Check]
}

func (c checksEndpointWebhookImpl) CreateWebhook(ctx context.Context, check CheckWebhook) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointWebhookImpl) UpdateWebhook(ctx context.Context, check CheckWebhook) (*Check, error) {
	return c.Update(ctx, check)
}

type CheckWHOIS struct {
	PK            int      `json:"pk,omitempty"`
	Name          string   `json:"name,omitempty"`
	ContactGroups []string `json:"contact_groups,omitempty"`
	Locations     []string `json:"locations,omitempty"`
	Tags          []string `json:"tags,omitempty"`
	IsPaused      bool     `json:"is_paused,omitempty"`
	Address       string   `json:"msp_address,omitempty"`
	ExpectString  string   `json:"msp_expect_string,omitempty"`
	Threshold     int      `json:"msp_threshold,omitempty"`
	NumRetries    int      `json:"msp_num_retries,omitempty"`
	UptimeSla     float64  `json:"msp_uptime_sla,omitempty"`
	Notes         string   `json:"msp_notes,omitempty"`
}

func (c CheckWHOIS) PrimaryKey() int {
	return c.PK
}

type checksEndpointWHOISImpl struct {
	EndpointCreator[CheckWHOIS, CheckResponse, Check]
	EndpointUpdater[CheckWHOIS, CheckResponse, Check]
}

func (c checksEndpointWHOISImpl) CreateWHOIS(ctx context.Context, check CheckWHOIS) (*Check, error) {
	return c.Create(ctx, check)
}

func (c checksEndpointWHOISImpl) UpdateWHOIS(ctx context.Context, check CheckWHOIS) (*Check, error) {
	return c.Update(ctx, check)
}
