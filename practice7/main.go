package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"
)

func DoWork() int {
	time.Sleep(time.Second * 1)
	n, err := rand.Int(rand.Reader, big.NewInt(25))
	if err != nil {
		return 0
	}
	return int(n.Int64())

}

func main() {
	ch := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				r := DoWork()
				ch <- r
			}()
		}
		wg.Wait()
		close(ch)
	}()

	for c := range ch {
		fmt.Printf("value: %v\n", c)
	}
}
