// Copyright 2013 The go-toggl AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package toggl

import (
	"time"
)

// ProjectUsersService handles communication with the workspace related
// methods of the Toggl API.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/workspaces.md
type ProjectUsersService struct {
	client *Client
}

// ProjectUser represents association between project and user.
type ProjectUser struct {
	ID          int        `json:"id,omitempty"`
	ProjectID   int        `json:"pid,omitempty"`
	UserID      int        `json:"uid,omitempty"`
	WorkspaceID int        `json:"wid,omitempty"`
	Manager     bool       `json:"manager,omitempty"`
	Rate        float64    `json:"rate,omitempty"`
	At          *time.Time `json:"time,omitempty"`
}

// ProjectUserResponse acts as a response wrapper where response returns
// in format of "data": Project's object.
type ProjectUserResponse struct {
	Data *ProjectUser `json:"data,omitempty"`
}

// ProjectUserCreate represents posted data to be sent to project users endpoint.
type ProjectUserCreate struct {
	ProjectUser *ProjectUser `json:"project_user,omitempty"`
}

// Create a project user.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/project_users.md#create-a-project-user
func (s *ProjectUsersService) Create(pu *ProjectUser) (*ProjectUser, error) {
	u := "project_users"
	puc := &ProjectUserCreate{ProjectUser: pu}
	req, err := s.client.NewRequest("POST", u, puc)
	if err != nil {
		return nil, err
	}

	data := new(ProjectUserResponse)
	_, err = s.client.Do(req, data)

	return data.Data, err
}
