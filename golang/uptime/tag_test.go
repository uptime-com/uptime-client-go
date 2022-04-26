package uptime

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestTagList(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/check-tags", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"count": 1, "results": [{"pk": 1}]}`)
	})

	opt := &TagListOptions{Page: 1, PageSize: 1000}
	tags, _, err := client.Tags.List(context.Background(), opt)
	if err != nil {
		t.Errorf("Tags.List returned error: %v", err)
	}

	want := []*Tag{{PK: 1}}
	if !reflect.DeepEqual(tags, want) {
		t.Errorf("Checks.List returned %+v, want %+v", tags, want)
	}
}

func TestTagCreate(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := &Tag{
		Tag:      "test tag",
		ColorHex: "#000000",
	}

	mux.HandleFunc("/check-tags", func(w http.ResponseWriter, r *http.Request) {
		v := new(Tag)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"results": {"tag": "test tag", "color_hex": "#000000"}}`)
	})

	tag, _, err := client.Tags.Create(context.Background(), input)
	if err != nil {
		t.Errorf("Tags.Create returned error: %v", err)
	}

	want := &Tag{Tag: "test tag", ColorHex: "#000000"}
	if !reflect.DeepEqual(tag, want) {
		t.Errorf("Tags.Create returned %+v, want %+v", tag, want)
	}
}

func TestTagGet(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/check-tags/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"pk": 1}`)
	})

	tag, _, err := client.Tags.Get(context.Background(), 1)
	if err != nil {
		t.Errorf("Tags.Get returned error: %v", err)
	}

	want := &Tag{
		PK: 1,
	}
	if !reflect.DeepEqual(tag, want) {
		t.Errorf("Tags.Get returned %+v, want %+v", tag, want)
	}
}

func TestTagUpdate(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := &Tag{
		PK:       1,
		Tag:      "test tag",
		ColorHex: "#000000",
	}

	mux.HandleFunc("/check-tags/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")

		v := &Tag{}
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		fmt.Fprint(w, `{"results": {"pk": 1}}`)
	})

	tag, _, err := client.Tags.Update(context.Background(), input)
	if err != nil {
		t.Errorf("Tags.Update returned error: %v", err)
	}

	want := &Tag{PK: 1}
	if !reflect.DeepEqual(tag, want) {
		t.Errorf("Tags.Update returned %+v, want %+v", tag, want)
	}
}

func TestTagDelete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/check-tags/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Tags.Delete(context.Background(), 1)
	if err != nil {
		t.Errorf("Tags.Delete returned error: %v", err)
	}
}
