package upapi

import "context"

type Contact struct {
	PK                       int      `json:"pk,omitempty"`
	Name                     string   `json:"name,omitempty"`
	SmsList                  []string `json:"sms_list,omitempty"`
	EmailList                []string `json:"email_list,omitempty"`
	PhonecallList            []string `json:"phonecall_list,omitempty"`
	Integrations             []string `json:"integrations,omitempty"`
	PushNotificationProfiles []string `json:"push_notification_profiles,omitempty"`
}

func (c Contact) PrimaryKey() int {
	return c.PK
}

type ContactListOptions struct {
	Page              int    `url:"page,omitempty"`
	PageSize          int    `url:"page_size,omitempty"`
	Search            string `url:"search,omitempty"`
	Ordering          string `url:"ordering,omitempty"`
	HasOnCallSchedule bool   `url:"has_on_call_schedule,omitempty"`
}

type ContactListResponse struct {
	Count   int       `json:"count,omitempty"`
	Results []Contact `json:"results,omitempty"`
}

func (r ContactListResponse) List() []Contact {
	return r.Results
}

type ContactResponse Contact

func (r ContactResponse) Item() *Contact {
	return (*Contact)(&r)
}

type ContactsEndpoint interface {
	List(context.Context, ContactListOptions) ([]Contact, error)
	Create(context.Context, Contact) (*Contact, error)
	Update(context.Context, Contact) (*Contact, error)
	Get(context.Context, PrimaryKey) (*Contact, error)
	Delete(context.Context, PrimaryKey) error
}

func NewContactsEndpoint(cbd CBD) ContactsEndpoint {
	const endpoint = "contacts"
	return &contactsEndpointImpl{
		EndpointLister:  NewEndpointLister[ContactListResponse, Contact, ContactListOptions](cbd, endpoint),
		EndpointCreator: NewEndpointCreator[Contact, ContactResponse, Contact](cbd, endpoint),
		EndpointUpdater: NewEndpointUpdater[Contact, ContactResponse, Contact](cbd, endpoint),
		EndpointGetter:  NewEndpointGetter[PrimaryKey, ContactResponse, Contact](cbd, endpoint),
		EndpointDeleter: NewEndpointDeleter[PrimaryKey](cbd, endpoint),
	}
}

type contactsEndpointImpl struct {
	EndpointLister[ContactListResponse, Contact, ContactListOptions]
	EndpointCreator[Contact, ContactResponse, Contact]
	EndpointUpdater[Contact, ContactResponse, Contact]
	EndpointGetter[PrimaryKey, ContactResponse, Contact]
	EndpointDeleter[PrimaryKey]
}
