// Copyright 2011 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package symbol

import (
	"github.com/google/mtail/internal/runtime/compiler/position"
	"github.com/google/mtail/internal/runtime/compiler/types"
)

// Kind enumerates the kind of a Symbol.
type Kind int

// Kind enumerates the kinds of symbols found in the program text.
const (
	VarSymbol     Kind = iota // Variables
	CaprefSymbol              // Capture group references
	DecoSymbol                // Decorators
	PatternSymbol             // Named pattern constants
	endSymbol                 // for testing
)

func (k Kind) String() string { _ = "STUB: not implemented"; return "" }

// Symbol describes a named program object.
type Symbol struct {
	Name    string             // identifier name
	Kind    Kind               // kind of program object
	Type    types.Type         // object's type
	Pos     *position.Position // Source file position of definition
	Binding interface{}        // binding to storage allocated in runtime
	Addr    int                // Address offset in another structure, object specific
	Used    bool               // Optional marker that this symbol is used after declaration.
}

// NewSymbol creates a record of a given symbol kind, named name, found at loc.
func NewSymbol(name string, kind Kind, pos *position.Position) (sym *Symbol) {
	_ = "STUB: not implemented"
	return nil
}

// Scope maintains a record of the identifiers declared in the current program
// scope, and a link to the parent scope.
type Scope struct {
	Parent  *Scope
	Symbols map[string]*Symbol
}

// NewScope creates a new scope within the parent scope.
func NewScope(parent *Scope) *Scope { _ = "STUB: not implemented"; return nil }

// Insert attempts to insert a symbol into the scope.  If the scope already
// contains an object alt with the same name, the scope is unchanged and the
// function returns alt.  Otherwise the symbol is inserted, and returns nil.
func (s *Scope) Insert(sym *Symbol) (alt *Symbol) { _ = "STUB: not implemented"; return nil }

// InsertAlias attempts to insert a duplicate name for an existing symbol into
// the scope.  If the scope already contains an object alt with the alias, the
// scope is unchanged and the function returns alt.  Otherwise, the symbol is
// inserted and the function returns nil.
func (s *Scope) InsertAlias(sym *Symbol, alias string) (alt *Symbol) {
	_ = "STUB: not implemented"
	return nil
}

// Lookup returns the symbol with the given name if it is found in this or any
// parent scope, otherwise nil.
func (s *Scope) Lookup(name string, kind Kind) *Symbol { _ = "STUB: not implemented"; return nil }

// String prints the current scope and all parents to a string, recursing up to
// the root scope.  This method is only used for debugging.
func (s *Scope) String() string { _ = "STUB: not implemented"; return "" }

// CopyFrom copies all the symbols from another scope object into this one.
// It recurses up the input scope copying all visible symbols into one.
func (s *Scope) CopyFrom(o *Scope) { _ = "STUB: not implemented"; return }
