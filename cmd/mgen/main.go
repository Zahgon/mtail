// Copyright 2013 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

// Command mgen generates mtail programs for fuzz testing by following a simple grammar.
package main

import (
	"flag"
)

var (
	randSeed      = flag.Int64("rand_seed", 1, "Seed to use for math.rand.")
	minIterations = flag.Int64("min_iterations", 5000, "Minimum number of iterations before stopping program generation.")
	dictionary    = flag.Bool("dictionary", false, "Generate a fuzz dictionary to stdout only.")
)

type node struct {
	alts [][]string
	term string
}

var table = map[string]node{
	"start":     {[][]string{{"stmt_list"}}, ""},
	"stmt_list": {[][]string{{""}, {"stmt_list", "stmt"}}, ""},
	"stmt": {[][]string{
		{"cond", "{", "stmt_list", "}"},
		{"expr"},
		{"decl"},
		{"def_spec"},
		{"deco_spec"},
		{"next"},
		{"const", "ID", "pattern_expr"},
	}, ""},
	"expr":          {[][]string{{"assign_expr"}}, ""},
	"assign_expr":   {[][]string{{"rel_expr"}, {"unary_expr", "=", "rel_expr"}, {"unary_expr", "+=", "rel_expr"}}, ""},
	"rel_expr":      {[][]string{{"additive_expr"}, {"additive_expr", "relop", "additive_expr"}}, ""},
	"relop":         {[][]string{{"<"}, {">"}, {"<="}, {">="}, {"=="}, {"!="}}, ""},
	"additive_expr": {[][]string{{"unary_expr"}, {"additive_expr", "+", "unary_expr"}, {"additive_expr", "-", "unary_expr"}}, ""},
	"unary_expr":    {[][]string{{"postfix_expr"}, {"BUILTIN", "(", "arg_expr_list", ")"}}, ""},
	"arg_expr_list": {[][]string{{""}, {"assign_expr"}, {"arg_expr_list", ",", "assign_expr"}}, ""},
	"postfix_expr":  {[][]string{{"primary_expr"}, {"postfix_expr", "++"}, {"postfix_expr", "[", "expr", "]"}}, ""},
	"primary_expr":  {[][]string{{"ID"}, {"CAPREF"}, {"STRING"}, {"(", "expr", ")"}, {"NUMERIC"}}, ""},
	"cond":          {[][]string{{"pattern_expr"}, {"rel_expr"}}, ""},
	"pattern_expr":  {[][]string{{"REGEX"}, {"pattern_expr", "+", "REGEX"}, {"pattern_expr", "+", "ID"}}, ""},
	"decl":          {[][]string{{"hide_spec", "type_spec", "declarator"}}, ""},
	"hide_spec":     {[][]string{{""}, {"hidden"}}, ""},
	"declarator":    {[][]string{{"declarator", "by_spec"}, {"declarator", "as_spec"}, {"ID"}, {"STRING"}}, ""},
	"type_spec":     {[][]string{{"counter"}, {"gauge"}, {"timer"}, {"text"}, {"histogram"}}, ""},
	"by_spec":       {[][]string{{"by", "by_expr_list"}}, ""},
	"by_expr_list":  {[][]string{{"ID"}, {"STRING"}, {"by_expr_list", ",", "ID"}, {"by_expr_list", ",", "STRING"}}, ""},
	"as_spec":       {[][]string{{"as", "STRING"}}, ""},
	"def_spec":      {[][]string{{"def", "ID", "{", "stmt_list", "}"}}, ""},
	"deco_spec":     {[][]string{{"deco", "{", "stmt_list", "}"}}, ""},

	"BUILTIN": {[][]string{{"strptime"}, {"timestamp"}, {"len"}, {"tolower"}}, ""},

	"CAPREF":  {[][]string{}, "$1"},
	"REGEX":   {[][]string{}, "/foo/"},
	"STRING":  {[][]string{}, "\"bar\""},
	"ID":      {[][]string{}, "quux"},
	"NUMERIC": {[][]string{}, "37"},
}

func emitter(c chan string) { _ = "STUB: not implemented"; return }

func generateProgram() { _ = "STUB: not implemented"; return }

// Initial state

// While the state stack is not empty

// Pop the next state

// fmt.Println("state", state, "states", states)

// Look for the state transition

// If there are state transition alternatives
// fmt.Println("n", n)

// Pick a state transition at random

// fmt.Println("a", a, n.alts[a], len(n.alts[a]))
// Push the states picked onto the stack (in reverse order)

// fmt.Println("i", i, n.alts[a][len(n.alts[a])-i-1])

// fmt.Println("states", states)

// If there is a terminal, emit it
// fmt.Println("(term)", state, n.term)

// If the state doesn't exist in the table, treat it as a terminal, and emit it.
// fmt.Println("(state)", state, state)

func generateDictionary() { _ = "STUB: not implemented"; return }

func main() {
	flag.Parse()

	if *dictionary {
		generateDictionary()
	} else {
		generateProgram()
	}
}
