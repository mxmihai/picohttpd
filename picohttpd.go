package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Define flags
	port := flag.Int("port", 80, "Port to listen on")
	answer := flag.String("answer", "OK", "Response string")
	flag.Parse()

	// Handler that returns the answer
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(*answer)))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(*answer))
	})

	addr := fmt.Sprintf(":%d", *port)
	log.Printf("PicoHTTPD listening on %s, answering: %q\n", addr, *answer)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
