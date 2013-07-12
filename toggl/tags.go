// Copyright 2013 The go-toggl AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package toggl

import (
	"errors"
	"fmt"
)

// TagsService handles communication with the tags related
// methods of the Toggl API.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/tags.md
type TagsService struct {
	client *Client
}

// Tag represents a tag.
type Tag struct {
	ID          int    `json:"id,omitempty"`
	WorkspaceID int    `json:"wid,omitempty"`
	Name        string `json:"name,omitempty"`
}

// TagResponse acts as a response wrapper where response returns
// in format of "data": Tag's object.
type TagResponse struct {
	Data *Tag `json:"data,omitempty"`
}

// TagCreate represents posted data to be sent to clients endpoint.
type TagCreate struct {
	Tag *Tag `json:"tag,omitempty"`
}

// Create a new tag.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/tags.md#create-tag
func (s *TagsService) Create(t *Tag) (*Tag, error) {
	u := "tags"
	tc := &TagCreate{t}
	req, err := s.client.NewRequest("POST", u, tc)
	if err != nil {
		return nil, err
	}

	data := new(TagResponse)
	_, err = s.client.Do(req, data)

	return data.Data, err
}

// Update a tag.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/tags.md#update-a-tag
func (s *TagsService) Update(t *Tag) (*Tag, error) {
	if t == nil {
		return nil, errors.New("Tag cannot be nil")
	}
	if t.ID <= 0 {
		return nil, errors.New("Invalid Tag.ID")
	}

	u := fmt.Sprintf("tags/%v", t.ID)

	tc := &TagCreate{t}
	req, err := s.client.NewRequest("PUT", u, tc)
	if err != nil {
		return nil, err
	}

	data := new(TagResponse)
	_, err = s.client.Do(req, data)

	return data.Data, err
}

// Delete a tag.
//
// Toggl API docs: https://github.com/toggl/toggl_api_docs/blob/master/chapters/tags.md#delete-a-tag
func (s *TagsService) Delete(id int) error {
	u := fmt.Sprintf("tags/%v", id)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}
