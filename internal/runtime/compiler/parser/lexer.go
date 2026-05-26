// Copyright 2011 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package parser

import (
	"bufio"
	"io"
	"strings"
)

// List of keywords.  Keep this list sorted!
var keywords = map[string]Kind{
	"after":     AFTER,
	"as":        AS,
	"buckets":   BUCKETS,
	"by":        BY,
	"const":     CONST,
	"counter":   COUNTER,
	"def":       DEF,
	"del":       DEL,
	"else":      ELSE,
	"gauge":     GAUGE,
	"hidden":    HIDDEN,
	"histogram": HISTOGRAM,
	"limit":     LIMIT,
	"next":      NEXT,
	"otherwise": OTHERWISE,
	"stop":      STOP,
	"text":      TEXT,
	"timer":     TIMER,
}

// List of builtin functions.  Keep this list sorted!
var builtins = []string{
	"bool",
	"float",
	"getfilename",
	"int",
	"len",
	"settime",
	"string",
	"strptime",
	"strtol",
	"subst",
	"timestamp",
	"tolower",
}

// Dictionary returns a list of all keywords and builtins of the language.
func Dictionary() (r []string) { _ = "STUB: not implemented"; return nil }

// A stateFn represents each state the scanner can be in.
type stateFn func(*Lexer) stateFn

// A lexer holds the state of the scanner.
type Lexer struct {
	name  string        // Name of program.
	input *bufio.Reader // Source program
	state stateFn       // Current state function of the lexer.

	// The "read cursor" in the input.
	rune  rune // The current rune.
	width int  // Width in bytes.
	line  int  // The line position of the current rune.
	col   int  // The column position of the current rune.

	InRegex bool // Context aware flag from parser to say we're in a regex

	// The currently being lexed token.
	startcol int             // Starting column of the current token.
	text     strings.Builder // the text of the current token

	tokens chan Token // Output channel for tokens emitted.
}

// NewLexer creates a new scanner type that reads the input provided.
func NewLexer(name string, input io.Reader) *Lexer { _ = "STUB: not implemented"; return nil }

// NextToken returns the next token in the input.  When no token is available
// to be returned it executes the next action in the state machine.
func (l *Lexer) NextToken() Token { _ = "STUB: not implemented"; return *new(Token) }

// emit passes a token to the client.
func (l *Lexer) emit(kind Kind) { _ = "STUB: not implemented"; return }

// Reset the current token

// Internal end of file value.
const eof rune = -1

// next returns the next rune in the input.
func (l *Lexer) next() rune { _ = "STUB: not implemented"; return 0 }

// backup indicates that we haven't yet dealt with the next rune. Use when
// terminating tokens on unknown runes.
func (l *Lexer) backup() { _ = "STUB: not implemented"; return }

// stepCursor moves the read cursor.
func (l *Lexer) stepCursor() { _ = "STUB: not implemented"; return }

// accept accepts the current rune and its position into the current token.
func (l *Lexer) accept() { _ = "STUB: not implemented"; return }

// skip does not accept the current rune into the current token's text, but
// does accept its position into the token. Use only at the start or end of a
// token.
func (l *Lexer) skip() {
	_ = "STUB: not implemented"

	// ignore skips over the current rune, removing it from the text of the token,
	// and resetting the start position of the current token. Use only between
	// tokens.
	return
}

func (l *Lexer) ignore() { _ = "STUB: not implemented"; return }

// errorf returns an error token and resets the scanner.
func (l *Lexer) errorf(format string, args ...interface{}) stateFn {
	_ = "STUB: not implemented"
	return *new(stateFn)
}

// Reset the current token

// State functions.

// lexProg starts lexing a program.
func lexProg(l *Lexer) stateFn { _ = "STUB: not implemented"; return *new(stateFn) }

// Stop the machine, we're done.

// Lex a comment.
func lexComment(l *Lexer) stateFn { _ = "STUB: not implemented"; return *new(stateFn) }

// Lex a numerical constant.
func lexNumeric(l *Lexer) stateFn { _ = "STUB: not implemented"; return *new(stateFn) }

func isDurationSuffix(r rune) bool { _ = "STUB: not implemented"; return false }

func lexDuration(l *Lexer) stateFn { _ = "STUB: not implemented"; return *new(stateFn) }

// Lex a quoted string.  The text of a quoted string does not include the '"' quotes.
func lexQuotedString(l *Lexer) stateFn {
	_ = "STUB: not implemented"
	// Skip leading quote
	return *new(stateFn)
}

// Skip trailing quote.

// Lex a capture group reference. These are local variable references to
// capture groups in the preceding regular expression.
func lexCapref(l *Lexer) stateFn {
	_ = "STUB: not implemented"
	// Skip the leading $
	return *new(stateFn)
}

// Lex an identifier, or builtin keyword.
func lexIdentifier(l *Lexer) stateFn { _ = "STUB: not implemented"; return *new(stateFn) }

// Lex a regular expression pattern. The text of the regular expression does
// not include the '/' quotes.
func lexRegex(l *Lexer) stateFn {
	_ = "STUB: not implemented"
	// Exit regex mode when leaving this function.
	return *new(stateFn)
}

// Backup trailing slash on successful parse

// Lex a decorator name. These are functiony templatey wrappers around blocks
// of rules.
func lexDecorator(l *Lexer) stateFn {
	_ = "STUB: not implemented"
	// Skip the leading @
	return *new(stateFn)
}

// Helper predicates.

// isAlpha reports whether r is an alphabetical rune.
func isAlpha(r rune) bool { _ = "STUB: not implemented"; return false }

// isAlnum reports whether r is an alphanumeric rune.
func isAlnum(r rune) bool { _ = "STUB: not implemented"; return false }

// isDigit reports whether r is a numerical rune.
func isDigit(r rune) bool { _ = "STUB: not implemented"; return false }

// isSpace reports whether r is whitespace.
func isSpace(r rune) bool { _ = "STUB: not implemented"; return false }
