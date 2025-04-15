package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	data := &Todo{
		ID:          1,
		Title:       "test",
		Description: "test",
		Status:      "test",
	}
	input, _ := json.Marshal(data)
	getReq := httptest.NewRequest("GET", "/todos", nil)
	router := NewTodoRouter()
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, getReq)
	fmt.Println(rr.Code)
	if rr.Code != 200 {
		t.Errorf("Expected 200, got %d", rr.Code)
	}
	fmt.Println(rr.Body)
	postReq := httptest.NewRequest("POST", "/todos", bytes.NewBuffer(input))
	router.ServeHTTP(rr, postReq)
	fmt.Printf("Post %d", rr.Code)
	if rr.Code != 200 {
		t.Errorf("Expected 200, got %d", rr.Code)
	}
	fmt.Println(rr.Body)
	router.ServeHTTP(rr, getReq)
	fmt.Println(rr.Code)
	if rr.Code != 200 {
		t.Errorf("Expected 200, got %d", rr.Code)
	}
	fmt.Println(rr.Body)
	Updateddata := &Todo{
		ID:          1,
		Title:       "test2",
		Description: "test2",
		Status:      "test2",
	}
	update, _ := json.Marshal(Updateddata)
	putReq := httptest.NewRequest("PUT", "/todos/1", bytes.NewBuffer(update))
	router.ServeHTTP(rr, putReq)
	fmt.Printf("Put %d", rr.Code)
	if rr.Code != 200 {
		t.Errorf("Expected 200, got %d", rr.Code)
	}
	fmt.Println(rr.Body)
	router.ServeHTTP(rr, getReq)
	fmt.Println(rr.Code)
	if rr.Code != 200 {
		t.Errorf("Expected 200, got %d", rr.Code)
	}
	fmt.Println(rr.Body)

}
