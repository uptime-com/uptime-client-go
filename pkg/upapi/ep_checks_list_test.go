package upapi

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChecksList_IsPausedIsOnlySentWhenSet(t *testing.T) {
	var queries []string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		queries = append(queries, r.URL.RawQuery)
		_, _ = io.WriteString(w, `{"count": 0, "results": []}`)
	}))
	defer srv.Close()
	api, err := New(WithBaseURL(srv.URL + "/api/v1/"))
	require.NoError(t, err)

	for _, opts := range []CheckListOptions{
		{Page: 1},
		{Page: 1, IsPaused: BoolPtr(false)},
		{Page: 1, IsPaused: BoolPtr(true)},
	} {
		_, err := api.Checks().List(context.Background(), opts)
		require.NoError(t, err)
	}

	require.Equal(t, []string{"page=1", "is_paused=false&page=1", "is_paused=true&page=1"}, queries)
}
