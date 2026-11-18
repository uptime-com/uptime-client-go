package upapi

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

const alertDownJSON = `{"pk":8720057,"url":"https://example.test/api/v1/alerts/8720057/","created_at":"2026-11-18T14:22:54.180447Z",
"ignore_alert_url":"https://example.test/api/v1/alerts/8720057/ignore/","check_pk":90726,"check_url":"https://example.test/api/v1/checks/90726/",
"check_address":"flakybin.org","check_name":"flakybin (HTTP) #239","check_monitoring_service_type":"HTTP","check_is_paused":true,
"state_is_up":false,"ignored":false,"num_locations_down":1,
"all_alerts":[{"pk":8652157,"url":"https://example.test/api/v1/alerts/alert/8652157/root-cause/","created_at":"2026-11-18T14:22:50Z",
"monitoring_server_name":"US-New York(Equinix)-2","monitoring_server_ipv4":"157.245.253.193",
"monitoring_server_ipv6":"2604:a880:400:d0::4a5f:5001","location":"US-NY-New York","output":"HTTP CRITICAL"}]}`

const alertUpJSON = `{"pk":8720058,"state_is_up":true,"num_locations_down":0,"all_alerts":[]}`

func requireDownAlert(t *testing.T, a AlertItem) {
	t.Helper()
	require.Equal(t, int64(8720057), a.PK)
	require.False(t, a.StateIsUp)
	require.True(t, a.CheckIsPaused)
	require.Equal(t, int64(1), a.NumLocationsDown)
	require.Equal(t, "https://example.test/api/v1/alerts/8720057/ignore/", a.IgnoreAlertURL)
	require.Len(t, a.AllAlerts, 1)
	loc := a.AllAlerts[0]
	require.Equal(t, int64(8652157), loc.PK)
	require.Equal(t, "US-NY-New York", loc.Location)
	require.Equal(t, "US-New York(Equinix)-2", loc.MonitoringServerName)
	require.Equal(t, "157.245.253.193", loc.MonitoringServerIPv4.String())
	require.Equal(t, "2604:a880:400:d0::4a5f:5001", loc.MonitoringServerIPv6.String())
	require.Equal(t, "HTTP CRITICAL", loc.Output)
}

func newAlertsTestAPI(t *testing.T, path, body string) API {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, path, r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, body)
	}))
	t.Cleanup(srv.Close)
	api, err := New(WithBaseURL(srv.URL+"/api/v1/"), WithToken("test"))
	require.NoError(t, err)
	return api
}

func TestAlertsEndpoint_List(t *testing.T) {
	body := `{"next":null,"previous":null,"results":[` + alertDownJSON + `,` + alertUpJSON + `]}`
	api := newAlertsTestAPI(t, "/api/v1/alerts/", body)

	result, err := api.Alerts().List(context.Background(), AlertListOptions{})
	require.NoError(t, err)
	require.Len(t, result.Items, 2)
	requireDownAlert(t, result.Items[0])
	require.True(t, result.Items[1].StateIsUp)
	require.Empty(t, result.Items[1].AllAlerts)
}

func TestAlertsEndpoint_Get(t *testing.T) {
	t.Run("down", func(t *testing.T) {
		api := newAlertsTestAPI(t, "/api/v1/alerts/8720057/", alertDownJSON)
		a, err := api.Alerts().Get(context.Background(), PrimaryKey(8720057))
		require.NoError(t, err)
		requireDownAlert(t, *a)
	})
	t.Run("up", func(t *testing.T) {
		api := newAlertsTestAPI(t, "/api/v1/alerts/8720058/", alertUpJSON)
		a, err := api.Alerts().Get(context.Background(), PrimaryKey(8720058))
		require.NoError(t, err)
		require.True(t, a.StateIsUp)
		require.Empty(t, a.AllAlerts)
	})
}
