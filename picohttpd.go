package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

const version = "1.1.0"

func printHelp() {
	fmt.Printf(`picohttpd %s - A tiny HTTP responder for diagnostics and scripting

Usage:
  picohttpd [options]

Options:
  -port <num>       Port to listen on (default: 80)
  -answer <text>    Response text or 'cmd:<command>' to execute (default: "OK")
  -path <string>    Path to respond to (default: "/")
  -v                Show version and exit
  -h, --help        Show this help message and exit

Examples:
  picohttpd
      → Responds with "OK" on port 80, path "/"

  picohttpd -port 8080 -answer "Hello"
      → Responds with "Hello" on http://localhost:8080/

  picohttpd -answer "cmd:uptime"
      → Responds with the result of 'uptime'

  picohttpd -path "/ping" -answer "cmd:uptime"
      → Responds with 'uptime' only at /ping
`, version)
}

func main() {
	// Handle manual -v and -h early (before flag.Parse)
	for _, arg := range os.Args[1:] {
		if arg == "-v" {
			fmt.Printf("picohttpd version %s\n", version)
			os.Exit(0)
		}
		if arg == "-h" || arg == "--help" {
			printHelp()
			os.Exit(0)
		}
	}
	
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
