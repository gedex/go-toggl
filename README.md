go-toggl
========

go-toggl is Go library for accessing Toggl API.

**Documentation:** <http://godoc.org/github.com/gedex/go-toggl/toggl>
**Build Status:** [![Build Status](https://travis-ci.org/gedex/go-toggl.png?branch=master)](https://travis-ci.org/gedex/go-toggl)

## Basic Usage

~~~go
c := toggl.NewClient("YOUR_API_TOKEN")
// Get list of workspaces
ws, err := c.Workspaces.List()
if err != nil {
	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
}
for _, w := range ws {
	fmt.Println(w.ID, w.Name)
}
~~~

Please see [examples](./examples) for a complete example.

## Credits

* [go-github](https://github.com/google/go-github) in which go-toggl mimics the structure.
* [Toggl API docs](https://github.com/toggl/toggl_api_docs/)

## License

This library is distributed under the BSD-style license found in the LICENSE.md file.
