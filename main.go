package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

var version = "<dev>"
var commit = ""
var date = ""

func main() {
	exitCode := run(os.Args, os.Stdin, os.Stdout, os.Stderr)
	os.Exit(exitCode)
}

func run(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) int {
	flag := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flag.SetOutput(stderr)
	flagVersion := flag.Bool("version", false, "print the version")
	flagHelp := flag.Bool("help", false, "print this help")
	err := flag.Parse(args[1:])
	if err != nil {
		fmt.Fprintf(stderr, "%v\n", err)
		return 2
	}

	if *flagVersion {
		fmt.Fprintf(stderr, "serve %s %s %s\n", version, commit, date)
		return 0
	}

	if *flagHelp {
		flag.Usage()
		return 0
	}

	log := log.New(stdout, "", log.LstdFlags)

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
		return 1
	}
	log.Printf("Listening on %s", ln.Addr())

	// serve local files
	fs := http.FileServer(http.Dir("."))

	// serve on the listening port
	srv := &http.Server{Handler: loggingMiddleware(fs)}
	err = srv.Serve(ln)
	if err != nil {
		log.Printf("%s", err)
		return 1
	}

	return 0
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
