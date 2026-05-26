// Copyright 2017 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package parser

import (
	"strings"

	"github.com/google/mtail/internal/runtime/compiler/ast"
	"github.com/google/mtail/internal/runtime/compiler/symbol"
)

// Sexp is for converting program syntax trees into typed s-expression for printing.
type Sexp struct {
	output strings.Builder // Accumulator for the result

	EmitTypes bool

	col  int // column to indent current line to
	line strings.Builder
}

func (s *Sexp) indent() { _ = "STUB: not implemented"; return }

func (s *Sexp) outdent() { _ = "STUB: not implemented"; return }

func (s *Sexp) prefix() (r string) { _ = "STUB: not implemented"; return "" }

func (s *Sexp) emit(str string) { _ = "STUB: not implemented"; return }

func (s *Sexp) newline() { _ = "STUB: not implemented"; return }

// VisitBefore implements the astNode Visitor interface.
func (s *Sexp) VisitBefore(n ast.Node) (ast.Visitor, ast.Node) {
	_ = "STUB: not implemented"
	return *new(ast.Visitor), *new(ast.Node)
}

// normal walk

// VisitAfter implements the astNode Visitor interface.
func (s *Sexp) VisitAfter(node ast.Node) ast.Node { _ = "STUB: not implemented"; return *new(ast.Node) }

func (s *Sexp) emitScope(scope *symbol.Scope) { _ = "STUB: not implemented"; return }

// Dump begins the dumping of the syntax tree, returning the s-expression as a single string.
func (s *Sexp) Dump(n ast.Node) string { _ = "STUB: not implemented"; return "" }
