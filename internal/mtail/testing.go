// Copyright 2019 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package mtail

import (
	"context"
	"expvar"
	"testing"
	"time"

	"github.com/google/mtail/internal/metrics/datum"
	"github.com/google/mtail/internal/waker"
)

const defaultDoOrTimeoutDeadline = 10 * time.Second

type TestServer struct {
	*Server

	streamWaker waker.Waker // for idle logstreams; others are polled explicitly in PollWatched
	// AwakenLogStreams wakes n log streams.  This acts as a barrier method,
	// synchronising the logstreams and the test.
	AwakenLogStreams waker.WakeFunc

	patternWaker waker.Waker // polling for new glob pattern matches
	// AwakenPatternPollers wakes n pattern pollers.  This acts as a barrier
	// method, synchronising the pattern poll with the test.
	AwakenPatternPollers waker.WakeFunc // the glob awakens

	tb testing.TB

	cancel context.CancelFunc

	// Set this to change the poll deadline when using DoOrTimeout within this TestServer.
	DoOrTimeoutDeadline time.Duration
}

// TestMakeServer makes a new TestServer for use in tests, but does not start
// the server.  If an error occurs during creation, a testing.Fatal is issued.
func TestMakeServer(tb testing.TB, patternWakers int, streamWakers int, options ...Option) *TestServer {
	_ = "STUB: not implemented"

	// Reset counters when running multiple tests.  Tests that use expvar
	// helpers cannot be made parallel.
	return nil
}

// TestStartServer creates a new TestServer and starts it running.  It returns
// the server, and a stop function.  `patternWakers` indicates the number of
// expected pattern wakers to wait for at this moment; usually 1 because the
// test server is started with a `LogPathPattern`.  `streamWakers` indiecates
// the number of expected stream wakers to wait for at this moment.  The value
// of this parameter shuld be the number of log files created in test
// (e.g. with `testutil.TestOpenFile`) before invoking this function.
func TestStartServer(tb testing.TB, patternWakers int, streamWakers int, options ...Option) (*TestServer, func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Start starts the TestServer and returns a stop function.
func (ts *TestServer) Start() func() { _ = "STUB: not implemented"; return nil }

func (ts *TestServer) LoadAllPrograms() { _ = "STUB: not implemented"; return }

// GetExpvar is a helper function on TestServer that acts like TestGetExpvar.
func (ts *TestServer) GetExpvar(name string) expvar.Var {
	_ = "STUB: not implemented"
	return *new(expvar.Var)
}

// ExpectExpvarDeltaWithDeadline returns a deferrable function which tests if the expvar metric with name has changed by delta within the given deadline, once the function begins.  Before returning, it fetches the original value for comparison.
func (ts *TestServer) ExpectExpvarDeltaWithDeadline(name string, want int64) func() {
	_ = "STUB: not implemented"
	return nil
}

// ExpectMapExpvarMetricDeltaWithDeadline returns a deferrable function which tests if the expvar map metric with name and key has changed by delta within the given deadline, once the function begins.  Before returning, it fetches the original value for comparison.
func (ts *TestServer) ExpectMapExpvarDeltaWithDeadline(name, key string, want int64) func() {
	_ = "STUB: not implemented"
	return nil
}

// GetProgramMetric fetches the datum of the program metric name.
func (ts *TestServer) GetProgramMetric(name, prog string) datum.Datum {
	_ = "STUB: not implemented"
	return *new(datum.Datum)
}

// ExpectProgMetricDeltaWithDeadline tests that a given program metric increases by want within the deadline.  It assumes that the named metric is an Int type datum.Datum.
func (ts *TestServer) ExpectProgMetricDeltaWithDeadline(name, prog string, want int64) func() {
	_ = "STUB: not implemented"
	return nil
}
