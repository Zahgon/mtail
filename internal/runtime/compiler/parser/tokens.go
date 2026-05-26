// Copyright 2011 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package parser

import (
	"github.com/google/mtail/internal/runtime/compiler/position"
)

// Kind enumerates the types of lexical tokens in a mtail program.
type Kind int

// String returns a readable name of the token Kind.
func (k Kind) String() string {
	_ = "STUB: not implemented"
	// 0xE000 is the magic offset for the first token ID in goyacc, and 2 is
	// the offset of the internal tokens in the token table.  Yes this is a
	// hack around what appears to be an original yacc bug.
	return ""
}

// Token describes a lexed Token from the input, containing its type, the
// original text of the Token, and its position in the input.
type Token struct {
	Kind     Kind
	Spelling string
	Pos      position.Position
}

// String returns a printable form of a Token.
func (t Token) String() string { _ = "STUB: not implemented"; return "" }
