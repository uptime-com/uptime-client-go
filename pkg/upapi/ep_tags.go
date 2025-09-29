package upapi

import "context"

// Tag represents a check tag in Uptime.com.
type Tag struct {
	PK       int64  `json:"pk,omitempty"`
	URL      string `json:"url,omitempty"`
	Tag      string `json:"tag,omitempty"`
	ColorHex string `json:"color_hex,omitempty"`
}

func (t Tag) PrimaryKey() PrimaryKey {
	return PrimaryKey(t.PK)
}

// TagListResponse represents a list of check tags.
type TagListResponse struct {
	Count    int64  `json:"count,omitempty"`
	Next     string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
	Results  []Tag  `json:"results,omitempty"`
}

func (r TagListResponse) List() []Tag {
	return r.Results
}

// TagListOptions specifies the optional parameters to tag listing API call.
type TagListOptions struct {
	Page     int64  `url:"page,omitempty"`
	PageSize int64  `url:"page_size,omitempty"`
	Search   string `url:"search,omitempty"`
	Ordering string `url:"ordering,omitempty"`
}

// TagItemResponse represents a response from the tagsImpl.
type TagItemResponse = Tag

func (t TagItemResponse) Item() Tag {
	return t
}

// TagCreateUpdateResponse represents a response from the tagsImpl
type TagCreateUpdateResponse struct {
	Results Tag `json:"results,omitempty"`
}

func (t TagCreateUpdateResponse) Item() Tag {
	return t.Results
}

type TagsEndpoint interface {
	List(context.Context, TagListOptions) ([]Tag, error)
	Create(context.Context, Tag) (*Tag, error)
	Get(context.Context, PrimaryKeyable) (*Tag, error)
	Update(context.Context, PrimaryKeyable, Tag) (*Tag, error)
	Delete(context.Context, PrimaryKeyable) error
}

func NewTagsEndpoint(cbd CBD) TagsEndpoint {
	const endpoint = "check-tags"
	return &tagsEndpointImpl{
		EndpointLister:  NewEndpointLister[TagListResponse, Tag, TagListOptions](cbd, endpoint),
		EndpointGetter:  NewEndpointGetter[TagItemResponse, Tag](cbd, endpoint),
		EndpointCreator: NewEndpointCreator[Tag, TagCreateUpdateResponse, Tag](cbd, endpoint),
		EndpointUpdater: NewEndpointUpdater[Tag, TagCreateUpdateResponse, Tag](cbd, endpoint),
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
