// Copyright 2013 The go-toggl AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
TODO: Document this example header.
*/

package toggl

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	// LibraryVersion represents this library version
	LibraryVersion = "0.1"

	// BaseURL represents Toggl API base URL
	BaseURL = "https://toggl.com/api/v8/"

	// UserAgent represents this client User-Agent
	UserAgent = "go-toggl/" + LibraryVersion
)

// Client manages communication with the Toggl API.
type Client struct {
	// HTTP client used to communicate with the API
	client *http.Client

	// base64 encoded authorization header.
	basicAuth string

	// Base URL for API requests.
	BaseURL *url.URL

	// UserAgent agent used when communicating with Toggl API.
	UserAgent string

	// Services used for talking to differents parts of the API.
	Users *UsersService
	// TOOD:
	Clients      *ClientsService
	Projects     *ProjectsService
	ProjectUsers *ProjectUsersService
	Tags         *TagsService
	Tasks        *TasksService
	TimeEntries  *TimeEntriesService
	// Users *UsersService
	Workspaces *WorkspacesService
	// WorkspaceUsers *WorkspaceUsersService
}

// NewClient returns a new Toggl API client. Expects user's api token
// to be provided. Api token can be found in https://www.toggl.com/user/edit
func NewClient(apiToken string) *Client {
	baseURL, _ := url.Parse(BaseURL)
	basicAuth := base64.StdEncoding.EncodeToString([]byte(apiToken + ":api_token"))
	client := http.DefaultClient

	c := &Client{
		client:    client,
		basicAuth: basicAuth,
		BaseURL:   baseURL,
		UserAgent: UserAgent,
	}
	c.Users = &UsersService{client: c}
	c.Clients = &ClientsService{client: c}
	c.Projects = &ProjectsService{client: c}
	c.ProjectUsers = &ProjectUsersService{client: c}
	c.Tags = &TagsService{client: c}
	c.Tasks = &TasksService{client: c}
	c.TimeEntries = &TimeEntriesService{client: c}
	// c.Users = &UsersService{client: c}
	c.Workspaces = &WorkspacesService{client: c}
	// c.WorkspaceUsers = &WorkspaceUsersService{client: c}

	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash.
// If specified, the value pointed to by body is JSON encoded and included
// as the request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	ref, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(ref)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.UserAgent)
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", c.basicAuth))

	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// decoded and stored in the value pointed by v, or returned as an error if
// and API error has occurred.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}
	return resp, err
}

// CheckResponse checks the API response for error, and returns the error
// if present. A response is considered an error if it has a status code outside
// the 200 range.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	message, _ := ioutil.ReadAll(r.Body)

	return fmt.Errorf("%v %v: %d %v", r.Request.Method, r.Request.URL, r.StatusCode, string(message))
}
