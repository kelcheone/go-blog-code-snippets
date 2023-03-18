package main

/*
type Event struct {
	id    int
	event string
	data  string
}

func main2() {

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

func handleEvents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleEvents")
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	// Set the headers related to event streaming.
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Send a plain text mesage over the connection.
	// This is just to show that the connection is still alive
	// and the server is still sending data.
	fmt.Fprintf(w, "data: %s \n", "Hello World")
	flusher.Flush()

	// Create a ticker that will tick every second.
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	// Send the event data as server-sent events over the
	// HTTP connection.
	for range ticker.C {
		fmt.Fprintf(w, "data: %s \n", "Hello World")
		flusher.Flush()
	}

}

func handleProgress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	for i := 0; i < 100; i++ {
		// Create a new event.
		event := Event{
			id:    i,
			event: "onProgress",
			data:  fmt.Sprintf(`{"progress": %d%%}`, i),
		}

		// Send the event to the client.
		sendEvent(w, event)
		fmt.Println("sendEvent", i)

		// Sleep for 100 milliseconds.
		if i%10 == 0 {
			time.Sleep(1000 * time.Millisecond)
		}
	}

	// Send the final event.
	event := Event{
		id:    100,
		event: "onComplete",
		data:  `{"progress": 100%}`,
	}
	sendEvent(w, event)

	// Close the connection.
	w.(http.Flusher).Flush()

}

func sendEvent(w http.ResponseWriter, event Event) {
	// Write the event ID.
	if event.id != 0 {
		fmt.Fprintf(w, "id: %d \n", event.id)
	}

	// Write the event name.
	if event.event != "" {
		fmt.Fprintf(w, "event: %s \n", event.event)
	}

	// Write the event data.
	if event.data != "" {
		fmt.Fprintf(w, "data: %s \n", event.data)
	}

	// Write a blank line to indicate the end of the event.
	fmt.Fprintf(w, "\n")

	// Flush the data immediately instead of buffering it for later.
	flusher, ok := w.(http.Flusher)
	if ok {
		flusher.Flush()
	}
}
*/
