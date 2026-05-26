// Copyright 2011 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package ast

// Visitor VisitBefore method is invoked for each node encountered by Walk.
// If the result Visitor v is not nil, Walk visits each of the children of that
// node with v.  VisitAfter is called on n at the end.
type Visitor interface {
	VisitBefore(n Node) (Visitor, Node)
	VisitAfter(n Node) Node
}

// convenience function.
func walknodelist(v Visitor, list []Node) []Node { _ = "STUB: not implemented"; return nil }

// Walk traverses (walks) an AST node with the provided Visitor v.
func Walk(v Visitor, node Node) Node { _ = "STUB: not implemented"; return *new(Node) }

// Returning nil from VisitBefore signals to Walk that the Visitor has
// handled the children of this node.  VisitAfter will not be called.

// These nodes are terminals, thus have no children to walk.
