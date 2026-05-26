// Copyright 2021 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package runtime

import (
	"io"
	"net/http"
)

const loaderTemplate = `
<h2 id="loader">Program Loader</h2>
<table border="1">
<tr>
<th>program name</th>
<th>errors</th>
<th>load errors</th>
<th>load successes</th>
<th>unloads</th>
<th>runtime errors</th>
<th>last runtime error</th>
</tr>
<tr>
{{range $name, $errors := $.Errors}}
<td>{{ if index $.ProgLoaded $name}}<a href="/progz?prog={{$name}}">{{$name}}</a>{{else}}{{$name}}{{end}}</td>
<td>
{{if $errors}}
{{$errors}}
{{else}}
No compile errors
{{end}}
</td>
<td>{{index $.Loaderrors $name}}</td>
<td>{{index $.Loadsuccess $name}}</td>
<td>{{index $.Unloads $name}}</td>
<td>{{index $.RuntimeErrors $name}}</td>
<td><pre>{{index $.RuntimeErrorString $name}}</pre></td>
</tr>
{{end}}
</table>
`

// WriteStatusHTML writes the current state of the loader as HTML to the given writer w.
func (r *Runtime) WriteStatusHTML(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func (r *Runtime) ProgzHandler(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}
