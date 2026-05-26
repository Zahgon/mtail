// Copyright 2021 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

// package opt has a compiler pass for making optimisations on the AST.
package opt

import (
	"github.com/google/mtail/internal/runtime/compiler/ast"
	"github.com/google/mtail/internal/runtime/compiler/errors"
)

func Optimise(n ast.Node) (ast.Node, error) { _ = "STUB: not implemented"; return *new(ast.Node), nil }

type optimiser struct {
	errors errors.ErrorList
}

func (o *optimiser) VisitBefore(node ast.Node) (ast.Visitor, ast.Node) {
	_ = "STUB: not implemented"
	return *new(ast.Visitor), *new(ast.Node)
}

func (o *optimiser) VisitAfter(node ast.Node) ast.Node {
	_ = "STUB: not implemented"
	return *new(ast.Node)
}
