package upapi

import (
	"context"
)

type Contact struct {
	PK                       int64    `json:"pk,omitempty"`
	URL                      string   `json:"url,omitempty"`
	Name                     string   `json:"name,omitempty"`
	SmsList                  []string `json:"sms_list,omitempty"`
	EmailList                []string `json:"email_list,omitempty"`
	PhonecallList            []string `json:"phonecall_list,omitempty"`
	Integrations             []string `json:"integrations,omitempty"`
	PushNotificationProfiles []string `json:"push_notification_profiles,omitempty"`
}

func (c Contact) PrimaryKey() PrimaryKey {
	return PrimaryKey(c.PK)
}

type ContactListOptions struct {
	Page              int64  `url:"page,omitempty"`
	PageSize          int64  `url:"page_size,omitempty"`
	Search            string `url:"search,omitempty"`
	Ordering          string `url:"ordering,omitempty"`
	HasOnCallSchedule bool   `url:"has_on_call_schedule,omitempty"`
}

type ContactListResponse struct {
	Count   int64     `json:"count,omitempty"`
	Results []Contact `json:"results,omitempty"`
}

func (r ContactListResponse) List() []Contact {
	return r.Results
}

func (r ContactListResponse) CountItems() int64 {
	return r.Count
}

type ContactResponse Contact

func (r ContactResponse) Item() Contact {
	return Contact(r)
}

type ContactCreateUpdateResponse struct {
	Results Contact `json:"results,omitempty"`
}

func (r ContactCreateUpdateResponse) Item() Contact {
	return r.Results
}

type ContactsEndpoint interface {
	List(context.Context, ContactListOptions) (*ListResult[Contact], error)
	Create(context.Context, Contact) (*Contact, error)
	Update(context.Context, PrimaryKeyable, Contact) (*Contact, error)
	Get(context.Context, PrimaryKeyable) (*Contact, error)
	Delete(context.Context, PrimaryKeyable) error
}

func NewContactsEndpoint(cbd CBD) ContactsEndpoint {
	const endpoint = "contacts"
	return &contactsEndpointImpl{
		EndpointLister:  NewEndpointLister[ContactListResponse, Contact, ContactListOptions](cbd, endpoint),
		EndpointCreator: NewEndpointCreator[Contact, ContactCreateUpdateResponse, Contact](cbd, endpoint),
		EndpointUpdater: NewEndpointUpdater[Contact, ContactCreateUpdateResponse, Contact](cbd, endpoint),
		EndpointGetter:  NewEndpointGetter[ContactResponse, Contact](cbd, endpoint),
		EndpointDeleter: NewEndpointDeleter(cbd, endpoint),
	}
}

type contactsEndpointImpl struct {
	EndpointLister[ContactListResponse, Contact, ContactListOptions]
	EndpointCreator[Contact, ContactCreateUpdateResponse, Contact]
	EndpointUpdater[Contact, ContactCreateUpdateResponse, Contact]
	EndpointGetter[ContactResponse, Contact]
	EndpointDeleter
}
