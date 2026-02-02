package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var (
		target = flag.String("target", "http://127.0.0.1:8080", "base URL")
		dur    = flag.Duration("dur", 20*time.Second, "duration")
		c      = flag.Int("c", runtime.GOMAXPROCS(0)*4, "concurrency")
	)
	flag.Parse()

	client := &http.Client{Timeout: 2 * time.Second}
	var ok uint64

	stop := time.Now().Add(*dur)
	var wg sync.WaitGroup
	wg.Add(*c)

	for i := 0; i < *c; i++ {
		go func(id int) {
			defer wg.Done()
			for time.Now().Before(stop) {
				u := fmt.Sprintf("%s/sum?a=%d&b=%d", *target, id, id+1)
				resp, err := client.Get(u)
				if err != nil {
					continue
				}
				io.Copy(io.Discard, resp.Body)
				resp.Body.Close()
				if resp.StatusCode == 200 {
					atomic.AddUint64(&ok, 1)
				}
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("ok=%d\n", ok)
}
