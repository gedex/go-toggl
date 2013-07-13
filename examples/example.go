// Copyright 2013 The go-toggl AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/gedex/go-toggl/toggl"
	"os"
)

const (
	apiToken = "YOUR_API_TOKEN"
)

func main() {
	c := toggl.NewClient(apiToken)
	ws, err := c.Workspaces.List()
	checkError(err)

	// List of workspaces
	for _, w := range ws {
		fmt.Println(w.ID, w.Name)

		// List of users in current workspace
		users, err := c.Workspaces.ListUsers(w.ID)
		checkError(err)
		fmt.Println("Users on", w.Name, "workspace")
		for _, u := range users {
			fmt.Println("-", u.ID, u.Fullname, u.Email)
		}

		// List of clients in current workspace
		clients, err := c.Workspaces.ListClients(w.ID)
		checkError(err)
		fmt.Println("Clients on", w.Name, "workspace")
		for _, _c := range clients {
			fmt.Println("-", _c.ID, _c.Name)
		}

		fmt.Println()
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
