package upapi

import (
	"context"
	"encoding/json"
	"fmt"
)

type SLAReportService struct {
	PK   int    `json:"pk,omitempty"`
	Name string `json:"name,omitempty"`
}

// UnmarshalJSON implements the json.Unmarshaler interface.
//
// This is workaround for the fact that the API can return either a string (name) or an int for the PK field.
func (s *SLAReportService) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 0 {
		return fmt.Errorf("empty data")
	}

	if data[0] == '"' {
		err = json.Unmarshal(data, &s.Name)
	} else {
		err = json.Unmarshal(data, &s.PK)
	}

	return
}

// MarshalJSON implements the json.Marshaler interface.
//
// This is workaround for the fact that the API can accept either a string (name) or an int for the PK field.
func (s SLAReportService) MarshalJSON() ([]byte, error) {
	if s.Name != "" && s.PK != 0 {
		return nil, fmt.Errorf("SLA service can only have a name or a PK, not both")
	}

	if s.PK != 0 {
		return json.Marshal(s.PK)
	}

	return json.Marshal(s.Name)
}

type SLAReport struct {
	PK                              int64               `json:"pk,omitempty"`
	URL                             string              `json:"url,omitempty"`
	StatsURL                        string              `json:"stats_url,omitempty"`
	Name                            string              `json:"name"`
	ServicesTags                    []string            `json:"services_tags,omitempty"`
	ServicesSelected                *[]SLAReportService `json:"services_selected,omitempty"`
	ReportingGroups                 *[]SLAReportGroup   `json:"reporting_groups,omitempty"`
	DefaultDateRange                string              `json:"default_date_range,omitempty"`
	FilterWithDowntime              bool                `json:"filter_with_downtime,omitempty"`
	FilterUptimeSLAViolations       bool                `json:"filter_uptime_sla_violations,omitempty"`
	FilterSlowest                   bool                `json:"filter_slowest,omitempty"`
	FilterResponseTimeSLAViolations bool                `json:"filter_response_time_sla_violations,omitempty"`
	ShowUptimeSection               bool                `json:"show_uptime_section,omitempty"`
	ShowUptimeSLA                   bool                `json:"show_uptime_sla,omitempty"`
	ShowResponseTimeSection         bool                `json:"show_response_time_section,omitempty"`
	ShowResponseTimeSLA             bool                `json:"show_response_time_sla,omitempty"`
	UptimeSectionSort               string              `json:"uptime_section_sort,omitempty"`
	ResponseTimeSectionSort         string              `json:"response_time_section_sort,omitempty"`
}

func (s SLAReport) PrimaryKey() PrimaryKey {
	return PrimaryKey(s.PK)
}

type SLAReportListOptions struct {
	Page     int64  `url:"page,omitempty"`
	PageSize int64  `url:"page_size,omitempty"`
	Search   string `url:"search,omitempty"`
	Ordering string `url:"ordering,omitempty"`
}

type SLAReportListResponse struct {
	Count   int64       `json:"count,omitempty"`
	Results []SLAReport `json:"results,omitempty"`
}

func (r SLAReportListResponse) List() []SLAReport {
	return r.Results
}

func (r SLAReportListResponse) CountItems() int64 {
	return r.Count
}

type SLAReportResponse SLAReport

func (r SLAReportResponse) Item() SLAReport {
	return SLAReport(r)
}

type SLAReportCreateUpdateResponse struct {
	Results SLAReport `json:"results,omitempty"`
}

func (r SLAReportCreateUpdateResponse) Item() SLAReport {
	return r.Results
}

type SLAReportsEndpoint interface {
	List(context.Context, SLAReportListOptions) (*ListResult[SLAReport], error)
	Create(context.Context, SLAReport) (*SLAReport, error)
	Update(context.Context, PrimaryKeyable, SLAReport) (*SLAReport, error)
	Get(context.Context, PrimaryKeyable) (*SLAReport, error)
	Delete(context.Context, PrimaryKeyable) error
	ReportingGroups(PrimaryKeyable) SLAReportsGroupsEndpoint
}

func NewSLAReportsEndpoint(cbd CBD) SLAReportsEndpoint {
	const endpoint = "sla-reports"
	return &slaReportEndpointImpl{
		cbd:             cbd,
		endpoint:        endpoint,
		EndpointLister:  NewEndpointLister[SLAReportListResponse, SLAReport, SLAReportListOptions](cbd, endpoint),
		EndpointCreator: NewEndpointCreator[SLAReport, SLAReportCreateUpdateResponse](cbd, endpoint),
		EndpointUpdater: NewEndpointUpdater[SLAReport, SLAReportCreateUpdateResponse](cbd, endpoint),
		EndpointGetter:  NewEndpointGetter[SLAReportResponse](cbd, endpoint),
		EndpointDeleter: NewEndpointDeleter(cbd, endpoint),
	}
}

type slaReportEndpointImpl struct {
	cbd      CBD
	endpoint string
	EndpointLister[SLAReportListResponse, SLAReport, SLAReportListOptions]
	EndpointCreator[SLAReport, SLAReportCreateUpdateResponse, SLAReport]
	EndpointUpdater[SLAReport, SLAReportCreateUpdateResponse, SLAReport]
	EndpointGetter[SLAReportResponse, SLAReport]
	EndpointDeleter
}

func (c *slaReportEndpointImpl) ReportingGroups(pk PrimaryKeyable) SLAReportsGroupsEndpoint {
	endpoint := fmt.Sprintf("%s/%d/groups", c.endpoint, pk.PrimaryKey())
	return &slaReportsGroupsEndpointImpl{
		EndpointLister:  NewEndpointLister[SLAReportGroupListResponse, SLAReportGroup, SLAReportGroupListOptions](c.cbd, endpoint),
		EndpointCreator: NewEndpointCreator[SLAReportGroup, SLAReportGroupCreateUpdateResponse](c.cbd, endpoint),
		EndpointGetter:  NewEndpointGetter[SLAReportGroupResponse](c.cbd, endpoint),
		EndpointDeleter: NewEndpointDeleter(c.cbd, endpoint),
	}
}
