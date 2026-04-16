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

// TestChecks_EncryptionTriState verifies the *string semantics for
// msp_encryption across every check struct that exposes the field:
// nil pointer omits the field (server applies its default), pointer to ""
// sends an explicit Off, and pointer to "SSL_TLS" sends explicit TLS-on.
// Adding a new struct with msp_encryption? Add it to the table below.
func TestChecks_EncryptionTriState(t *testing.T) {
	type encCase struct {
		name string
		make func(enc *string) any
	}
	cases := []encCase{
		{"http", func(e *string) any { return CheckHTTP{Name: "test", Encryption: e} }},
		{"tcp", func(e *string) any { return CheckTCP{Name: "test", Encryption: e} }},
		{"imap", func(e *string) any { return CheckIMAP{Name: "test", Encryption: e} }},
		{"pop", func(e *string) any { return CheckPOP{Name: "test", Encryption: e} }},
		{"smtp", func(e *string) any { return CheckSMTP{Name: "test", Encryption: e} }},
		{"check", func(e *string) any { return Check{Name: "test", Encryption: e} }},
	}

	marshal := func(t *testing.T, v any) map[string]any {
		t.Helper()
		data, err := json.Marshal(v)
		if err != nil {
			t.Fatal(err)
		}
		return unmarshalToMap(t, data)
	}

	for _, c := range cases {
		t.Run(c.name+"/nil_pointer_omits_field", func(t *testing.T) {
			m := marshal(t, c.make(nil))
			if _, ok := m["msp_encryption"]; ok {
				t.Errorf("expected 'msp_encryption' to be omitted when pointer is nil")
			}
		})
		t.Run(c.name+"/empty_string_sends_off", func(t *testing.T) {
			empty := ""
			m := marshal(t, c.make(&empty))
			v, ok := m["msp_encryption"]
			if !ok {
				t.Fatalf("expected 'msp_encryption' to be present when pointer is non-nil")
			}
			if v != "" {
				t.Errorf("expected msp_encryption=\"\", got %v", v)
			}
		})
		t.Run(c.name+"/SSL_TLS_sends_tls_on", func(t *testing.T) {
			tls := "SSL_TLS"
			m := marshal(t, c.make(&tls))
			if v := m["msp_encryption"]; v != "SSL_TLS" {
				t.Errorf("expected msp_encryption=\"SSL_TLS\", got %v", v)
			}
		})
	}
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
