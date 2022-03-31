package memo

type result struct {
	value interface{}
	err   error
}

type MemCache struct {
	f     MemoFunc
	cache map[string]*result
}

func (m *MemCache) Call(s string) *result {
	r, ok := m.cache[s]
	if !ok {
		data, err := m.f(s)
		r = &result{data, err}
		m.cache[s] = r
	}
	return r
}

type MemoFunc func(string) (interface{}, error)

func New(mf MemoFunc) *MemCache {
	return &MemCache{
		f:     mf,
		cache: make(map[string]*result),
	}
}
