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
	accessToken = "9aeebc52e0e4651f38b93cc93d7a655c"
)

func main() {
	var wid int

	c := toggl.NewClient("9aeebc52e0e4651f38b93cc93d7a655c")

	// We need workspace ID to be able creating a new project
	ws, err := c.Workspaces.List()
	checkError(err)
	if len(ws) == 0 {
		fmt.Println("No workspace")
		os.Exit(0)
	}
	wid = ws[0].ID

	// Create a project
	p := &toggl.Project{Name: "Test project", WorkspaceID: wid}
	p, err = c.Projects.Create(p)
	checkError(err)
	fmt.Println("Newly created project's ID", p.ID)

	// Update newly created project
	p.Name = "Test project v2"
	p, err = c.Projects.Update(p)
	checkError(err)
	fmt.Println("Newly created project's name is updated")

	// Get information about updated project
	p, err = c.Projects.Get(p.ID)
	checkError(err)
	fmt.Println("Newly created project's name", p.Name)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
