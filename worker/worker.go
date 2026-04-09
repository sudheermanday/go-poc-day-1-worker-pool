package worker

import (
	"context"
	"fmt"
	"sync"

	"go-poc-day-1-worker-pool/task"
)

type Worker struct {
	ID int `json:"id"`
}

func (w Worker) Start(ctx context.Context, tasks <-chan task.Task, wg *sync.WaitGroup, retry int) {
	defer wg.Done()
	for {
		select {
		case <- ctx.Done():
			fmt.Println("Worker",w.ID,"stopping...")
			return
		case t,ok := <-tasks:
			if !ok {
				fmt.Println("Worker",w.ID,"no more tasks, stopping...")
				return
			}
			fmt.Println("Worker", w.ID, "processing task", t.ID)
			done,result:= t.Process()
			for i:=0;!done && i<retry;i++{
				fmt.Println("Worker", w.ID, "retrying task", t.ID, "Attempt:", i+1)
				done,result = t.Process()
			}
			if done {
				fmt.Println("Worker", w.ID, "completed task", t.ID, "Result:", result)
			}else{
				fmt.Println("Worker", w.ID, "failed to process task", t.ID, "Error:", result)
			}
			fmt.Println(result)
		}
	}
} 