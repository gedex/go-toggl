// Copyright 2013 The go-toggl AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package toggl

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestWorkspacesService_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/workspaces", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprintf(w, `[{"id": 1}]`)
	})

	result, err := client.Workspaces.List()
	if err != nil {
		t.Errorf("Workspaces.List returned error: %v", err)
	}

	want := []Workspace{Workspace{ID: 1}}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Workspaces.List returned %v, want %v", result, want)
	}
}

func TestWorkspacesService_ListUsers(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/workspaces/1/users", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprintf(w, `[{"id": 1}]`)
	})

	result, err := client.Workspaces.ListUsers(1)
	if err != nil {
		t.Errorf("Workspaces.ListUsers returned error: %v", err)
	}

	want := []User{{ID: 1}}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Workspaces.ListUsers returned %v, want %v", result, want)
	}
}

func TestWorkspacesService_ListClients(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/workspaces/1/clients", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprintf(w, `[{"id": 1}]`)
	})

	result, err := client.Workspaces.ListClients(1)
	if err != nil {
		t.Errorf("Workspaces.ListClients returned error: %v", err)
	}

	want := []WorkspaceClient{{ID: 1}}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Workspaces.ListClients returned %v, want %v", result, want)
	}
}

func TestWorkspacesService_ListProjects(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/workspaces/1/projects", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprintf(w, `[{"id": 1}]`)
	})

	result, err := client.Workspaces.ListProjects(1, "")
	if err != nil {
		t.Errorf("Workspaces.ListProjects returned error: %v", err)
	}

	want := []Project{{ID: 1}}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Workspaces.ListProjects returned %v, want %v", result, want)
	}
}

func TestWorkspacesService_ListTasks(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/workspaces/1/tasks", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprintf(w, `[{"id": 1}]`)
	})

	result, err := client.Workspaces.ListTasks(1, "")
	if err != nil {
		t.Errorf("Workspaces.ListTasks returned error: %v", err)
	}

	want := []Task{{ID: 1}}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Workspaces.ListTasks returned %v, want %v", result, want)
	}
}
