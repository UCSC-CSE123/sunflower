package cmd

import (
	"net/http"
	"time"

	"github.com/UCSC-CSE123/sunflower/api/sf"
)

func addRoutes() {
	// Add all routes like this:
	// http.HandleFunc("/api/endpoint", func)
	http.HandleFunc("/api/state", sf.Serve(10, 100, 25, 500*time.Millisecond))
}
