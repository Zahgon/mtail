// Copyright 2020 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package mtail

import (
	"net/http"
)

const statusTemplate = `
<!DOCTYPE html>
<html>
<head>
<title>mtail on {{.BindAddress}}</title>
</head>
<body>
<h1>mtail on {{.BindAddress}}</h1>
<p>Build: {{.BuildInfo}}</p>
<p>Metrics: <a href="/json">json</a>, <a href="/graphite">graphite</a>, <a href="/metrics">prometheus</a></p>
<p>Info: {{ if .HTTPInfoEndpoints }}<a href="/varz">varz</a>, <a href="/progz">progz</a> <a href="/tracez">tracez</a></p>{{ else }} disabled {{ end }}</p>
<p>Debug: {{ if .HTTPDebugEndpoints }}<a href="/debug/pprof">debug/pprof</a>, <a href="/debug/vars">debug/vars</a>{{ else }} disabled {{ end }}</p>
`

const statusTemplateEnd = `
</body>
</html>
`

// ServeHTTP satisfies the http.Handler interface, and is used to serve the
// root page of mtail for online status reporting.
func (m *Server) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"
	return
}

// FaviconHandler is used to serve up the favicon.ico for mtail's http server.
func FaviconHandler(w http.ResponseWriter, _ *http.Request) { _ = "STUB: not implemented"; return }
