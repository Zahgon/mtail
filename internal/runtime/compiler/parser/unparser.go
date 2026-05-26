// Copyright 2011 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package parser

import (
	"strings"

	"github.com/google/mtail/internal/runtime/compiler/ast"
)

// Unparser is for converting program syntax trees back to program text.
type Unparser struct {
	pos       int
	output    strings.Builder
	line      strings.Builder
	emitTypes bool
}

func (u *Unparser) indent() { _ = "STUB: not implemented"; return }

func (u *Unparser) outdent() { _ = "STUB: not implemented"; return }

func (u *Unparser) prefix() (s string) { _ = "STUB: not implemented"; return "" }

func (u *Unparser) emit(s string) { _ = "STUB: not implemented"; return }

func (u *Unparser) newline() { _ = "STUB: not implemented"; return }

// VisitBefore implements the ast.Visitor interface.
func (u *Unparser) VisitBefore(n ast.Node) (ast.Visitor, ast.Node) {
	_ = "STUB: not implemented"
	return *new(ast.Visitor), *new(ast.Node)
}

// VisitAfter implements the ast.Visitor interface.
func (u *Unparser) VisitAfter(n ast.Node) ast.Node {
	_ = "STUB: not implemented"

	// Unparse begins the unparsing of the syntax tree, returning the program text as a single string.
	return *new(ast.Node)
}

func (u *Unparser) Unparse(n ast.Node) string { _ = "STUB: not implemented"; return "" }
