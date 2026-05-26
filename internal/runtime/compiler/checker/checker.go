// Copyright 2016 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package checker

import (
	"strings"

	"github.com/google/mtail/internal/runtime/compiler/ast"
	"github.com/google/mtail/internal/runtime/compiler/errors"
	"github.com/google/mtail/internal/runtime/compiler/symbol"
)

const (
	defaultMaxRegexpLength   = 1024
	defaultMaxRecursionDepth = 100
)

// checker holds data for a semantic checker.
type checker struct {
	scope *symbol.Scope // the current scope

	decoScopes []*symbol.Scope // A stack of scopes used for resolving symbols in decorated nodes

	errors errors.ErrorList

	depth             int
	tooDeep           bool
	maxRecursionDepth int
	maxRegexLength    int
	noRegexSymbols    bool
}

// Check performs a semantic check of the astNode, and returns a potentially
// modified astNode and either a list of errors found, or nil if the program is
// semantically valid.  At the completion of Check, the symbol table and type
// annotation are also complete.
func Check(node ast.Node, maxRegexpLength int, maxRecursionDepth int) (ast.Node, error) {
	_ = "STUB: not implemented"
	// set defaults
	return *new(ast.Node), nil
}

// VisitBefore performs most of the symbol table construction, so that symbols
// are guaranteed to exist before their use.
func (c *checker) VisitBefore(node ast.Node) (ast.Visitor, ast.Node) {
	_ = "STUB: not implemented"
	return *new(ast.Visitor), *new(ast.Node)
}

// TODO(jaq): This should be a numeric type, unless we want to
// enforce more specific rules like "Counter can only be Int."

// One type per key

// and one for the value.

// Apply a terribly bad heuristic to choose a suggestion.

// If the string is all uppercase, pretend it was a const
// pattern because that's what the docs do.

// Append a scope placeholder for the recursion into the block.  It has no parent, it'll be cloned when the decorator is instantiated.

// Create a new scope for the decorator instantiation.

// Clone the DecoDecl scope zygote into this scope.

// checkSymbolTable emits errors if any eligible symbols in the current scope
// are not marked as used or have an invalid type.
func (c *checker) checkSymbolTable() { _ = "STUB: not implemented"; return }

// Users don't have control over the patterns given from decorators
// so this should never be an error; but it can be useful to know
// if a program is doing unnecessary work.

// Don't warn about the zeroth capture group; it's not user-defined.

// VisitAfter performs the type annotation and checking, once the child nodes
// of expressions have been annotated and checked.  Within this function,
// gotType refers to the types inferred in the AST, and wantType is the type
// expected for this expression.  After unification, uType is the concrete type
// of the expression, and the visitor should set any node Types as appropriate.
//
// The notation for type inference used comes from the 2010 lecture notes for
// Stanford's CS413 class.
// https://web.stanford.edu/class/cs143/lectures/lecture09.pdf
func (c *checker) VisitAfter(node ast.Node) ast.Node {
	_ = "STUB: not implemented"
	return *new(ast.Node)
}

// Pop the scope

// OK as conditions

// If the parser saw an IDTerm with type Pattern, then we know it's really a pattern constant and need to wrap it in an unary match in this context.

// Pop the scope.

// Don't check symbol usage here because the decorator is only partially defined.
// Pop the scope.

// The last element in this list will be the empty stack created by the
// DecoDecl on the way in.  If there's no last element, then we can't
// have entered a DecoDecl yet.

// Merge the current scope into it.

// Pop the scope off the list, and insert it into this node.

// Store the zygote from the scope stack on this declaration.

// Arithmetic: e1 OP e2
// O ⊢ e1 : Tl, O ⊢ e2 : Tr
// Tl <= Tr , Tr <= Tl
// ⇒ O ⊢ e : lub(Tl, Tr)

// First handle the Tl <= Tr and vice versa.

// Change the type mismatch error to make more sense in this context.

// Implicit type conversion for non-comparisons, promoting each
// half to the return type of the op.

// bitwise: e1 OP e2
// O ⊢ e1 : Int, O ⊢ e2 : Int
// ⇒ O ⊢ e : Int

// If the parser saw an IDTerm with type Pattern, then we know it's really a pattern constant and need to wrap it in an unary match in this context.

// Likewise for the RHS

// logical: e1 OP e2
// O ⊢ e1 : Bool, O ⊢ e2 : Bool
// ⇒ O ⊢ e : Bool

// comparable, logical: e2 OP e2
// O ⊢ e1 : Tl, O ⊢ e2 : Tr
// Tl <= Tr , Tr <= Tl
// ⇒ O ⊢ e : Bool

// First handle the Tl <= Tr and vice versa.

// Implicit type conversion: Promote types if the ast types are not
// the same as the expression type.

// e1 = e2; e1 += e2
// O ⊢ e1 : Tl, O ⊢ e2 : Tr
// Tr <= Tl
// ⇒ O ⊢ e : Tl

// TODO(jaq): the rT <= lT relationship is not correctly encoded here.

// If the LHS is assignable, mark it as an lvalue, otherwise error.

// e1 =~ e2, e1 !~ e2
// O ⊢ e1 : String , O ⊢ e2 : Pattern
// ⇒ O ⊢ e : Bool
// TODO(jaq): We're not correctly encoding this.

// Implicit conversion of the RHS to a PatternExpr if not already.

// !e1
// O ⊢ e1 : Int
// ⇒ O ⊢ e : Bool

// e1++ , e1--
// O ⊢ e1 : Int
// ⇒ O ⊢ e : Int

// TODO we do this backwards versus ADD_ASSIGN above, why

// If the expr is assignable, mark it as an lvalue, otherwise error.

// After unification, the expr still has to be of Int type.

// Implicit match expressions, an expression of type Pattern returning Bool
// /e1/
// O ⊢ e1 : Pattern
// ⇒ O ⊢ e : Bool

// (e1, e2, ...)
// ⇒ O ⊢ e: e1⨯e1⨯...

// e1[e2, e3, ..., en]
// O ⊢ e1 : T1⨯T2⨯...Tn⨯Tr
// O ⊢ e2,e3,...,en : T1,T2,...,Tn
// ⇒ O ⊢ e : Tr

// prune this node to n.LHS if Index is nil.  Leave 0 length exprlist as that's a type error.

// undefined, already caught (where?)

// We now have enough information to tell that something the
// parser thought was an IDTerm is really a pattern constant,
// so we can rewrite the AST here.  We can't yet wrap the
// pattern expression with Unary Match because we don't know
// the context yet, but see CondExpr and BinaryExpr's
// logical-op.

// it's a Dimension, continue after switch

// T1,T2,...,Tn

// Tr

// Having typechecked the expression against the expected types, and
// have detected mismatched keylengths, we have a well-formed
// expression, so can now fold to just IDTerm if there's no ExprList.

// f(e1, e2, ..., en)
// O ⊢ f : T1⨯T2⨯...Tn⨯Tr
// O ⊢ e1,e2,...,en : T1,T2,...,Tn
// ⇒ O ⊢ e : Tr
// TODO: recall the syntax for subst a fresh type above

// Second argument to strptime is the format string.  If it is
// defined at compile time, we can verify it can be use as a format
// string by parsing itself.

// Layout strings can contain an underscore to indicate a digit
// field if the layout field can contain two digits; but they
// won't parse themselves.  Zulu Timezones in the layout need
// to be converted to offset in the parsed time.

// Evaluate the expression.

// Evaluate the expression.

// checkRegex is a helper method to compile and check a regular expression, and
// to generate its capture groups as symbols.
func (c *checker) checkRegex(pattern string, n ast.Node) { _ = "STUB: not implemented"; return }

// We reserve the names of the capturing groups as declarations
// of those symbols, so that future CAPREF tokens parsed can
// retrieve their value.  By recording them in the symbol table, we
// can warn the user about unknown capture group references.

// No return, let this loop collect all errors

// No return, let this loop collect all errors

// patternEvaluator is a helper that performs concatenation of pattern
// fragments so that they can be compiled as whole regular expression patterns.
type patternEvaluator struct {
	scope   *symbol.Scope
	errors  *errors.ErrorList
	pattern strings.Builder
}

func (p *patternEvaluator) VisitBefore(n ast.Node) (ast.Visitor, ast.Node) {
	_ = "STUB: not implemented"
	return *new(ast.Visitor), *new(ast.Node)
}

// Already looked up sym, if still nil then undefined.

func (p *patternEvaluator) VisitAfter(n ast.Node) ast.Node {
	_ = "STUB: not implemented"
	return *new(ast.Node)
}
