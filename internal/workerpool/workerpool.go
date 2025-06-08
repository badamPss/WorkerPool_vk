package workerpool

type WorkerPool struct {
	jobs    chan string
	results chan<- string
	workers map[int]*Worker
	nextID  int
}

func NewWorkerPool(jobs chan string, result chan<- string) *WorkerPool {
	return &WorkerPool{
		jobs:    jobs,
		results: result,
		workers: make(map[int]*Worker)}
}

func (wp *WorkerPool) Add() {
	worker := NewWorker(wp.nextID, wp.jobs, wp.results)
	wp.workers[wp.nextID] = worker
	worker.Run()
	wp.nextID++
}

func (wp *WorkerPool) Remove(id int) {
	if worker, ok := wp.workers[id]; ok {
		worker.Stop()
		delete(wp.workers, id)
	}
}

func (wp *WorkerPool) Start() {
	for _, worker := range wp.workers {
		go worker.Run()
	}
}

func (wp *WorkerPool) Stop() {
	close(wp.jobs)
}

func (wp *WorkerPool) Results() chan<- string {
	return wp.results
}
