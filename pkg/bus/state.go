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

type Auto struct {
	ID    string `json:"ID"`
	Count int    `json:"Count"`
}

type State struct {
	NumAutos int    `json:"NumAutos"`
	Autos    []Auto `json:"Autos"`
}

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

func (bus *Auto) UpdateCount(delta int) {
	if bus.Count+delta < 0 {
		bus.Count = 0
		return
	}
	bus.Count += delta
}
