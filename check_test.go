package uptime

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCheckList(t *testing.T) {
	client, mux, _, teardown := testSetup()
	defer teardown()

	mux.HandleFunc("/checks", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"count": 1, "results": [{"pk": 1}]}`)
	})

	opt := &CheckListOptions{Page: 1, PageSize: 1000}
	checks, _, err := client.Checks.List(context.Background(), opt)
	if err != nil {
		t.Errorf("Checks.List returned error: %v", err)
	}

	want := []*Check{{PK: 1}}
	if !reflect.DeepEqual(checks, want) {
		t.Errorf("Checks.List returned %+v, want %+v", checks, want)
	}
}

func TestCheckCreate(t *testing.T) {
	client, mux, _, teardown := testSetup()
	defer teardown()

	input := &Check{
		CheckType:     "WHOIS",
		Address:       "uptime.com",
		Interval:      1,
		ContactGroups: []string{"Default"},
	}

	mux.HandleFunc("/checks/add-whois", func(w http.ResponseWriter, r *http.Request) {
		v := new(Check)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"results": {"msp_address": "http://www.uptime.com"}}`)
	})

	check, _, err := client.Checks.Create(context.Background(), input)
	if err != nil {
		t.Errorf("Checks.Create returned error: %v", err)
	}

	want := &Check{Address: "http://www.uptime.com"}
	if !reflect.DeepEqual(check, want) {
		t.Errorf("Checks.Create returned %+v, want %+v", check, want)
	}
}

func TestCheckGet(t *testing.T) {
	client, mux, _, teardown := testSetup()
	defer teardown()

	mux.HandleFunc("/checks/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"pk": 1}`)
	})

	check, _, err := client.Checks.Get(context.Background(), 1)
	if err != nil {
		t.Errorf("Checks.Get returned error: %v", err)
	}

	want := &Check{
		PK: 1,
	}
	if !reflect.DeepEqual(check, want) {
		t.Errorf("Checks.Get returned %+v, want %+v", check, want)
	}
}

func TestCheckUpdate(t *testing.T) {
	client, mux, _, teardown := testSetup()
	defer teardown()

	input := &Check{
		PK:        1,
		Name:      "Check Name",
		Locations: []string{"US East"},
	}

	mux.HandleFunc("/checks/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")

		v := &Check{}
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		fmt.Fprint(w, `{"results": {"pk": 1}}`)
	})

	check, _, err := client.Checks.Update(context.Background(), input)
	if err != nil {
		t.Errorf("Checks.Update returned error: %v", err)
	}

	want := &Check{PK: 1}
	if !reflect.DeepEqual(check, want) {
		t.Errorf("Checks.Update returned %+v, want %+v", check, want)
	}
}

func TestCheckDelete(t *testing.T) {
	client, mux, _, teardown := testSetup()
	defer teardown()

	mux.HandleFunc("/checks/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Checks.Delete(context.Background(), 1)
	if err != nil {
		t.Errorf("Checks.Delete returned error: %v", err)
	}
}

func TestCheckStats(t *testing.T) {
	client, mux, _, teardown := testSetup()
	defer teardown()
	pk := 1
	opt := &CheckStatsOptions{
		StartDate:              "2020-01-01T00:00:00Z",
		EndDate:                "2020-01-02T00:00:00Z",
		Location:               "US East",
		LocationsResponseTimes: true,
		IncludeAlerts:          true,
		Download:               false,
		PDF:                    false,
	}
	want := &CheckStatsResponse{
		StartDate: opt.StartDate,
		EndDate:   opt.EndDate,
		Statistics: []*CheckStats{
			{
				Date:         "2020-01-01",
				Outages:      0,
				DowntimeSecs: 0,
			},
		},
		Totals: CheckStatsTotals{
			Outages:      0,
			DowntimeSecs: 0,
		},
	}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		wantURL := fmt.Sprintf("/checks/%d/stats/", pk)
		gotURL := r.URL.EscapedPath()
		if !cmp.Equal(wantURL, gotURL) {
			t.Errorf("Request URL: %s", cmp.Diff(wantURL, gotURL))
		}
		wantValues := url.Values{
			"start_date":               []string{want.StartDate},
			"end_date":                 []string{want.EndDate},
			"location":                 []string{opt.Location},
			"locations_response_times": []string{"true"},
			"include_alerts":           []string{"true"},
			"download":                 []string{"false"},
			"pdf":                      []string{"false"},
		}
		if !cmp.Equal(wantValues, r.URL.Query()) {
			t.Errorf("Request URL: %s", cmp.Diff(wantValues, r.URL.Query()))
		}
		data, err := os.Open("testdata/stats.json")
		if err != nil {
			t.Fatal(err)
		}
		defer data.Close()
		io.Copy(w, data)
	})
	stats, _, err := client.Checks.Stats(context.Background(), pk, opt)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, stats) {
		t.Error(cmp.Diff(want, stats))
	}
}
