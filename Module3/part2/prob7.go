package part2

import "fmt"

func Worker(jobs, results chan int){
	for job := range jobs{
		fmt.Println("Processing the job...")
		result := job * 2
		results <- result
	}
}

func Prob7(){
	jobs := make(chan int, 5)
	results := make(chan int, 5)


	for i := 1; i <= 3; i++{
		go Worker(jobs, results)
	}

	for i := 1; i <=5; i++{
		jobs <- i
	}
	close(jobs)

	for i:=1; i<=5; i++{
		fmt.Print(<-results)
	}
}