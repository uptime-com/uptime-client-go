package upapi

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChecksEndpointListCloudStatusGroups(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "/api/v1/checks/cloudstatus-groups/", r.URL.Path)
		require.Equal(t, "aws", r.URL.Query().Get("search"))
		_, _ = io.WriteString(w, `{
			"count": 2,
			"results": [
				{"id": 1, "name": "Amazon Web Services"},
				{"id": 2, "name": "AWS Marketplace"}
			]
		}`)
	}))

	cbd, err := WithBaseURL(srv.URL + "/api/v1/")(testCBD)
	require.NoError(t, err)

	ep := NewChecksEndpoint(cbd)

	result, err := ep.ListCloudStatusGroups(ctx, CloudStatusGroupListOptions{Search: "aws"})
	require.NoError(t, err)
	require.Equal(t, int64(2), result.TotalCount)
	require.Len(t, result.Items, 2)
	require.Equal(t, int64(1), result.Items[0].ID)
	require.Equal(t, "Amazon Web Services", result.Items[0].Name)
	require.Equal(t, int64(2), result.Items[1].ID)
}

func TestChecksEndpointListCloudStatusGroupsEmpty(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "/api/v1/checks/cloudstatus-groups/", r.URL.Path)
		_, _ = io.WriteString(w, `{"count": 0, "results": []}`)
	}))

	cbd, err := WithBaseURL(srv.URL + "/api/v1/")(testCBD)
	require.NoError(t, err)

	ep := NewChecksEndpoint(cbd)

	result, err := ep.ListCloudStatusGroups(ctx, CloudStatusGroupListOptions{})
	require.NoError(t, err)
	require.Equal(t, int64(0), result.TotalCount)
	require.Empty(t, result.Items)
}
