package strategy

import "fmt"

type EvictionAlgo interface {
	Evict(c *Cache)
}

type Fifo struct{}

func (l *Fifo) Evict(c *Cache) {
	fmt.Println("Evicting by fifo strtegy")
}

type Lru struct{}

func (l *Lru) Evict(c *Cache) {
	fmt.Println("Evicting by lru strtegy")
}

type Lfu struct{}

func (l *Lfu) Evict(c *Cache) {
	fmt.Println("Evicting by lfu strtegy")
}
