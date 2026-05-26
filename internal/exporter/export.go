// Copyright 2011 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

// Package exporter provides the interface for getting metrics out of mtail,
// into your monitoring system of choice.
package exporter

import (
	"context"
	"expvar"
	"flag"
	"io"
	"sync"
	"time"

	"github.com/google/mtail/internal/metrics"
	"github.com/pkg/errors"
)

// Commandline Flags.
var (
	writeDeadline = flag.Duration("metric_push_write_deadline", 10*time.Second, "Time to wait for a push to succeed before exiting with an error.")
)

// Exporter manages the export of metrics to passive and active collectors.
type Exporter struct {
	ctx            context.Context
	cancelFunc     context.CancelFunc
	wg             sync.WaitGroup
	store          *metrics.Store
	pushInterval   time.Duration
	hostname       string
	omitProgLabel  bool
	emitTimestamp  bool
	exportDisabled bool
	pushTargets    []pushOptions
	initDone       chan struct{}
	shutdownDone   chan struct{}
}

// Option configures a new Exporter.
type Option func(*Exporter) error

// Hostname specifies the mtail hostname to use in exported metrics.
func Hostname(hostname string) Option { _ = "STUB: not implemented"; return *new(Option) }

// OmitProgLabel sets the Exporter to not put program names in metric labels.
func OmitProgLabel() Option { _ = "STUB: not implemented"; return *new(Option) }

// EmitTimestamp instructs the exporter to send metric's timestamps to collectors.
func EmitTimestamp() Option { _ = "STUB: not implemented"; return *new(Option) }

func PushInterval(opt time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func DisableExport() Option { _ = "STUB: not implemented"; return *new(Option) }

var ErrNeedsStore = errors.New("exporter needs a Store")

// New creates a new Exporter.
func New(ctx context.Context, store *metrics.Store, options ...Option) (*Exporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// defaults after options have been set

// This routine manages shutdown of the Exporter.

// Wait for the context to be completed before waiting for subroutines.

// Stop instructs the exporter to shut down.  The function returns once the exporter has finished.
func (e *Exporter) Stop() { _ = "STUB: not implemented"; return }

// SetOption takes one or more option functions and applies them in order to Exporter.
func (e *Exporter) SetOption(options ...Option) error { _ = "STUB: not implemented"; return nil }

// formatLabels converts a metric name and key-value map of labels to a single
// string for exporting to the correct output format for each export target.
// ksep and sep mark what to use for key/val separator, and between label separators respoectively.
// If not empty, rep is used to replace cases of ksep and sep in the original strings.
func formatLabels(name string, m map[string]string, ksep, sep, rep string) string {
	_ = "STUB: not implemented"
	return ""
}

// Format a LabelSet into a string to be written to one of the timeseries
// sockets.
type formatter func(string, *metrics.Metric, *metrics.LabelSet, time.Duration) string

func (e *Exporter) writeSocketMetrics(c io.Writer, f formatter, exportTotal *expvar.Int, exportSuccess *expvar.Int) error {
	_ = "STUB: not implemented"
	return nil
}

// Don't try to send text metrics to any push service.

// PushMetrics sends metrics to each of the configured services.
func (e *Exporter) PushMetrics() { _ = "STUB: not implemented"; return }

// StartMetricPush pushes metrics to the configured services each interval.
func (e *Exporter) StartMetricPush() { _ = "STUB: not implemented"; return }

type pushOptions struct {
	net, addr      string
	f              formatter
	total, success *expvar.Int
}

// RegisterPushExport adds a push export connection to the Exporter.  Items in
// the list must describe a Dial()able connection and will have all the metrics
// pushed to each pushInterval.
func (e *Exporter) RegisterPushExport(p pushOptions) { _ = "STUB: not implemented"; return }
