// Copyright 2011 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package compiler

import (
	"io"

	"github.com/google/mtail/internal/runtime/code"
)

type Compiler struct {
	emitAst             bool
	emitAstTypes        bool
	maxRegexpLength     int
	maxRecursionDepth   int
	disableOptimisation bool
}

func New(options ...Option) (*Compiler, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Compiler) SetOption(options ...Option) error { _ = "STUB: not implemented"; return nil }

// Option configures a new Compiler.
type Option func(*Compiler) error

// EmitAst emits the AST after the parse phase.
func EmitAst() Option { _ = "STUB: not implemented"; return *new(Option) }

// EmitAstTypes emits the AST with types after the type checking phase.
func EmitAstTypes() Option { _ = "STUB: not implemented"; return *new(Option) }

// MaxRegexpLength sets the maximum allowable length of a regular expression.
func MaxRegexpLength(maxRegexpLength int) Option { _ = "STUB: not implemented"; return *new(Option) }

// MaxRecursionDepth sets the maximum allowable depth of the AST.
func MaxRecursionDepth(maxRecursionDepth int) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// DisableOptimisation disables the optimisation phase.
func DisableOptimisation() Option { _ = "STUB: not implemented"; return *new(Option) }

// Compile compiles a program from the input into bytecode and data stored in an Object, or a list
// of compile errors.
func (c *Compiler) Compile(name string, input io.Reader) (obj *code.Object, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
