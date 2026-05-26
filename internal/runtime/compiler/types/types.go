// Copyright 2016 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package types

import (
	"errors"
	"regexp/syntax"
	"sync"
)

// Type represents a type in the mtail program.
type Type interface {
	// Root returns an exemplar Type after unification occurs.  If the type
	// system is complete after unification, Root will be a TypeOperator.  Root
	// is the equivalent of Find in the union-find algorithm.
	Root() Type

	// String returns a string representation of a Type.
	String() string
}

// TypeError describes an error in which a type was expected, but another was encountered.
type TypeError struct {
	error    error
	expected Type
	received Type
}

var (
	ErrRecursiveUnification = errors.New("recursive unification error")
	ErrTypeMismatch         = errors.New("type mismatch")
	ErrInternal             = errors.New("internal error")
)

func (e *TypeError) Root() Type { _ = "STUB: not implemented"; return *new(Type) }

func (e *TypeError) String() string { _ = "STUB: not implemented"; return "" }

func (e TypeError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *TypeError) Unwrap() error {
	_ = "STUB: not implemented"

	// AsTypeError behaves like `errors.As`, attempting to cast the type `t` into a
	// provided `target` TypeError and returning if it was successful.
	return nil
}

func AsTypeError(t Type, target **TypeError) (ok bool) { _ = "STUB: not implemented"; return false }

// IsTypeError behaves like `errors.Is`, indicating that the type is a TypeError.
func IsTypeError(t Type) bool { _ = "STUB: not implemented"; return false }

var (
	nextVariableIDMu sync.Mutex
	nextVariableID   int
)

// Variable represents an unbound type variable in the type system.
type Variable struct {
	ID int

	// Instance is set if this variable has been bound to a type.
	instanceMu sync.RWMutex
	Instance   Type
}

// NewVariable constructs a new unique TypeVariable.
func NewVariable() *Variable { _ = "STUB: not implemented"; return nil }

// Root returns an exemplar of this TypeVariable, in this case the root of the unification tree.
func (t *Variable) Root() Type { _ = "STUB: not implemented"; return *new(Type) }

func (t *Variable) String() string { _ = "STUB: not implemented"; return "" }

// SetInstance sets the exemplar instance of this TypeVariable, during
// unification.  SetInstance is the equivalent of Union in the Union-Find
// algorithm.
func (t *Variable) SetInstance(t1 Type) { _ = "STUB: not implemented"; return }

// Operator represents a type scheme in the type system.
type Operator struct {
	// Name is a common name for this operator
	Name string
	// Args is the sequence of types that are parameters to this type.  They
	// may be fully bound type operators, or partially defined (i.e. contain
	// TypeVariables) in which case they represent polymorphism in the operator
	// they are arguments to.
	Args []Type
}

// Root returns an exemplar of a TypeOperator, i.e. itself.
func (t *Operator) Root() Type { _ = "STUB: not implemented"; return *new(Type) }

func (t *Operator) String() (s string) { _ = "STUB: not implemented"; return "" }

const (
	functionName  = "→"
	dimensionName = "⨯"
	alternateName = "|"
)

// Function is a convenience method, which instantiates a new Function type
// scheme, with the given args as parameters.
func Function(args ...Type) *Operator { _ = "STUB: not implemented"; return nil }

// IsFunction returns true if the given type is a Function type.
func IsFunction(t Type) bool { _ = "STUB: not implemented"; return false }

// Dimension is a convenience method which instantiates a new Dimension type
// scheme, with the given args as the dimensions of the type.  (This type looks
// a lot like a Product type.)
func Dimension(args ...Type) *Operator { _ = "STUB: not implemented"; return nil }

// IsDimension returns true if the given type is a Dimension type.
func IsDimension(t Type) bool { _ = "STUB: not implemented"; return false }

// Alternate is a convenience method which instantiates a new Alternate type
// scheme, with the given args as the possible types this type may take.  (You
// might know this sort of type by the name Sum type.)
func Alternate(args ...Type) *Operator { _ = "STUB: not implemented"; return nil }

// IsAlternate returns true if the given type is an Alternate type.
func IsAlternate(t Type) bool { _ = "STUB: not implemented"; return false }

// IsComplete returns true if the type and all its arguments have non-variable exemplars.
func IsComplete(t Type) bool { _ = "STUB: not implemented"; return false }

// Builtin type constants.
var (
	Error         = &TypeError{}
	InternalError = &TypeError{error: ErrInternal}
	Undef         = &Operator{"Undef", []Type{}}
	None          = &Operator{"None", []Type{}}
	Bool          = &Operator{"Bool", []Type{}}
	Int           = &Operator{"Int", []Type{}}
	Float         = &Operator{"Float", []Type{}}
	String        = &Operator{"String", []Type{}}
	Pattern       = &Operator{"Pattern", []Type{}}
	// TODO(jaq): use composite type so we can typecheck the bucket directly, e.g. hist[j] = i.
	Buckets = &Operator{"Buckets", []Type{}}

	// Numeric types can be either Int or Float.
	Numeric = Alternate(Int, Float)
)

// Builtins is a mapping of the builtin language functions to their type definitions.
var Builtins = map[string]Type{
	"int":         Function(NewVariable(), Int),
	"bool":        Function(NewVariable(), Bool),
	"float":       Function(NewVariable(), Float),
	"string":      Function(NewVariable(), String),
	"timestamp":   Function(Int),
	"len":         Function(String, Int),
	"settime":     Function(Int, None),
	"strptime":    Function(String, String, None),
	"strtol":      Function(String, Int, Int),
	"tolower":     Function(String, String),
	"getfilename": Function(String),
	"subst":       Function(Pattern, String, String, String),
}

// FreshType returns a new type from the provided type scheme, replacing any
// unbound type variables with new type variables.
func FreshType(t Type) Type {
	_ = "STUB: not implemented"
	// mappings keeps track of replaced variables in this type so that t -> t
	// becomes q -> q not q -> r
	return *new(Type)
}

// occursIn returns true if `v` is in any of `types`.
func OccursIn(v Type, types []Type) bool { _ = "STUB: not implemented"; return false }

// occursInType returns true if `v` is `t2` or recursively contained within `t2`.
func occursInType(v Type, t2 Type) bool { _ = "STUB: not implemented"; return false }

// Equals compares two types, testing for equality.
func Equals(t1, t2 Type) bool { _ = "STUB: not implemented"; return false }

// Unify performs type unification of both parameter Types.  It returns the
// least upper bound of both types, the most general type that is capable of
// representing both parameters.  If either type is a type variable, then that
// variable is unified with the LUB.  In reporting errors, it is assumed that a
// is the expected type and b is the type observed.
func Unify(a, b Type) Type { _ = "STUB: not implemented"; return *new(Type) }

// reverse args, to recurse the pattern above

// Re-reverse from the recursion

// We flipped the args, flip them back.

// Both are Alternates, find intersection of type arguments.

type TypeCoercion struct {
	sub, sup Type
}

// type coercions for builtin types
var typeCoercions = []TypeCoercion{
	{Bool, Int},
	{Bool, Float}, // contentious
	{Int, Float},  // contentious
	{Bool, String},
	{Int, String},
	{Float, String},
	{String, Pattern},
	{Int, Bool}, // an integer using C style cast to bool
}

// LeastUpperBound returns the smallest type that may contain both parameter types.
func LeastUpperBound(a, b Type) Type { _ = "STUB: not implemented"; return *new(Type) }

// If either is a TypeVariable, the other is the lub

// If either is Undef, other is the lub

// Easy substitutions

// Patterns imply match status, which is boolean.

// A Numeric can be an Int, or a Float, but not vice versa.

// A string can be a pattern, but not vice versa.

// A pattern and an Int are Bool

// inferCaprefType determines a type for the nth capturing group in re, based on contents
// of that capture group.
func InferCaprefType(re *syntax.Regexp, n int) Type { _ = "STUB: not implemented"; return *new(Type) }

func inferGroupType(group *syntax.Regexp) Type { _ = "STUB: not implemented"; return *new(Type) }

// Must be at least one digit in the group.

// Only one decimal point allowed.

// getCaptureGroup returns the Regexp node of the capturing group numbered cgID
// in re.
func getCaptureGroup(re *syntax.Regexp, cgID int) *syntax.Regexp {
	_ = "STUB: not implemented"
	return nil
}

// groupOnlyMatches returns true iff group only matches runes in s.
func groupOnlyMatches(group *syntax.Regexp, s string) bool { _ = "STUB: not implemented"; return false }
