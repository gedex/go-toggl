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

// TasksService handles communication with the tasks related
// methods of the Toggl API.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/tasks.md
type TasksService struct {
	client *Client
}

// Task represents a task.
type Task struct {
	ID               int        `json:"id,omitempty"`
	Name             string     `json:"name,omitempty"`
	ProjectID        int        `json:"pid,omitempty"`
	UserID           int        `json:"uid,omitempty"`
	WorkspaceID      int        `json:"wid,omitempty"`
	EstimatedSeconds int        `json:"estimated_seconds,omitempty"`
	Active           bool       `json:"active,omitempty"`
	At               *time.Time `json:"at,omitempty"`
}

// TaskResponse acts as a response wrapper where response returns
// in format of "data": Tasks's object.
type TaskResponse struct {
	Data *Task `json:"data,omitempty"`
}

// TaskMassResponse acts as a response wrapper where response returns
// in format of "data": [ ... ].
type TaskMaskResponse struct {
	Data []Task `json:"data,omitempty"`
}

// TaskCreate represents posted data to be sent to task endpoint.
type TaskCreate struct {
	Task *Task `json:"task,omitempty"`
}

// Create a task.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/tasks.md#actions-for-single-project-user
func (s *TasksService) Create(t *Task) (*Task, error) {
	u := "tasks"
	tc := &TaskCreate{t}
	req, err := s.client.NewRequest("POST", u, tc)
	if err != nil {
		return nil, err
	}

	data := new(TaskResponse)
	_, err = s.client.Do(req, data)

	return data.Data, err
}

// Get task details by task_id.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/tasks.md#get-task-details
func (s *TasksService) Get(id int) (*Task, error) {
	u := fmt.Sprintf("tasks/%v", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	data := new(TaskResponse)
	_, err = s.client.Do(req, data)

	return data.Data, err
}

// Update a task.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/tasks.md#update-a-task
func (s *TasksService) Update(t *Task) (*Task, error) {
	if t == nil {
		return nil, errors.New("Task cannot be nil")
	}
	if t.ID <= 0 {
		return nil, errors.New("Invalid Task.ID")
	}

	u := fmt.Sprintf("tasks/%v", t.ID)

	tc := &TaskCreate{t}
	req, err := s.client.NewRequest("PUT", u, tc)
	if err != nil {
		return nil, err
	}

	data := new(TaskResponse)
	_, err = s.client.Do(req, data)

	return data.Data, err
}

// MassUpdate mass update tasks.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/tasks.md#update-multiple-tasks
func (s *TasksService) MassUpdate(ids string, t *Task) ([]Task, error) {
	u := fmt.Sprintf("tasks/%v", ids)

	tc := &TaskCreate{t}
	req, err := s.client.NewRequest("PUT", u, tc)
	if err != nil {
		return nil, err
	}

	data := new(TaskMaskResponse)
	_, err = s.client.Do(req, data)

	return data.Data, err
}

// Delete a task.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/tasks.md#delete-a-task
func (s *TasksService) Delete(id int) error {
	u := fmt.Sprintf("tasks/%v", id)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}

// MassDelete mass delete tasks.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/tasks.md#delete-multiple-tasks
func (s *TasksService) MassDelete(ids string) error {
	u := fmt.Sprintf("tasks/%v", ids)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}
