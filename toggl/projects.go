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

// ProjectsService handles communication with the projects related
// methods of the Toggl API.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/projects.md
type ProjectsService struct {
	client *Client
}

// Project represents project on a workspace.
type Project struct {
	ID          int        `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	WorkspaceID int        `json:"wid,omitempty"`
	ClientID    int        `json:"cid,omitempty"`
	Active      bool       `json:"active,omitempty"`
	IsPrivate   bool       `json:"is_private,omitempty"`
	Template    bool       `json:"template,omitempty"`
	TemplateID  int        `json:"template_id,omitempty"`
	Billable    bool       `json:"billable,omitempty"`
	At          *time.Time `json:"at,omitempty"`
}

// ProjectResponse acts as a response wrapper where response returns
// in format of "data": Project's object.
type ProjectResponse struct {
	Data *Project `json:"data,omitempty"`
}

// ProjectCreate represents posted data to be sent to projects endpoint.
type ProjectCreate struct {
	Project *Project `json:"project,omitempty"`
}

// Create a project.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/projects.md#create-project
func (s *ProjectsService) Create(p *Project) (*Project, error) {
	u := "projects"
	pc := &ProjectCreate{Project: p}
	req, err := s.client.NewRequest("POST", u, pc)
	if err != nil {
		return nil, err
	}

	data := new(ProjectResponse)
	_, err = s.client.Do(req, data)

	return data.Data, err
}

// Get project data.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/projects.md#get-project-data
func (s *ProjectsService) Get(id int) (*Project, error) {
	u := fmt.Sprintf("projects/%v", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	data := new(ProjectResponse)
	_, err = s.client.Do(req, data)

	return data.Data, err
}

// Update project data.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/clients.md#update-a-client
func (s *ProjectsService) Update(p *Project) (*Project, error) {
	if p == nil {
		return nil, errors.New("Project cannot be nil")
	}
	if p.ID <= 0 {
		return nil, errors.New("Invalid Project.ID")
	}

	u := fmt.Sprintf("projects/%v", p.ID)

	pc := &ProjectCreate{p}
	req, err := s.client.NewRequest("PUT", u, pc)
	if err != nil {
		return nil, err
	}

	data := new(ProjectResponse)
	_, err = s.client.Do(req, data)

	return data.Data, err
}

// ProjectUsers gets project users.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/projects.md#get-project-users
func (s *ProjectsService) ProjectUsers(id int) ([]ProjectUser, error) {
	u := fmt.Sprintf("projects/%v/project_users", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	data := new([]ProjectUser)
	_, err = s.client.Do(req, data)
	return *data, err
}
