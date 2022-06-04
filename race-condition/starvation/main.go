package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Value struct {
	mu       sync.Mutex
	total    int64
	greedy   int64
	generous int64
}

func main() {

	var v Value
	executionTime := 1 * time.Second
	var wg sync.WaitGroup

	greedyRoutine := func() {
		defer wg.Done()
		for start := time.Now(); time.Since(start) < executionTime; {
			v.mu.Lock()
			time.Sleep(6 * time.Nanosecond)
			v.total++
			v.greedy++
			v.mu.Unlock()
		}
	}
	generousRoutine := func() {
		defer wg.Done()
		for start := time.Now(); time.Since(start) < executionTime; {
			v.mu.Lock()
			time.Sleep(2 * time.Nanosecond)
			v.total++
			v.mu.Unlock()

			v.mu.Lock()
			time.Sleep(2 * time.Nanosecond)
			v.mu.Unlock()

			v.mu.Lock()
			time.Sleep(2 * time.Nanosecond)
			v.generous++
			v.mu.Unlock()
		}
	}

	wg.Add(2)
	go greedyRoutine()
	go generousRoutine()
	wg.Wait()
	fmt.Printf("Greedy executions   = %d (%d %% aprox.)\n", atomic.LoadInt64(&v.greedy), (atomic.LoadInt64(&v.greedy) * 100 / atomic.LoadInt64(&v.total)))
	fmt.Printf("Generous executions = %d (%d %% aprox.)\n", atomic.LoadInt64(&v.generous), (atomic.LoadInt64(&v.generous) * 100 / atomic.LoadInt64(&v.total)))
	fmt.Printf("Total executions    = %d (100 %%)\n", atomic.LoadInt64(&v.total))
}
