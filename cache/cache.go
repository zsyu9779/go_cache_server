package cache

type Cache interface {
	Get(k string) ([]byte, error)
	Set(k string, v []byte) error
	Del(k string) error
	GetStat() Stat
}

type Stat struct {
	Count     int64
	KeySize   int64
	ValueSize int64
}

func (s *Stat) add(k string, v []byte) {
	s.Count++
	s.KeySize += int64(len(k))
	s.ValueSize += int64(len(v))
}

func (s *Stat) del(k string, v []byte) {
	s.Count--
	s.KeySize -= int64(len(k))
	s.ValueSize -= int64(len(v))
}

func New(typ string) Cache {
	var c Cache
	if typ =="inmemory" {
		c = newInMemoryCache()
	}
	return c
}



