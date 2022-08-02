// Usage:
// Put me together with your main.
// Build with `pprof` added to `-tags`.

//go:build pprof
// +build pprof

package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

// Add `-X main.PPROF_SEVER=` with a new address to `-ldflags` to change this.
var PPROF_SEVER = "localhost:6060"

func init() {
	go func() {
		log.Println(http.ListenAndServe(PPROF_SEVER, nil))
	}()
}
