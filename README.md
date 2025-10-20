# picohttpd

**Tiny HTTP server written in Go**

`picohttpd` is a minimal HTTP server that responds with a simple text string to any HTTP request. It is designed to be **lightweight, fast, and self-contained**.

---

## Features

- Responds with a configurable string for any request.
- Configurable port and response string via command-line flags:
  ```bash
  ./picohttpd -port 8080 -answer "PONG"
  ```
- Default behavior: listens on **port 80** and responds `"OK"`.
- Extremely low CPU and memory footprint.
- Single executable, no extra files or dependencies required.

---

## Usage

```bash
# Default
./picohttpd

# Custom port and answer
./picohttpd -port 8080 -answer "PONG"
```

---

## Build

- The source code is in `picohttpd.go`.
- You can build it using Go (version >= 1.21):

```bash
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o picohttpd picohttpd.go
```

- For convenience, **pre-built Linux binaries** are available in the `build` folder (compiled for Linux x86_64, compatible with Rocky Linux, RHEL, CentOS, Debian, Ubuntu).

---

## Systemd Service Example (Optional)

You can run `picohttpd` as a service:

```ini
[Unit]
Description=Tiny picohttpd HTTP server
After=network.target

[Service]
Type=simple
ExecStart=/usr/local/bin/picohttpd -port 80 -answer "OK"
Restart=always
RestartSec=2
User=nobody
Group=nobody

[Install]
WantedBy=multi-user.target
```

Enable and start:

```bash
sudo systemctl daemon-reload
sudo systemctl enable --now picohttpd.service
```

---

## Notes

- This server is intended for **health checks, testing, or simple HTTP responses**.
- No HTML or additional files are required; everything is served by the binary.
- For permanent availability of the binary, consider attaching it to a **GitHub Release** or hosting it elsewhere.
