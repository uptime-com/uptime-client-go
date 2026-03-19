package upapi

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAccountUsageEndpoint_List(t *testing.T) {
	ctx := context.Background()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "/api/v1/auth/account-usage/", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `[{"name":"HTTP Checks","current":5,"max":100},{"name":"Group Checks","current":11,"max":11}]`)
	}))
	defer srv.Close()

	api, err := New(WithBaseURL(srv.URL+"/api/v1/"), WithToken("test"))
	require.NoError(t, err)

	result, err := api.AccountUsage().List(ctx)
	require.NoError(t, err)
	require.Equal(t, int64(2), result.TotalCount)
	require.Len(t, result.Items, 2)
	require.Equal(t, AccountUsageItem{Name: "HTTP Checks", Current: 5, Max: 100}, result.Items[0])
	require.Equal(t, AccountUsageItem{Name: "Group Checks", Current: 11, Max: 11}, result.Items[1])
}