package cmd

import (
	"net/http"

	"github.com/UCSC-CSE123/sunflower/api/sf"
)

func addRoutes(inputs args) {
	// Add all routes like this:
	// http.HandleFunc("/api/endpoint", func)
	http.HandleFunc("/api/state", sf.Serve(inputs.Autos, inputs.InitialCount, inputs.Delta, inputs.Duration))
}
