package main

import (
	"fmt"
	"sync"
	"time"
)

type sender struct {
	message chan string
}

type receiver struct {
	Id      int
	message chan string
}

var wg sync.WaitGroup

func (r *receiver) start() {

	go func() {
		for m := range r.message {
			fmt.Printf("sender id %d and received %s\n", r.Id, m)

		}
		wg.Done()

	}()
	time.Sleep(time.Millisecond * 1)
}
func (s *sender) start() {

	//go func() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		// waitgroup is for messgaes and not for go routines
		// you add waitgroup so you ask the routine to wait till you complete the task
		// mainly for synchronisation
		//go func(i int) {
		s.message <- fmt.Sprintf(" message for user %d", i+1)
		fmt.Printf("message sent\n")
		//}(i)
	}
	//}()
	time.Sleep(time.Millisecond * 1)
}

func main() {
	ch := make(chan string)
	s := &sender{
		message: ch,
	}
	r := &receiver{
		Id:      1,
		message: ch,
	}
	r.start()
	s.start()

	wg.Wait()
	close(ch)
	fmt.Println("done")

}
