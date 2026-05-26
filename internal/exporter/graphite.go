// Copyright 2011 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package exporter

import (
	"expvar"
	"flag"
	"net/http"
	"time"

	"github.com/google/mtail/internal/metrics"
)

var (
	graphiteHostPort = flag.String("graphite_host_port", "",
		"Host:port to graphite carbon server to write metrics to.")
	graphitePrefix = flag.String("graphite_prefix", "",
		"Prefix to use for graphite metrics.")

	graphiteExportTotal   = expvar.NewInt("graphite_export_total")
	graphiteExportSuccess = expvar.NewInt("graphite_export_success")
)

func (e *Exporter) HandleGraphite(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// metricToGraphite encodes a metric in the graphite text protocol format.  The
// metric lock is held before entering this function.
func metricToGraphite(_ string, m *metrics.Metric, l *metrics.LabelSet, _ time.Duration) string {
	_ = "STUB: not implemented"
	return ""
}
