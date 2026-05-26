// Copyright 2021 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package runtime

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// Option configures a new program Runtime.
type Option func(*Runtime) error

// OverrideLocation sets the timezone location for the VM.
func OverrideLocation(loc *time.Location) Option { _ = "STUB: not implemented"; return *new(Option) }

// CompileOnly sets the Runtime to compile programs only, without executing them.
func CompileOnly() Option { _ = "STUB: not implemented"; return *new(Option) }

// ErrorsAbort sets the Runtime to abort the Runtime on compile errors.
func ErrorsAbort() Option { _ = "STUB: not implemented"; return *new(Option) }

// DumpAst emits the AST after program compilation.
func DumpAst() Option { _ = "STUB: not implemented"; return *new(Option) }

// DumpAstTypes emits the AST after type checking.
func DumpAstTypes() Option { _ = "STUB: not implemented"; return *new(Option) }

// DumpBytecode instructs the loader to print the compiled bytecode after code generation.
func DumpBytecode() Option { _ = "STUB: not implemented"; return *new(Option) }

// SyslogUseCurrentYear instructs the VM to annotate yearless timestamps with the current year.
func SyslogUseCurrentYear() Option { _ = "STUB: not implemented"; return *new(Option) }

// MaxRegexpLength sets the maximum length an mtail regular expression can have, in terms of characters.
func MaxRegexpLength(maxRegexpLength int) Option { _ = "STUB: not implemented"; return *new(Option) }

// MaxRecursionDepth sets the maximum depth the abstract syntax tree built during lexation can have.
func MaxRecursionDepth(maxRecursionDepth int) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// OmitMetricSource instructs the Runtime to not annotate metrics with their program source when added to the metric store.
func OmitMetricSource() Option { _ = "STUB: not implemented"; return *new(Option) }

// PrometheusRegisterer passes in a registry for setting up exported metrics.
func PrometheusRegisterer(reg prometheus.Registerer) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// LogRuntimeErrors instructs the VM to emit runtime errors into the log.
func LogRuntimeErrors() Option { _ = "STUB: not implemented"; return *new(Option) }

func TraceExecution() Option { _ = "STUB: not implemented"; return *new(Option) }
