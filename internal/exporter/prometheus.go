// Copyright 2015 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package exporter

import (
	"expvar"
	"io"

	"github.com/google/mtail/internal/metrics"
	"github.com/google/mtail/internal/metrics/datum"
	"github.com/prometheus/client_golang/prometheus"
)

var metricExportTotal = expvar.NewInt("metric_export_total")

func noHyphens(s string) string { _ = "STUB: not implemented"; return "" }

// Describe implements the prometheus.Collector interface.
func (e *Exporter) Describe(c chan<- *prometheus.Desc) { _ = "STUB: not implemented"; return }

// Collect implements the prometheus.Collector interface.
func (e *Exporter) Collect(c chan<- prometheus.Metric) { _ = "STUB: not implemented"; return }

/* #nosec G104 always retursn nil */

// We don't have a way of converting text metrics to prometheus format.

// By default no timestamp is emitted to Prometheus. Setting a
// timestamp is not recommended. It can lead to unexpected results
// if the timestamp is not updated or moved fowarded enough to avoid
// triggering Promtheus staleness handling.
// Read more in docs/faq.md

// Write is used to write Prometheus metrics to an io.Writer.
func (e *Exporter) Write(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func promTypeForKind(k metrics.Kind) prometheus.ValueType {
	_ = "STUB: not implemented"
	return *new(prometheus.ValueType)
}

func promValueForDatum(d datum.Datum) float64 { _ = "STUB: not implemented"; return 0 }
