package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

func TestHandlerGet(t *testing.T) {
	server := NewServer()
	httpTester := httptest.NewServer(http.HandlerFunc(server.handleGetUser))
	nreq := 1000
	wg := &sync.WaitGroup{}
	for i := 0; i < nreq; i++ {
		wg.Add(1)
		go func(i int) {
			id := i%100 + 1
			url := fmt.Sprintf("%s/?id=%d", httpTester.URL, id)
			resp, err := http.Get(url)

			if err != nil {
				t.Error(err)
			}
			user := &user{}
			if err := json.NewDecoder(resp.Body).Decode(user); err != nil {
				t.Error(err)
			}
			fmt.Println(user)
			wg.Done()
		}(i)

		time.Sleep(time.Millisecond * 1)
	}
	wg.Wait()
	fmt.Printf("Database hits %d", server.dbhit)

}
