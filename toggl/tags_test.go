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

func TestTagsService_Create(t *testing.T) {
	setup()
	defer teardown()

	input := &Tag{Name: "billed"}

	mux.HandleFunc("/tags", func(w http.ResponseWriter, r *http.Request) {
		v := new(TagCreate)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v.Tag, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"data":{"id": 1, "name": "billed"}}`)
	})

	result, err := client.Tags.Create(input)
	if err != nil {
		t.Errorf("Tags.Create returned error: %v", err)
	}

	want := &Tag{ID: 1, Name: "billed"}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Tags.Create returned %v, want %v", result, want)
	}
}

func TestTagsService_Update(t *testing.T) {
	setup()
	defer teardown()

	input := &Tag{ID: 1, Name: "not billed"}

	mux.HandleFunc("/tags/1", func(w http.ResponseWriter, r *http.Request) {
		v := new(TagCreate)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "PUT")
		if !reflect.DeepEqual(v.Tag, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"data":{"id": 1, "name": "not billed"}}`)
	})

	result, err := client.Tags.Update(input)
	if err != nil {
		t.Errorf("Tags.Update returned error: %v", err)
	}

	want := &Tag{ID: 1, Name: "not billed"}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Tags.Update returned %v, want %v", result, want)
	}
}

func TestTagsService_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/tags/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	err := client.Tags.Delete(1)
	if err != nil {
		t.Errorf("Tags.Delete returned error: %v", err)
	}
}
