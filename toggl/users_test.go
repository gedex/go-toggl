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

func TestUsersService_Me(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"data":{"id": 1}}`)
	})

	result, err := client.Users.Me(true)
	if err != nil {
		t.Errorf("Users.Me returned error: %v", err)
	}

	want := &User{ID: 1}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Users.Me returned %v, want %v", result, want)
	}
}

func TestUsersService_Signup(t *testing.T) {
	setup()
	defer teardown()

	input := &UserCredential{"email@example.com", "password"}

	mux.HandleFunc("/signups", func(w http.ResponseWriter, r *http.Request) {
		v := new(UserSignup)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v.User, input) {
			t.Errorf("Requst body = %+v, want %+v", v.User, input)
		}

		fmt.Fprint(w, `{"data":{"id": 1}}`)
	})

	newUser, err := client.Users.Signup(input)
	if err != nil {
		t.Errorf("Users.Signup returned error: %v", err)
	}

	want := &User{ID: 1}
	if !reflect.DeepEqual(newUser, want) {
		t.Errorf("Users.Signup returned %+v, want %+v", newUser, want)
	}
}
