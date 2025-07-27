package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Hello World %v \n", "System")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	go doSomething(ctx)

	time.Sleep(time.Second * 6)
	fmt.Println("Finished")
}

func doSomething(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Context done ", ctx.Err())
	case <-time.After(time.Second * 5):
		fmt.Println("Done process")
	}
}
