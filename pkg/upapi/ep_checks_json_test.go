package upapi

import (
	"encoding/json"
	"testing"
)

func TestCheckHTTP_OmitsZeroValues(t *testing.T) {
	check := CheckHTTP{
		Name: "test",
	}
	data, err := json.Marshal(check)
	if err != nil {
		t.Fatal(err)
	}
	m := unmarshalToMap(t, data)

	if _, ok := m["name"]; !ok {
		t.Error("expected 'name' to be present")
	}
	for _, field := range []string{"msp_address", "msp_status_code", "msp_encryption", "is_paused", "msp_include_in_global_metrics"} {
		if _, ok := m[field]; ok {
			t.Errorf("expected %q to be omitted for zero value, but it was present", field)
		}
	}
}

// TestCheckHTTP_EncryptionTriState verifies the *string semantics for
// msp_encryption: nil pointer omits the field (server applies its default),
// pointer to "" sends an explicit Off, pointer to "SSL_TLS" sends explicit
// TLS-on. This is the only way to distinguish "use server default" from
// "explicitly disable encryption".
func TestCheckHTTP_EncryptionTriState(t *testing.T) {
	t.Run("nil pointer omits field", func(t *testing.T) {
		data, err := json.Marshal(CheckHTTP{Name: "test", Encryption: nil})
		if err != nil {
			t.Fatal(err)
		}
		m := unmarshalToMap(t, data)
		if _, ok := m["msp_encryption"]; ok {
			t.Error("expected 'msp_encryption' to be omitted when pointer is nil")
		}
	})
	t.Run("pointer to empty string sends Off", func(t *testing.T) {
		empty := ""
		data, err := json.Marshal(CheckHTTP{Name: "test", Encryption: &empty})
		if err != nil {
			t.Fatal(err)
		}
		m := unmarshalToMap(t, data)
		v, ok := m["msp_encryption"]
		if !ok {
			t.Fatal("expected 'msp_encryption' to be present when pointer is non-nil")
		}
		if v != "" {
			t.Errorf("expected msp_encryption to be empty string, got %v", v)
		}
	})
	t.Run("pointer to SSL_TLS sends TLS-on", func(t *testing.T) {
		tls := "SSL_TLS"
		data, err := json.Marshal(CheckHTTP{Name: "test", Encryption: &tls})
		if err != nil {
			t.Fatal(err)
		}
		m := unmarshalToMap(t, data)
		if v := m["msp_encryption"]; v != "SSL_TLS" {
			t.Errorf("expected msp_encryption=\"SSL_TLS\", got %v", v)
		}
	})
}

func TestCheckHTTP_BoolPtrFalse(t *testing.T) {
	check := CheckHTTP{
		Name:     "test",
		IsPaused: BoolPtr(false),
	}
	data, err := json.Marshal(check)
	if err != nil {
		t.Fatal(err)
	}
	m := unmarshalToMap(t, data)

	v, ok := m["is_paused"]
	if !ok {
		t.Fatal("expected 'is_paused' to be present when set to false via pointer")
	}
	if v != false {
		t.Errorf("expected 'is_paused' to be false, got %v", v)
	}
}

func TestCheckHTTP_BoolPtrTrue(t *testing.T) {
	check := CheckHTTP{
		Name:                   "test",
		IsPaused:               BoolPtr(true),
		IncludeInGlobalMetrics: BoolPtr(true),
	}
	data, err := json.Marshal(check)
	if err != nil {
		t.Fatal(err)
	}
	m := unmarshalToMap(t, data)

	for _, field := range []string{"is_paused", "msp_include_in_global_metrics"} {
		v, ok := m[field]
		if !ok {
			t.Errorf("expected %q to be present", field)
			continue
		}
		if v != true {
			t.Errorf("expected %q to be true, got %v", field, v)
		}
	}
}

func TestCheckSSLCertConfig_BoolPtrFields(t *testing.T) {
	t.Run("omits nil bools", func(t *testing.T) {
		cfg := CheckSSLCertConfig{
			Protocol: "https",
		}
		data, err := json.Marshal(cfg)
		if err != nil {
			t.Fatal(err)
		}
		m := unmarshalToMap(t, data)

		for _, field := range []string{"ssl_cert_crl", "ssl_cert_first_element_only"} {
			if _, ok := m[field]; ok {
				t.Errorf("expected %q to be omitted when nil", field)
			}
		}
	})

	t.Run("includes false bools via pointer", func(t *testing.T) {
		cfg := CheckSSLCertConfig{
			CRL:              BoolPtr(false),
			FirstElementOnly: BoolPtr(false),
		}
		data, err := json.Marshal(cfg)
		if err != nil {
			t.Fatal(err)
		}
		m := unmarshalToMap(t, data)

		for _, field := range []string{"ssl_cert_crl", "ssl_cert_first_element_only"} {
			v, ok := m[field]
			if !ok {
				t.Errorf("expected %q to be present when set to false via pointer", field)
				continue
			}
			if v != false {
				t.Errorf("expected %q to be false, got %v", field, v)
			}
		}
	})
}

func TestCheckRDAP_SendResolvedNotifications(t *testing.T) {
	check := CheckRDAP{
		Name:                      "test",
		SendResolvedNotifications: BoolPtr(false),
	}
	data, err := json.Marshal(check)
	if err != nil {
		t.Fatal(err)
	}
	m := unmarshalToMap(t, data)

	v, ok := m["msp_send_resolved_notifications"]
	if !ok {
		t.Fatal("expected 'msp_send_resolved_notifications' to be present when set to false via pointer")
	}
	if v != false {
		t.Errorf("expected false, got %v", v)
	}
}

func TestBoolPtr(t *testing.T) {
	tr := BoolPtr(true)
	if tr == nil || *tr != true {
		t.Error("BoolPtr(true) should return pointer to true")
	}
	fa := BoolPtr(false)
	if fa == nil || *fa != false {
		t.Error("BoolPtr(false) should return pointer to false")
	}
	if tr == fa {
		t.Error("BoolPtr should return distinct pointers")
	}
}

func unmarshalToMap(t *testing.T, data []byte) map[string]any {
	t.Helper()
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}
	return m
}
