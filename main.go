package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)

var count int
var countFile int
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

	fmt.Fprintf(w, "Count: %d\n", currentCount)
}

func storeCountHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	mutex.Lock()
	countFile++
	fileName := fmt.Sprintf("count_%d.txt", countFile)
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	l, err := f.WriteString(fmt.Sprintf("Count: %d\n", count))
	if err != nil {
		fmt.Println(err)
		f.Close()
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
	}
	mutex.Unlock()

	fmt.Fprintf(w, "Store count file to %s\n", fileName)
}

func main() {
	http.HandleFunc("/count", countHandler)
	http.HandleFunc("/storecount", storeCountHandler)

	// Start the server on port 8080
	println("Starting server at port 1535")
	if err := http.ListenAndServe(":1535", nil); err != nil {
		panic(err)
	}
}
