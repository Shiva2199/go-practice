package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type user struct {
	Name string
	Id   int
}
type server struct {
	db    map[int]*user
	cache map[int]*user
	dbhit int
}

func NewServer() *server {
	db := make(map[int]*user)
	for i := 0; i < 100; i++ {
		db[i+1] = &user{Name: fmt.Sprintf("user_%d", i+1),
			Id: i + 1,
		}
	}
	return &server{
		db:    db,
		cache: make(map[int]*user),
	}
}
func (s *server) handleGetUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	user, ok := s.cache[id]
	if ok {
		json.NewEncoder(w).Encode(user)
		return
	}
	if err != nil {
		json.NewEncoder(w).Encode("Error converting user id")
	}
	user, ok = s.db[id]
	if !ok {
		json.NewEncoder(w).Encode("user not found")
	}
	s.dbhit++
	s.cache[id] = user
	json.NewEncoder(w).Encode(user)

}

func main() {

}
