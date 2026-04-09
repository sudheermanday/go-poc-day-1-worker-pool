package main

import (
	"context"
	"flag"
	"fmt"
	"go-poc-day-1-worker-pool/task"
	"go-poc-day-1-worker-pool/worker"
	"go-poc-day-1-worker-pool/config"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	configPath := flag.String("config","config.json","path to config file")
	flag.Parse()
	fmt.Println("Reading config from:", *configPath)
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal("Error loading config:", err)
	}
	rand.Seed(time.Now().UnixNano())
	numWorkers := cfg.NumWorkers
	numTasks := cfg.NumTasks
	retryLimit := cfg.RetryCount
	
	tasks := make(chan task.Task, numTasks)
	
	ctx,cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var wg sync.WaitGroup

	// Start workers
	for i:=1; i<=numWorkers; i++{
		wg.Add(1)
		w := worker.Worker{ID: i}
		go w.Start(ctx,tasks,&wg, retryLimit)
	}

	// Send tasks
	for i:=1;i<=numTasks;i++{
		tasks <- task.Task{ID: i, Data: fmt.Sprintf("Sample data for task:",i)}
	}

	close(tasks)

	wg.Wait()
}