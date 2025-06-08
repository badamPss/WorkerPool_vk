package main

import (
	"strconv"
	"workerpool/internal/workerpool"
)

func main() {
	const numJobs = 10
	const numWorkers = 5
	jobs := make(chan string, numJobs)
	results := make(chan string, numJobs)

	workerpool := workerpool.NewWorkerPool(jobs, results)

	for i := 0; i < numWorkers; i++ {
		workerpool.Add()
	}

	workerpool.Remove(1)
	workerpool.Remove(2)

	workerpool.Add()

	for j := 1; j <= numJobs; j++ {
		jobs <- "task " + strconv.Itoa(j)
	}

	for a := 1; a <= numJobs; a++ {
		<-results
	}
	workerpool.Stop()
}
