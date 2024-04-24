package cache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	users map[string]cacheItem
	mu    *sync.RWMutex
}

type cacheItem struct {
	value interface{}
	ttl   time.Time
}

func New(mu *sync.RWMutex) Cache {

	return Cache{
		users: make(map[string]cacheItem),
		mu:    mu,
	}
}

func (c *Cache) Set(key string, item interface{}, ttl time.Duration) {
	c.mu.Lock()
	c.users[key] = cacheItem{value: item, ttl: time.Now().Add(ttl)}
	c.mu.Unlock()

}

func (c *Cache) Get(key string) interface{} {

	c.mu.Lock()
	defer c.mu.Unlock()

	data, ex := c.users[key]
	if !ex {
		fmt.Println(" net takogo")
		return nil
	}

	if time.Now().After(data.ttl){
		delete(c.users, key)
		return nil
	}

	return data.value

}

func (c *Cache) Delete(key string) {

	c.mu.Lock()
	delete(c.users, key)
	c.mu.Unlock()
}
