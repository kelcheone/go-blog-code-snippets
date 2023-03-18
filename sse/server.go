// Server for the SSE example
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Event struct {
	id    int
	event string
	data  string
}

var (
	eventChan = make(chan Event)
)

func handleEvents(w http.ResponseWriter, r *http.Request) {
	// Set the headers related to event streaming.
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Listen to connection close and un-register messageChan
	notify := w.(http.CloseNotifier).CloseNotify()
	go func() {
		<-notify
		log.Println("Client closed connection")
	}()

	// Send 10 messages as server side events
	for i := 0; i < 10; i++ {
		fmt.Fprintf(w, "id: %d event: %s data: %s\n\n", i, "message", "Hello world")
		w.(http.Flusher).Flush()
		time.Sleep(1 * time.Second)
	}
}

func handleProgress(w http.ResponseWriter, r *http.Request) {
	// Set the headers related to event streaming.
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Listen to connection close and un-register messageChan
	notify := w.(http.CloseNotifier).CloseNotify()
	go func() {
		<-notify
		log.Println("Client closed connection")
	}()

	// Send 10 messages as server side events
	for i := 0; i < 10; i++ {
		fmt.Fprintf(w, "id: %d event: %s data: %s\n\n", i, "message", "Hello world")
		w.(http.Flusher).Flush()
		time.Sleep(1 * time.Second)
	}
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/events", handleEvents)
	router.HandleFunc("/progress", handleProgress)

	log.Println("Listening on :8080")

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		log.Fatal(http.ListenAndServe(":8080", router))
		wg.Done()
	}()
	wg.Wait()
}
