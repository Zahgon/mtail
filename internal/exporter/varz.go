// Copyright 2015 Google Inc.  All Rights Reserved.
// This file is available under the Apache license.

package exporter

import (
	"expvar"
	"net/http"

	"github.com/google/mtail/internal/metrics"
)

var exportVarzTotal = expvar.NewInt("exporter_varz_total")

const varzFormat = "%s{%s} %s\n"

// HandleVarz exports the metrics in Varz format via HTTP.
func (e *Exporter) HandleVarz(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func metricToVarz(m *metrics.Metric, l *metrics.LabelSet, omitProgLabel bool, hostname string) string {
	_ = "STUB: not implemented"
	return ""
}
