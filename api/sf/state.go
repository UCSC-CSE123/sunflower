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

// Serve initializes a state with the given parameters.
// That state is then asynchronously updated according the given interval.
// Both reads and writes are thread safe.
func Serve(numBuses, initialCount, delta int, interval time.Duration) http.HandlerFunc {
	state = bus.NewState(numBuses, initialCount)
	rand.Seed(time.Now().UnixNano())
	max := delta
	min := -delta
	go func() {
		for range time.Tick(interval) {
			mutex.Lock()
			for i := range state.Autos {
				deltaNaught := rand.Intn(max-min) + min
				state.Autos[i].UpdateCount(deltaNaught)
			}
			mutex.Unlock()
		}
	}()

	return accessState
}

func accessState(w http.ResponseWriter, r *http.Request) {
	mutex.RLock()
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(state)
	mutex.RUnlock()
}
