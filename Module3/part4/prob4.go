package part4

import (
	"fmt"
	"sync"
	"time"
)

type Config struct{
	data map[string]string
	mu sync.RWMutex
}

func (c *Config) Get(key string){
	c.mu.RLock()
	defer c.mu.RUnlock()
	fmt.Println(c.data[key])
}

func (c *Config) Set(key, value string){
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func Prob4(){
	c1 := Config{
		data: make(map[string]string),
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		time.Sleep(time.Second * 1)
		c1.Set("a1", "value1")
		wg.Done()
	}()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(){
			time.Sleep(time.Millisecond * 1000)
			c1.Get("a1")
			wg.Done()
		}()
	}
	wg.Wait()
}