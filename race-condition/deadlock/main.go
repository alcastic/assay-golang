package main

import (
	"sync"
	"time"
)

type Numb struct {
	value int
	mu    sync.Mutex
}

func Plus(a, b *Numb) *Numb {
	a.mu.Lock()
	defer a.mu.Unlock()

	time.Sleep(1 * time.Second)

	b.mu.Lock()
	defer b.mu.Unlock()

	return &Numb{
		value: a.value + b.value,
		mu:    sync.Mutex{},
	}
}

func main() {
	wg := sync.WaitGroup{}

	a := &Numb{
		value: 2,
		mu:    sync.Mutex{},
	}
	b := &Numb{
		value: 3,
		mu:    sync.Mutex{},
	}

	go func(wg *sync.WaitGroup, a, b *Numb) {
		wg.Add(1)
		Plus(a, b)
		defer wg.Done()
	}(&wg, a, b)

	go func(wg *sync.WaitGroup, a, b *Numb) {
		wg.Add(1)
		Plus(a, b)
		defer wg.Done()
	}(&wg, b, a)

	wg.Wait()
}
