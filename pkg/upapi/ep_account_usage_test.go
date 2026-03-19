package upapi

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAccountUsageEndpoint_Get(t *testing.T) {
	ctx := context.Background()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "/api/v1/auth/account-usage/", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `[{"Account":"Test","Checks Used":18,"Checks Allocated":100,"Content Matching Available":true}]`)
	}))
	defer srv.Close()

	api, err := New(WithBaseURL(srv.URL+"/api/v1/"), WithToken("test"))
	require.NoError(t, err)

	result, err := api.AccountUsage().Get(ctx)
	require.NoError(t, err)
	require.Equal(t, "Test", (*result)["Account"])
	require.Equal(t, float64(18), (*result)["Checks Used"])
	require.Equal(t, float64(100), (*result)["Checks Allocated"])
	require.Equal(t, true, (*result)["Content Matching Available"])
}
