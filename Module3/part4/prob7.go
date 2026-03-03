package part4

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	storage map[string]string
	mu      sync.RWMutex // Use ONE RWMutex for both operations
}

func (c *Cache) Read(key string) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val := c.storage[key]
	if val != ""{
		fmt.Println(c.storage[key])
	} else {
		fmt.Println("Empty")
	}
}

func (c *Cache) Write(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.storage[key] = value
}

func Prob7() {
	c1 := &Cache{
		storage: make(map[string]string),
	}
	var wg sync.WaitGroup

	// 1. Launch the Writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		fmt.Println("Writer: Writing 'America' to cache...")
		c1.Write("Server", "America")
	}()

	// 2. Launch 5 Readers that loop until they find the value
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(sec int) {
			defer wg.Done()
			val := 990 + (sec * 10)
			time.Sleep(time.Millisecond * time.Duration(val))
			c1.Read("Server")
		}(i)
	}

	wg.Wait()
	fmt.Println("Closing the program.")
}