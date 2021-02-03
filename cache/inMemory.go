package cache

import "sync"

func newInMemoryCache() *inMemoryCache {
	return &inMemoryCache{
		c:     make(map[string][]byte),
		mutex: sync.RWMutex{},
		Stat:  Stat{},
	}
}
type inMemoryCache struct {
	c     map[string][]byte
	mutex sync.RWMutex
	Stat
}

func (m *inMemoryCache) Set(k string, v []byte) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	tmp, exist := m.c[k]
	if exist {
		m.del(k, tmp)
	}
	m.c[k] = v
	m.add(k, v)
	return nil
}

func (m *inMemoryCache) Get(k string) ([]byte, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.c[k], nil
}

func (m *inMemoryCache)Del(k string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	tmp, exist := m.c[k]
	if exist {
		delete(m.c, k)
		m.del(k, tmp)
	}
	return nil
}
func (m *inMemoryCache)GetStat()Stat {
	return m.Stat
}