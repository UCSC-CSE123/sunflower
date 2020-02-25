package cmd

import (
	"log"
	"net"
	"net/http"
)

// Execute is the starting point for all
// CLI's.
func Execute() error {
	args := getFlags()

	if err := serverSetup(args.Host, args.Port); err != nil {
		return err
	}

	return nil
}

// Initializes the default http server using
// the given host and port.
func serverSetup(host, port string) error {
	// Add all the routes.
	addRoutes()

	// Start the server.
	addr := net.JoinHostPort(host, port)
	log.Println("starting server at", addr)

	// This should only return on an error.
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		return err
	}

	return nil
}
