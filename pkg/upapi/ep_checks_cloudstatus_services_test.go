package upapi

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChecksEndpointListCloudStatusServices(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "/api/v1/checks/cloudstatus-services/", r.URL.Path)
		require.Equal(t, "1", r.URL.Query().Get("group"))
		_, _ = io.WriteString(w, `{
			"count": 2,
			"results": [
				{
					"id": 10,
					"name": "ec2-us-east-1",
					"title": "EC2",
					"sub_title": "US East (N. Virginia)",
					"group_id": 1,
					"group": "Amazon Web Services"
				},
				{
					"id": 11,
					"name": "s3-us-east-1",
					"title": "S3",
					"sub_title": "US East (N. Virginia)",
					"group_id": 1,
					"group": "Amazon Web Services"
				}
			]
		}`)
	}))

	cbd, err := WithBaseURL(srv.URL + "/api/v1/")(testCBD)
	require.NoError(t, err)

	ep := NewChecksEndpoint(cbd)

	result, err := ep.ListCloudStatusServices(ctx, CloudStatusServiceListOptions{Group: "1"})
	require.NoError(t, err)
	require.Equal(t, int64(2), result.TotalCount)
	require.Len(t, result.Items, 2)
	require.Equal(t, int64(10), result.Items[0].ID)
	require.Equal(t, "ec2-us-east-1", result.Items[0].Name)
	require.Equal(t, "EC2", result.Items[0].Title)
	require.Equal(t, "US East (N. Virginia)", result.Items[0].SubTitle)
	require.Equal(t, int64(1), result.Items[0].GroupID)
	require.Equal(t, "Amazon Web Services", result.Items[0].Group)
}

func TestChecksEndpointListCloudStatusServicesFilterByGroupName(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "/api/v1/checks/cloudstatus-services/", r.URL.Path)
		require.Equal(t, "Amazon", r.URL.Query().Get("group"))
		_, _ = io.WriteString(w, `{"count": 0, "results": []}`)
	}))

	cbd, err := WithBaseURL(srv.URL + "/api/v1/")(testCBD)
	require.NoError(t, err)

	ep := NewChecksEndpoint(cbd)

	result, err := ep.ListCloudStatusServices(ctx, CloudStatusServiceListOptions{Group: "Amazon"})
	require.NoError(t, err)
	require.Equal(t, int64(0), result.TotalCount)
	require.Empty(t, result.Items)
}
