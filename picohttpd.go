package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func main() {
	// Define flags
	port := flag.Int("port", 80, "Port to listen on")
	answer := flag.String("answer", "OK", "Response text or 'cmd:<command>' to execute")
	path := flag.String("path", "/", "Path to respond to (default: /)")
	flag.Parse()

	// Normalize path
	if *path == "" || !strings.HasPrefix(*path, "/") {
		*path = "/"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != *path {
			http.NotFound(w, r)
			return
		}

		var output string

		if strings.HasPrefix(*answer, "cmd:") {
			cmdStr := strings.TrimPrefix(*answer, "cmd:")
			cmdStr = strings.TrimSpace(cmdStr)
			cmd := exec.Command("bash", "-c", cmdStr)
			out, err := cmd.CombinedOutput()
			if err != nil {
				output = fmt.Sprintf("Error: %v\n%s", err, out)
			} else {
				output = string(out)
			}
		} else {
			output = *answer
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(output)))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(output))
	})

	addr := fmt.Sprintf(":%d", *port)
	log.Printf("PicoHTTPD listening on %s, path %q, answering: %q\n", addr, *path, *answer)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
