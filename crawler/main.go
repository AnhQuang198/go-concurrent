package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	n := 10000
	maxWorkers := 5

	wg := new(sync.WaitGroup)

	queueCh := make(chan int, n)

	for i := 1; i <= n; i++ {
		queueCh <- i
	}
	close(queueCh)

	for i := 1; i <= maxWorkers; i++ {
		wg.Add(1)
		go func(count int) {
			for v := range queueCh {
				time.Sleep(time.Millisecond * 10)
				fmt.Printf("Worker %d is crawling web url %d \n", count, v)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
}
