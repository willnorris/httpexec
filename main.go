// Copyright 2021 Will Norris
// SPDX-License-Identifier: MIT-0

// httpexec executes a command in response to HTTP requests. Requests can
// optionally require a password using HTTP BasicAuth.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

var (
	addr     = flag.String("addr", "localhost:8080", "TCP address to listen on")
	password = flag.String("password", "", "required HTTP Basic Auth password")
	command  = flag.String("command", "/bin/false", "command to execute")
)

func main() {
	flag.Parse()

	server := &http.Server{
		Addr:    *addr,
		Handler: http.HandlerFunc(serveHTTP),
	}
	fmt.Printf("Listening on %s\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}

func serveHTTP(w http.ResponseWriter, r *http.Request) {
	if *password != "" {
		_, pass, _ := r.BasicAuth()
		if *password != pass {
			w.WriteHeader(http.StatusForbidden)
			log.Print("403 Forbidden")
			return
		}
	}
	args := strings.Fields(*command)
	cmd := exec.Command(args[0], args[1:]...)
	log.Printf("executing %s", cmd)
	output, err := cmd.CombinedOutput()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(output)
}
