package uptime

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestOutageList(t *testing.T) {
	client, mux, _, teardown := testSetup()
	defer teardown()

	mux.HandleFunc("/outages", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"count": 1, "results": [{"pk": 1}]}`)
	})

	opt := &OutageListOptions{Page: 1, PageSize: 1000, Search: "", Ordering: "", CheckMonitoringServiceType: ""}
	outages, _, err := client.Outages.List(context.Background(), opt)
	if err != nil {
		t.Errorf("Outages.List returned error: %v", err)
	}

	want := []*Outage{{PK: 1}}
	if !reflect.DeepEqual(outages, want) {
		t.Errorf("Outages.List returned %+v, want %+v", outages, want)
	}
}

func TestOutageGet(t *testing.T) {
	client, mux, _, teardown := testSetup()
	defer teardown()

	mux.HandleFunc("/outages/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"pk": 1, "all_alerts": [{"pk": 2}]}`)
	})

	o, _, err := client.Outages.Get(context.Background(), "1")
	if err != nil {
		t.Errorf("Outages.Get returned error: %v", err)
	}

	want := &Outage{
		PK: 1,
		AllAlerts: &[]Alert{{
			PK: 2,
		}},
	}
	if !reflect.DeepEqual(o, want) {
		t.Errorf("Outages.Get returned %+v, want %+v", o, want)
	}
}
