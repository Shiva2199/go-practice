Sure! Here's another problem to test your Go development skills:

### Problem Statement

**Title:** Implement a RESTful API for a To-Do List Application

**Description:**
You are tasked with building a RESTful API for a simple To-Do list application. The API should allow users to create, read, update, and delete to-do items. Each to-do item should have an ID, title, description, and status (e.g., pending, completed).

**Requirements:**
1. **HTTP Methods:** Implement endpoints for the following operations:
   - `POST /todos`: Create a new to-do item.
   - `GET /todos`: Retrieve all to-do items.
   - `GET /todos/{id}`: Retrieve a specific to-do item by ID.
   - `PUT /todos/{id}`: Update a specific to-do item by ID.
   - `DELETE /todos/{id}`: Delete a specific to-do item by ID.
2. **Data Storage:** Use an in-memory data structure to store the to-do items.
3. **Error Handling:** Properly handle errors such as invalid input or non-existent IDs.
4. **Concurrency:** Ensure thread-safe operations on the to-do items.

### Example

Given the following operations:
1. `POST /todos` with body `{"title": "Buy groceries", "description": "Milk, Bread, Eggs", "status": "pending"}`
2. `GET /todos`
3. `PUT /todos/1` with body `{"title": "Buy groceries", "description": "Milk, Bread, Eggs", "status": "completed"}`
4. `DELETE /todos/1`

### Skeleton Code

Here's a skeleton to get you started:

```go
package main

import (
    "encoding/json"
    "net/http"
    "sync"
)

type Todo struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Status      string `json:"status"`
}

type Server struct {
    todos map[int]Todo
    mu    sync.Mutex
    nextID int
}

func NewServer() *Server {
    return &Server{todos: make(map[int]Todo), nextID: 1}
}

func (s *Server) createTodoHandler(w http.ResponseWriter, r *http.Request) {
    var todo Todo
    if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    s.mu.Lock()
    todo.ID = s.nextID
    s.nextID++
    s.todos[todo.ID] = todo
    s.mu.Unlock()
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(todo)
}

func (s *Server) getTodosHandler(w http.ResponseWriter, r *http.Request) {
    s.mu.Lock()
    todos := make([]Todo, 0, len(s.todos))
    for _, todo := range s.todos {
        todos = append(todos, todo)
    }
    s.mu.Unlock()
    json.NewEncoder(w).Encode(todos)
}

func (s *Server) getTodoHandler(w http.ResponseWriter, r *http.Request) {
    // Implement logic to retrieve a specific to-do item by ID
}

func (s *Server) updateTodoHandler(w http.ResponseWriter, r *http.Request) {
    // Implement logic to update a specific to-do item by ID
}

func (s *Server) deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
    // Implement logic to delete a specific to-do item by ID
}

func main() {
    server := NewServer()
    http.HandleFunc("/todos", server.createTodoHandler)
    http.HandleFunc("/todos", server.getTodosHandler)
    http.HandleFunc("/todos/{id}", server.getTodoHandler)
    http.HandleFunc("/todos/{id}", server.updateTodoHandler)
    http.HandleFunc("/todos/{id}", server.deleteTodoHandler)
    http.ListenAndServe(":8080", nil)
}
```

### Instructions:
1. Complete the handlers for retrieving, updating, and deleting to-do items.
2. Ensure proper error handling and logging.
3. Test your solution with different requests.

Feel free to ask if you need any help or have questions about this problem!