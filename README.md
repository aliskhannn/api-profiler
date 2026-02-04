# api-profiler

Цель: простой HTTP API на Go, создание нагрузки и оптимизация по CPU/памяти с использованием pprof и trace, а также бенчмарков + benchstat.

## Endpoints

- `GET /sum?a=1&b=2` -> `{"sum":3}`
- `GET /json` -> небольшой JSON-ответ
- pprof UI: `http://127.0.0.1:6060/debug/pprof/` (отдельный pprof-сервер)

## Run API

В одном терминале:

```bash
go run ./cmd/api
```

По умолчанию:
- API: `http://localhost:8080`
- pprof: `http://127.0.0.1:6060/debug/pprof/`

Проверка:

```bash
curl 'http://localhost:8080/sum?a=1&b=2'
curl 'http://localhost:8080/json'
```

## Run load generator

Во втором терминале (пока API запущен):

```bash
go run ./cmd/loadgen -target=http://127.0.0.1:8080 -dur=30s -c=32
```

## Profiling (pprof + trace)

pprof/trace доступны через HTTP-эндпоинты `net/http/pprof` (например, `/debug/pprof/profile`, `/debug/pprof/heap`, `/debug/pprof/trace`).

Открой список профилей в браузере:  
`http://127.0.0.1:6060/debug/pprof/`

### CPU profile (30s)

Снять CPU-профиль на 30 секунд (параметр `seconds` поддерживается):

```bash
curl 'http://127.0.0.1:6060/debug/pprof/profile?seconds=30' -o cpu.pprof
```

Посмотреть интерактивно:

```bash
go tool pprof -http=:0 ./cpu.pprof
```

### Heap profile

Снять heap профиль:

```bash
curl 'http://127.0.0.1:6060/debug/pprof/heap' -o heap.pprof
```

Посмотреть:

```bash
go tool pprof -http=:0 ./heap.pprof
```

### Trace (10s)

Снять execution trace (параметр `seconds` поддерживается):

```bash
curl 'http://127.0.0.1:6060/debug/pprof/trace?seconds=10' -o trace.out
```

Открыть trace:

```bash
go tool trace trace.out
```

## Benchmarks + benchstat

`benchstat` делает статистическое A/B сравнение результатов бенчмарков (обычно “до/после”), и типичный сценарий — прогонять `go test -bench` с `-count=10` для обоих состояний.

### Run benchmarks

Baseline (до оптимизаций):

```bash
GOMAXPROCS=4 go test -run '^$' -bench . -benchmem -count=10 ./... > old.txt
```

После оптимизаций:

```bash
GOMAXPROCS=4 go test -run '^$' -bench . -benchmem -count=10 ./... > new.txt
```

### Install benchstat

```bash
go install golang.org/x/perf/cmd/benchstat@latest
```

### Compare results

```bash
benchstat old.txt new.txt
```

### Results (old -> new)

```text
goos: linux
goarch: amd64
pkg: github.com/aliskhannn/api-profiler/internal/api
cpu: Intel(R) Pentium(R) Gold 7505 @ 2.00GHz
              │   old.txt    │               new.txt                │
              │    sec/op    │    sec/op     vs base                │
SumHandler-4    1.626µ ± 23%   1.481µ ±  2%   -8.89% (p=0.001 n=10)
JSONHandler-4   1.913µ ± 15%   1.066µ ± 20%  -44.28% (p=0.000 n=10)
geomean         1.763µ         1.256µ        -28.75%

              │   old.txt    │               new.txt                │
              │     B/op     │     B/op      vs base                │
SumHandler-4    1.672Ki ± 0%   1.664Ki ± 0%   -0.47% (p=0.000 n=10)
JSONHandler-4   1.680Ki ± 0%   1.211Ki ± 0%  -27.91% (p=0.000 n=10)
geomean         1.676Ki        1.420Ki       -15.29%

              │  old.txt   │              new.txt               │
              │ allocs/op  │ allocs/op   vs base                │
SumHandler-4    21.00 ± 0%   19.00 ± 0%   -9.52% (p=0.000 n=10)
JSONHandler-4   21.00 ± 0%   14.00 ± 0%  -33.33% (p=0.000 n=10)
geomean         21.00        16.31       -22.34%
```

## What changed (high level)

- `Sum`: убран `fmt.Sprintf`, сборка JSON без форматтера (меньше CPU и allocs).
- `JSON`: убран `map + json.Encoder`, возвращается статический JSON.

## Commit history

```text
* 87a2c39 (HEAD -> main) chore: ignore benchmark output files
* d0b7859 perf(json): avoid map+json encoder, use static response
* 8bc7b18 perf(sum): remove fmt.Sprintf and reduce allocations
* d8763d6 docs: add profiling and benchmarking commands
* db32c61 (origin/main, origin/HEAD) chore: change go module, add .gitignore
* 989b676 tool: add load generator for profiling
* dffb47d test: add benchmarks for handlers
* 72fc4d7 chore: add dedicated pprof server
* 713f69a feat: baseline http api with /sum and /json
* 3362319 chore: init module and README skeleton
* 943377a chore: init go module
```