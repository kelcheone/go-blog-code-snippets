package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync/atomic"
	"syscall"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Printf("Received request: %s %s, from %s", r.Method, r.URL.Path, r.RemoteAddr)
	fmt.Fprintf(w, "Hello, World!")
}

var healthy int32

func heathCheck(w http.ResponseWriter, r *http.Request) {
	if atomic.LoadInt32(&healthy) == 1 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.WriteHeader(http.StatusServiceUnavailable)
}

type key int

const (
	requestIDKey key = 0
)

func tracing(nextReuestID func() string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				requestID = nextReuestID()
			}
			ctx := context.WithValue(r.Context(), requestIDKey, requestID)
			w.Header().Set("X-Request-Id", requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func logging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				requestID, ok := r.Context().Value(requestIDKey).(string)
				if !ok {
					requestID = "unknown"
				}

				logger.Println(requestID, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func main() {
	http.HandleFunc("GET /", handler)
	http.HandleFunc("GET /healthz", heathCheck)
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	nextRequestID := func() string {
		return strconv.FormatInt(time.Now().UnixNano(), 10)
	}
	server := &http.Server{
		Addr:         ":8080",
		Handler:      tracing(nextRequestID)(logging(logger)(http.DefaultServeMux)),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// set server to healthy
	atomic.StoreInt32(&healthy, 1)
	go func() {
		<-quit
		log.Println("Server is shutting down ....")
		atomic.StoreInt32(&healthy, 0)
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

		defer cancel()
		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			logger.Fatalf("Could not gracefully shutdown the server %+v\n", err)
		}
		close(done)
	}()

	logger.Println("Server starting at port 8080 ...")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on :8080 %+v\n", err)
	}

	<-done
	logger.Println("Server stopped")
}
