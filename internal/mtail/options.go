// Copyright 2011 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package mtail

import (
	"errors"
	"time"

	"github.com/google/mtail/internal/exporter"
	"github.com/google/mtail/internal/runtime"
	"github.com/google/mtail/internal/tailer"
	"github.com/google/mtail/internal/waker"
)

// Option configures mtail.Server.
type Option interface {
	apply(*Server) error
}

// ProgramPath sets the path to find mtail programs in the Server.
type ProgramPath string

func (opt ProgramPath) apply(m *Server) error { _ = "STUB: not implemented"; return nil }

// LogPathPatterns sets the patterns to find log paths in the Server.
func LogPathPatterns(patterns ...string) Option { _ = "STUB: not implemented"; return *new(Option) }

type logPathPatterns []string

func (opt logPathPatterns) apply(m *Server) error { _ = "STUB: not implemented"; return nil }

// IgnoreRegexPattern sets the regex pattern to ignore files.
type IgnoreRegexPattern string

func (opt IgnoreRegexPattern) apply(m *Server) error { _ = "STUB: not implemented"; return nil }

// BindAddress sets the HTTP server address in Server.
func BindAddress(address, port string) Option { _ = "STUB: not implemented"; return *new(Option) }

type bindAddress struct {
	address, port string
}

var ErrDuplicateHTTPBindAddress = errors.New("HTTP server bind address already supplied")

func (opt bindAddress) apply(m *Server) error { _ = "STUB: not implemented"; return nil }

// BindUnixSocket sets the UNIX socket path in Server.
type BindUnixSocket string

func (opt BindUnixSocket) apply(m *Server) error { _ = "STUB: not implemented"; return nil }

// SetBuildInfo sets the mtail program build information in the Server.
type SetBuildInfo BuildInfo

func (opt SetBuildInfo) apply(m *Server) error { _ = "STUB: not implemented"; return nil }

// OverrideLocation sets the timezone location for log timestamps without any such information.
func OverrideLocation(loc *time.Location) Option { _ = "STUB: not implemented"; return *new(Option) }

type overrideLocation struct {
	*time.Location
}

func (opt overrideLocation) apply(m *Server) error { _ = "STUB: not implemented"; return nil }

// LogPatternPollWaker triggers polls on the filesystem for new logs that match the log glob patterns.
func LogPatternPollWaker(w waker.Waker) Option { _ = "STUB: not implemented"; return *new(Option) }

type logPatternPollWaker struct {
	waker.Waker
}

func (opt logPatternPollWaker) apply(m *Server) error { _ = "STUB: not implemented"; return nil }

// LogstreamPollWaker triggers polls on the filesystem for new logs that match the log glob streams.
func LogstreamPollWaker(w waker.Waker) Option { _ = "STUB: not implemented"; return *new(Option) }

type logstreamPollWaker struct {
	waker.Waker
}

func (opt logstreamPollWaker) apply(m *Server) error { _ = "STUB: not implemented"; return nil }

type niladicOption struct {
	applyfunc func(m *Server) error
}

func (n *niladicOption) apply(m *Server) error { _ = "STUB: not implemented"; return nil }

// OneShot sets one-shot mode in the Server.
var OneShot = &niladicOption{
	func(m *Server) error {
		m.rOpts = append(m.rOpts, runtime.ErrorsAbort())
		m.tOpts = append(m.tOpts, tailer.OneShot)
		m.eOpts = append(m.eOpts, exporter.DisableExport())
		m.oneShot = true
		return nil
	},
}

// CompileOnly sets compile-only mode in the Server.
var CompileOnly = &niladicOption{
	func(m *Server) error {
		m.rOpts = append(m.rOpts, runtime.CompileOnly())
		m.eOpts = append(m.eOpts, exporter.DisableExport())
		m.compileOnly = true
		return nil
	},
}

// DumpAst instructs the Server's compiler to print the AST after parsing.
var DumpAst = &niladicOption{
	func(m *Server) error {
		m.rOpts = append(m.rOpts, runtime.DumpAst())
		return nil
	},
}

// DumpAstTypes instructs the Server's compiler to print the AST after type checking.
var DumpAstTypes = &niladicOption{
	func(m *Server) error {
		m.rOpts = append(m.rOpts, runtime.DumpAstTypes())
		return nil
	},
}

// DumpBytecode instructs the Server's compiuler to print the program bytecode after code generation.
var DumpBytecode = &niladicOption{
	func(m *Server) error {
		m.rOpts = append(m.rOpts, runtime.DumpBytecode())
		return nil
	},
}

// HttpDebugEndpoints enables debug http endpoints
var HTTPDebugEndpoints = &niladicOption{
	func(m *Server) error {
		m.httpDebugEndpoints = true
		return nil
	},
}

// HttpInfoEndpoints enables info http endpoints
var HTTPInfoEndpoints = &niladicOption{
	func(m *Server) error {
		m.httpInfoEndpoints = true
		return nil
	},
}

// SyslogUseCurrentYear instructs the Server to use the current year for year-less log timestamp during parsing.
var SyslogUseCurrentYear = &niladicOption{
	func(m *Server) error {
		m.rOpts = append(m.rOpts, runtime.SyslogUseCurrentYear())
		return nil
	},
}

// OmitProgLabel sets the Server to not put the program name as a label in exported metrics.
var OmitProgLabel = &niladicOption{
	func(m *Server) error {
		m.eOpts = append(m.eOpts, exporter.OmitProgLabel())
		return nil
	},
}

// OmitMetricSource sets the Server to not link created metrics to their source program.
var OmitMetricSource = &niladicOption{
	func(m *Server) error {
		m.rOpts = append(m.rOpts, runtime.OmitMetricSource())
		return nil
	},
}

// EmitMetricTimestamp tells the Server to export the metric's timestamp.
var EmitMetricTimestamp = &niladicOption{
	func(m *Server) error {
		m.eOpts = append(m.eOpts, exporter.EmitTimestamp())
		return nil
	},
}

// LogRuntimeErrors instructs the VM to emit runtime errors to the log.
var LogRuntimeErrors = &niladicOption{
	func(m *Server) error {
		m.rOpts = append(m.rOpts, runtime.LogRuntimeErrors())
		return nil
	},
}

// JaegerReporter creates a new jaeger reporter that sends to the given Jaeger endpoint address.
type JaegerReporter string

func (opt JaegerReporter) apply(_ *Server) error { _ = "STUB: not implemented"; return nil }

// MetricPushInterval sets the interval between metrics pushes to passive collectors.
type MetricPushInterval time.Duration

func (opt MetricPushInterval) apply(m *Server) error { _ = "STUB: not implemented"; return nil }

// MaxRegexpLength sets the maximum length an mtail regular expression can have, in terms of characters.
type MaxRegexpLength int

func (opt MaxRegexpLength) apply(m *Server) error { _ = "STUB: not implemented"; return nil }

// MaxRecursionDepth sets the maximum depth the abstract syntax tree built during lexation can have.
type MaxRecursionDepth int

func (opt MaxRecursionDepth) apply(m *Server) error { _ = "STUB: not implemented"; return nil }
