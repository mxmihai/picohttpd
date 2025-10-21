# picohttpd

**Tiny HTTP server written in Go**

`picohttpd` is a minimal, self-contained HTTP responder written in Go.  
It’s designed for **diagnostics, automation scripts, and health checks** — lightweight, fast, and with no dependencies.

---

## Features

- Single executable — no config files or dependencies.
- Responds with either:
  - A static text message (default: `"OK"`), or
  - The output of a command via `cmd:<command>`
- Configurable:
  - **Port** with `-port` (default: 80)
  - **Response text** with `-answer` (default: `"OK"`)
  - **Path** with `-path` (responds only to the specified path; default: /)
- Minimal resource usage — suitable for embedded and cloud environments.
- Provides `-v` (version info) and `-h` (help).

---

## Usage

```bash
# Default (port 80, path "/", answer "OK")
./picohttpd

# Custom port and response
./picohttpd -port 8080 -answer "OK"

# Respond only at /ping
./picohttpd -path "/ping" -answer "pong"

# Execute a command and return its output
./picohttpd -answer "cmd:uptime"

# Execute any bash command and return its output
./picohttpd -answer "cmd:echo Hello from $(hostname)"

# Display version
./picohttpd -v

# Display help
./picohttpd -h
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
ExecStart=/usr/local/bin/picohttpd -port 80 -path / -answer "OK"
Restart=always
RestartSec=2
#User=nobody
#Group=nobody

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
- Runs entirely in memory — no filesystem or HTML needed, everything is served by the binary.
- Extremely low CPU and RAM footprint.
- For permanent availability of the binary, consider attaching it to a **GitHub Release** or hosting it elsewhere.
