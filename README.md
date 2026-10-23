
# btest

> A lightweight, fast, and low-end-friendly LAN speed test utility built with Go and modern web standards.

[![Go Version](https://img.shields.io/badge/Go-1.22%2B-blue.svg)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

---

## ● Overview

**btest** is designed to measure network performance (Download and Upload speeds) across a local area network (LAN). Powered by a high-performance Go backend with **zero external dependencies** and an optimized browser frontend, it avoids heavy memory overhead, making it ideal for running on low-end hardware, single-board computers (like Raspberry Pi), or home routers.

---

## ● Features

* **Zero External Dependencies:** Built entirely using Go's standard library (`net/http`, `crypto/rand`, `embed`).
* **Embedded Frontend:** The HTML/CSS/JS dashboard is compiled directly into the binary via `go:embed`, requiring only a single executable file to run.
* **Flexible Test Modes:**
  * **Standard Test (10s):** Automated 10-second download and upload benchmark.
  * **Limitless Tests:** Continuous stream testing for stress-testing network links.
* **Low-End Client Optimizations:** Utilizes native JavaScript `Blob` and `WritableStream` implementations to prevent V8 memory bloat and garbage collection spikes during heavy throughput.

---

## ● Prerequisites

* [Go (v1.26.5 or higher)](https://golang.org/doc/install) installed on your host machine.

---

## ● Installation & Running

1. **Clone the repository:**
```bash
git clone https://github.com/tin-studio/btest.git
cd btest

```

2. **Run the server directly:**
```bash
go run main.go

```


3. **Build a standalone binary (optional):**
```bash
go build -o btest main.go
./btest

```


**Or just download binaries/executables at releases page**

*if encountered error especially on older windows, please consider to install and enable [VxKex](https://github.com/i486/VxKex) to the executable*

The server will start on port **8083** by default (`http://0.0.0.0:8083`).

---

## ● Usage

1. Open your browser and navigate to `http://<your-server-ip>:8083`.


2. Choose your preferred test mode from the dashboard:


* **Run Standard Test (10s):** Measures standard throughput over a 10-second window.


* **Limitless Download/Upload:** Continuously pushes or pulls data until manually stopped.




3. View real-time speeds in **MBps** directly on screen.



---

## ● API Endpoints

* `GET /` — Serves the embedded web dashboard.
* `GET /download` — Generates a high-speed pseudo-random byte stream (`256 KB` chunks) with caching disabled for accurate downlink testing.


* `POST /upload` — Accepts raw binary data streams, efficiently discards them via `io.Discard`, and returns JSON confirmation with the total bytes received.



---

## ● Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request for improvements.

---

## ● License

Distributed under the MIT License. See `LICENSE` for details.

