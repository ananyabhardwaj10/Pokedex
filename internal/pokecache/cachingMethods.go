package pokecache
import(
	"time"
)

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		data: make(map[string]cacheEntry),
		interval: interval,
	}
	go c.reapLoop()

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	ce := cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.data[key] = ce
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now()
		c.mu.Lock()

		for k, v := range c.data {
			if now.Sub(v.createdAt) > c.interval {
				delete(c.data, k)
			}
		}
		c.mu.Unlock()
	}
}