package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

var mu sync.RWMutex

type TodoList []Todo

var todoList TodoList

func handleGet(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()
	if len(todoList) == 0 {
		json.NewEncoder(w).Encode("No Items")
		return
	}
	json.NewEncoder(w).Encode(todoList)
}
func handlePost(w http.ResponseWriter, r *http.Request) {
	var item Todo
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		fmt.Printf("failed to covert request body with Error: %s\n", err)
		json.NewEncoder(w).Encode("Failed to  Save Item")
		return

	}
	mu.Lock()
	defer mu.Unlock()
	todoList = append(todoList, item)
	json.NewEncoder(w).Encode("Item Saved")
}
func handlePut(w http.ResponseWriter, r *http.Request) {
	///todos/{id}
	vars := mux.Vars(r)
	strId := vars["id"]

	// r.URL.
	intId, err := strconv.Atoi(strId)
	if err != nil {
		fmt.Printf("failed to convert ID to int with Error: %s\n", err)
		json.NewEncoder(w).Encode("Failed to Save Item")
		return
	}
	var item Todo
	json.NewDecoder(r.Body).Decode(&item)
	var updated bool
	mu.Lock()
	defer mu.Unlock()
	for i := range todoList {
		if (todoList)[i].ID != intId {
			continue
		}
		(todoList)[i].Description = item.Description
		(todoList)[i].Status = item.Status
		(todoList)[i].Title = item.Title
		updated = true
		break
	}
	if updated {
		fmt.Printf("Item Saved %v", item)

		json.NewEncoder(w).Encode("Item Saved")
	} else {
		fmt.Printf("Item failed to Update Not found")
		json.NewEncoder(w).Encode("Item failed to Update Not found")

	}

}
func NewTodoRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/todos", handleGet).Methods("GET")
	router.HandleFunc("/todos", handlePost).Methods("POST")
	router.HandleFunc("/todos/{id}", handlePut).Methods("PUT")

	return router
}
func main() {

	router := NewTodoRouter()
	server := &http.Server{
		Handler: router,
		Addr:    ":8443",
	}
	server.ListenAndServe()

}
