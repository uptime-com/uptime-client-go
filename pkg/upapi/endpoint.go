package upapi

import (
	"context"
	"fmt"
	"net/http"
)

type Itemable[Item any] interface {
	Item() Item
}

type Listable[Item any] interface {
	List() []Item
	CountItems() int64
}

// ListResult contains items and total count from list responses.
type ListResult[Item any] struct {
	Items      []Item
	TotalCount int64
}

type PrimaryKeyable interface {
	PrimaryKey() PrimaryKey
}

type PrimaryKey int

func (p PrimaryKey) PrimaryKey() PrimaryKey {
	return p
}

// EndpointGetter is a generic interface for getting a single item from an endpoint.
type EndpointGetter[ResponseType Itemable[ItemType], ItemType any] interface {
	Get(ctx context.Context, pk PrimaryKeyable) (*ItemType, error)
}

func NewEndpointGetter[ResponseType Itemable[ItemType], ItemType any](cbd CBD, endpoint string) EndpointGetter[ResponseType, ItemType] {
	return &endpointGetterImpl[ResponseType, ItemType]{cbd, endpoint}
}

type endpointGetterImpl[ResponseType Itemable[ItemType], ItemType any] struct {
	CBD
	endpoint string
}

func (p *endpointGetterImpl[ResponseType, ItemType]) Get(ctx context.Context, pk PrimaryKeyable) (*ItemType, error) {
	rq, err := p.BuildRequest(ctx, http.MethodGet, fmt.Sprintf("%s/%d/", p.endpoint, pk.PrimaryKey()), nil, nil)
	if err != nil {
		return nil, err
	}
	rs, err := p.Do(rq)
	if err != nil {
		return nil, err
	}
	defer rs.Body.Close()
	if rs.StatusCode != http.StatusOK {
		return nil, ErrorFromResponse(rs)
	}
	var data ResponseType
	err = p.DecodeResponse(rs, &data)
	if err != nil {
		return nil, err
	}
	item := data.Item()
	return &item, nil
}

// EndpointLister is a generic interface for listing items from an endpoint.
type EndpointLister[ResponseType Listable[ItemType], ItemType any, OptionsType any] interface {
	List(ctx context.Context, opts OptionsType) (*ListResult[ItemType], error)
}

func NewEndpointLister[ResponseType Listable[ItemType], ItemType any, OptionsType any](cbd CBD, endpoint string) EndpointLister[Listable[ItemType], ItemType, OptionsType] {
	return &endpointListerImpl[ResponseType, ItemType, OptionsType]{cbd, endpoint}
}

type endpointListerImpl[ResponseType Listable[ItemType], ItemType any, OptionsType any] struct {
	CBD
	endpoint string
}

func (p *endpointListerImpl[ResponseType, ItemType, OptionsType]) List(ctx context.Context, opts OptionsType) (*ListResult[ItemType], error) {
	rq, err := p.BuildRequest(ctx, http.MethodGet, fmt.Sprintf("%s/", p.endpoint), opts, nil)
	if err != nil {
		return nil, err
	}
	rs, err := p.Do(rq)
	if err != nil {
		return nil, err
	}
	defer rs.Body.Close()
	if rs.StatusCode != http.StatusOK {
		return nil, ErrorFromResponse(rs)
	}
	var data ResponseType
	err = p.DecodeResponse(rs, &data)
	if err != nil {
		return nil, err
	}
	return &ListResult[ItemType]{
		Items:      data.List(),
		TotalCount: data.CountItems(),
	}, nil
}

// EndpointCreator is a generic interface for creating an item from an endpoint.
type EndpointCreator[RequestType any, ResponseType Itemable[ItemType], ItemType any] interface {
	Create(ctx context.Context, arg RequestType) (*ItemType, error)
}

func NewEndpointCreator[RequestType any, ResponseType Itemable[ItemType], ItemType any](cbd CBD, endpoint string) EndpointCreator[RequestType, ResponseType, ItemType] {
	return &endpointCreatorImpl[RequestType, ResponseType, ItemType]{cbd, endpoint}
}

type endpointCreatorImpl[RequestType any, ResponseType Itemable[ItemType], ItemType any] struct {
	CBD
	endpoint string
}

func (p *endpointCreatorImpl[RequestType, ResponseType, ItemType]) Create(ctx context.Context, arg RequestType) (*ItemType, error) {
	rq, err := p.BuildRequest(ctx, http.MethodPost, fmt.Sprintf("%s/", p.endpoint), nil, arg)
	if err != nil {
		return nil, err
	}
	rs, err := p.Do(rq)
	if err != nil {
		return nil, err
	}
	defer rs.Body.Close()
	if rs.StatusCode != http.StatusOK {
		return nil, ErrorFromResponse(rs)
	}
	var data ResponseType
	err = p.DecodeResponse(rs, &data)
	if err != nil {
		return nil, err
	}
	item := data.Item()
	return &item, nil
}

// EndpointUpdater is a generic interface for updating an item from an endpoint.
type EndpointUpdater[ArgumentType any, ResultType Itemable[ItemType], ItemType any] interface {
	Update(ctx context.Context, pk PrimaryKeyable, arg ArgumentType) (*ItemType, error)
}

func NewEndpointUpdater[RequestType any, ResponseType Itemable[ItemType], ItemType any](cbd CBD, endpoint string) EndpointUpdater[RequestType, ResponseType, ItemType] {
	return &endpointUpdaterImpl[RequestType, ResponseType, ItemType]{cbd, endpoint}
}

type endpointUpdaterImpl[RequestType any, ResponseType Itemable[ItemType], ItemType any] struct {
	CBD
	endpoint string
}

func (p *endpointUpdaterImpl[RequestType, ResponseType, ItemType]) Update(ctx context.Context, pk PrimaryKeyable, arg RequestType) (*ItemType, error) {
	rq, err := p.BuildRequest(ctx, http.MethodPatch, fmt.Sprintf("%s/%d/", p.endpoint, pk.PrimaryKey()), nil, arg)
	if err != nil {
		return nil, err
	}
	rs, err := p.Do(rq)
	if err != nil {
		return nil, err
	}
	defer rs.Body.Close()
	if rs.StatusCode != http.StatusOK {
		return nil, ErrorFromResponse(rs)
	}
	var data ResponseType
	err = p.DecodeResponse(rs, &data)
	if err != nil {
		return nil, err
	}
	item := data.Item()
	return &item, nil
}

// EndpointDeleter is a generic interface for deleting an item from an endpoint.
type EndpointDeleter interface {
	Delete(ctx context.Context, pk PrimaryKeyable) error
}

func NewEndpointDeleter(cbd CBD, endpoint string) EndpointDeleter {
	return &endpointDeleterImpl{cbd, endpoint}
}

type endpointDeleterImpl struct {
	CBD
	endpoint string
}

func (p *endpointDeleterImpl) Delete(ctx context.Context, pk PrimaryKeyable) error {
	rq, err := p.BuildRequest(ctx, http.MethodDelete, fmt.Sprintf("%s/%d/", p.endpoint, pk.PrimaryKey()), nil, nil)
	if err != nil {
		return err
	}
	rs, err := p.Do(rq)
	if err != nil {
		return err
	}
	defer rs.Body.Close()
	if rs.StatusCode != http.StatusOK {
		return ErrorFromResponse(rs)
	}
	return nil
}
