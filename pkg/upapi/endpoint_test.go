package upapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hashicorp/go-cleanhttp"
	"github.com/stretchr/testify/require"
)

type testItem struct {
	Foo string `json:"foo"`
	Bar int    `json:"bar"`
	Baz bool   `json:"baz"`
}

func (i testItem) PrimaryKey() PrimaryKey {
	return 1
}

type testItemResponse struct {
	Results *testItem `json:"results"`
}

func (r testItemResponse) Item() *testItem {
	return r.Results
}

type testListResponse struct {
	Count   int        `json:"count"`
	Results []testItem `json:"results"`
}

func (r testListResponse) List() []testItem {
	return r.Results
}

type testOptions struct {
	Page int `url:"page,omitempty"`
}

var testCBD CBD = struct {
	Doer
	RequestBuilder
	ResponseDecoder
}{
	cleanhttp.DefaultClient(),
	&requestBuilderImpl{},
	&responseDecoderImpl{},
}

func TestEndpointGetter(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "/api/v1/items/1/", r.URL.Path)
		_, _ = io.WriteString(w, `{"results": {"foo": "foo", "bar": 1, "baz": true}}`)
	}))

	cbd, err := WithBaseURL(srv.URL + "/api/v1/")(testCBD)
	require.NoError(t, err)

	ep := NewEndpointGetter[PrimaryKey, testItemResponse, testItem](cbd, "items")

	obj, err := ep.Get(ctx, PrimaryKey(1))
	require.NoError(t, err)
	require.Equal(t, &testItem{"foo", 1, true}, obj)
}

func TestEndpointLister(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "/api/v1/items/", r.URL.Path)
		_, _ = io.WriteString(w, `
		{
			"count": 2,
			"results": [
				{"foo": "foo1", "bar": 1, "baz": true},
				{"foo": "foo2", "bar": 2, "baz": false}
			]
		}`)
	}))
	c, err := WithBaseURL(srv.URL + "/api/v1/")(testCBD)
	require.NoError(t, err)

	ep := NewEndpointLister[testListResponse, testItem, testOptions](c, "items")

	list, err := ep.List(ctx, testOptions{Page: 100500})
	require.NoError(t, err)
	require.Len(t, list, 2)
	require.Equal(t, testItem{"foo1", 1, true}, list[0])
	require.Equal(t, testItem{"foo2", 2, false}, list[1])
}

func TestEndpointCreator(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/api/v1/items/", r.URL.Path)
		buf := bytes.NewBuffer(nil)
		_, err := buf.ReadFrom(r.Body)
		require.NoError(t, err)
		require.Equal(t, `{"foo":"foo","bar":1,"baz":true}`, strings.TrimSpace(buf.String()))
		_, _ = io.WriteString(w, `{"results": {"foo": "foo", "bar": 1, "baz": true}}`)
	}))

	c, err := WithBaseURL(srv.URL + "/api/v1/")(testCBD)
	require.NoError(t, err)

	ep := NewEndpointCreator[testItem, testItemResponse, testItem](c, "items")

	obj, err := ep.Create(ctx, testItem{"foo", 1, true})
	require.NoError(t, err)
	require.Equal(t, &testItem{"foo", 1, true}, obj)
}

func TestEndpointUpdater(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		require.Equal(t, http.MethodPatch, r.Method)
		require.Equal(t, "/api/v1/items/1/", r.URL.Path)
		buf := bytes.NewBuffer(nil)
		_, err := buf.ReadFrom(r.Body)
		require.NoError(t, err)
		require.Equal(t, `{"foo":"foo","bar":1,"baz":true}`, strings.TrimSpace(buf.String()))
		_, _ = io.WriteString(w, `{"results": {"foo": "foo", "bar": 1, "baz": true}}`)
	}))

	c, err := WithBaseURL(srv.URL + "/api/v1/")(testCBD)
	require.NoError(t, err)

	ep := NewEndpointUpdater[testItem, testItemResponse, testItem](c, "items")

	obj, err := ep.Update(ctx, testItem{"foo", 1, true})
	require.NoError(t, err)
	require.Equal(t, &testItem{"foo", 1, true}, obj)
}

func TestEndpointDeleter(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		require.Equal(t, "/api/v1/items/1/", r.URL.Path)
	}))

	c, err := WithBaseURL(srv.URL + "/api/v1/")(testCBD)
	require.NoError(t, err)

	ep := NewEndpointDeleter[PrimaryKey](c, "items")

	err = ep.Delete(ctx, PrimaryKey(1))
	require.NoError(t, err)
}
