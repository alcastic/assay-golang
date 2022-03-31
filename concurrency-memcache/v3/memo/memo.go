package memo

import (
	"fmt"
	"sync"
	"time"
)

type result struct {
	value interface{}
	err   error
}

type MemCache struct {
	f     MemoFunc
	cache map[string]result
	mutex sync.Mutex
}

func (m *MemCache) Call(s string) result {
	now := time.Now()

	m.mutex.Lock()
	r, cached := m.cache[s]
	m.mutex.Unlock()

	defer func(st time.Time, s string, c bool) {
		fmt.Printf("url: %v, cached: %t, total_time: %v\n", s, c, time.Since(st))
	}(now, s, cached)

	if !cached {
		r.value, r.err = m.f(s)
		m.mutex.Lock()
		m.cache[s] = r
		m.mutex.Unlock()
	}
	return r
}

type MemoFunc func(string) (interface{}, error)

func New(mf MemoFunc) *MemCache {
	return &MemCache{
		f:     mf,
		cache: make(map[string]result),
		mutex: sync.Mutex{},
	}
}
