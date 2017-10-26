package main

import (
	"log"
	"math/rand"
	"strconv"
	"time"
)

func client(index int, c *Cache) {
	for {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
		key := strconv.Itoa(rand.Intn(100000))
		value := key + " value"

		log.Printf("client %d setting %s=%s", index, key, value)
		c.Set(key, value, time.Second*10)
		if value2, found := c.Get(key); !found || value2 != value {
			log.Fatal("incorrect value returned from cache")
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	c := NewCache()

	for i := 0; i < 10; i++ {
		//'go' keyword runs the function on the right of the keyword, but on a new go routine
		//concurrent running
		//**when the main function exits, all other go routines exit
		//so if we don't do anything, all the routines start, but immediately exits
		go client(i, c)
	}

	time.Sleep(time.Minute)

}
