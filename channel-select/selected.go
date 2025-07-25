package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	for {
		time.Sleep(6 * time.Second)
		ch <- "This is from server1!"
	}
}

func server2(ch chan string) {
	for {
		time.Sleep(3 * time.Second)
		ch <- "This is from server2!"
	}
}

func main() {
	fmt.Println("Select with channels")
	fmt.Println("--------------------")

	channel1 := make(chan string)
	channel2 := make(chan string)

	go server1(channel1)
	go server2(channel2)

	for {
		select {
		case msg1 := <-channel1:
			fmt.Println("Case 1: ", msg1)
		case msg1 := <-channel1:
			fmt.Println("Case 2: ", msg1)
		case msg2 := <-channel2:
			fmt.Println("Case 3: ", msg2)
		case msg2 := <-channel2:
			fmt.Println("Case 4: ", msg2)
			//default:
			//	fmt.Println("Nothing to see here")
		}
	}
}
