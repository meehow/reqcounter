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
		id := limitLength(r.URL.Query().Get("id"))
		fmt.Fprintln(w, stats.Counter[id])
		return
	}
	var test Test
	if err := json.NewDecoder(r.Body).Decode(&test); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := limitLength(test.ID)
	stats.Lock()
	stats.Counter[id]++
	stats.Unlock()
	w.WriteHeader(http.StatusAccepted)
}

func limitLength(s string) string {
	if len(s) > 40 {
		return s[:40]
	}
	return s
}
