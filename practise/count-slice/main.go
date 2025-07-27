package main

import "fmt"

func main() {
	//TODO: chia 2 nửa của 1 mảng, dùng goroutince để tính tổng từng phần và gộp kết quả
	fmt.Println("Start to count slice number")
	lstNums := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	position := len(lstNums) / 2

	lstNumPrev := lstNums[:position]
	lstNumNext := lstNums[position:]

	fmt.Println("lstNum previous is", lstNumPrev)
	fmt.Println("lstNum next is", lstNumNext)

	sum := make(chan int32)
	go sumPart(lstNumPrev, sum)
	go sumPart(lstNumNext, sum)

	totalSumPrev := <-sum
	totalSumNext := <-sum

	totalSum := totalSumPrev + totalSumNext
	fmt.Println("total sum is", totalSum)
}

func sumPart(lstNums []int32, sum chan int32) {
	var total int32
	for _, v := range lstNums {
		total += v
	}
	sum <- total
}
