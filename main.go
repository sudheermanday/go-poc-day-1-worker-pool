package main

import (
	"context"
	"fmt"
	"go-poc-day-1-worker-pool/task"
	"go-poc-day-1-worker-pool/worker"
	"sync"
	"time"

)

func main() {
	numWorkers := 3
	numTasks := 10
	
	tasks := make(chan task.Task, numTasks)
	
	ctx,cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var wg sync.WaitGroup

	// Start workers
	for i:=1; i<=numWorkers; i++{
		wg.Add(1)
		w := worker.Worker{ID: i}
		go w.Start(ctx,tasks,&wg)
	}

	// Send tasks
	for i:=1;i<=numTasks;i++{
		tasks <- task.Task{ID: i, Data: fmt.Sprintf("Sample data for task:",i)}
	}

	close(tasks)

	wg.Wait()
}