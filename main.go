package main

import (
	"memory_cache/cache"
	"fmt"
	"sync"
	"time"
)

func main() {

	var mu sync.RWMutex
	cache := cache.New(&mu)
	cache.Set("Mouse", 45, time.Second*5)
	get := cache.Get("Mouse")
	fmt.Println(get)
	cache.Delete("Mouse")
	get = cache.Get("Mouse")
	fmt.Println(get)

}
