package upapi

import (
	"context"
	"fmt"
	"net/http"
)

type ItemGetter[Item any] interface {
	Item() *Item
}

type ListGetter[Item any] interface {
	List() []Item
}

type PrimaryKeyGetter interface {
	PrimaryKey() int
}

type PrimaryKey int

func (p PrimaryKey) PrimaryKey() int {
	return int(p)
}

// EndpointGetter is a generic interface for getting a single item from an endpoint.
type EndpointGetter[RequestType PrimaryKeyGetter, ResponseType ItemGetter[Item], Item any] interface {
	Get(ctx context.Context, arg RequestType) (*Item, error)
}

func NewEndpointGetter[RequestType PrimaryKeyGetter, ResponseType ItemGetter[ItemType], ItemType any](cbd CBD, endpoint string) EndpointGetter[RequestType, ResponseType, ItemType] {
	return &endpointGetterImpl[RequestType, ResponseType, ItemType]{cbd, endpoint}
}

type endpointGetterImpl[RequestType PrimaryKeyGetter, ResponseType ItemGetter[ItemType], ItemType any] struct {
	CBD
	endpoint string
}

func (p *endpointGetterImpl[RequestType, ResponseType, ItemType]) Get(ctx context.Context, arg RequestType) (*ItemType, error) {
	rq, err := p.BuildRequest(ctx, http.MethodGet, fmt.Sprintf("%s/%d/", p.endpoint, arg.PrimaryKey()), nil, nil)
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
	return data.Item(), nil
}

// EndpointLister is a generic interface for listing items from an endpoint.
type EndpointLister[ResponseType ListGetter[ItemType], ItemType any, OptionsType any] interface {
	List(ctx context.Context, opts OptionsType) ([]ItemType, error)
}

func NewEndpointLister[ResponseType ListGetter[ItemType], ItemType any, OptionsType any](cbd CBD, endpoint string) EndpointLister[ListGetter[ItemType], ItemType, OptionsType] {
	return &endpointListerImpl[ResponseType, ItemType, OptionsType]{cbd, endpoint}
}

type endpointListerImpl[ResponseType ListGetter[ItemType], ItemType any, OptionsType any] struct {
	CBD
	endpoint string
}

func (p *endpointListerImpl[ResponseType, ItemType, OptionsType]) List(ctx context.Context, opts OptionsType) ([]ItemType, error) {
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
	return data.List(), nil
}

// EndpointCreator is a generic interface for creating an item from an endpoint.
type EndpointCreator[RequestType any, ResponseType ItemGetter[ItemType], ItemType any] interface {
	Create(ctx context.Context, arg RequestType) (*ItemType, error)
}

func NewEndpointCreator[RequestType any, ResponseType ItemGetter[ItemType], ItemType any](cbd CBD, endpoint string) EndpointCreator[RequestType, ResponseType, ItemType] {
	return &endpointCreatorImpl[RequestType, ResponseType, ItemType]{cbd, endpoint}
}

type endpointCreatorImpl[RequestType any, ResponseType ItemGetter[ItemType], ItemType any] struct {
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
	return data.Item(), nil
}

// EndpointUpdater is a generic interface for updating an item from an endpoint.
type EndpointUpdater[RequestType PrimaryKeyGetter, ResponseType ItemGetter[ItemType], ItemType any] interface {
	Update(ctx context.Context, arg RequestType) (*ItemType, error)
}

func NewEndpointUpdater[RequestType PrimaryKeyGetter, ResponseType ItemGetter[ItemType], ItemType any](cbd CBD, endpoint string) EndpointUpdater[RequestType, ResponseType, ItemType] {
	return &endpointUpdaterImpl[RequestType, ResponseType, ItemType]{cbd, endpoint}
}

type endpointUpdaterImpl[RequestType PrimaryKeyGetter, ResponseType ItemGetter[ItemType], ItemType any] struct {
	CBD
	endpoint string
}

func (p *endpointUpdaterImpl[RequestType, ResponseType, ItemType]) Update(ctx context.Context, arg RequestType) (*ItemType, error) {
	rq, err := p.BuildRequest(ctx, http.MethodPatch, fmt.Sprintf("%s/%d/", p.endpoint, arg.PrimaryKey()), nil, arg)
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
	return data.Item(), nil
}

// EndpointDeleter is a generic interface for deleting an item from an endpoint.
type EndpointDeleter[RequestType PrimaryKeyGetter] interface {
	Delete(ctx context.Context, arg RequestType) error
}

func NewEndpointDeleter[RequestType PrimaryKeyGetter](cbd CBD, endpoint string) EndpointDeleter[RequestType] {
	return &endpointDeleterImpl[RequestType]{cbd, endpoint}
}

type endpointDeleterImpl[RequestType PrimaryKeyGetter] struct {
	CBD
	endpoint string
}

func (p *endpointDeleterImpl[RequestType]) Delete(ctx context.Context, arg RequestType) error {
	rq, err := p.BuildRequest(ctx, http.MethodDelete, fmt.Sprintf("%s/%d/", p.endpoint, arg.PrimaryKey()), nil, nil)
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
