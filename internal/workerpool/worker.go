package workerpool

import (
	"fmt"
	"time"
)

type Worker struct {
	id      int
	jobs    <-chan string
	results chan<- string
	quit    chan struct{}
}

func NewWorker(id int, jobs <-chan string, results chan<- string) *Worker {
	return &Worker{
		id:      id,
		jobs:    jobs,
		results: results,
		quit:    make(chan struct{})}
}

func (w *Worker) Run() {
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
				time.Sleep(time.Second)
				fmt.Println("worker", w.id, "finished job", job)
				w.results <- job
			}
		}
	}()
}

func (w *Worker) Stop() {
	close(w.quit)
}
