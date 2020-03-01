package sf

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/UCSC-CSE123/sunflower/pkg/bus"
)

// Serve initializes a state with the given parameters.
// That state is then asynchronously updated according the given interval.
// Both reads and writes are thread safe.
// Instantiating this function n times yields n distinct states that are accessible via the return http.HandlerFunc paramater.
func Serve(numBuses, initialCount, delta int, interval time.Duration) http.HandlerFunc {
	var mutex sync.RWMutex
	state := bus.NewState(numBuses, initialCount)

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

	return func(w http.ResponseWriter, r *http.Request) {
		mutex.RLock()
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(state)
		mutex.RUnlock()
	}
}
