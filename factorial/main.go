package main

import (
	"fmt"
	"sync"
)

type Job struct {
	ID    int
	Input int
}

func worker(id int, jobs <-chan Job, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("worker %d processing job %d\n", id, j.ID)
		result := factorial(j.Input)
		results <- result
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}

func main() {
	jobs := make(chan Job, 100)
	results := make(chan int, 100)
	var wg sync.WaitGroup

	numWorkers := 3
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for j := 1; j <= 9; j++ {
		jobs <- Job{ID: j, Input: j}
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()
	fmt.Println("waiting for results")
	for a := range results {
		fmt.Println(a)
	}
}
