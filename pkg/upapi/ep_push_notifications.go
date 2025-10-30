package upapi

import (
	"context"
)

type PushNotificationProfile struct {
	PK            int64    `json:"pk,omitempty"`
	URL           string   `json:"url,omitempty"`
	CreatedAt     string   `json:"created_at,omitempty"`
	ModifiedAt    string   `json:"modified_at,omitempty"`
	UUID          string   `json:"uuid,omitempty"`
	User          string   `json:"user,omitempty"`
	DeviceName    string   `json:"device_name,omitempty"`
	DisplayName   string   `json:"display_name,omitempty"`
	ContactGroups []string `json:"contact_groups,omitempty"`
}

func (p PushNotificationProfile) PrimaryKey() PrimaryKey {
	return PrimaryKey(p.PK)
}

type PushNotificationProfileListResponse struct {
	Count    int64                     `json:"count,omitempty"`
	Next     string                    `json:"next,omitempty"`
	Previous string                    `json:"previous,omitempty"`
	Results  []PushNotificationProfile `json:"results,omitempty"`
}

func (r PushNotificationProfileListResponse) List() []PushNotificationProfile {
	return r.Results
}

type PushNotificationProfileListOptions struct {
	Page     int64  `url:"page,omitempty"`
	PageSize int64  `url:"page_size,omitempty"`
	Search   string `url:"search,omitempty"`
	Ordering string `url:"ordering,omitempty"`
}

type PushNotificationProfileItemResponse struct {
	PushNotificationProfile `json:",inline"`
}

func (p PushNotificationProfileItemResponse) Item() PushNotificationProfile {
	return p.PushNotificationProfile
}

type PushNotificationProfileCreateUpdateResponse struct {
	PushNotificationProfile `json:",inline"`
}

func (p PushNotificationProfileCreateUpdateResponse) Item() PushNotificationProfile {
	return p.PushNotificationProfile
}

type PushNotificationProfileCreateRequest struct {
	AppKey        string   `json:"app_key" flag:"app-key"`
	UUID          string   `json:"uuid,omitempty" flag:"uuid"`
	DeviceName    string   `json:"device_name" flag:"device-name"`
	ContactGroups []string `json:"contact_groups" flag:"contact-groups"`
}

type PushNotificationProfileUpdateRequest struct {
	DeviceName    string   `json:"device_name,omitempty" flag:"device-name"`
	ContactGroups []string `json:"contact_groups,omitempty" flag:"contact-groups"`
}

type PushNotificationsEndpoint interface {
	List(context.Context, PushNotificationProfileListOptions) ([]PushNotificationProfile, error)
	Create(context.Context, PushNotificationProfileCreateRequest) (*PushNotificationProfile, error)
	Get(context.Context, PrimaryKeyable) (*PushNotificationProfile, error)
	Update(context.Context, PrimaryKeyable, PushNotificationProfileUpdateRequest) (*PushNotificationProfile, error)
	Delete(context.Context, PrimaryKeyable) error
}

func NewPushNotificationsEndpoint(cbd CBD) PushNotificationsEndpoint {
	const endpoint = "push-notifications"
	return &pushNotificationsEndpointImpl{
		CBD:             cbd,
		endpoint:        endpoint,
		EndpointLister:  NewEndpointLister[PushNotificationProfileListResponse, PushNotificationProfile, PushNotificationProfileListOptions](cbd, endpoint),
		EndpointGetter:  NewEndpointGetter[PushNotificationProfileItemResponse, PushNotificationProfile](cbd, endpoint),
		EndpointCreator: NewEndpointCreator[PushNotificationProfileCreateRequest, PushNotificationProfileCreateUpdateResponse, PushNotificationProfile](cbd, endpoint),
		EndpointUpdater: NewEndpointUpdater[PushNotificationProfileUpdateRequest, PushNotificationProfileCreateUpdateResponse, PushNotificationProfile](cbd, endpoint),
		EndpointDeleter: NewEndpointDeleter(cbd, endpoint),
	}
}

type pushNotificationsEndpointImpl struct {
	CBD
	endpoint string
	EndpointLister[PushNotificationProfileListResponse, PushNotificationProfile, PushNotificationProfileListOptions]
	EndpointGetter[PushNotificationProfileItemResponse, PushNotificationProfile]
	EndpointCreator[PushNotificationProfileCreateRequest, PushNotificationProfileCreateUpdateResponse, PushNotificationProfile]
	EndpointUpdater[PushNotificationProfileUpdateRequest, PushNotificationProfileCreateUpdateResponse, PushNotificationProfile]
	EndpointDeleter
}
