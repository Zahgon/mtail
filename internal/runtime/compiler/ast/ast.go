// Copyright 2011 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package ast

import (
	"sync"
	"time"

	"github.com/google/mtail/internal/metrics"
	"github.com/google/mtail/internal/runtime/compiler/position"
	"github.com/google/mtail/internal/runtime/compiler/symbol"
	"github.com/google/mtail/internal/runtime/compiler/types"
)

type Node interface {
	Pos() *position.Position // Returns the position of the node from the original source
	Type() types.Type        // Returns the type of the expression in this node
}

type StmtList struct {
	Scope    *symbol.Scope // Pointer to the local scope for this enclosing block
	Children []Node
}

func (n *StmtList) Pos() *position.Position { _ = "STUB: not implemented"; return nil }

func (n *StmtList) Type() types.Type { _ = "STUB: not implemented"; return *new(types.Type) }

type ExprList struct {
	Children []Node

	typMu sync.RWMutex
	typ   types.Type
}

func (n *ExprList) Pos() *position.Position { _ = "STUB: not implemented"; return nil }

func (n *ExprList) Type() types.Type { _ = "STUB: not implemented"; return *new(types.Type) }

func (n *ExprList) SetType(t types.Type) { _ = "STUB: not implemented"; return }

type CondStmt struct {
	Cond  Node
	Truth Node
	Else  Node
	Scope *symbol.Scope // a conditional expression can cause new variables to be defined
}

func (n *CondStmt) Pos() *position.Position { _ = "STUB: not implemented"; return nil }

func (n *CondStmt) Type() types.Type { _ = "STUB: not implemented"; return *new(types.Type) }

type IDTerm struct {
	P      position.Position
	Name   string
	Symbol *symbol.Symbol
	Lvalue bool // If set, then this node appears on the left side of an
	// assignment and needs to have its address taken only.
}

func (n *IDTerm) Pos() *position.Position { _ = "STUB: not implemented"; return nil }

func (n *IDTerm) Type() types.Type { _ = "STUB: not implemented"; return *new(types.Type) }

// id not defined

type CaprefTerm struct {
	P       position.Position
	Name    string
	IsNamed bool // true if the capref is a named reference, not positional
	Symbol  *symbol.Symbol
}

func (n *CaprefTerm) Pos() *position.Position { _ = "STUB: not implemented"; return nil }

func (n *CaprefTerm) Type() types.Type { _ = "STUB: not implemented"; return *new(types.Type) }

// sym not defined due to undefined capref error

type BuiltinExpr struct {
	P    position.Position
	Name string
	Args Node

	typMu sync.RWMutex
	typ   types.Type
}

func (n *BuiltinExpr) Pos() *position.Position { _ = "STUB: not implemented"; return nil }

func (n *BuiltinExpr) Type() types.Type { _ = "STUB: not implemented"; return *new(types.Type) }

func (n *BuiltinExpr) SetType(t types.Type) { _ = "STUB: not implemented"; return }

type BinaryExpr struct {
	LHS, RHS Node
	Op       int

	typMu sync.RWMutex
	typ   types.Type
}

func (n *BinaryExpr) Pos() *position.Position { _ = "STUB: not implemented"; return nil }

func (n *BinaryExpr) Type() types.Type { _ = "STUB: not implemented"; return *new(types.Type) }

func (n *BinaryExpr) SetType(t types.Type) { _ = "STUB: not implemented"; return }

type UnaryExpr struct {
	P    position.Position // pos is the position of the op
	Expr Node
	Op   int

	typMu sync.RWMutex
	typ   types.Type
}

func (n *UnaryExpr) Pos() *position.Position { _ = "STUB: not implemented"; return nil }

func (n *UnaryExpr) Type() types.Type { _ = "STUB: not implemented"; return *new(types.Type) }

func (n *UnaryExpr) SetType(t types.Type) { _ = "STUB: not implemented"; return }

type IndexedExpr struct {
	LHS, Index Node

	typMu sync.RWMutex
	typ   types.Type
}

func (n *IndexedExpr) Pos() *position.Position { _ = "STUB: not implemented"; return nil }

func (n *IndexedExpr) Type() types.Type { _ = "STUB: not implemented"; return *new(types.Type) }

func (n *IndexedExpr) SetType(t types.Type) { _ = "STUB: not implemented"; return }

type VarDecl struct {
	P            position.Position
	Name         string
	Hidden       bool
	Keys         []string
	Limit        int64
	Buckets      []float64
	Kind         metrics.Kind
	ExportedName string
	Symbol       *symbol.Symbol
}

func (n *VarDecl) Pos() *position.Position { _ = "STUB: not implemented"; return nil }

func (n *VarDecl) Type() types.Type { _ = "STUB: not implemented"; return *new(types.Type) }

type StringLit struct {
	P    position.Position
	Text string
}

func (n *StringLit) Pos() *position.Position { _ = "STUB: not implemented"; return nil }

func (n *StringLit) Type() types.Type { _ = "STUB: not implemented"; return *new(types.Type) }

type IntLit struct {
	P position.Position
	I int64
}

func (n *IntLit) Pos() *position.Position { _ = "STUB: not implemented"; return nil }

func (n *IntLit) Type() types.Type { _ = "STUB: not implemented"; return *new(types.Type) }

type FloatLit struct {
	P position.Position
	F float64
}

func (n *FloatLit) Pos() *position.Position { _ = "STUB: not implemented"; return nil }

func (n *FloatLit) Type() types.Type {
	_ = "STUB: not implemented"

	// PatternExpr is the top of a pattern expression.
	return *new(types.Type)
}

type PatternExpr struct {
	Expr    Node
	Pattern string // if not empty, the fully defined pattern after typecheck
	Index   int    // reference to the compiled object offset after codegen
}

func (n *PatternExpr) Pos() *position.Position { _ = "STUB: not implemented"; return nil }

func (n *PatternExpr) Type() types.Type {
	_ = "STUB: not implemented"
	return *

	// PatternLit holds inline constant pattern fragments.
	new(types.Type)
}

type PatternLit struct {
	P       position.Position
	Pattern string
}

func (n *PatternLit) Pos() *position.Position { _ = "STUB: not implemented"; return nil }

func (n *PatternLit) Type() types.Type {
	_ = "STUB: not implemented"
	return *

	// PatternFragment holds a named pattern part.
	new(types.Type)
}

type PatternFragment struct {
	ID      Node
	Expr    Node
	Symbol  *symbol.Symbol // Optional Symbol for a named pattern
	Pattern string         // If not empty, contains the complete evaluated pattern of the expr
}

func (n *PatternFragment) Pos() *position.Position { _ = "STUB: not implemented"; return nil }

func (n *PatternFragment) Type() types.Type { _ = "STUB: not implemented"; return *new(types.Type) }

type DecoDecl struct {
	P      position.Position
	Name   string
	Block  Node
	Symbol *symbol.Symbol
	Scope  *symbol.Scope // The declaration creates its own scope, as a zygote to be instantiated later.
}

func (n *DecoDecl) Pos() *position.Position { _ = "STUB: not implemented"; return nil }

func (n *DecoDecl) Type() types.Type { _ = "STUB: not implemented"; return *new(types.Type) }

type DecoStmt struct {
	P     position.Position
	Name  string
	Block Node
	Decl  *DecoDecl     // Pointer to the declaration of the decorator this statement invokes.
	Scope *symbol.Scope // Instantiated with a copy of the Def's Scope.
}

func (n *DecoStmt) Pos() *position.Position { _ = "STUB: not implemented"; return nil }

func (n *DecoStmt) Type() types.Type { _ = "STUB: not implemented"; return *new(types.Type) }

type NextStmt struct {
	P position.Position
}

func (n *NextStmt) Pos() *position.Position { _ = "STUB: not implemented"; return nil }

func (n *NextStmt) Type() types.Type { _ = "STUB: not implemented"; return *new(types.Type) }

type OtherwiseStmt struct {
	P position.Position
}

func (n *OtherwiseStmt) Pos() *position.Position { _ = "STUB: not implemented"; return nil }

func (n *OtherwiseStmt) Type() types.Type { _ = "STUB: not implemented"; return *new(types.Type) }

type DelStmt struct {
	P      position.Position
	N      Node
	Expiry time.Duration
}

func (n *DelStmt) Pos() *position.Position { _ = "STUB: not implemented"; return nil }

func (n *DelStmt) Type() types.Type { _ = "STUB: not implemented"; return *new(types.Type) }

type ConvExpr struct {
	N Node

	mu  sync.RWMutex
	typ types.Type
}

func (n *ConvExpr) Pos() *position.Position { _ = "STUB: not implemented"; return nil }

func (n *ConvExpr) Type() types.Type { _ = "STUB: not implemented"; return *new(types.Type) }

func (n *ConvExpr) SetType(t types.Type) { _ = "STUB: not implemented"; return }

type Error struct {
	P        position.Position
	Spelling string
}

func (n *Error) Pos() *position.Position { _ = "STUB: not implemented"; return nil }

func (n *Error) Type() types.Type { _ = "STUB: not implemented"; return *new(types.Type) }

type StopStmt struct {
	P position.Position
}

func (n *StopStmt) Pos() *position.Position { _ = "STUB: not implemented"; return nil }

func (n *StopStmt) Type() types.Type {
	_ = "STUB: not implemented"

	// mergepositionlist is a helper that merges the positions of all the nodes in a list.
	return *new(types.Type)
}

func mergepositionlist(l []Node) *position.Position { _ = "STUB: not implemented"; return nil }
