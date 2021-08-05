package hitcounter

import (
	"sync"
	"testing"
)

func TestHit(t *testing.T) {
	expected := uint(1)

	h := &Hit{}
	h.Hit(expected)

	if h.Counter[expected] != expected {
		t.Errorf("unexpected result, expecting %v got %v", expected, h.Counter[expected])
	}
}

func TestGetHits(t *testing.T) {
	expected := uint(1)

	h := &Hit{}
	h.Hit(expected)

	ret := h.GetHits(expected)

	if ret != expected {
		t.Errorf("unexpected result, expecting %v got %v", expected, ret)
	}
}

func TestHitThreads(t *testing.T) {
	var wg sync.WaitGroup

	h := &Hit{}

	for i := uint(0); i < 500; i++ {
		//Define the add outside of thread
		//or else race condition.
		wg.Add(1)

		go func(i uint) {
			h.Hit(i)
			wg.Done()
		}(i)

	}

	//Wait for all threads to complete.
	wg.Wait()

	ret := h.GetHits(400)
	if ret != uint(300) {
		t.Errorf("unexpected result, expecting %v, got %v", 300, ret)
	}
}

func CreateWorkEnv() *Hit {
	var wg sync.WaitGroup

	h := &Hit{}

	for i := uint(0); i < 500; i++ {
		//Define the add outside of thread
		//or else race condition.
		wg.Add(1)

		go func(i uint) {
			h.Hit(i)
			wg.Done()
		}(i)

	}

	//Wait for all threads to complete.
	wg.Wait()

	return h
}

func TestWorkLoop(t *testing.T) {
	h := CreateWorkEnv()

	start := uint(0)
	end := uint(300)

	var ret uint
	ret = getLoop(nil, start, end)
	if ret != uint(0) {
		t.Error("unexpected value with nil map")
	}

	ret = getLoop(h.Counter, end, start)
	if ret != uint(0) {
		t.Error("unexpected value with start > end")
	}

	ret = getLoop(h.Counter, start, end)
	if ret == uint(0) {
		t.Error("unexpected value with valid data")
	}
}

func BenchmarkWorkLoop(b *testing.B) {
	start := uint(0)
	end := uint(300)

	data := map[uint]uint{
		0: 1,
		2: 65,
		3: 35,
		4: 6435,
		5: 64345,
	}

	for i := 0; i < b.N; i++ {
		ret := getLoop(data, start, end)
		if ret == uint(0) {
			b.Error("unexpected result")
		}
	}
}
