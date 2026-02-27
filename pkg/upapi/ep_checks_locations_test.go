package upapi

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChecksEndpointLocations(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "/api/v1/checks/locations/", r.URL.Path)
		_, _ = io.WriteString(w, `{"locations": ["US-NY-New York", "US-CA-Los Angeles", "United Kingdom-London"]}`)
	}))

	cbd, err := WithBaseURL(srv.URL + "/api/v1/")(testCBD)
	require.NoError(t, err)

	ep := NewChecksEndpoint(cbd)

	result, err := ep.ListLocations(ctx)
	require.NoError(t, err)
	require.Equal(t, int64(3), result.TotalCount)
	require.Equal(t, []string{
		"US-NY-New York",
		"US-CA-Los Angeles",
		"United Kingdom-London",
	}, result.Items)
}

func TestChecksEndpointLocationsEmpty(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "/api/v1/checks/locations/", r.URL.Path)
		_, _ = io.WriteString(w, `{"locations": []}`)
	}))

	cbd, err := WithBaseURL(srv.URL + "/api/v1/")(testCBD)
	require.NoError(t, err)

	ep := NewChecksEndpoint(cbd)

	result, err := ep.ListLocations(ctx)
	require.NoError(t, err)
	require.Equal(t, int64(0), result.TotalCount)
	require.Empty(t, result.Items)
}
