// pkg/upapi/ep_maintenance_notifications_test.go
package upapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseDjangoDuration(t *testing.T) {
	// Wire strings below are the exact forms observed from the live API.
	cases := map[string]int64{
		"00:00:00":    0,
		"01:00:00":    3600,
		"01:30:00":    5400,
		"-1 23:55:00": -300,
		"-1 23:30:00": -1800,
		"1 01:00:00":  90000,
		"3600":        3600, // defensive integer fallback
		"-1800":       -1800,
		"":            0,
	}
	for in, want := range cases {
		got, err := parseDjangoDuration(in)
		require.NoErrorf(t, err, "input %q", in)
		require.Equalf(t, want, got, "input %q", in)
	}
}

func TestMaintenanceNotificationDecode(t *testing.T) {
	const body = `{
		"id": 101,
		"schedule_id": 7,
		"offset": "-1 23:30:00",
		"event": "START",
		"contact_groups": [{"id": 10, "name": "On-Call"}],
		"created_at": "2026-06-15T14:35:00Z",
		"modified_at": "2026-06-15T14:35:00Z"
	}`
	var r MaintenanceNotificationResponse
	require.NoError(t, json.Unmarshal([]byte(body), &r))
	n := r.Item()
	require.Equal(t, int64(101), n.PK)
	require.Equal(t, int64(7), n.ScheduleID)
	require.Equal(t, int64(-1800), n.Offset) // parsed in Item()
	require.Equal(t, "START", n.Event)
	require.Len(t, n.ContactGroups, 1)
	require.Equal(t, int64(10), n.ContactGroups[0].PK)
}

func TestMaintenanceNotificationsList(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/maintenance/notifications/", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		_, _ = w.Write([]byte(`{"count":1,"next":null,"previous":null,"results":[{"id":101,"schedule_id":7,"offset":"-1 23:30:00","event":"START","contact_groups":[{"id":10,"name":"x"}]}]}`))
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	api, err := New(WithToken("t"), WithBaseURL(srv.URL+"/api/v1/"))
	require.NoError(t, err)

	res, err := api.MaintenanceNotifications().List(context.Background(), MaintenanceNotificationListOptions{})
	require.NoError(t, err)
	require.Len(t, res.Items, 1)
	require.Equal(t, int64(-1800), res.Items[0].Offset) // parsed on the list path
}

func TestMaintenanceNotificationsCRUD(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/maintenance/notifications/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var in MaintenanceScheduleNotificationInput
			require.NoError(t, json.NewDecoder(r.Body).Decode(&in))
			require.Equal(t, int64(7), in.ScheduleID)
			require.Equal(t, int64(-1800), in.Offset)
			require.Equal(t, "START", in.Event)
			require.Equal(t, []int64{10}, in.ContactGroups)
			// The generic endpoints require HTTP 200 on EVERY verb. Do NOT use 201/204.
			_, _ = w.Write([]byte(`{"results":{"id":101,"schedule_id":7,"offset":"-1 23:30:00","event":"START","contact_groups":[{"id":10,"name":"On-Call"}]}}`))
		default:
			t.Fatalf("unexpected method %s", r.Method)
		}
	})
	mux.HandleFunc("/api/v1/maintenance/notifications/101/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			_, _ = w.Write([]byte(`{"id":101,"schedule_id":7,"offset":"-1 23:30:00","event":"START","contact_groups":[{"id":10,"name":"On-Call"}]}`))
		case http.MethodPatch:
			var in MaintenanceScheduleNotificationInput
			require.NoError(t, json.NewDecoder(r.Body).Decode(&in))
			require.Equal(t, "END", in.Event)
			require.Equal(t, int64(5400), in.Offset)
			_, _ = w.Write([]byte(`{"results":{"id":101,"schedule_id":7,"offset":"01:30:00","event":"END","contact_groups":[{"id":10,"name":"On-Call"}]}}`))
		case http.MethodDelete:
			w.WriteHeader(http.StatusOK) // SDK deleter requires 200, not 204
		default:
			t.Fatalf("unexpected method %s", r.Method)
		}
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	api, err := New(WithToken("t"), WithBaseURL(srv.URL+"/api/v1/"))
	require.NoError(t, err)
	ep := api.MaintenanceNotifications()

	created, err := ep.Create(context.Background(), MaintenanceScheduleNotificationInput{
		ScheduleID: 7, Offset: -1800, Event: "START", ContactGroups: []int64{10},
	})
	require.NoError(t, err)
	require.Equal(t, int64(101), created.PK)
	require.Equal(t, int64(-1800), created.Offset)

	got, err := ep.Get(context.Background(), PrimaryKey(101))
	require.NoError(t, err)
	require.Equal(t, int64(-1800), got.Offset)
	require.Equal(t, "START", got.Event)

	upd, err := ep.Update(context.Background(), PrimaryKey(101), MaintenanceScheduleNotificationInput{
		ScheduleID: 7, Offset: 5400, Event: "END", ContactGroups: []int64{10},
	})
	require.NoError(t, err)
	require.Equal(t, int64(5400), upd.Offset)
	require.Equal(t, "END", upd.Event)

	require.NoError(t, ep.Delete(context.Background(), PrimaryKey(101)))
}
