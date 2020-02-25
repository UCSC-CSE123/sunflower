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

func init() {
	state, _ = bus.NewState(10, 25)
	rand.Seed(time.Now().UnixNano())
	max := 10 + 1
	min := -10
	go func() {
		for range time.Tick(10 * time.Second) {

			mutex.Lock()

			for i := range state {
				state[i].UpdateCount(rand.Intn(max-min) + min)
			}

			mutex.Unlock()

		}
	}()
}

func AccessState(w http.ResponseWriter, r *http.Request) {
	mutex.RLock()
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(state)
	mutex.RUnlock()
}
