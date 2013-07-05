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
