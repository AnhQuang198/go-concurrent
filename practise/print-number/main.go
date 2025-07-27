package main

import (
	"fmt"
	"sync"
)

func main() {
	//TODO: Dùng 2 goroutince: 1 in số chẵn, 1 in số lẻ từ 1 đến 10
	fmt.Println("Start to print number")
	lstNums := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	wg := new(sync.WaitGroup)

	var lstNumEven []int32
	var lstNumOdd []int32

	for _, num := range lstNums {
		if num%2 == 0 {
			lstNumEven = append(lstNumEven, num)
		} else {
			lstNumOdd = append(lstNumOdd, num)
		}
	}

	wg.Add(2)
	go printNumber(lstNumEven, true, wg)
	go printNumber(lstNumOdd, false, wg)
	wg.Wait()
}

func printNumber(lstNums []int32, isEvent bool, wg *sync.WaitGroup) {
	for _, num := range lstNums {
		if isEvent {
			fmt.Printf("%d is even\n", num)
		} else {
			fmt.Printf("%d is odd\n", num)
		}
	}
	wg.Done()
}
