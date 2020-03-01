package sf

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/UCSC-CSE123/sunflower/pkg/bus"
)

// ServeInstantChange initializes a state with the given parameters.
// That state is then asynchronously updated according the given interval.
// Both reads and writes are thread safe.
// Instantiating this function n times yields n distinct states that are accessible via the return http.HandlerFunc paramater.
func ServeInstantChange(numBuses, initialCount, delta int, interval time.Duration) http.HandlerFunc {

	updateFunc := func(internalState *bus.State) {
		rand.Seed(time.Now().UnixNano())
		max := delta
		min := -delta
		for i := range internalState.Autos {
			deltaNaught := rand.Intn(max-min) + min
			internalState.Autos[i].UpdateCount(deltaNaught)
		}
	}

	return CustomServe(numBuses, initialCount, interval, updateFunc)
}

// CustomServe instantiate a bus state with numBuses and initialCount.
// It then updates that state asynchronously every interval according to updateFunc.
// It returns an http.HandlerFunc that has access to that state.
func CustomServe(numBuses, initialCount int, interval time.Duration, updateFunc func(*bus.State)) http.HandlerFunc {
	var mutex sync.RWMutex
	state := bus.NewState(numBuses, initialCount)

	go func() {
		for range time.Tick(interval) {
			mutex.Lock()
			updateFunc(&state)
			mutex.Unlock()
		}
	}()

	return func(w http.ResponseWriter, r *http.Request) {
		mutex.RLock()
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(state)
		mutex.RUnlock()
	}
}
