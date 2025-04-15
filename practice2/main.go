package main

import (
	"fmt"
	"time"
)

func main() {
	// unbuffered channel
	ch := make(chan int)
	go func(ch chan int) {
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Println("unbuffered channel ", i)
		}
		close(ch)
	}(ch)
	// for i := range ch {
	// 	fmt.Println(i)
	// }
	time.Sleep(5 * time.Second)
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("Unbuffered channel done")
	fmt.Println("Buffered channel  ")
	// buffered channel
	buffCh := make(chan int, 5)
	go func(buffCh chan int) {
		for i := 0; i < 10; i++ {
			buffCh <- i
			fmt.Println("Buffered channel ", i)
		}
		close(buffCh)
	}(buffCh)
	time.Sleep(5 * time.Second)
	fmt.Println("Buffered channel started ")

	for buf := range buffCh {
		fmt.Println(buf)
	}
	fmt.Println("Buffered channel done ")

}
