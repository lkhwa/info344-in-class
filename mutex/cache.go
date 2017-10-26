package main

import (
	"fmt"
	"sync"
	"time"
)

type entry struct {
	value     string
	expiresAt time.Time
}

//Cache is a TTL (Time to Live) cache that is safe for concurrent use
//add something to cache, keep for ___ time, then the entry is removed automaticallys
type Cache struct {
	entries map[string]*entry
	//TODO: protect this for concurrent use!
	mx sync.RWMutex
}

//NewCache constructs a new Cache object
func NewCache() *Cache {
	c := &Cache{
		entries: map[string]*entry{},
	}
	go c.janitor()
	return c
}

//Set adds a key/value to the cache
func (c *Cache) Set(key string, value string, timeToLive time.Duration) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.entries[key] = &entry{value, time.Now().Add(timeToLive)}
}

//Get gets the value associated with a key
func (c *Cache) Get(key string) (string, bool) {
	//TODO: implement this
	c.mx.RLock()         //read
	defer c.mx.RUnlock() //read
	entry, found := c.entries[key]
	if !found {
		return "", false
	}
	return entry.value, true
}

func (c *Cache) janitor() {
	for {
		time.Sleep(time.Second)
		now := time.Now()
		c.mx.Lock()
		fmt.Println("janitor is running")
		for key, entry := range c.entries {
			if entry.expiresAt.Before(now) {
				fmt.Printf("purging key %s\n", key)
				delete(c.entries, key)
			}
		}
		c.mx.Unlock() //can't use defer because this function is never going to exit, and defer only happens when fnction exits
	}
}
