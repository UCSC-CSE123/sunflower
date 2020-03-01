package bus

import "github.com/google/uuid"

func getUUID() string {
	return uuid.New().String()
}

const (
	// Loading represents the loading state.
	Loading = "Loading"
	// InTransit represents the in transit state.
	InTransit = "In Transit"
)

// Auto represents an automobile with an ID and
// a person count.
type Auto struct {
	ID     string `json:"ID"`
	Count  int    `json:"Count"`
	Status string `json:"Status"`
}

// State represents the number of autos in a state along
// with the autos themeselves.
type State struct {
	NumAutos int    `json:"NumAutos"`
	Autos    []Auto `json:"Autos"`
}

// NewState makes a new slice of autos with the given parameters.
func NewState(nAutos, initCount int) State {
	tAutos := make([]Auto, nAutos)

	for i := range tAutos {
		tAutos[i].ID = getUUID()

		tAutos[i].Count = initCount
		tAutos[i].Status = InTransit
	}

	return State{
		Autos:    tAutos,
		NumAutos: nAutos,
	}
}

// UpdateCount updates the internal count of an auto.
// If the the update causes the auto's count to fall below zero,
// it's count is then set to zero instead of going negative.
func (bus *Auto) UpdateCount(delta int) {
	if bus.Count+delta < 0 {
		bus.Count = 0
		return
	}
	bus.Count += delta
}
