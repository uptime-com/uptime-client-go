package upapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaintenanceScheduleDecode(t *testing.T) {
	const body = `{
		"id": 7,
		"name": "Weekly Patching",
		"schedule_type": "RRULE",
		"starts_at": "2026-06-20T02:00:00Z",
		"ends_at": "2026-09-20T02:00:00Z",
		"rrule": "FREQ=WEEKLY;BYDAY=SA",
		"duration_minutes": 120,
		"is_active": true,
		"pause_checks_during_maintenance": true,
		"services": [{"id": 42, "name": "API", "address": "api.example.com", "type": "HTTP"}],
		"tags": [{"id": 5, "tag": "prod", "color_hex": "#FF0000", "usage_count": 12}],
		"created_at": "2026-06-15T14:32:00Z",
		"modified_at": "2026-06-15T14:32:00Z"
	}`
	var m MaintenanceSchedule
	require.NoError(t, json.Unmarshal([]byte(body), &m))
	require.Equal(t, int64(7), m.PK)
	require.Equal(t, "RRULE", m.ScheduleType)
	require.Equal(t, int64(120), *m.DurationMinutes)
	require.Len(t, m.Services, 1)
	require.Equal(t, int64(42), m.Services[0].PK)
	require.Len(t, m.Tags, 1)
	require.Equal(t, int64(5), m.Tags[0].PK)
	require.Equal(t, PrimaryKey(7), m.PrimaryKey())
}

func TestMaintenanceSchedulesCRUD(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/maintenance/schedules/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var in MaintenanceScheduleInput
			require.NoError(t, json.NewDecoder(r.Body).Decode(&in))
			require.Equal(t, "ONE_OFF", in.ScheduleType)
			require.Equal(t, []int64{42}, in.Services)
			// The generic endpoints require HTTP 200 on EVERY verb (endpoint.go checks
			// rs.StatusCode != http.StatusOK and errors otherwise). Do NOT use 201/204.
			_, _ = w.Write([]byte(`{"results":{"id":7,"schedule_type":"ONE_OFF","is_active":true}}`))
		default:
			t.Fatalf("unexpected method %s", r.Method)
		}
	})
	mux.HandleFunc("/api/v1/maintenance/schedules/7/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			_, _ = w.Write([]byte(`{"id":7,"name":"x","schedule_type":"ONE_OFF","is_active":true}`))
		case http.MethodPatch:
			var in MaintenanceScheduleInput
			require.NoError(t, json.NewDecoder(r.Body).Decode(&in))
			require.False(t, in.IsActive)
			require.Equal(t, "y", in.Name)
			_, _ = w.Write([]byte(`{"results":{"id":7,"name":"y","schedule_type":"ONE_OFF","is_active":false}}`))
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
	ep := api.MaintenanceSchedules()

	created, err := ep.Create(context.Background(), MaintenanceScheduleInput{
		Name: "x", ScheduleType: "ONE_OFF", StartsAt: "2026-06-20T02:00:00Z",
		EndsAt: nil, IsActive: true, Services: []int64{42}, Tags: []int64{},
	})
	require.NoError(t, err)
	require.Equal(t, int64(7), created.PK)

	got, err := ep.Get(context.Background(), PrimaryKey(7))
	require.NoError(t, err)
	require.Equal(t, "x", got.Name)

	upd, err := ep.Update(context.Background(), PrimaryKey(7), MaintenanceScheduleInput{
		Name: "y", ScheduleType: "ONE_OFF", IsActive: false, Services: []int64{}, Tags: []int64{},
	})
	require.NoError(t, err)
	require.False(t, upd.IsActive)

	require.NoError(t, ep.Delete(context.Background(), PrimaryKey(7)))
}
