// Copyright 2015 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package exporter

import (
	"expvar"
	"net/http"
)

var exportJSONErrors = expvar.NewInt("exporter_json_errors")

// HandleJSON exports the metrics in JSON format via HTTP.
func (e *Exporter) HandleJSON(w http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"
	return
}
