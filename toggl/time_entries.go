// Copyright 2013 The go-toggl AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package toggl

import (
	"errors"
	"fmt"
	"net/url"
	"time"
)

// TimeEntriesService handles communication with the time entries related
// methods of the Toggl API.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/time_entries.md
type TimeEntriesService struct {
	client *Client
}

// TimeEntry represents a time entry
type TimeEntry struct {
	ID          int        `json:"id,omitempty"`
	WorkspaceID int        `json:"wid,omitempty"`
	ProjectID   int        `json:"pid,omitempty"`
	TaskID      int        `json:"tid,omitempty"`
	Billable    bool       `json:"billable,omitempty"`
	Start       *time.Time `json:"start,omitempty"`
	Stop        *time.Time `json:"stop,omitempty"`
	Duration    int        `json:"duration,omitempty"`
	CreatedWith string     `json:"created_with,omitempty"`
	Tags        []string   `json:"tags,omitempty"`
	Duronly     bool       `json:"duronly,omitempty"`
	At          int        `json:"at,omitempty"`
}

// TimeEntryResponse acts as a response wrapper where response returns
// in format of "data": TimeEntry's object.
type TimeEntryResponse struct {
	Data *TimeEntry `json:"data,omitempty"`
}

// TimeEntryCreate represents posted data to be sent to time entries endpoint.
type TimeEntryCreate struct {
	TimeEntry *TimeEntry `json:"time_entry,omitempty"`
}

// Create a time entry.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/time_entries.md#create-a-time-entry
func (s *TimeEntriesService) Create(te *TimeEntry) (*TimeEntry, error) {
	u := "time_entries"
	tec := &TimeEntryCreate{te}
	req, err := s.client.NewRequest("POST", u, tec)
	if err != nil {
		return nil, err
	}

	data := new(TimeEntryResponse)
	_, err = s.client.Do(req, data)

	return data.Data, err
}

// Start a time entry.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/time_entries.md#start-a-time-entry
func (s *TimeEntriesService) Start(te *TimeEntry) (*TimeEntry, error) {
	u := "time_entries/start"
	tec := &TimeEntryCreate{te}
	req, err := s.client.NewRequest("POST", u, tec)
	if err != nil {
		return nil, err
	}

	data := new(TimeEntryResponse)
	_, err = s.client.Do(req, data)

	return data.Data, err
}

// Stop a time entry.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/time_entries.md#stop-a-time-entry
func (s *TimeEntriesService) Stop(id int) (*TimeEntry, error) {
	u := fmt.Sprintf("time_entries/%v/stop", id)

	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	data := new(TimeEntryResponse)
	_, err = s.client.Do(req, data)

	return data.Data, err
}

// Get time entry details.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/time_entries.md#get-time-entry-details
func (s *TimeEntriesService) Get(id int) (*TimeEntry, error) {
	u := fmt.Sprintf("time_entries/%v", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	data := new(TimeEntryResponse)
	_, err = s.client.Do(req, data)

	return data.Data, err
}

// Update a time entry.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/time_entries.md#update-a-time-entry
func (s *TimeEntriesService) Update(te *TimeEntry) (*TimeEntry, error) {
	if te == nil {
		return nil, errors.New("TimeEntry cannot be nil")
	}
	if te.ID <= 0 {
		return nil, errors.New("Invalid TimeEntry.ID")
	}

	u := fmt.Sprintf("time_entries/%v", te.ID)

	tec := &TimeEntryCreate{te}
	req, err := s.client.NewRequest("PUT", u, tec)
	if err != nil {
		return nil, err
	}

	data := new(TimeEntryResponse)
	_, err = s.client.Do(req, data)

	return data.Data, err
}

// Delete a time entry.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/time_entries.md#delete-a-time-entry
func (s *TimeEntriesService) Delete(id int) error {
	u := fmt.Sprintf("time_entries/%v", id)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}

// List time entries. With start and end parameters you can specify
// the date range of the time entries returned. If start and end
// are not specified, time entries started during the last 9 days
// are returned. start and end must be ISO 8601 date and time strings.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/time_entries.md#get-time-entries-started-in-a-specific-time-range
func (s *TimeEntriesService) List(start, end *time.Time) ([]TimeEntry, error) {
	u := "time_entries"
	params := url.Values{}
	if start != nil {
		params.Add("start_date", start.Format(time.RFC3339))
	}
	if end != nil {
		params.Add("end_date", end.Format(time.RFC3339))
	}
	u += "?" + params.Encode()

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	data := new([]TimeEntry)
	_, err = s.client.Do(req, data)

	return *data, err
}

// Get running time entry.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/time_entries.md#get-running-time-entry
func (s *TimeEntriesService) Current() (*TimeEntry, error) {
	u := "time_entries/current"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	data := new(TimeEntryResponse)
	_, err = s.client.Do(req, data)

	return data.Data, err
}
