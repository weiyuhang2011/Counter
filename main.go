package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

// AppState holds the shared state of the application.
type AppState struct {
	count int
	mu    sync.Mutex // Protects the count variable
}

// Increment the count safely.
func (app *AppState) incrementCount() int {
	app.mu.Lock()
	defer app.mu.Unlock()
	app.count++
	return app.count
}

// Safe retrieval of the count value.
func (app *AppState) getCount() int {
	app.mu.Lock()
	defer app.mu.Unlock()
	return app.count
}

// HTTP handler to increment the count.
func countHandler(app *AppState) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		updatedCount := app.incrementCount()
		fmt.Fprintf(w, "Count updated: %d\n", updatedCount)
	}
}

func main() {
	appState := &AppState{}

	// Start a background goroutine to write the count to a file every second.
	go func() {
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		// Open file in append mode.
		f, err := os.OpenFile("count-writing.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
		if err != nil {
			log.Fatalf("Error opening file: %v", err)
		}
		defer f.Close()

		for range ticker.C {
			currentCount := appState.getCount()

			// Write the current count to the file with a newline for each entry.
			_, err := f.WriteString(strconv.Itoa(currentCount) + "\n")
			log.Printf("Count written to file: %d\n", currentCount)
			if err != nil {
				log.Printf("Error writing to file: %v", err)
			}
			f.Sync() // Ensure that the write is flushed to the disk.
		}
	}()

	// Set up HTTP server
	http.HandleFunc("/count", countHandler(appState))
	log.Println("Server started on :1535")
	log.Fatal(http.ListenAndServe(":1535", nil))
}
