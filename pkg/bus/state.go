package bus

import (
	"os/exec"
)

func getUUID() ([]byte, error) {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		return nil, err
	}
	return out, nil
}

type Auto struct {
	ID    []byte `json:"ID"`
	Count int    `json:"Count"`
}

type State []Auto

func NewState(nAutos, initCount int) (State, error) {
	tAutos := make([]Auto, nAutos)
	var err error

	for i := range tAutos {
		tAutos[i].ID, err = getUUID()
		if err != nil {
			return nil, err
		}
		tAutos[i].Count = initCount
	}

	return tAutos, nil
}

func (bus *Auto) UpdateCount(delta int) {
	if bus.Count+delta < 0 {
		bus.Count = 0
		return
	}
	bus.Count += delta
}
