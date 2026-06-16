// pkg/upapi/ep_maintenance_notifications.go
package upapi

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

// parseDjangoDuration converts a Django DurationField string to total seconds.
// The wire format (verified against the live API) is duration_string:
// "[-][D ]HH:MM:SS[.ffffff]" with a SPACE day-separator. It also accepts a plain
// integer seconds string defensively. Negatives use a day borrow, so
// total = days*86400 + H*3600 + M*60 + S, where the hms part is always non-negative
// and the sign lives on the days field (e.g. "-1 23:30:00" = -86400 + 84600 = -1800).
func parseDjangoDuration(s string) (int64, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, nil
	}
	if n, err := strconv.ParseInt(s, 10, 64); err == nil {
		return n, nil
	}
	var days int64
	timePart := s
	if i := strings.IndexByte(s, ' '); i >= 0 {
		d, err := strconv.ParseInt(strings.TrimSpace(s[:i]), 10, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid day field %q: %w", s[:i], err)
		}
		days = d
		timePart = strings.TrimSpace(s[i+1:])
	}
	hms := strings.Split(timePart, ":")
	if len(hms) != 3 {
		return 0, fmt.Errorf("invalid time part %q", timePart)
	}
	hours, err := strconv.ParseInt(hms[0], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid hours %q: %w", hms[0], err)
	}
	minutes, err := strconv.ParseInt(hms[1], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid minutes %q: %w", hms[1], err)
	}
	secsFloat, err := strconv.ParseFloat(hms[2], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid seconds %q: %w", hms[2], err)
	}
	return days*86400 + hours*3600 + minutes*60 + int64(secsFloat), nil
}

// MaintenanceContactGroup is the nested contact-group object in read responses.
type MaintenanceContactGroup struct {
	PK   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// MaintenanceScheduleNotification is the read representation. Offset (seconds) is
// derived from the Django duration string OffsetRaw via withParsedOffset(), which is
// applied in every response type's Item()/List() conversion below.
//
// IMPORTANT: do NOT add a custom UnmarshalJSON to this type. The detail/create/update
// responses decode into the NAMED types MaintenanceNotificationResponse /
// MaintenanceNotificationCreateUpdateResponse, which do not inherit methods from this
// underlying type, so an UnmarshalJSON here would silently never run on those paths.
// Parsing inside Item()/List() covers all paths uniformly.
type MaintenanceScheduleNotification struct {
	PK            int64                     `json:"id,omitempty"`
	ScheduleID    int64                     `json:"schedule_id,omitempty"`
	Offset        int64                     `json:"-"`
	OffsetRaw     string                    `json:"offset,omitempty"`
	Event         string                    `json:"event,omitempty"`
	ContactGroups []MaintenanceContactGroup `json:"contact_groups,omitempty"`
	CreatedAt     string                    `json:"created_at,omitempty"`
	ModifiedAt    string                    `json:"modified_at,omitempty"`
}

// withParsedOffset returns a copy with Offset filled from OffsetRaw. A malformed
// server value falls back to 0 (the server controls this field, so this is defensive).
func (n MaintenanceScheduleNotification) withParsedOffset() MaintenanceScheduleNotification {
	secs, err := parseDjangoDuration(n.OffsetRaw)
	if err == nil {
		n.Offset = secs
	}
	return n
}

func (n MaintenanceScheduleNotification) PrimaryKey() PrimaryKey {
	return PrimaryKey(n.PK)
}

// MaintenanceScheduleNotificationInput is the write representation (offset in seconds).
type MaintenanceScheduleNotificationInput struct {
	ScheduleID    int64   `json:"schedule_id"`
	Offset        int64   `json:"offset"`
	Event         string  `json:"event"`
	ContactGroups []int64 `json:"contact_groups"`
}

type MaintenanceNotificationListOptions struct {
	Page     int64  `url:"page,omitempty"`
	PageSize int64  `url:"page_size,omitempty"`
	Ordering string `url:"ordering,omitempty"`
}

type MaintenanceNotificationListResponse struct {
	Count   int64                             `json:"count,omitempty"`
	Results []MaintenanceScheduleNotification `json:"results,omitempty"`
}

func (r MaintenanceNotificationListResponse) List() []MaintenanceScheduleNotification {
	out := make([]MaintenanceScheduleNotification, len(r.Results))
	for i := range r.Results {
		out[i] = r.Results[i].withParsedOffset()
	}
	return out
}
func (r MaintenanceNotificationListResponse) CountItems() int64 { return r.Count }

type MaintenanceNotificationResponse MaintenanceScheduleNotification

func (r MaintenanceNotificationResponse) Item() MaintenanceScheduleNotification {
	return MaintenanceScheduleNotification(r).withParsedOffset()
}

type MaintenanceNotificationCreateUpdateResponse struct {
	Results MaintenanceScheduleNotification `json:"results,omitempty"`
}

func (r MaintenanceNotificationCreateUpdateResponse) Item() MaintenanceScheduleNotification {
	return r.Results.withParsedOffset()
}

type MaintenanceNotificationsEndpoint interface {
	List(context.Context, MaintenanceNotificationListOptions) (*ListResult[MaintenanceScheduleNotification], error)
	Create(context.Context, MaintenanceScheduleNotificationInput) (*MaintenanceScheduleNotification, error)
	Update(context.Context, PrimaryKeyable, MaintenanceScheduleNotificationInput) (*MaintenanceScheduleNotification, error)
	Get(context.Context, PrimaryKeyable) (*MaintenanceScheduleNotification, error)
	Delete(context.Context, PrimaryKeyable) error
}

func NewMaintenanceNotificationsEndpoint(cbd CBD) MaintenanceNotificationsEndpoint {
	const endpoint = "maintenance/notifications"
	return &maintenanceNotificationsEndpointImpl{
		EndpointLister:  NewEndpointLister[MaintenanceNotificationListResponse, MaintenanceScheduleNotification, MaintenanceNotificationListOptions](cbd, endpoint),
		EndpointCreator: NewEndpointCreator[MaintenanceScheduleNotificationInput, MaintenanceNotificationCreateUpdateResponse, MaintenanceScheduleNotification](cbd, endpoint),
		EndpointUpdater: NewEndpointUpdater[MaintenanceScheduleNotificationInput, MaintenanceNotificationCreateUpdateResponse, MaintenanceScheduleNotification](cbd, endpoint),
		EndpointGetter:  NewEndpointGetter[MaintenanceNotificationResponse, MaintenanceScheduleNotification](cbd, endpoint),
		EndpointDeleter: NewEndpointDeleter(cbd, endpoint),
	}
}

type maintenanceNotificationsEndpointImpl struct {
	EndpointLister[MaintenanceNotificationListResponse, MaintenanceScheduleNotification, MaintenanceNotificationListOptions]
	EndpointCreator[MaintenanceScheduleNotificationInput, MaintenanceNotificationCreateUpdateResponse, MaintenanceScheduleNotification]
	EndpointUpdater[MaintenanceScheduleNotificationInput, MaintenanceNotificationCreateUpdateResponse, MaintenanceScheduleNotification]
	EndpointGetter[MaintenanceNotificationResponse, MaintenanceScheduleNotification]
	EndpointDeleter
}
