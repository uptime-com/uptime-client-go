package uptime

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestIntegrationList(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/integrations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"count": 1, "results": [{"pk": 1}]}`)
	})

	opt := &IntegrationListOptions{Page: 1, PageSize: 1000}
	integrations, _, err := client.Integrations.List(context.Background(), opt)
	if err != nil {
		t.Errorf("Integrations.List returned error: %v", err)
	}

	want := []*Integration{{PK: 1}}
	if !reflect.DeepEqual(integrations, want) {
		t.Errorf("Integrations.List returned %+v, want %+v", integrations, want)
	}
}

func TestIntegrationCreate(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := &Integration{
		Name:          "test integration",
		APIEndpoint:   "https://api.opsgenie.com/v1/json/uptime",
		APIKey:        "53c84605-27b6-4a6d-a161-2dcb9bce6a4f",
		ContactGroups: []string{"Default"},
		Module:        "opsgenie",
	}

	mux.HandleFunc("/integrations/add-opsgenie", func(w http.ResponseWriter, r *http.Request) {
		v := new(Integration)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"results": {"module": "opsgenie"}}`)
	})

	integration, _, err := client.Integrations.Create(context.Background(), input)
	if err != nil {
		t.Errorf("Integrations.Create returned error: %v", err)
	}

	want := &Integration{Module: "opsgenie"}
	if !reflect.DeepEqual(integration, want) {
		t.Errorf("Integrations.Create returned %+v, want %+v", integration, want)
	}
}

func TestIntegrationGet(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/integrations/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"pk": 1}`)
	})

	integration, _, err := client.Integrations.Get(context.Background(), 1)
	if err != nil {
		t.Errorf("Integrations.Get returned error: %v", err)
	}

	want := &Integration{
		PK: 1,
	}
	if !reflect.DeepEqual(integration, want) {
		t.Errorf("Integrations.Get returned %+v, want %+v", integration, want)
	}
}

func TestIntegrationUpdate(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := &Integration{
		PK:            1,
		Name:          "Integration Name",
		ContactGroups: []string{"Default"},
	}

	mux.HandleFunc("/integrations/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")

		v := &Integration{}
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		fmt.Fprint(w, `{"results": {"pk": 1}}`)
	})

	integration, _, err := client.Integrations.Update(context.Background(), input)
	if err != nil {
		t.Errorf("Integrations.Update returned error: %v", err)
	}

	want := &Integration{PK: 1}
	if !reflect.DeepEqual(integration, want) {
		t.Errorf("Integrations.Update returned %+v, want %+v", integration, want)
	}
}

func TestIntegrationDelete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/integrations/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Integrations.Delete(context.Background(), 1)
	if err != nil {
		t.Errorf("Integration.Delete returned error: %v", err)
	}
}
