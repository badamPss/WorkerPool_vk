package workerpool

import (
	"fmt"
)

type Worker struct {
	id   int
	jobs <-chan string
	quit chan struct{}
}

func NewWorker(id int, jobs <-chan string) *Worker {
	return &Worker{
		id:   id,
		jobs: jobs,
		quit: make(chan struct{})}
}

func (w *Worker) Run(results chan<- string) {
	fmt.Printf("Starting worker %d\n", w.id)
	go func() {
		for {
			select {
			case <-w.quit:
				fmt.Println("worker", w.id, "stopping")
				return
			case job, ok := <-w.jobs:
				if !ok {
					fmt.Println("worker", w.id, "stopping: jobs channel closed")
					return
				}
				fmt.Println("worker", w.id, "started job", job)
				results <- job
				fmt.Println("worker", w.id, "finished job", job)
			}
		}
	}()
}

func (w *Worker) Stop() {
	close(w.quit)
}
