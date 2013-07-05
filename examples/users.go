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
	c := toggl.NewClient("9aeebc52e0e4651f38b93cc93d7a655c")
	me, err := c.Users.Me(false)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}

	fmt.Println("My email", me.Email)
}
