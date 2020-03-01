package sf

import "github.com/UCSC-CSE123/sunflower/pkg/bus"

type DebugInfo struct {
	StopPeriodicity string `json:"StopPeriodicity"`
	InitialCount    int    `json:"InitialCount"`
}

type StateWDebug struct {
	bus.State `json:"State"`
	Info      DebugInfo `json:"DebugInfo"`
}
