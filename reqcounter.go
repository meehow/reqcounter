package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	http.ListenAndServe("127.0.0.1:3001", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		counter, ok := stats.Counter[r.URL.Query().Get("id")]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "ID not found")
			return
		}
		fmt.Fprintf(w, "%d\n", counter)
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