package bus

import (
	"os/exec"
)

func getUUID() (string, error) {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		return "", err
	}
	out = out[:len(out)-1]
	return string(out), nil
}

// Auto represents an automobile with an ID and
// a person count.
type Auto struct {
	ID    string `json:"ID"`
	Count int    `json:"Count"`
}

// State represents the number of autos in a state along
// with the autos themeselves.
type State struct {
	NumAutos int    `json:"NumAutos"`
	Autos    []Auto `json:"Autos"`
}

// NewState makes a new slice of autos with the given parameters.
// An error may occur if the uuidgen command is not present.
func NewState(nAutos, initCount int) (State, error) {
	tAutos := make([]Auto, nAutos)
	var err error

	for i := range tAutos {
		tAutos[i].ID, err = getUUID()
		if err != nil {
			return State{}, err
		}
		tAutos[i].Count = initCount
	}

	return State{
		Autos:    tAutos,
		NumAutos: nAutos,
	}, nil
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
