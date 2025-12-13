package upapi

import (
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type cbdMock struct {
	mock.Mock
	Doer
	RequestBuilder
	ResponseDecoder
}

func (m *cbdMock) Do(rq *http.Request) (*http.Response, error) {
	args := m.Called(rq)
	if err := args.Error(1); err != nil {
		return nil, err
	}
	return args.Get(0).(*http.Response), nil
}

func (m *cbdMock) BuildRequest(ctx context.Context, method, path string, arg, data any) (*http.Request, error) {
	args := m.Called(ctx, method, path, arg, data)
	if err := args.Error(1); err != nil {
		return nil, err
	}
	return args.Get(0).(*http.Request), nil
}

func (m *cbdMock) DecodeResponse(r *http.Response, v interface{}) error {
	args := m.Called(r, v)
	return args.Error(0)
}

// TestBuildRequestSetsContentLength verifies that BuildRequest sets ContentLength
// when a body is provided. This is critical for HTTP/2 servers that may require
// Content-Length to be set for POST/PATCH requests.
// See: https://github.com/uptime-com/uptime-client-go/issues/40
func TestBuildRequestSetsContentLength(t *testing.T) {
	builder := &requestBuilderImpl{}

	type testData struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	data := testData{Name: "test", Value: 42}

	req, err := builder.BuildRequest(context.Background(), http.MethodPost, "/test", nil, data)
	require.NoError(t, err)
	require.NotNil(t, req)

	// The key assertion: ContentLength must be set
	require.Greater(t, req.ContentLength, int64(0), "ContentLength should be set when body is provided")

	// GetBody should also be set for retry support
	require.NotNil(t, req.GetBody, "GetBody should be set for retry support")

	// Verify body is readable and contains expected content
	body, err := io.ReadAll(req.Body)
	require.NoError(t, err)
	require.Contains(t, string(body), `"name":"test"`)
	require.Contains(t, string(body), `"value":42`)
}

// TestBuildRequestNoContentLengthForNilBody verifies ContentLength is 0 when no body
func TestBuildRequestNoContentLengthForNilBody(t *testing.T) {
	builder := &requestBuilderImpl{}

	req, err := builder.BuildRequest(context.Background(), http.MethodGet, "/test", nil, nil)
	require.NoError(t, err)
	require.NotNil(t, req)

	require.Equal(t, int64(0), req.ContentLength, "ContentLength should be 0 for GET without body")
	require.Nil(t, req.Body)
}
