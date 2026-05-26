// Copyright 2015 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

package errors

import (
	"github.com/google/mtail/internal/runtime/compiler/position"
)

type compileError struct {
	pos position.Position
	msg string
}

func (e compileError) Error() string { _ = "STUB: not implemented"; return "" }

// ErrorList contains a list of compile errors.
type ErrorList []*compileError

// Add appends an error at a position to the list of errors.
func (p *ErrorList) Add(pos *position.Position, msg string) { _ = "STUB: not implemented"; return }

// Append puts an ErrorList on the end of this ErrorList.
func (p *ErrorList) Append(l ErrorList) { _ = "STUB: not implemented"; return }

// ErrorList implements the error interface.
func (p ErrorList) Error() string { _ = "STUB: not implemented"; return "" }

func Errorf(format string, args ...interface{}) error { _ = "STUB: not implemented"; return nil }
