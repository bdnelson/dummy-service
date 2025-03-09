package main

import (
	"context"
	"dummyservice"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const DefaultPort = 9292
const DefaultAddress = ""

func configureFromEnv() string {
	port := DefaultPort
	portVar, ok := os.LookupEnv("DUMMY_PORT")
	if ok {
		var err error
		port, err = strconv.Atoi(portVar)
		if err != nil {
			port = DefaultPort
		}
	}

	listenOn, ok := os.LookupEnv("DUMMY_ADDRESS")
	if !ok {
		listenOn = DefaultAddress
	}

	return fmt.Sprintf("%s:%d", listenOn, port)
}

type logWriter struct {
	http.ResponseWriter
	code int
}

func (lw *logWriter) WriteHeader(code int) {
	lw.code = code
	lw.ResponseWriter.WriteHeader(code)
}

// Unwrap supports http.ResponseController.
func (lw *logWriter) Unwrap() http.ResponseWriter { return lw.ResponseWriter }

func wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lw := &logWriter{w, http.StatusOK}
		next.ServeHTTP(lw, r)
		log.Printf("%d: %s\n", lw.code, r.URL.Path)
	})
}

func main() {
	addr := configureFromEnv()
	mux := dummyservice.CreateRouter()
	server := &http.Server{
		Addr:    addr,
		Handler: wrap(mux),
	}

	log.Printf("Hello.  Starting dummyservice.service listening on \"%s\"...\n", addr)

	// Channel to listen for signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	log.Printf("Server started on %s", addr)

	// Block until signal is received
	sig := <-sigChan
	log.Printf("Received signal: %v", sig)

	// Shutdown server gracefully
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server stopped")
}
