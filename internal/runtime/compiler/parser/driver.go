// Copyright 2016 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

// Build the parser:
//go:generate goyacc -v y.output -o parser.go -p mtail parser.y

// Package parser implements the parse phase of the mtail program compilation.
// The parser itself is defined in parser.y, and goyacc generates the program
// code and token definitions.  The parser fetches tokens from the lexer, which
// scans the input converting the program source into a token stream.  The
// driver code wraps the generated parser and marshals the ast and errors back
// to the caller.
//
// Two pretty-printers are used for debugging: the unparser, which converts an
// ast back into program text, and an approximation of an s-expression printer,
// which tries to model in indented text the structure of the ast.
package parser

import (
	"flag"
	"io"

	"github.com/google/mtail/internal/runtime/compiler/ast"
	"github.com/google/mtail/internal/runtime/compiler/errors"
	"github.com/google/mtail/internal/runtime/compiler/position"
)

// Parse reads the program named name from the input, and if successful returns
// an ast.Node for the root of the AST, otherwise parser errors.
func Parse(name string, input io.Reader) (ast.Node, error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

// EOF is a marker for end of file.  It has the same value as the goyacc internal Kind `$end`.
const EOF = 0

// parser defines the data structure for parsing an mtail program.
type parser struct {
	name   string
	root   ast.Node
	errors errors.ErrorList
	l      *Lexer
	t      Token             // Most recently lexed token.
	pos    position.Position // Optionally contains the position of the start of a production
}

func newParser(name string, input io.Reader) *parser { _ = "STUB: not implemented"; return nil }

func (p *parser) ErrorP(s string, pos *position.Position) { _ = "STUB: not implemented"; return }

func (p *parser) Error(s string) { _ = "STUB: not implemented"; return }

// Lex reads the next token from the Lexer, turning it into a form useful for the goyacc generated parser.
// The variable lval is modified to carry token information, and the token type is returned.
func (p *parser) Lex(lval *mtailSymType) int { _ = "STUB: not implemented"; return 0 }

func (p *parser) inRegex() { _ = "STUB: not implemented"; return }

func init() {
	// Initialise globals defined in generated parser.go, defaults to 0 and false
	flag.IntVar(&mtailDebug, "mtailDebug", 0, "Set parser debug level.")
	mtailErrorVerbose = true
}
