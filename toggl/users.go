// Copyright 2013 The go-toggl AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package toggl

import (
	"time"
)

// UsersService handles communication with the user related
// methods of the Toggl API.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/users.md
type UsersService struct {
	client *Client
}

// UserResponse acts as a response wrapper where response returns
// in format of "data": User's object.
type UserResponse struct {
	Since int   `json:"since,omitempty"`
	Data  *User `json:"data,omitempty"`
}

// User represents Toggl's user.
type User struct {
	// User's ID
	ID int `json:"id,omitempty"`

	// User's fullname
	Fullname string `json:"fullname,omitempty"`

	// User API token in https://www.toggl.com/user/edit
	APIToken string `json:"api_token,omitempty"`

	// Default workspace ID
	DefautWID int `json:"default_wid,omitempty"`

	// User's email
	Email string `json:"email,omitempty"`

	// Time format like "h:mm A"
	TimeOfDayFormat string `json:"timeofday_format,omitempty"`

	// Date format like "MM/DD/YYYY"
	DateFormat string `json:"date_format,omitempty"`

	// Whether start and stop time are saved on time entry
	StoreStartAndStopTime bool `json:"store_start_and_stop_time,omitempty"`

	// Sunday=0
	BeginningOfWeek int `json:"beginning_of_week,omitempty"`

	// User's language such as "en_US"
	Language string `json:"language,omitempty"`

	// URL with the user's profile picture
	ImageURL string `json:"image_url,omitempty"`

	// Should a piechart be shown on the sidebar
	SidebarPiechart bool `json:"sidebar_piechart,omitempty"`

	// Timestamp of last changes, e.g. "2013-03-06T12:18:42+00:00"
	At *time.Time `json:"at,omitempty"`

	TimeEntries []TimeEntry `json:"time_entries,omitempty"`
	Projects    []Project   `json:"projects,omitempty"`
	Tags        []Tag       `json:"tags,omitempty"`
	// workspaces []
	// clients []
}

// UserSignup represents posted data to be sent to Signup endpoint.
type UserSignup struct {
	User *UserCredential `json:"user,omitempty"`
}

// UserCredential represents user credential to signup a new user.
type UserCredential struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

// Me returns current user data. If related data is set to true,
// a more complete response will be returned.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/users.md#get-current-user-data
func (s *UsersService) Me(withRelatedData bool) (*User, error) {
	u := "me"
	if withRelatedData {
		u += "?with_related_data=true"
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	data := new(UserResponse)
	_, err = s.client.Do(req, data)

	return data.Data, err
}

// Signup new user.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/users.md#sign-up-new-user
func (s *UsersService) Signup(uc *UserCredential) (*User, error) {
	u := "signups"
	us := &UserSignup{uc}
	req, err := s.client.NewRequest("POST", u, us)
	if err != nil {
		return nil, err
	}

	data := new(UserResponse)
	_, err = s.client.Do(req, data)

	return data.Data, err
}
