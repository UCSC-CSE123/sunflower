package sf

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/UCSC-CSE123/sunflower/pkg/bus"
)

var state bus.State
var mutex sync.RWMutex

func Serve(numBuses, initialCount, delta int, interval time.Duration) (http.HandlerFunc, error) {
	var err error
	state, err = bus.NewState(numBuses, initialCount)
	if err != nil {
		return nil, err
	}
	rand.Seed(time.Now().UnixNano())
	max := delta
	min := -delta
	go func() {
		for range time.Tick(interval) {
			mutex.Lock()
			for i := range state.Autos {
				deltaNot := rand.Intn(max-min) + min
				state.Autos[i].UpdateCount(deltaNot)
			}
			mutex.Unlock()
		}
	}()

	return accessState, nil
}

func accessState(w http.ResponseWriter, r *http.Request) {
	mutex.RLock()
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(state)
	mutex.RUnlock()
}
