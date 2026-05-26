// Copyright 2016 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package codegen

import (
	"github.com/google/mtail/internal/runtime/code"
	"github.com/google/mtail/internal/runtime/compiler/ast"
	"github.com/google/mtail/internal/runtime/compiler/errors"
	"github.com/google/mtail/internal/runtime/compiler/parser"
	"github.com/google/mtail/internal/runtime/compiler/position"
	"github.com/google/mtail/internal/runtime/compiler/types"
)

// codegen represents a code generator.
type codegen struct {
	name string // Name of the program.

	errors errors.ErrorList // Any compile errors detected are accumulated here.
	obj    code.Object      // The object to return, if successful.

	l     []int           // Label table for recording jump destinations.
	decos []*ast.DecoStmt // Decorator stack to unwind when entering decorated blocks.
}

// CodeGen is the function that compiles the program to bytecode and data.
func CodeGen(name string, n ast.Node) (*code.Object, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *codegen) errorf(pos *position.Position, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *codegen) emit(n ast.Node, opcode code.Opcode, operand interface{}) {
	_ = "STUB: not implemented"
	return
}

// newLabel creates a new label to jump to.
func (c *codegen) newLabel() (l int) { _ = "STUB: not implemented"; return 0 }

// setLabel points a label to the next instruction.
func (c *codegen) setLabel(l int) { _ = "STUB: not implemented"; return }

// pc returns the program offset of the last instruction.
func (c *codegen) pc() int { _ = "STUB: not implemented"; return 0 }

func (c *codegen) VisitBefore(node ast.Node) (ast.Visitor, ast.Node) {
	_ = "STUB: not implemented"
	return *new(ast.Visitor), *new(ast.Node)
}

// If the Type is not in the map, then default to metrics.Int.  This is
// a hack for metrics that no type can be inferred, retaining
// historical behaviour.

// Scalar counters can be initialized to zero.  Dimensioned counters we
// don't know the values of the labels yet.  Gauges and Timers we can't
// assume start at zero.

// Calling GetDatum here causes the storage to be allocated.

// Initialize to zero at the zero time.

// Calling GetDatum here causes the storage to be allocated.

// int is int64 only on 64bit platforms.  To be fair MaxInt is a
// ridiculously excessive size for this anyway, you're going to use 2GiB
// x sizeof(datum) in a single metric.

// Set matched flag false for children.

// Re-set matched flag to true for rest of current block.

// Store the location of this regular expression in the PatternExpr

// Skip, const pattern fragments are concatenated into PatternExpr storage, not executable.

// rn.index contains the index of the compiled regular expression object
// in the re slice of the object code

// n.Symbol.Addr is the capture group offset

// Do nothing, defs are inlined.

// Put the current block on the stack

// then iterate over the decorator's nodes

// Visit the 'next' block on the decorated block stack

// overwrite the dload instruction

// Double-emit the lhs so that it can be assigned to

// Didn't handle it, let normal walk proceed

var typedOperators = map[int]map[types.Type]code.Opcode{
	parser.PLUS: {
		types.Int:     code.Iadd,
		types.Float:   code.Fadd,
		types.String:  code.Cat,
		types.Pattern: code.Cat,
	},
	parser.MINUS: {
		types.Int:   code.Isub,
		types.Float: code.Fsub,
	},
	parser.MUL: {
		types.Int:   code.Imul,
		types.Float: code.Fmul,
	},
	parser.DIV: {
		types.Int:   code.Idiv,
		types.Float: code.Fdiv,
	},
	parser.MOD: {
		types.Int:   code.Imod,
		types.Float: code.Fmod,
	},
	parser.POW: {
		types.Int:   code.Ipow,
		types.Float: code.Fpow,
	},
	parser.ASSIGN: {
		types.Int:    code.Iset,
		types.Float:  code.Fset,
		types.String: code.Sset,
	},
}

func getOpcodeForType(op int, opT types.Type) (code.Opcode, error) {
	_ = "STUB: not implemented"
	return *new(code.Opcode), nil
}

var builtin = map[string]code.Opcode{
	"getfilename": code.Getfilename,
	"len":         code.Length,
	"settime":     code.Settime,
	"strptime":    code.Strptime,
	"strtol":      code.S2i,
	"subst":       code.Subst,
	"timestamp":   code.Timestamp,
	"tolower":     code.Tolower,
}

func (c *codegen) VisitAfter(node ast.Node) ast.Node {
	_ = "STUB: not implemented"
	return *new(ast.Node)
}

// TODO(jaq): Nothing, no support in VM yet.

// len args should be 1

// When operand is not nil, inc pops the delta from the stack.

// Already walked the lhs and rhs of this expression

// And a second lhs

func (c *codegen) emitConversion(n ast.Node, inType, outType types.Type) error {
	_ = "STUB: not implemented"
	return nil
}

// nothing, pattern is implicit bool

// Nothing; no-op.

func (c *codegen) writeJumps() { _ = "STUB: not implemented"; return }
