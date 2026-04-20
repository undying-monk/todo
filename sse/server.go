package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func sseHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Set required headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*") // For local testing

	// 2. Access the Flusher to push data immediately
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	fmt.Println("Client connected")

	// 3. Loop to send data at intervals
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-r.Context().Done():
			// Handle client disconnection
			fmt.Println("Client disconnected")
			return
		case t := <-ticker.C:
			// Prepare JSON data
			msg := map[string]string{
				"time": t.Format("15:04:05"),
				"info": "Server update from Go",
			}
			jsonData, _ := json.Marshal(msg)

			// 4. Format according to SSE spec: "data: <content>\n\n"
			fmt.Fprintf(w, "data: %s\n\n", jsonData)

			// 5. Flush the buffer to the client
			flusher.Flush()
		}
	}
}

func main() {
	http.HandleFunc("/events", sseHandler)

	fmt.Println("SSE Server running at http://localhost:3000/events")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}