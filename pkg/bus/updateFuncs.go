package bus

import (
	"math/rand"
	"sync"
	"time"
)

func returnTrueWithPercent(n int) bool {
	return rand.Intn(100) < n
}

// SemiRealisticSimWithoutAutoAdditions returns a function updater that
// updates the state as follows:
//
// * Autos have a `probability` probability to stop when the function is executed (see: CustomServe).
// * If an auto does stop it will take `waitTime` time to complete the stop.
// * Once a stop is completed the auto population will be changed by rand(-passengerDelta, passengerDelta).
// * Autos stop independently of each other (ie one auto does NOT have to wait for another auto to finish it stop before it can stop).
// * Autos are NOT added nor removed from the state.
// * However, autos do stop (if the probability is in their favor) at the same time.
func SemiRealisticSimWithoutAutoAdditions(waitTime time.Duration, passengerDelta, probability int) func(*State, *sync.RWMutex) {

	return func(iState *State, mutex *sync.RWMutex) {

		// We'll need a wait group.
		var jobs sync.WaitGroup

		// and calculate the min and max delta.
		min := -passengerDelta
		max := passengerDelta

		// Run through all the buses in the state.
		for i := range iState.Autos {

			// Roll a weighted die to see if we'll actually stop.
			if !returnTrueWithPercent(probability) {
				// If the die is not in our favor,
				// we'll just skip this stop and go to the next stop.
				continue
			}

			// Otherwise...
			//
			// We have a job to do!
			jobs.Add(1)

			// We'll have to update the state
			// here to say that it's loading
			// which requires a write lock
			mutex.Lock()

			// Update the bus' status.
			iState.Autos[i].Status = Loading

			// We can unlock the mutex here since we'll be sleeping for some time.
			mutex.Unlock()

			// We want buses to stop in parallel so we'll execute this in a go routine
			// so that the next bus can stop (or not) as well without having to wait for us.
			go func(index int) {
				// Simulate the wait time by sleeping.
				time.Sleep(waitTime)

				// Once the sleep is done update the info
				// this requires a lock.
				mutex.Lock()

				// Update the count with a random delta [-delta,delta]
				deltaNaught := rand.Intn(max-min) + min
				iState.Autos[index].UpdateCount(deltaNaught)
				iState.Autos[index].Status = InTransit

				// We'll need to signal that this job is done.
				jobs.Done()

				// Then unlock the mutex so another go routine can update it's state as well.
				mutex.Unlock()

			}(i) // Capturing a loop variable in a go routine is a no-no. So we'll pass it in as a parameter.
		}

		// Now we just wait
		jobs.Wait()
	}
}
