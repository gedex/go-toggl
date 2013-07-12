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

func TestClientsService_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/clients", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id": 1}]`)
	})

	result, err := client.Clients.List()
	if err != nil {
		t.Errorf("Clients.List returned error: %v", err)
	}

	want := []WorkspaceClient{WorkspaceClient{ID: 1}}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Clients.List returned %v, want %v", result, want)
	}
}

func TestClientsService_ListClientProjects(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/clients/1/projects", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id": 1}]`)
	})

	result, err := client.Clients.ListClientProjects(1)
	if err != nil {
		t.Errorf("Clients.ListClientProjects returned error: %v", err)
	}

	want := []Project{Project{ID: 1}}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Clients.ListClientProjects returned %v, want %v", result, want)
	}
}

func TestClientsService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/clients/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data":{"id": 1}}`)
	})

	result, err := client.Clients.Get(1)
	if err != nil {
		t.Errorf("Clients.Get returned error: %v", err)
	}

	want := &WorkspaceClient{ID: 1}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Clients.Get returned %v, want %v", result, want)
	}
}

func TestClientsService_Create(t *testing.T) {
	setup()
	defer teardown()

	input := &WorkspaceClient{ID: 1}

	mux.HandleFunc("/clients", func(w http.ResponseWriter, r *http.Request) {
		v := new(WorkspaceClientCreate)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v.Client, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"data":{"id": 1}}`)
	})

	result, err := client.Clients.Create(input)
	if err != nil {
		t.Errorf("Clients.Create returned error: %v", err)
	}

	want := &WorkspaceClient{ID: 1}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Clients.Create returned %v, want %v", result, want)
	}
}

func TestClientsService_Update(t *testing.T) {
	setup()
	defer teardown()

	input := &WorkspaceClient{ID: 1, Name: "name"}

	mux.HandleFunc("/clients/1", func(w http.ResponseWriter, r *http.Request) {
		v := new(WorkspaceClientCreate)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "PUT")
		if !reflect.DeepEqual(v.Client, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"data":{"id": 1, "name": "name"}}`)
	})

	result, err := client.Clients.Update(input)
	if err != nil {
		t.Errorf("Clients.Update returned error: %v", err)
	}

	want := &WorkspaceClient{ID: 1, Name: "name"}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Clients.Update returned %v, want %v", result, want)
	}
}

func TestClientsService_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/clients/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	err := client.Clients.Delete(1)
	if err != nil {
		t.Errorf("Clients.Delete returned error: %v", err)
	}
}
