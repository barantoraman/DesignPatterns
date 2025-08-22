package strategy

type Store interface {
	SetEvictionAlgo(e EvictionAlgo)
	Add(key, val string)
	Get(key string)
	Evict()
}

type Cache struct {
	Storage      map[string]string
	EvictionAlgo EvictionAlgo
	Cap          int
	MaxCap       int
}

func InitCache(e EvictionAlgo) *Cache {
	return &Cache{
		Storage:      make(map[string]string),
		EvictionAlgo: e,
		Cap:          0,
		MaxCap:       2,
	}
}

func (c *Cache) SetEvictionAlgo(e EvictionAlgo) {
	c.EvictionAlgo = e
}

func (c *Cache) Add(key, val string) {
	if c.Cap == c.MaxCap {
		c.Evict()
	}
	c.Cap++
	c.Storage[key] = val
}

func (c *Cache) Get(key string) {
	delete(c.Storage, key)
}

func (c *Cache) Evict() {
	c.EvictionAlgo.Evict(c)
	c.Cap--
}
