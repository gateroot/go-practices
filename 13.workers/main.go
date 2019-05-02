package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(taskCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		task, ok := <-taskCh
		if !ok {
			return
		}
		d := time.Duration(task) * time.Millisecond
		time.Sleep(d)
		fmt.Println("processing task", task)
	}
}

func pool(wg *sync.WaitGroup, workers, tasks int) {
	taskCh := make(chan int)

	for i := 0; i < workers; i++ {
		go worker(taskCh, wg)
	}

	for i := 0; i < tasks; i++ {
		taskCh <- i
	}

	close(taskCh)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(36)
	go pool(&wg, 36, 50)
	wg.Wait()
}
