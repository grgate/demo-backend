package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	commitSha string
	version   string
	feature   string
)

func main() {
	feature = os.Getenv("ENABLE_FEATURE")
	if feature == "" {
		feature = "A"
	}

	go serveProbe()
	serveHTTP()
}

func serveHTTP() {
    fmt.Println("starting")
	fmt.Printf("serving http :8081 | version: %s | commit sha: %s | feature: "+
		"%s\n", version, commitSha, feature)

	m := http.NewServeMux()

	m.HandleFunc("/", rootHandler)

	s := http.Server{
		Addr:    "0.0.0.0:8081",
		Handler: m,
	}

	log.Fatal(s.ListenAndServe())
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", feature)
	log.Println("GET /")
}

func serveProbe() {
	fmt.Println("serving probe :8091")

	m := http.NewServeMux()

	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	s := http.Server{
		Addr:    "0.0.0.0:8091",
		Handler: m,
	}

	log.Fatal(s.ListenAndServe())
}
