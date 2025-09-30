# Simple Go API with Profiling and Benchmarks

This project is a simple Go HTTP API that performs basic operations (sum of two numbers) and provides **CPU and memory profiling** via Go’s `pprof`. It includes benchmarking and tools to optimize the API for performance.

---

## Project Architecture

```
.
├── cmd/
│   └── server/              # Main application entry point
├── internal/
│   └── api/
│       └── handlers/
│           └── sum/         # Sum HTTP handler
├── go.mod
├── Makefile
└── README.md
```

**Layers:**

* **cmd/server** — entry point, starts the HTTP server.
* **internal/api/handlers/sum** — HTTP handler for sum operation.
* **bin/** — folder for compiled binaries.

---

## Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/simple-api.git
cd simple-api
```

2. Install dependencies:

```bash
go mod tidy
```

3. Build the project:

```bash
make build
```

4. Run the server:

```bash
make run
```

Server runs at **[http://localhost:8080](http://localhost:8080)**.

---

## API Endpoint

**Sum endpoint:**

```
GET /sum?a=1&b=2
```

**Example response:**

```json
{
  "result": 3
}
```

---

## Load Testing

You can generate load with `hey`:

```bash
hey -n 10000 -c 50 "http://localhost:8080/sum?a=1&b=2"
```

**Parameters:**

* `-n 10000` — total number of requests
* `-c 50` — concurrent clients

---

## Profiling

Go pprof endpoints are automatically available:

* CPU profile: `/debug/pprof/profile?seconds=30`
* Heap profile: `/debug/pprof/heap`

Collect profiles using Makefile:

```bash
make cpu-profile
make heap-profile
```

Analyze profiles:

```bash
go tool pprof cpu.pprof
(pprof) top
(pprof) web
```

---

## Benchmarks

Run benchmarks with memory stats:

```bash
make bench
```

* Use `benchstat` to compare performance before/after optimization.

---

## Makefile Commands

| Command             | Description                           |
| ------------------- | ------------------------------------- |
| `make build`        | Build the binary                      |
| `make run`          | Run the server                        |
| `make clean`        | Remove binary                         |
| `make bench`        | Run benchmarks with memory stats      |
| `make cpu-profile`  | Collect 30s CPU profile via pprof     |
| `make heap-profile` | Collect memory heap profile via pprof |

---

This setup allows you to:

* Run and test the API
* Generate load and benchmark performance
* Profile CPU and memory usage with pprof
* Optimize your API based on real performance metrics