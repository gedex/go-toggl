// Copyright 2013 The go-toggl AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package toggl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestTimeEntriesService_Create(t *testing.T) {
	setup()
	defer teardown()

	input := &TimeEntry{ID: 1}

	mux.HandleFunc("/time_entries", func(w http.ResponseWriter, r *http.Request) {
		v := new(TimeEntryCreate)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v.TimeEntry, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"data":{"id": 1}}`)
	})

	result, err := client.TimeEntries.Create(input)
	if err != nil {
		t.Errorf("TimeEntries.Create returned error: %v", err)
	}

	want := &TimeEntry{ID: 1}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("TimeEntries.Create returned %v, want %v", result, want)
	}
}

func TestTimeEntriesService_Start(t *testing.T) {
	setup()
	defer teardown()

	input := &TimeEntry{ID: 1}

	mux.HandleFunc("/time_entries/start", func(w http.ResponseWriter, r *http.Request) {
		v := new(TimeEntryCreate)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v.TimeEntry, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"data":{"id": 1}}`)
	})

	result, err := client.TimeEntries.Start(input)
	if err != nil {
		t.Errorf("TimeEntries.Start returned error: %v", err)
	}

	want := &TimeEntry{ID: 1}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("TimeEntries.Start returned %v, want %v", result, want)
	}
}

func TestTimeEntriesService_Stop(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/time_entries/1/stop", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		fmt.Fprint(w, `{"data":{"id": 1}}`)
	})

	result, err := client.TimeEntries.Stop(1)
	if err != nil {
		t.Errorf("TimeEntries.Stop returned error: %v", err)
	}

	want := &TimeEntry{ID: 1}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("TimeEntries.Stop returned %v, want %v", result, want)
	}
}

func TestTimeEntriesService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/time_entries/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data":{"id": 1}}`)
	})

	result, err := client.TimeEntries.Get(1)
	if err != nil {
		t.Errorf("TimeEntries.Get returned error: %v", err)
	}

	want := &TimeEntry{ID: 1}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("TimeEntries.Get returned %v, want %v", result, want)
	}
}

func TestTimeEntriesService_Update(t *testing.T) {
	setup()
	defer teardown()

	input := &TimeEntry{ID: 1}

	mux.HandleFunc("/time_entries/1", func(w http.ResponseWriter, r *http.Request) {
		v := new(TimeEntryCreate)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "PUT")
		if !reflect.DeepEqual(v.TimeEntry, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"data":{"id": 1}}`)
	})

	result, err := client.TimeEntries.Update(input)
	if err != nil {
		t.Errorf("TimeEntries.Update returned error: %v", err)
	}

	want := &TimeEntry{ID: 1}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("TimeEntries.Update returned %v, want %v", result, want)
	}
}

func TestTimeEntriesService_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/time_entries/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	err := client.TimeEntries.Delete(1)
	if err != nil {
		t.Errorf("TimeEntries.Delete returned error: %v", err)
	}
}

func TestTimeEntriesService_List(t *testing.T) {
	setup()
	defer teardown()

	start := time.Date(2013, time.July, 13, 0, 0, 0, 0, time.UTC)
	end := time.Date(2013, time.July, 15, 0, 0, 0, 0, time.UTC)

	mux.HandleFunc("/time_entries", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"start_date": start.Format(time.RFC3339),
			"end_date":   end.Format(time.RFC3339),
		})
		fmt.Fprint(w, `[{"id": 1}]`)
	})

	result, err := client.TimeEntries.List(&start, &end)
	if err != nil {
		t.Errorf("TimeEntries.List returned error: %v", err)
	}

	want := []TimeEntry{TimeEntry{ID: 1}}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("TimeEntries.List returned %v, want %v", result, want)
	}
}

func TestTimeEntriesService_Current(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/time_entries/current", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data": {"id": 1}}`)
	})

	result, err := client.TimeEntries.Current()
	if err != nil {
		t.Errorf("TimeEntries.Current returned error: %v", err)
	}

	want := &TimeEntry{ID: 1}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("TimeEntries.Current returned %v, want %v", result, want)
	}
}
