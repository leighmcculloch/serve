package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	// listen on designated port or random port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		ln, err = net.Listen("tcp", ":0")
	}
	if err != nil {
		log.Printf("%s", err)
	}
	log.Printf("Listening on %s", ln.Addr())

	// serve local files
	fs := http.FileServer(http.Dir("."))

	// serve on the listening port
	srv := &http.Server{Handler: loggingMiddleware(fs)}
	err = srv.Serve(ln)
	if err != nil {
		log.Printf("%s", err)
	}
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
