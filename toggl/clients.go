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

// ClientsService handles communication with the client related
// methods of the Toggl API.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/clients.md
type ClientsService struct {
	client *Client
}

// Client represents client of user's workspace.
type WorkspaceClient struct {
	ID          int        `json:"id,omitempty"`
	WorkspaceID int        `json:"wid,omitempty"`
	Name        string     `json:"name,omitempty"`
	Notes       string     `json:"notes,omitempty"`
	HourlyRate  float64    `json:"hrate,omitempty"`
	Currency    string     `json:"cur,omitempty"`
	At          *time.Time `json:"at,omitempty"` // indicates the time client was last updated
}

// WorkspaceClientResponse acts as a response wrapper where response returns
// in format of "data": Client's object.
type WorkspaceClientResponse struct {
	Data *WorkspaceClient `json:"data,omitempty"`
}

// WorkspaceClientCreate represents posted data to be sent to clients endpoint.
type WorkspaceClientCreate struct {
	Client *WorkspaceClient `json:"client,omitempty"`
}

// List visible clients to the user.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/clients.md#get-clients-visible-to-user
func (s *ClientsService) List() ([]WorkspaceClient, error) {
	u := "clients"

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	data := new([]WorkspaceClient)
	_, err = s.client.Do(req, data)

	return *data, err
}

// ListClientProjects lists client projects.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/clients.md#get-client-projects
func (s *ClientsService) ListClientProjects(id int) ([]Project, error) {
	u := fmt.Sprintf("clients/%v/projects", id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	data := new([]Project)
	_, err = s.client.Do(req, data)

	return *data, err
}

// Get client details by client_id.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/clients.md#get-client-details
func (s *ClientsService) Get(id int) (*WorkspaceClient, error) {
	u := fmt.Sprintf("clients/%v", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	data := new(WorkspaceClientResponse)
	_, err = s.client.Do(req, data)

	return data.Data, err
}

// Create a new client in specified workspace.
//
// https://github.com/toggl/toggl_api_docs/blob/master/chapters/clients.md#create-a-client
func (s *ClientsService) Create(c *WorkspaceClient) (*WorkspaceClient, error) {
	u := "clients"
	us := &WorkspaceClientCreate{c}
	req, err := s.client.NewRequest("POST", u, us)
	if err != nil {
		return nil, err
	}

	data := new(WorkspaceClientResponse)
	_, err = s.client.Do(req, data)

	return data.Data, err
}

// Update a client.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/clients.md#update-a-client
func (s *ClientsService) Update(c *WorkspaceClient) (*WorkspaceClient, error) {
	if c == nil {
		return nil, errors.New("WorkspaceClient cannot be nil")
	}
	if c.ID <= 0 {
		return nil, errors.New("Invalid WorkspaceClient.ID")
	}

	u := fmt.Sprintf("clients/%v", c.ID)

	us := &WorkspaceClientCreate{c}
	req, err := s.client.NewRequest("PUT", u, us)
	if err != nil {
		return nil, err
	}

	data := new(WorkspaceClientResponse)
	_, err = s.client.Do(req, data)

	return data.Data, err
}

// Delete a client.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/clients.md#delete-a-client
func (s *ClientsService) Delete(id int) error {
	u := fmt.Sprintf("clients/%v", id)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}
