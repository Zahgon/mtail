// Copyright 2018 Google Inc. All Rights Reserved.
// This file is available under the Apache license.

/*
Command mdot turns an mtail program AST into a graphviz graph on standard output.

To use, run it like (assuming your shell is in the same directory as this file)

	go run github.com/google/mtail/cmd/mdot --prog ../../examples/dhcpd.mtail | xdot -

or

	go run github.com/google/mtail/cmd/mdot --prog ../../examples/dhcpd.mtail --http_port 8080

to view the dot output visit http://localhost:8080

You'll need the graphviz `dot' command installed.
*/
package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/golang/glog"
	"github.com/google/mtail/internal/mtail"
	"github.com/google/mtail/internal/runtime/compiler/ast"
)

var (
	prog     = flag.String("prog", "", "Name of the program source to parse.")
	httpPort = flag.String("http_port", "", "Port number to run HTTP server on.")
)

type dotter struct {
	w        io.Writer
	id       int
	parentID []int // id of the parent node
}

func (d *dotter) nextID() int { _ = "STUB: not implemented"; return 0 }

func (d *dotter) emitNode(id int, node ast.Node) { _ = "STUB: not implemented"; return }

func (d *dotter) emitLine(src, dst int) { _ = "STUB: not implemented"; return }

func (d *dotter) VisitBefore(node ast.Node) (ast.Visitor, ast.Node) {
	_ = "STUB: not implemented"
	return *new(ast.Visitor), *new(ast.Node)
}

func (d *dotter) VisitAfter(node ast.Node) ast.Node {
	_ = "STUB: not implemented"
	return *new(ast.Node)
}

func makeDot(name string, w io.Writer) error { _ = "STUB: not implemented"; return nil }

func main() {
	flag.Parse()

	if *prog == "" {
		glog.Exitf("No -prog given")
	}

	if *httpPort == "" {
		glog.Exit(makeDot(*prog, os.Stdout))
	}

	http.HandleFunc("/",
		func(w http.ResponseWriter, _ *http.Request) {
			dot := exec.Command("dot", "-Tsvg")
			in, err := dot.StdinPipe()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			out, err := dot.StdoutPipe()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = dot.Start()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = makeDot(*prog, in)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = in.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Add("Content-type", "image/svg+xml")
			w.WriteHeader(http.StatusOK)
			_, err = io.Copy(w, out)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			err = dot.Wait()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
	http.HandleFunc("/favicon.ico", mtail.FaviconHandler)
	glog.Info(http.ListenAndServe(fmt.Sprintf(":%s", *httpPort), nil))
}
