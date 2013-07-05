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
	clients, err := c.Clients.List()
	checkError(err)

	// List of clients
	for _, _c := range clients {
		fmt.Println(_c.ID, _c.Name)
		wid = _c.WorkspaceID
	}

	// Create new client.
	nc := &toggl.WorkspaceClient{Name: "ACME", WorkspaceID: wid}
	nc, err = c.Clients.Create(nc)
	checkError(err)
	fmt.Println("Newly created client's ID", nc.ID)

	// Update client.
	nc.Name = "ACME v2"
	nc, err = c.Clients.Update(nc)
	checkError(err)
	fmt.Println("Newly created client's name is updated")

	// Get information about updated client.
	nc, err = c.Clients.Get(nc.ID)
	checkError(err)
	fmt.Println("Newly created client's name", nc.Name)

	// Delete newly created client
	err = c.Clients.Delete(nc.ID)
	checkError(err)
	fmt.Println("Client ID", nc.ID, "has been deleted")
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
