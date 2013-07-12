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

func TestProjectsService_Create(t *testing.T) {
	setup()
	defer teardown()

	input := &Project{Name: "name"}

	mux.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		v := new(ProjectCreate)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v.Project, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"data":{"id": 1, "name": "name"}}`)
	})

	result, err := client.Projects.Create(input)
	if err != nil {
		t.Errorf("Projects.Create returned error: %v", err)
	}

	want := &Project{ID: 1, Name: "name"}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Projects.Create returned %v, want %v", result, want)
	}
}

func TestProjectsService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/projects/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data":{"id": 1}}`)
	})

	result, err := client.Projects.Get(1)
	if err != nil {
		t.Errorf("Projects.Get returned error: %v", err)
	}

	want := &Project{ID: 1}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Projects.Get returned %v, want %v", result, want)
	}
}

func TestProjectsService_Update(t *testing.T) {
	setup()
	defer teardown()

	input := &Project{ID: 1, Name: "name"}

	mux.HandleFunc("/projects/1", func(w http.ResponseWriter, r *http.Request) {
		v := new(ProjectCreate)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "PUT")
		if !reflect.DeepEqual(v.Project, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"data":{"id": 1, "name": "name"}}`)
	})

	result, err := client.Projects.Update(input)
	if err != nil {
		t.Errorf("Projects.Update returned error: %v", err)
	}

	want := &Project{ID: 1, Name: "name"}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Projects.Update returned %v, want %v", result, want)
	}
}

func TestProjectsService_ProjectUsers(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/projects/1/project_users", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id": 1}]`)
	})

	result, err := client.Projects.ProjectUsers(1)
	if err != nil {
		t.Errorf("Projects.ProjectUsers returned error: %v", err)
	}

	want := []ProjectUser{ProjectUser{ID: 1}}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Projects.ProjectUsers returned %v, want %v", result, want)
	}
}
