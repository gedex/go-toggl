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
)

func TestTasksService_Create(t *testing.T) {
	setup()
	defer teardown()

	input := &Task{ID: 1}

	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		v := new(TaskCreate)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v.Task, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"data":{"id": 1}}`)
	})

	result, err := client.Tasks.Create(input)
	if err != nil {
		t.Errorf("Tasks.Create returned error: %v", err)
	}

	want := &Task{ID: 1}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Tasks.Create returned %v, want %v", result, want)
	}
}

func TestTasksService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/tasks/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data":{"id": 1}}`)
	})

	result, err := client.Tasks.Get(1)
	if err != nil {
		t.Errorf("Tasks.Get returned error: %v", err)
	}

	want := &Task{ID: 1}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Tasks.Get returned %v, want %v", result, want)
	}
}

func TestTasksService_Update(t *testing.T) {
	setup()
	defer teardown()

	input := &Task{ID: 1}

	mux.HandleFunc("/tasks/1", func(w http.ResponseWriter, r *http.Request) {
		v := new(TaskCreate)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "PUT")
		if !reflect.DeepEqual(v.Task, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"data":{"id": 1}}`)
	})

	result, err := client.Tasks.Update(input)
	if err != nil {
		t.Errorf("Tasks.Update returned error: %v", err)
	}

	want := &Task{ID: 1}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Tasks.Update returned %v, want %v", result, want)
	}
}

func TestTasksService_MassUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &Task{Active: false}

	mux.HandleFunc("/tasks/1,2", func(w http.ResponseWriter, r *http.Request) {
		v := new(Task)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "PUT")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"data":[{"id": 1, "active": false},{"id": 2, "active": false}]}`)
	})

	result, err := client.Tasks.MassUpdate("1,2", input)
	if err != nil {
		t.Errorf("Tasks.MassUpdate returned error: %v", err)
	}

	want := []Task{{ID: 1, Active: false}, {ID: 2, Active: false}}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Tasks.MassUpdate returned %v, want %v", result, want)
	}
}

func TestTasksService_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/tasks/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	err := client.Tasks.Delete(1)
	if err != nil {
		t.Errorf("Tasks.Delete returned error: %v", err)
	}
}

func TestTasksService_MassDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/tasks/1,2", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	err := client.Tasks.MassDelete("1,2")
	if err != nil {
		t.Errorf("Tasks.MassDelete returned error: %v", err)
	}
}
