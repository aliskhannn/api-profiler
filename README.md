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

## Profiling (pprof)

Открой список профилей в браузере:  
`http://127.0.0.1:6060/debug/pprof/`

### CPU profile (30s)

Снять CPU-профиль на 30 секунд:

```bash
curl 'http://127.0.0.1:6060/debug/pprof/profile?seconds=30' -o cpu.pprof
```

Посмотреть интерактивно:

```bash
go tool pprof -http=:0 ./cpu.pprof
```

`profile?seconds=N` — стандартный способ получить CPU профиль через pprof HTTP-эндпоинт.

### Heap profile

Снять heap профиль:

```bash
curl 'http://127.0.0.1:6060/debug/pprof/heap' -o heap.pprof
```

Посмотреть:

```bash
go tool pprof -http=:0 ./heap.pprof
```

`/debug/pprof/heap` — стандартный heap профиль, отдаваемый `net/http/pprof`.

### Trace (10s)

Снять execution trace:

```bash
curl 'http://127.0.0.1:6060/debug/pprof/trace?seconds=10' -o trace.out
```

Открыть trace:

```bash
go tool trace trace.out
```

`/debug/pprof/trace?seconds=N` отдает trace, который затем анализируется через `go tool trace`.

## Benchmarks + benchstat

### Run benchmarks

Снять baseline (до изменений) и сохранить в файл:

```bash
go test -run '^$' -bench . -benchmem -count=10 ./... > old.txt
```

После изменений (после оптимизации) снять снова:

```bash
go test -run '^$' -bench . -benchmem -count=10 ./... > new.txt
```

Флаг `-count=10` нужен, чтобы benchstat сравнивал результаты на основе нескольких прогонов, а не одного случайного измерения.

### Install benchstat

```bash
go install golang.org/x/perf/cmd/benchstat@latest
```

Команда `benchstat` находится в `golang.org/x/perf/cmd/benchstat`.

### Compare results

```bash
benchstat old.txt new.txt
```

`benchstat` предназначен для статистического A/B сравнения результатов бенчмарков (обычно “до/после”).