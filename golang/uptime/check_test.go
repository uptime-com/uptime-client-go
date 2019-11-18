package uptime

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestCheckList(t *testing.T) {
	client, mux, _, teardown := setup()
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
	client, mux, _, teardown := setup()
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
	client, mux, _, teardown := setup()
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
	client, mux, _, teardown := setup()
	defer teardown()

	input := &Check{
		PK:        1,
		Name:      "Check Name",
		Locations: []string{"US-East"},
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
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/checks/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Checks.Delete(context.Background(), 1)
	if err != nil {
		t.Errorf("Checks.Delete returned error: %v", err)
	}
}
