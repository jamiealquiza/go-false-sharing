package main

import (
	"sync"
	"testing"
)

type paddedMutex struct {
	p0 [64]byte
	m  *sync.Mutex
}

type paddedCounter struct {
	p0 [64]byte
	c  uint32
}

/*
From https://golang.org/src/sync/mutex.go:
type Mutex struct {
	state int32
	sema  uint32
}
*/

func BenchmarkNonPaddedMutex(b *testing.B) {
	b.StopTimer()

	ms := [4]*sync.Mutex{}
	ms[0], ms[1] = &sync.Mutex{}, &sync.Mutex{}
	ms[2], ms[3] = &sync.Mutex{}, &sync.Mutex{}

	wg := &sync.WaitGroup{}
	wg.Add(4)

	b.StartTimer()

	go func(w *sync.WaitGroup) {
		for i := 0; i < b.N; i++ {
			if i%2 == 0 {
				ms[0].Lock()
			} else {
				ms[0].Unlock()
			}
		}
		wg.Done()
	}(wg)

	go func(w *sync.WaitGroup) {
		for i := 0; i < b.N; i++ {
			if i%2 == 0 {
				ms[1].Lock()
			} else {
				ms[1].Unlock()
			}
		}
		wg.Done()
	}(wg)

	go func(w *sync.WaitGroup) {
		for i := 0; i < b.N; i++ {
			if i%2 == 0 {
				ms[2].Lock()
			} else {
				ms[2].Unlock()
			}
		}
		wg.Done()
	}(wg)

	go func(w *sync.WaitGroup) {
		for i := 0; i < b.N; i++ {
			if i%2 == 0 {
				ms[3].Lock()
			} else {
				ms[3].Unlock()
			}
		}
		wg.Done()
	}(wg)

	wg.Wait()
}

func BenchmarkPaddedMutex(b *testing.B) {
	b.StopTimer()

	ms := [4]*paddedMutex{}
	ms[0], ms[1] = &paddedMutex{m: &sync.Mutex{}}, &paddedMutex{m: &sync.Mutex{}}
	ms[2], ms[3] = &paddedMutex{m: &sync.Mutex{}}, &paddedMutex{m: &sync.Mutex{}}

	wg := &sync.WaitGroup{}
	wg.Add(4)

	b.StartTimer()

	go func(w *sync.WaitGroup) {
		for i := 0; i < b.N; i++ {
			if i%2 == 0 {
				ms[0].m.Lock()
			} else {
				ms[0].m.Unlock()
			}
		}
		wg.Done()
	}(wg)

	go func(w *sync.WaitGroup) {
		for i := 0; i < b.N; i++ {
			if i%2 == 0 {
				ms[1].m.Lock()
			} else {
				ms[1].m.Unlock()
			}
		}
		wg.Done()
	}(wg)

	go func(w *sync.WaitGroup) {
		for i := 0; i < b.N; i++ {
			if i%2 == 0 {
				ms[2].m.Lock()
			} else {
				ms[2].m.Unlock()
			}
		}
		wg.Done()
	}(wg)

	go func(w *sync.WaitGroup) {
		for i := 0; i < b.N; i++ {
			if i%2 == 0 {
				ms[3].m.Lock()
			} else {
				ms[3].m.Unlock()
			}
		}
		wg.Done()
	}(wg)

	wg.Wait()
}

func BenchmarkNonPaddedCounter(b *testing.B) {
	b.StopTimer()

	c := [4]uint32{}

	wg := &sync.WaitGroup{}
	wg.Add(4)

	b.StartTimer()

	go func(w *sync.WaitGroup) {
		for i := 0; i < b.N; i++ {
			c[0] = uint32(b.N)
		}
		wg.Done()
	}(wg)

	go func(w *sync.WaitGroup) {
		for i := 0; i < b.N; i++ {
			c[1] = uint32(b.N)
		}
		wg.Done()
	}(wg)

	go func(w *sync.WaitGroup) {
		for i := 0; i < b.N; i++ {
			c[2] = uint32(b.N)
		}
		wg.Done()
	}(wg)

	go func(w *sync.WaitGroup) {
		for i := 0; i < b.N; i++ {
			c[3] = uint32(b.N)
		}
		wg.Done()
	}(wg)

	wg.Wait()
}

func BenchmarkPaddedCounter(b *testing.B) {
	b.StopTimer()

	c := [4]*paddedCounter{&paddedCounter{}, &paddedCounter{}, &paddedCounter{}, &paddedCounter{}}

	wg := &sync.WaitGroup{}
	wg.Add(4)

	b.StartTimer()

	go func(w *sync.WaitGroup) {
		for i := 0; i < b.N; i++ {
			c[0].c = uint32(b.N)
		}
		wg.Done()
	}(wg)

	go func(w *sync.WaitGroup) {
		for i := 0; i < b.N; i++ {
			c[1].c = uint32(b.N)
		}
		wg.Done()
	}(wg)

	go func(w *sync.WaitGroup) {
		for i := 0; i < b.N; i++ {
			c[2].c = uint32(b.N)
		}
		wg.Done()
	}(wg)

	go func(w *sync.WaitGroup) {
		for i := 0; i < b.N; i++ {
			c[3].c = uint32(b.N)
		}
		wg.Done()
	}(wg)

	wg.Wait()
}
