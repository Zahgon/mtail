// Copyright 2015 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package runtime

// mtail programs may be created, updated, and deleted while mtail is running, and they will be
// reloaded without having to restart the mtail process -- mtail will handle these on a HUP signal.

import (
	"expvar"
	"io"
	"sync"
	"time"

	"github.com/google/mtail/internal/logline"
	"github.com/google/mtail/internal/metrics"
	"github.com/google/mtail/internal/runtime/compiler"
	"github.com/google/mtail/internal/runtime/vm"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// LineCount counts the number of lines received by the program loader.
	LineCount = expvar.NewInt("lines_total")
	// ProgLoads counts the number of program load events.
	ProgLoads = expvar.NewMap("prog_loads_total")
	// ProgUnloads counts the number of program unload events.
	ProgUnloads = expvar.NewMap("prog_unloads_total")
	// ProgLoadErrors counts the number of program load errors.
	ProgLoadErrors = expvar.NewMap("prog_load_errors_total")
)

const (
	fileExt = ".mtail"
)

// LoadAllPrograms loads all programs in a directory and starts watching the
// directory for filesystem changes.  Any compile errors are stored for later retrieival.
// This function returns an error if an internal error occurs.
func (r *Runtime) LoadAllPrograms() error { _ = "STUB: not implemented"; return nil }

// LoadProgram loads or reloads a program from the full pathname programPath.  The name of
// the program is the basename of the file.
func (r *Runtime) LoadProgram(programPath string) error { _ = "STUB: not implemented"; return nil }

// CompileAndRun compiles a program read from the input, starting execution if
// it succeeds.  If an existing virtual machine of the same name already
// exists, the previous virtual machine is terminated and the new loaded over
// it.  If the new program fails to compile, any existing virtual machine with
// the same name remains running.
func (r *Runtime) CompileAndRun(name string, input io.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

// Load the metrics from the compilation into the global metric storage for export.

// Terminates the existing vm.

type vmHandle struct {
	contentHash []byte
	vm          *vm.VM
	lines       chan *logline.LogLine
}

// Runtime handles the lifecycle of programs and virtual machines, by watching
// the configured program source directory, compiling changes to programs, and
// managing the virtual machines.
type Runtime struct {
	wg sync.WaitGroup // used to await vm shutdown

	ms  *metrics.Store        // pointer to metrics.Store to pass to compiler
	reg prometheus.Registerer // plce to reg metrics

	cOpts []compiler.Option // options for constructing `c`
	c     *compiler.Compiler

	programPath string // Path that contains mtail programs.

	handleMu sync.RWMutex         // guards accesses to handles
	handles  map[string]*vmHandle // map of program names to virtual machines

	programErrorMu sync.RWMutex     // guards access to programErrors
	programErrors  map[string]error // errors from the last compile attempt of the program

	overrideLocation     *time.Location // Instructs the vm to override the timezone with the specified zone.
	compileOnly          bool           // Only compile programs and report errors, do not load VMs.
	errorsAbort          bool           // Compiler errors abort the loader.
	dumpBytecode         bool           // Instructs the loader to dump to stdout the compiled program after compilation.
	syslogUseCurrentYear bool           // Instructs the VM to overwrite zero years with the current year in a strptime instruction.
	omitMetricSource     bool
	logRuntimeErrors     bool // Instruct the VM to emit runtime errors to the log.
	trace                bool // Trace execution of each VM.

	signalQuit chan struct{} // When closed stops the signal handler goroutine.
}

var (
	ErrNeedsStore     = errors.New("loader needs a store")
	ErrNeedsWaitgroup = errors.New("loader needs a WaitGroup")
)

// New creates a new program loader that reads programs from programPath.
func New(lines <-chan *logline.LogLine, wg *sync.WaitGroup, programPath string, store *metrics.Store, options ...Option) (*Runtime, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Defer shutdown handling to avoid a race on r.wg.

// This goroutine is the main consumer/producer loop.

// signal to owner we're done

// Create one goroutine that handles reload signals.

// Guarantee all existing programmes get loaded before we leave.

// SetOption takes one or more option functions and applies them in order to Runtime.
func (r *Runtime) SetOption(options ...Option) error { _ = "STUB: not implemented"; return nil }

// UnloadProgram removes the named program, any currently running VM goroutine.
func (r *Runtime) UnloadProgram(pathname string) { _ = "STUB: not implemented"; return }
