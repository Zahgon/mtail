// Copyright 2011 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package mtail

import (
	"context"
	"net"
	"sync"

	"github.com/google/mtail/internal/exporter"
	"github.com/google/mtail/internal/logline"
	"github.com/google/mtail/internal/metrics"
	"github.com/google/mtail/internal/runtime"
	"github.com/google/mtail/internal/tailer"
	"github.com/prometheus/client_golang/prometheus"
)

// Server contains the state of the main mtail program.
type Server struct {
	ctx    context.Context
	cancel context.CancelFunc

	wg sync.WaitGroup // wait for main processes to shutdown

	store *metrics.Store // Metrics storage

	tOpts []tailer.Option    // options for constructing `t`
	t     *tailer.Tailer     // t manages log patterns and log streams, which sends lines to the VMs
	rOpts []runtime.Option   // options for constructing `r`
	r     *runtime.Runtime   // r loads programs and manages the VM lifecycle
	eOpts []exporter.Option  // options for constructing `e`
	e     *exporter.Exporter // e manages the export of metrics from the store

	lines chan *logline.LogLine // primary communication channel, owned by Tailer.

	reg *prometheus.Registry

	listener net.Listener // Configured with bind address.

	buildInfo BuildInfo // go build information

	programPath        string // path to programs to load
	oneShot            bool   // if set, mtail reads log files from the beginning, once, then exits
	compileOnly        bool   // if set, mtail compiles programs then exit
	httpDebugEndpoints bool   // if set, mtail will enable debug endpoints
	httpInfoEndpoints  bool   // if set, mtail will enable info endpoints for progz and varz
}

// We can only copy the build info once to the version library.  Protects tests from data races.
var buildInfoOnce sync.Once

// initRuntime constructs a new runtime and performs the initial load of program files in the program directory.
func (m *Server) initRuntime() (err error) { _ = "STUB: not implemented"; return nil }

// initExporter sets up an Exporter for this Server.
func (m *Server) initExporter() (err error) { _ = "STUB: not implemented"; return nil }

// Create mtail_build_info metric.

// initTailer sets up and starts a Tailer for this Server.
func (m *Server) initTailer() (err error) { _ = "STUB: not implemented"; return nil }

// initHTTPServer begins the http server.
func (m *Server) initHTTPServer() error { _ = "STUB: not implemented"; return nil }

// This goroutine runs the http server.

// This goroutine manages http server shutdown.

// Wait for the Serve routine to exit.

// New creates a Server from the supplied Options.  The Server is started by
// the time New returns, it watches the LogPatterns for files, starts tailing
// their changes and sends any new lines found to the virtual machines loaded
// from ProgramPath. If OneShot mode is enabled, it will exit after reading
// each log file from start to finish.
// TODO(jaq): this doesn't need to be a constructor anymore, it could start and
// block until quit, once TestServer.PollWatched is addressed.
func New(ctx context.Context, store *metrics.Store, options ...Option) (*Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Using a non-pedantic registry means we can be looser with metrics that
// are not fully specified at startup.

// TODO(jaq): Should these move to initExporter?

// internal/tailer/file.go

// internal/runtime/loader.go

// Prefix all expvar metrics with 'mtail_'

//nolint:contextcheck // TODO

//nolint:contextcheck // TODO

// SetOption takes one or more option functions and applies them in order to MtailServer.
func (m *Server) SetOption(options ...Option) error { _ = "STUB: not implemented"; return nil }

// Run awaits mtail's shutdown.
// TODO(jaq): remove this once the test server is able to trigger polls on the components.
func (m *Server) Run() error { _ = "STUB: not implemented"; return nil }
