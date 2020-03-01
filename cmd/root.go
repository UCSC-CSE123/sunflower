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

	if err := serverSetup(args); err != nil {
		return err
	}

	return nil
}

// Initializes the default http server using
// the given host and port.
func serverSetup(inputs args) error {
	// Add all the routes.
	addRoutes(inputs)

	// Start the server.
	addr := net.JoinHostPort(inputs.Host, inputs.Port)
	log.Printf("server parameters: %+v\n", inputs)
	log.Printf("starting server at http://%s/api/state", addr)

	// This should only return on an error.
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		return err
	}

	return nil
}
