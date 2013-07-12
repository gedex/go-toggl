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

func TestProjectUsersService_Create(t *testing.T) {
	setup()
	defer teardown()

	input := &ProjectUser{ID: 1}

	mux.HandleFunc("/project_users", func(w http.ResponseWriter, r *http.Request) {
		v := new(ProjectUserCreate)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v.ProjectUser, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"data":{"id": 1}}`)
	})

	result, err := client.ProjectUsers.Create(input)
	if err != nil {
		t.Errorf("ProjectUsers.Create returned error: %v", err)
	}

	want := &ProjectUser{ID: 1}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("ProjectUsers.Create returned %v, want %v", result, want)
	}
}

func TestProjectUsersService_MassCreate(t *testing.T) {
	setup()
	defer teardown()

	input := &ProjectUserMultipleUserID{UserID: "1,2"}

	mux.HandleFunc("/project_users", func(w http.ResponseWriter, r *http.Request) {
		v := new(ProjectUserMassCreate)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v.ProjectUser, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"data":[{"uid": 1},{"uid": 2}]}`)
	})

	result, err := client.ProjectUsers.MassCreate(input)
	if err != nil {
		t.Errorf("ProjectUsers.MassCreate returned error: %v", err)
	}

	want := []ProjectUser{ProjectUser{UserID: 1}, ProjectUser{UserID: 2}}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("ProjectUsers.MassCreate returned %v, want %v", result, want)
	}
}

func TestProjectUsersService_Update(t *testing.T) {
	setup()
	defer teardown()

	input := &ProjectUser{ID: 1}

	mux.HandleFunc("/project_users/1", func(w http.ResponseWriter, r *http.Request) {
		v := new(ProjectUserCreate)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "PUT")
		if !reflect.DeepEqual(v.ProjectUser, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"data":{"id": 1}}`)
	})

	result, err := client.ProjectUsers.Update(input)
	if err != nil {
		t.Errorf("ProjectUsers.Update returned error: %v", err)
	}

	want := &ProjectUser{ID: 1}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("ProjectUsers.Update returned %v, want %v", result, want)
	}
}

func TestProjectUsersService_MassUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &ProjectUser{Manager: false}

	mux.HandleFunc("/project_users/1,2", func(w http.ResponseWriter, r *http.Request) {
		v := new(ProjectUser)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "PUT")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"data":[{"id": 1, "manager": false},{"id": 2, "manager": false}]}`)
	})

	result, err := client.ProjectUsers.MassUpdate("1,2", input)
	if err != nil {
		t.Errorf("ProjectUsers.MassUpdate returned error: %v", err)
	}

	want := []ProjectUser{{ID: 1, Manager: false}, {ID: 2, Manager: false}}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("ProjectUsers.MassUpdate returned %v, want %v", result, want)
	}
}

func TestProjectUsersService_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/project_users/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	err := client.ProjectUsers.Delete(1)
	if err != nil {
		t.Errorf("ProjectUsers.Delete returned error: %v", err)
	}
}

func TestProjectUsersService_MassDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/project_users/1,2", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	err := client.ProjectUsers.MassDelete("1,2")
	if err != nil {
		t.Errorf("ProjectUsers.MassDelete returned error: %v", err)
	}
}
