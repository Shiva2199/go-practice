package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
		fmt.Printf("Hello World")
	})
	timeoutContext, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		wg.Done()

		if err := http.ListenAndServe("localhost:8080", nil); err != nil {
			panic(err)
		}
		time.Sleep(time.Second * 4)
	}()

	wg.Wait()
	req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, "http://localhost:8080/hello", nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Body)

}
