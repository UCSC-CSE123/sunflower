package cmd

import (
	"net/http"

	"github.com/UCSC-CSE123/sunflower/api/sf"
)

func addRoutes(inputs args) error {
	// Add all routes like this:
	// http.HandleFunc("/api/endpoint", func)
	stateHandler, err := sf.Serve(inputs.Autos, inputs.InitialCount, inputs.Delta, inputs.Duration)
	if err != nil {
		return err
	}
	http.HandleFunc("/api/state", stateHandler)
	return nil
}
