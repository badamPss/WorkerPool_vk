package main

import "strconv"

func main() {
	const numJobs = 5
	const numWorkers = 3
	jobs := make(chan string, numJobs)
	results := make(chan string, numJobs)

	workerpool := NewWorkerPool(jobs, results)

	for i := 0; i < numWorkers; i++ {
		workerpool.Add()
	}

	workerpool.Start()

	workerpool.Remove(1)

	for j := 1; j <= numJobs; j++ {
		jobs <- strconv.Itoa(j)
	}

	for a := 1; a <= numJobs; a++ {
		<-results
	}
	workerpool.Stop()
}
