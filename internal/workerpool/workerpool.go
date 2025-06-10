package workerpool

import (
	"fmt"
	"sync"
)

type WorkerPool struct {
	jobs    chan string
	results chan<- string
	workers map[int]*Worker
	nextID  int
	mu      sync.Mutex
}

func NewWorkerPool(jobs chan string, result chan<- string) *WorkerPool {
	return &WorkerPool{
		jobs:    jobs,
		results: result,
		workers: make(map[int]*Worker)}
}

func (wp *WorkerPool) Add() {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	worker := NewWorker(wp.nextID, wp.jobs)
	wp.workers[wp.nextID] = worker
	fmt.Printf("Worker %d added\n", wp.nextID)
	wp.nextID++
	worker.Run(wp.results)
}

func (wp *WorkerPool) Remove(id int) {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	if worker, ok := wp.workers[id]; ok {
		worker.Stop()
		delete(wp.workers, id)
		fmt.Printf("Worker %d removed\n", id)
	}
}

func (wp *WorkerPool) Stop() {
	close(wp.jobs)
	fmt.Println("Worker pool stopped")
}
