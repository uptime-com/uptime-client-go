package upapi

import (
	"context"
)

type ScheduledReport struct {
	PK              int64    `json:"pk,omitempty"`
	URL             string   `json:"url,omitempty"`
	Name            string   `json:"name"`
	ScheduledReport string   `json:"sla_report,omitempty"`
	RecipientUsers  []string `json:"recipient_users,omitempty"`
	RecipientEmails []string `json:"recipient_emails,omitempty"`
	FileType        string   `json:"file_type,omitempty"`
	Recurrence      string   `json:"recurrence,omitempty"`
	OnWeekday       int32    `json:"on_weekday,omitempty"`
	AtTime          int32    `json:"at_time"`
	IsEnabled       bool     `json:"is_enabled,omitempty"`
}

func (s ScheduledReport) PrimaryKey() PrimaryKey {
	return PrimaryKey(s.PK)
}

type ScheduledReportListOptions struct {
	Page     int64  `url:"page,omitempty"`
	PageSize int64  `url:"page_size,omitempty"`
	Search   string `url:"search,omitempty"`
	Ordering string `url:"ordering,omitempty"`
}

type ScheduledReportListResponse struct {
	Count   int64             `json:"count,omitempty"`
	Results []ScheduledReport `json:"results,omitempty"`
}

func (r ScheduledReportListResponse) List() []ScheduledReport {
	return r.Results
}

type ScheduledReportResponse ScheduledReport

func (r ScheduledReportResponse) Item() ScheduledReport {
	return ScheduledReport(r)
}

type ScheduledReportCreateUpdateResponse struct {
	Results ScheduledReport `json:"results,omitempty"`
}

func (r ScheduledReportCreateUpdateResponse) Item() ScheduledReport {
	return r.Results
}

type ScheduledReportsEndpoint interface {
	List(context.Context, ScheduledReportListOptions) ([]ScheduledReport, error)
	Create(context.Context, ScheduledReport) (*ScheduledReport, error)
	Update(context.Context, PrimaryKeyable, ScheduledReport) (*ScheduledReport, error)
	Get(context.Context, PrimaryKeyable) (*ScheduledReport, error)
	Delete(context.Context, PrimaryKeyable) error
}

func NewScheduledReportsEndpoint(cbd CBD) ScheduledReportsEndpoint {
	const endpoint = "scheduled-reports"
	return &scheduledReportEndpointImpl{
		EndpointLister:  NewEndpointLister[ScheduledReportListResponse, ScheduledReport, ScheduledReportListOptions](cbd, endpoint),
		EndpointCreator: NewEndpointCreator[ScheduledReport, ScheduledReportCreateUpdateResponse](cbd, endpoint),
		EndpointUpdater: NewEndpointUpdater[ScheduledReport, ScheduledReportCreateUpdateResponse](cbd, endpoint),
		EndpointGetter:  NewEndpointGetter[ScheduledReportResponse](cbd, endpoint),
		EndpointDeleter: NewEndpointDeleter(cbd, endpoint),
	}
}

type scheduledReportEndpointImpl struct {
	EndpointLister[ScheduledReportListResponse, ScheduledReport, ScheduledReportListOptions]
	EndpointCreator[ScheduledReport, ScheduledReportCreateUpdateResponse, ScheduledReport]
	EndpointUpdater[ScheduledReport, ScheduledReportCreateUpdateResponse, ScheduledReport]
	EndpointGetter[ScheduledReportResponse, ScheduledReport]
	EndpointDeleter
}
