// Example Hit Counter in Go-Lang with Mutex.
// API Has two functions, Hit(uint) and GetHits(uint)
// Has one struct Hit
// Hit should be called, with uint as the timestamp, on when a hit occurs
// GetHits should return the number of hits at that timestamp and 5 minutes before.
// uint represents seconds, so uint-300.
// should be thread safe.
// should not worry about timestamp value and relative system time.
// should not worry about overflowing uint.

package hitcounter

import (
	"sync"
)

const ()

var ()

type Hit struct {
	Counter map[uint]uint
	mtx     sync.RWMutex
}

// Hit adds one hit to the counter at TimeStamp
func (h *Hit) Hit(ts uint) {

	//Since we intend to write to a shared resource
	//Locking with writing mutex.
	h.mtx.Lock()
	defer h.mtx.Unlock()

	// Initialize Map if nil
	if h.Counter == nil {
		h.Counter = make(map[uint]uint)
	}

	h.Counter[ts] += 1
}

func (h *Hit) GetHits(ts uint) (ret uint) {
	var start uint //defaults to zero

	// address underflow problem
	if ts > 300 {
		start = ts - 300
	}

	//Since we intend to read a shared resource
	//Locking with read mutex.
	h.mtx.RLock()
	defer h.mtx.RUnlock()

	if h.Counter == nil {
		return
	}

	return getLoop(h.Counter, start, ts)
}

// Main work here, seperated out for benchmark tests.
func getLoop(data map[uint]uint, start, end uint) (ret uint) {
	if start > end || data == nil {
		//invalid options, return 0
		return
	}

	for k, v := range data {
		if k > start && k <= end {
			ret += v
		}
	}

	return
}
