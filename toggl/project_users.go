// Copyright 2013 The go-toggl AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package toggl

import (
	"errors"
	"fmt"
	"time"
)

// ProjectUsersService handles communication with the project_users related
// methods of the Toggl API.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/project_users.md
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
	At          *time.Time `json:"at,omitempty"`
}

// ProjectUserMultipleUserID represents a project user where UID is a string which can hold
// multiple IDs separated by comma.
type ProjectUserMultipleUserID struct {
	ID          int        `json:"id,omitempty"`
	ProjectID   int        `json:"pid,omitempty"`
	UserID      string     `json:"uid,omitempty"`
	WorkspaceID int        `json:"wid,omitempty"`
	Manager     bool       `json:"manager,omitempty"`
	Rate        float64    `json:"rate,omitempty"`
	At          *time.Time `json:"at,omitempty"`
}

// ProjectUserResponse acts as a response wrapper where response returns
// in format of "data": ProjectUser's object.
type ProjectUserResponse struct {
	Data *ProjectUser `json:"data,omitempty"`
}

// ProjectUserMassResponse acts as a response wrapper where response returns
// in format of "data": [ ... ].
type ProjectUserMassResponse struct {
	Data []ProjectUser `json:"data,omitempty"`
}

// ProjectUserCreate represents posted data to be sent to project users endpoint.
type ProjectUserCreate struct {
	ProjectUser *ProjectUser `json:"project_user,omitempty"`
}

// ProjectUserMassCreate represents posted data to create multipe project users
// for single project.
type ProjectUserMassCreate struct {
	ProjectUser *ProjectUserMultipleUserID `json:"project_user,omitempty"`
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

// MassCreate creates multiple project users for single project.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/project_users.md#create-multiple-project-users-for-single-project
func (s *ProjectUsersService) MassCreate(pu *ProjectUserMultipleUserID) ([]ProjectUser, error) {
	u := "project_users"
	puc := &ProjectUserMassCreate{ProjectUser: pu}
	req, err := s.client.NewRequest("POST", u, puc)
	if err != nil {
		return nil, err
	}

	data := new(ProjectUserMassResponse)
	_, err = s.client.Do(req, data)

	return data.Data, err
}

// Update a project user.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/project_users.md#update-a-project-user
func (s *ProjectUsersService) Update(pu *ProjectUser) (*ProjectUser, error) {
	if pu == nil {
		return nil, errors.New("ProjectUser cannot be nil")
	}
	if pu.ID <= 0 {
		return nil, errors.New("Invalid ProjectUser.ID")
	}

	u := fmt.Sprintf("project_users/%v", pu.ID)

	puc := &ProjectUserCreate{pu}
	req, err := s.client.NewRequest("PUT", u, puc)
	if err != nil {
		return nil, err
	}

	data := new(ProjectUserResponse)
	_, err = s.client.Do(req, data)

	return data.Data, err
}

// MassUpdate mass update project users.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/project_users.md#mass-update-for-project-users
func (s *ProjectUsersService) MassUpdate(pids string, pu *ProjectUser) ([]ProjectUser, error) {
	u := fmt.Sprintf("project_users/%v", pids)

	puc := &ProjectUserCreate{pu}
	req, err := s.client.NewRequest("PUT", u, puc)
	if err != nil {
		return nil, err
	}

	data := new(ProjectUserMassResponse)
	_, err = s.client.Do(req, data)

	return data.Data, err
}

// Delete a project user.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/project_users.md#delete-a-project-user
func (s *ProjectUsersService) Delete(id int) error {
	u := fmt.Sprintf("project_users/%v", id)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}

// MassDelete a project user.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/project_users.md#delete-multiple-project-users
func (s *ProjectUsersService) MassDelete(pids string) error {
	u := fmt.Sprintf("project_users/%v", pids)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}
