// Copyright 2013 The go-toggl AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package toggl

import (
	"errors"
	"fmt"
)

// WorkspacesUsersService handles communication with the workspace users related
// methods of the Toggl API.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/workspace_users.md
type WorkspaceUsersService struct {
	client *Client
}

// WorkspaceUsers represents workspace user.
type WorkspaceUser struct {
	ID          int  `json:"id,omitempty"`
	UserID      int  `json:"uid,omitempty"`
	WorkspaceID int  `json:"wid,omitempty"`
	Admin       bool `json:"admin,omitempty"`
	Active      bool `json:"active,omitempty"`
}

// WorkspaceUserResponse acts as a response wrapper where response returns
// in format of "data": WorkspaceUser's object.
type WorkspaceUserResponse struct {
	Data *WorkspaceUser `json:"data,omitempty"`
}

// WorkspaceUserCreate represents posted data to be sent to workspace users endpoint.
type WorkspaceUserCreate struct {
	WorkspaceUser *WorkspaceUser `json:"workspace_user,omitempty"`
}

// Update workspace user. Only flag admin can be changed.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/workspace_users.md#update-workspace-user
func (s *WorkspaceUsersService) Update(wu *WorkspaceUser) (*WorkspaceUser, error) {
	if wu == nil {
		return nil, errors.New("WorkspacesUsers cannot be nil")
	}
	if wu.ID <= 0 {
		return nil, errors.New("Invalid WorkspacesUser.ID")
	}

	u := fmt.Sprintf("workspace_users/%v", wu.ID)

	wuc := &WorkspaceUserCreate{wu}
	req, err := s.client.NewRequest("PUT", u, wuc)
	if err != nil {
		return nil, err
	}

	data := new(WorkspaceUserResponse)
	_, err = s.client.Do(req, data)

	return data.Data, err
}

// Delete workspace user.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/workspace_users.md#delete-workspace-user
func (s *WorkspaceUsersService) Delete(id int) error {
	u := fmt.Sprintf("workspace_users/%v", id)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}
