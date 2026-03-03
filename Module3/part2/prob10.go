package part2

import (
	"fmt"
	"sync"
	"time"
)

func Prob10(){
	tasks := make(chan string, 3)
	var wg sync.WaitGroup
	
	tasks <- "task1"
	tasks <- "task2"
	tasks <- "task3"

	fmt.Println("All the tasks have been sent.")

	close(tasks)

	wg.Add(1)
	go func(){
		defer wg.Done()
		for task := range tasks{
			time.Sleep(time.Second * 5)
			fmt.Println(task)
	}
	}()

	wg.Wait()
	fmt.Println("All tasks are done.")
}