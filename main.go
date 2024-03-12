package main

import (
	"fmt"
	"net/http"
	"sync"
)

var count int
var mutex sync.Mutex

func countHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	mutex.Lock()
	count++
	currentCount := count
	mutex.Unlock()

	fmt.Fprintf(w, "Count: %d", currentCount)
}

func main() {
	http.HandleFunc("/count", countHandler)

	// Start the server on port 8080
	println("Starting server at port 1535")
	if err := http.ListenAndServe(":1535", nil); err != nil {
		panic(err)
	}
}
