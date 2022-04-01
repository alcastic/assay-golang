package memo

import (
	"fmt"
	"time"
)

type MemoFunc func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed whe value is stored
}

func (e *entry) call(f MemoFunc, s string) {
	e.res.value, e.res.err = f(s)
	close(e.ready)
}

func (e *entry) deliver(req request, st time.Time, cached bool) {
	<-e.ready
	fmt.Printf("url: %v, cached: %t, total_time: %v\n", req.url, cached, time.Since(st))
	req.response <- e.res
}

type request struct {
	url      string
	response chan result
}

type MemCache struct {
	fetch    MemoFunc
	requests chan request
}

func (m *MemCache) monitor() {
	cache := make(map[string]*entry)
	for req := range m.requests {
		st := time.Now()
		cached := false
		ent := cache[req.url]
		if ent == nil {
			ent = &entry{ready: make(chan struct{})}
			cache[req.url] = ent
			go ent.call(m.fetch, req.url)
		} else {
			cached = true
		}
		go ent.deliver(req, st, cached)
	}
}

func (mc *MemCache) Call(url string) result {
	req := request{
		url:      url,
		response: make(chan result),
	}
	mc.requests <- req
	r := <-req.response
	close(req.response)
	return r
}

func New(fetch MemoFunc) *MemCache {
	memo := &MemCache{
		fetch:    fetch,
		requests: make(chan request),
	}
	go memo.monitor()
	return memo
}

func Close(m *MemCache) {
	close(m.requests)
}
