// Copyright 2013 The go-toggl AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package toggl

import (
	"fmt"
	"time"
)

// WorkspacesService handles communication with the workspace related
// methods of the Toggl API.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/workspaces.md
type WorkspacesService struct {
	client *Client
}

// Workspace represents workspace of Toggl's user.
type Workspace struct {
	ID      int        `json:"id,omitempty"`
	Name    string     `json:"name,omitempty"`
	Premium bool       `json:"premium,omitempty"`
	At      *time.Time `json:"at,omitempty"`
}

// List user's workspace.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/workspaces.md#get-workspaces
func (s *WorkspacesService) List() ([]Workspace, error) {
	u := "workspaces"

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	data := new([]Workspace)
	_, err = s.client.Do(req, data)

	return *data, err
}

// ListUsers returns list of users on specified workspace id.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/workspaces.md#get-workspace-users
func (s *WorkspacesService) ListUsers(id int) ([]User, error) {
	u := fmt.Sprintf("workspaces/%v/users", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	data := new([]User)
	_, err = s.client.Do(req, data)

	return *data, err
}

// ListClients returns list of clients on specified workspace id.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/workspaces.md#get-workspace-clients
func (s *WorkspacesService) ListClients(id int) ([]WorkspaceClient, error) {
	u := fmt.Sprintf("workspaces/%v/clients", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	data := new([]WorkspaceClient)
	_, err = s.client.Do(req, data)

	return *data, err
}

// TODO:
// Workspaces.ListProjects
// Workspaces.ListTasks
