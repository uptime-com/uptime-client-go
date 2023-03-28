package upapi

import "context"

// Tag represents a check tag in Uptime.com.
type Tag struct {
	PK       int    `json:"pk,omitempty"`
	URL      string `json:"url,omitempty"`
	Tag      string `json:"tag,omitempty"`
	ColorHex string `json:"color_hex,omitempty"`
}

func (t Tag) PrimaryKey() PrimaryKey {
	return PrimaryKey(t.PK)
}

// TagListResponse represents a list of check tags.
type TagListResponse struct {
	Count    int    `json:"count,omitempty"`
	Next     string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
	Results  []Tag  `json:"results,omitempty"`
}

func (r TagListResponse) List() []Tag {
	return r.Results
}

// TagListOptions specifies the optional parameters to tag listing API call.
type TagListOptions struct {
	Page     int    `url:"page,omitempty"`
	PageSize int    `url:"page_size,omitempty"`
	Search   string `url:"search,omitempty"`
	Ordering string `url:"ordering,omitempty"`
}

// TagItemResponse represents a response from the tagsImpl.
type TagItemResponse = Tag

func (t TagItemResponse) Item() *Tag {
	return &t
}

type TagsEndpoint interface {
	List(context.Context, TagListOptions) ([]Tag, error)
	Get(context.Context, PrimaryKeyable) (*Tag, error)
	Create(context.Context, Tag) (*Tag, error)
	Update(context.Context, Tag) (*Tag, error)
	Delete(context.Context, PrimaryKeyable) error
}

func NewTagsEndpoint(cbd CBD) TagsEndpoint {
	const endpoint = "check-tags"
	return &tagsEndpointImpl{
		EndpointLister:  NewEndpointLister[TagListResponse, Tag, TagListOptions](cbd, endpoint),
		EndpointGetter:  NewEndpointGetter[TagItemResponse, Tag](cbd, endpoint),
		EndpointCreator: NewEndpointCreator[Tag, TagItemResponse, Tag](cbd, endpoint),
		EndpointUpdater: NewEndpointUpdater[Tag, TagItemResponse, Tag](cbd, endpoint),
		EndpointDeleter: NewEndpointDeleter(cbd, endpoint),
	}
}

type tagsEndpointImpl struct {
	EndpointLister[TagListResponse, Tag, TagListOptions]
	EndpointGetter[TagItemResponse, Tag]
	EndpointCreator[Tag, TagItemResponse, Tag]
	EndpointUpdater[Tag, TagItemResponse, Tag]
	EndpointDeleter
}
