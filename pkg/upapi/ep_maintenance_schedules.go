package upapi

import "context"

// MaintenanceService is the nested service object returned in read responses.
type MaintenanceService struct {
	PK      int64  `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
	Type    string `json:"type,omitempty"`
}

// MaintenanceTag is the nested tag object returned in read responses.
type MaintenanceTag struct {
	PK       int64  `json:"id,omitempty"`
	Tag      string `json:"tag,omitempty"`
	ColorHex string `json:"color_hex,omitempty"`
}

// MaintenanceSchedule is the read/result representation of a maintenance schedule.
type MaintenanceSchedule struct {
	PK                           int64                `json:"id,omitempty"`
	Name                         string               `json:"name,omitempty"`
	ScheduleType                 string               `json:"schedule_type,omitempty"`
	StartsAt                     string               `json:"starts_at,omitempty"`
	EndsAt                       string               `json:"ends_at,omitempty"`
	RRule                        string               `json:"rrule,omitempty"`
	DurationMinutes              *int64               `json:"duration_minutes,omitempty"`
	IsActive                     bool                 `json:"is_active"`
	PauseChecksDuringMaintenance bool                 `json:"pause_checks_during_maintenance"`
	Services                     []MaintenanceService `json:"services,omitempty"`
	Tags                         []MaintenanceTag     `json:"tags,omitempty"`
	CreatedAt                    string               `json:"created_at,omitempty"`
	ModifiedAt                   string               `json:"modified_at,omitempty"`
}

func (m MaintenanceSchedule) PrimaryKey() PrimaryKey {
	return PrimaryKey(m.PK)
}

// MaintenanceScheduleInput is the write representation (services/tags as ID lists).
// No omitempty on scalar fields: Terraform always sends the full desired state, and
// PATCH must overwrite (e.g. is_active=false must not be dropped).
type MaintenanceScheduleInput struct {
	Name                         string  `json:"name"`
	ScheduleType                 string  `json:"schedule_type"`
	StartsAt                     string  `json:"starts_at"`
	EndsAt                       *string `json:"ends_at"`
	RRule                        string  `json:"rrule"`
	DurationMinutes              *int64  `json:"duration_minutes"`
	IsActive                     bool    `json:"is_active"`
	PauseChecksDuringMaintenance bool    `json:"pause_checks_during_maintenance"`
	Services                     []int64 `json:"services"`
	Tags                         []int64 `json:"tags"`
}

type MaintenanceScheduleListOptions struct {
	Page     int64  `url:"page,omitempty"`
	PageSize int64  `url:"page_size,omitempty"`
	Search   string `url:"search,omitempty"`
	Ordering string `url:"ordering,omitempty"`
}

type MaintenanceScheduleListResponse struct {
	Count   int64                 `json:"count,omitempty"`
	Results []MaintenanceSchedule `json:"results,omitempty"`
}

func (r MaintenanceScheduleListResponse) List() []MaintenanceSchedule { return r.Results }
func (r MaintenanceScheduleListResponse) CountItems() int64           { return r.Count }

type MaintenanceScheduleResponse MaintenanceSchedule

func (r MaintenanceScheduleResponse) Item() MaintenanceSchedule {
	return MaintenanceSchedule(r)
}

type MaintenanceScheduleCreateUpdateResponse struct {
	Results MaintenanceSchedule `json:"results,omitempty"`
}

func (r MaintenanceScheduleCreateUpdateResponse) Item() MaintenanceSchedule {
	return r.Results
}

type MaintenanceSchedulesEndpoint interface {
	List(context.Context, MaintenanceScheduleListOptions) (*ListResult[MaintenanceSchedule], error)
	Create(context.Context, MaintenanceScheduleInput) (*MaintenanceSchedule, error)
	Update(context.Context, PrimaryKeyable, MaintenanceScheduleInput) (*MaintenanceSchedule, error)
	Get(context.Context, PrimaryKeyable) (*MaintenanceSchedule, error)
	Delete(context.Context, PrimaryKeyable) error
}

func NewMaintenanceSchedulesEndpoint(cbd CBD) MaintenanceSchedulesEndpoint {
	const endpoint = "maintenance/schedules"
	return &maintenanceSchedulesEndpointImpl{
		EndpointLister:  NewEndpointLister[MaintenanceScheduleListResponse, MaintenanceSchedule, MaintenanceScheduleListOptions](cbd, endpoint),
		EndpointCreator: NewEndpointCreator[MaintenanceScheduleInput, MaintenanceScheduleCreateUpdateResponse, MaintenanceSchedule](cbd, endpoint),
		EndpointUpdater: NewEndpointUpdater[MaintenanceScheduleInput, MaintenanceScheduleCreateUpdateResponse, MaintenanceSchedule](cbd, endpoint),
		EndpointGetter:  NewEndpointGetter[MaintenanceScheduleResponse, MaintenanceSchedule](cbd, endpoint),
		EndpointDeleter: NewEndpointDeleter(cbd, endpoint),
	}
}

type maintenanceSchedulesEndpointImpl struct {
	EndpointLister[MaintenanceScheduleListResponse, MaintenanceSchedule, MaintenanceScheduleListOptions]
	EndpointCreator[MaintenanceScheduleInput, MaintenanceScheduleCreateUpdateResponse, MaintenanceSchedule]
	EndpointUpdater[MaintenanceScheduleInput, MaintenanceScheduleCreateUpdateResponse, MaintenanceSchedule]
	EndpointGetter[MaintenanceScheduleResponse, MaintenanceSchedule]
	EndpointDeleter
}
