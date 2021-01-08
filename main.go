package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fs := http.FileServer(http.Dir("."))
	address := ":" + port
	log.Printf("Listening on %s", address)
	err := http.ListenAndServe(address, loggingMiddleware(fs))
	log.Printf("%s", err.Error())
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := responseWriter{ResponseWriter: w}
		t := time.Now()
		next.ServeHTTP(&rw, r)
		log.Printf("%s %s %d %fms", r.Method, r.URL.Path, rw.StatusCode, time.Since(t).Seconds()*1000)
	})
}

type responseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (w *responseWriter) WriteHeader(code int) {
	w.StatusCode = code
	w.ResponseWriter.WriteHeader(code)
}
