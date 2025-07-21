package main

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {
	for {
		i := <-ch
		fmt.Println("Got ", i, "from channel")

		//simulate doing a lot of work
		time.Sleep(1 * time.Second)
	}
}

func main() {
	//Buffered channed cho phép gửi hoặc nhận giá trị vào channel theo size đã init.
	//Có thể gửi số lượng value max đến size đã cấu hình, sau đó mới cần đợi giá trị được đọc ra rồi mới gửi vào tiếp
	//Mặc định với Unbuffered channel khi send giá trị vào channel thì phải đợi giá trị được đọc ra rồi mới send tiếp được
	ch := make(chan int, 10)

	go listenToChan(ch)

	for i := 0; i < 100; i++ {
		fmt.Println("Sending ", i, "to channel...")
		ch <- i
		fmt.Println("Sent ", i, "to channel")
	}

	fmt.Println("Done!")
	close(ch)
}
