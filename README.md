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

- Check the **pre-built Linux binaries** available in the [Latest Release](https://github.com/mxmihai/picohttpd/releases) (compiled for Linux x86_64, compatible with Rocky Linux, RHEL, CentOS, Debian, Ubuntu). 

---

## Systemd Service Example (Optional)

You can run `picohttpd` as a service:

```ini
[Unit]
Description=Tiny picohttpd HTTP server
After=network.target

[Service]
Type=simple
ExecStart=/usr/local/bin/picohttpd -port 8080 -path / -answer "OK"
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
- Runs entirely in memory — no filesystem or HTML needed, everything is served by the binary.
- Extremely low CPU and RAM footprint.
- For permanent availability of the binary, consider attaching it to a **GitHub Release** or hosting it elsewhere.

---

# Security

`picohttpd` is designed to be **lightweight and minimal**, but because it can optionally execute shell commands via the `cmd:` feature, there are important security considerations:

## ⚠️ Key Points

- **Non-root operation recommended:**  
  Bind to low ports (like 80) using `setcap` instead of running as root, and use a dedicated system user (`pico`) for the service.

- **Limit network exposure:**  
  Bind to `127.0.0.1` if only local access is needed, or restrict ports via firewall rules to trusted hosts.

- **Command execution risks (`cmd:`):**  
  - Avoid running arbitrary shell commands unless necessary.  
  - Prefer whitelisting allowed commands to prevent accidental or malicious execution.  
  - Commands run on each request; avoid exposing sensitive files or operations.

- **Systemd hardening recommended:**  
  Use options such as `NoNewPrivileges`, `ProtectSystem=full`, `ProtectHome=yes`, and `CapabilityBoundingSet=CAP_NET_BIND_SERVICE` to minimize potential damage in case of compromise.

- **Timeouts and output limits:**  
  When running commands, consider using timeouts and truncating output to prevent resource exhaustion.

- **Monitoring:**  
  Enable logging and monitor requests to detect unusual activity.

- **Containerization (optional):**  
  Running `picohttpd` inside a container or sandbox can further isolate it from the host system.

## ✅ Safe Usage Recommendations

1. Use static responses (`-answer "OK"`) wherever possible.  
2. Run as a non-privileged user.  
3. Limit exposure to trusted networks.  
4. If using commands, whitelist them and monitor usage.  
5. Regularly update the host system and Go runtime.

---

# Disclaimer

**Important Notice:**  

By using `picohttpd`, you acknowledge and agree to the following terms:

1. **No Liability:**  
   The author **does not accept any responsibility or liability** for any direct, indirect, incidental, special, or consequential damages resulting from the use, misuse, or inability to use this software.

2. **Use at Your Own Risk:**  
   `picohttpd` is provided **as-is**. You are fully responsible for how you deploy, configure, and operate it.

3. **No Warranty:**  
   There is **no warranty** of any kind, either express or implied, including but not limited to fitness for a particular purpose, security, or reliability.

4. **Security Responsibility:**  
   It is your responsibility to implement proper security measures when using `picohttpd`. The author is **not responsible** for any security breaches, data loss, or unauthorized access.

5. **Legal Compliance:**  
   You must comply with all applicable laws and regulations when using this software. The author assumes no responsibility for your compliance or violations.

---

By using this software, you accept that **all risks are yours** and you release the author from any liability or claims.
