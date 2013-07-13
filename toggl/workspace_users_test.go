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

func TestWorkspaceUsersService_Update(t *testing.T) {
	setup()
	defer teardown()

	input := &WorkspaceUser{ID: 1}

	mux.HandleFunc("/workspace_users/1", func(w http.ResponseWriter, r *http.Request) {
		v := new(WorkspaceUserCreate)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "PUT")
		if !reflect.DeepEqual(v.WorkspaceUser, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"data":{"id": 1}}`)
	})

	result, err := client.WorkspaceUsers.Update(input)
	if err != nil {
		t.Errorf("WorkspaceUsers.Update returned error: %v", err)
	}

	want := &WorkspaceUser{ID: 1}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("WorkspaceUsers.Update returned %v, want %v", result, want)
	}
}

func TestWorkspaceUsersService_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/workspace_users/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	err := client.WorkspaceUsers.Delete(1)
	if err != nil {
		t.Errorf("WorkspaceUsers.Delete returned error: %v", err)
	}
}
