package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
)

type Test struct {
	ID string
	// Counter uint64
}

type Stats struct {
	sync.Mutex
	Counter map[string]uint64
}

var stats = Stats{Counter: make(map[string]uint64)}

func main() {
	http.HandleFunc("/", handler)
	addr := ":" + os.Getenv("PORT")
	if addr == ":" {
		addr = "127.0.0.1:3001"
	}
	fmt.Println("Listening on", addr)
	http.ListenAndServe(addr, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "%d\n", stats.Counter[r.URL.Query().Get("id")])
		return
	}
	var test Test
	if err := json.NewDecoder(r.Body).Decode(&test); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	stats.Lock()
	stats.Counter[test.ID]++
	stats.Unlock()
	w.WriteHeader(http.StatusAccepted)
}
