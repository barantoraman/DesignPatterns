package main

import "github.com/barantoraman/DesignPatternsTraining/internal/behavioral/strategy"

func main() {
	lfu := &strategy.Lfu{}
	cache := strategy.InitCache(lfu)

	cache.Add("a", "1")
	cache.Add("b", "2")

	cache.Add("c", "3")

	lru := &strategy.Lru{}
	cache.SetEvictionAlgo(lru)

	cache.Add("d", "4")

	fifo := &strategy.Fifo{}
	cache.SetEvictionAlgo(fifo)

	cache.Add("e", "5")
}
