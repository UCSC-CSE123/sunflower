package cmd

import (
	"net/http"

	"github.com/UCSC-CSE123/sunflower/api/sf"
	"github.com/UCSC-CSE123/sunflower/pkg/bus"
)

func addRoutes(inputs args) {
	// Add all routes like this:
	// http.HandleFunc("/api/endpoint", func)
	updateFunc := bus.SemiRealisticSimWithoutAutoAdditions(inputs.Seed, inputs.StopDuration, inputs.Delta, inputs.StopProbability)
	handler := sf.CustomServe(inputs.Autos, inputs.InitialCount, inputs.StopPeriod, updateFunc)
	http.HandleFunc("/api/state", handler)
}
