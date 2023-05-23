// Helloserver is an HTTP server that serves the message "hello, world" on "/".
// The query parameter "whom" sets the "world" part of the message.
// The -port flag sets the server's listen port, 8080 by default.
package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	port = flag.Int("port", 8080, "the port on which the server listens for HTTP requests")
)

func main() {
	//go:guide HandleFunc registers a function as an HTTP handler with the default HTTP mux.
	// This HandleFunc call registers a handler for the "/" path.
	// The handler is provided inline as a function literal.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Read the "whom" query parameter. If it's empty, use "world".
		whom := r.URL.Query().Get("whom")
		if whom == "" {
			whom = "world"
		}
		fmt.Fprint(w, "hello, ", whom)
	})
	//go:guide ListenAndServe starts a server that listens for HTTP requests on the provided address
	// and serves requests using a provided handler. This ListenAndServe call listens on *port and
	// serves requests with the handlers registered on the default mux, as with HandleFunc above.
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
