package main

import (
	"fmt"
	"net/http"
	"time"
)

type HelloHandler struct{}

func (hh HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// name := r.URL.Query().Get("name")
	// if name == "" {
	// 	name = "stranger"
	// }

	// w.Header().Set("Content-Type", "text/plain")
	// w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, "Hello, %s!\n", name)
	fmt.Fprintf(w, "Hello!\n")
}

type ByeHandler struct{}

func (bh ByeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bye!\n")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/hello", HelloHandler{})
	mux.Handle("/bye", ByeHandler{})
	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}
	fmt.Println("listening on :8080")
	err := s.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
