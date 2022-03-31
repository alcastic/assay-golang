package memo

import (
	"fmt"
	"sync"
	"time"
)

type entry struct {
	res   result
	ready chan struct{} // closed whe value is stored

}

type result struct {
	value interface{}
	err   error
}

type MemCache struct {
	f     MemoFunc
	mutex sync.Mutex
	cache map[string]*entry
}

func (m *MemCache) Call(s string) result {
	st := time.Now()
	cached := false

	m.mutex.Lock()
	e := m.cache[s]

	if e == nil {
		e = &entry{ready: make(chan struct{})}
		m.cache[s] = e
		m.mutex.Unlock()

		e.res.value, e.res.err = m.f(s)
		close(e.ready)
	} else {
		m.mutex.Unlock()
		<-e.ready
		cached = true
	}

	fmt.Printf("url: %v, cached: %t, total_time: %v\n", s, cached, time.Since(st))
	return e.res
}

type MemoFunc func(string) (interface{}, error)

func New(mf MemoFunc) *MemCache {
	return &MemCache{
		f:     mf,
		cache: make(map[string]*entry),
		mutex: sync.Mutex{},
	}
}
