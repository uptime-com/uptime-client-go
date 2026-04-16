package upapi

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCloudStatusGroup(t *testing.T) {
	t.Run("marshals as integer ID", func(t *testing.T) {
		g := CloudStatusGroup{ID: 3742, Name: "Amazon"}
		data, err := json.Marshal(&g)
		require.NoError(t, err)
		require.Equal(t, "3742", string(data))
	})
	t.Run("unmarshals from integer", func(t *testing.T) {
		var g CloudStatusGroup
		require.NoError(t, json.Unmarshal([]byte(`3742`), &g))
		require.Equal(t, int64(3742), g.ID)
		require.Empty(t, g.Name)
	})
	t.Run("unmarshals from object", func(t *testing.T) {
		var g CloudStatusGroup
		require.NoError(t, json.Unmarshal([]byte(`{"id": 3742, "name": "Amazon"}`), &g))
		require.Equal(t, int64(3742), g.ID)
		require.Equal(t, "Amazon", g.Name)
	})
	t.Run("unmarshals null", func(t *testing.T) {
		var g CloudStatusGroup
		require.NoError(t, json.Unmarshal([]byte(`null`), &g))
		require.Equal(t, int64(0), g.ID)
	})
	t.Run("round-trips inside CheckCloudStatusConfig response", func(t *testing.T) {
		body := []byte(`{
			"monitoring_type": "ALL",
			"group": {"id": 3742, "name": "Amazon"},
			"services": [],
			"service_titles": []
		}`)
		var cfg CheckCloudStatusConfig
		require.NoError(t, json.Unmarshal(body, &cfg))
		require.NotNil(t, cfg.Group)
		require.Equal(t, int64(3742), cfg.Group.ID)
		require.Equal(t, "Amazon", cfg.Group.Name)
		require.Equal(t, "ALL", cfg.MonitoringType)
	})
	t.Run("marshals request with numeric group ID", func(t *testing.T) {
		cfg := CheckCloudStatusConfig{
			Group:          &CloudStatusGroup{ID: 3742},
			MonitoringType: "ALL",
		}
		data, err := json.Marshal(cfg)
		require.NoError(t, err)
		require.JSONEq(t, `{"group":3742,"monitoring_type":"ALL"}`, string(data))
	})
}
