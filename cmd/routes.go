package cmd

import (
	"net/http"

	"github.com/UCSC-CSE123/sunflower/api/sf"
)

func addRoutes() {
	// Add all routes like this:
	// http.HandleFunc("/api/endpoint", func)
	http.HandleFunc("/api/state", sf.AccessState)
}
