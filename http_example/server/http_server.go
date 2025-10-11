package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	mux := http.NewServeMux()

	// Long-running handler (10 seconds)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request started")
		time.Sleep(10 * time.Second) // Simulate long work
		fmt.Fprintln(w, "Request finished")
		log.Println("Request finished")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Start server in a goroutine
	go func() {
		fmt.Println("Server started on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for Ctrl+C (interrupt signal)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop // Block until signal received

	fmt.Println("\nShutdown initiated...")

	// Context with 5-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	start := time.Now()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	fmt.Printf("Server exited properly after %.2f seconds\n", time.Since(start).Seconds())
}
