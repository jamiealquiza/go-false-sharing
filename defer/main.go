package main

import (
	"fmt"
	"sync"
	"time"
)

type SharedThing struct {
	sync.Mutex
	someValue bool
}

func main() {
	mySharedThing := SharedThing{}

	start := time.Now()

	for i := 0; i < 50000000; i++ {
		mySharedThing.updateThingDefer(i)
	}

	fmt.Printf("Time with defer: %s\n", time.Since(start))

	start = time.Now()
	for i := 0; i < 50000000; i++ {
		mySharedThing.updateThing(i)
	}

	fmt.Printf("Time without defer: %s\n", time.Since(start))
}

func (t SharedThing) updateThingDefer(n int) {
	t.Lock()
	defer t.Unlock()

	if n%2 == 1 {
		t.someValue = true
		return
	}

	t.someValue = false
}

func (t SharedThing) updateThing(n int) {
	t.Lock()

	if n%2 == 1 {
		t.someValue = true
		t.Unlock()
		return
	}

	t.someValue = false
	t.Unlock()
}
