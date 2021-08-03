package main

import (
	"fmt"
	LRU "lrucache/lru"
)

func main() {
	LRUcache_v1 := LRU.NewCache(2)
	LRUcache_v1.Put(1, 10)
	LRUcache_v1.Put(2, 20)
	LRUcache_v1.Put(3, 30)
	LRUcache_v1.Put(1, 110)
	fmt.Println(LRUcache_v1.Get(1))

	LRUcache_v2 := LRU.New(2)
	LRUcache_v2.Put(1, 10)
	LRUcache_v2.Put(2, 20)
	LRUcache_v2.Put(3, 30)
	LRUcache_v2.Put(1, 110)
	fmt.Println(LRUcache_v1.Get(1))

}
