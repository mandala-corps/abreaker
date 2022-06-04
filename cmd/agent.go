package cmd

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/mandala-corps/abreaker/internal/dto"
)

func AgentExecute(ctx context.Context, config *dto.Config) {
	// TODO add defer function for recover panic
	for {
		var wg sync.WaitGroup
		// watcher all aims
		for _, w := range config.Agent.Watchers {
			wg.Add(1)
			go coroutineWatch(ctx, &wg, w)
		}
		// wait for all coroutines done
		wg.Wait()
		// wait for next request
		time.Sleep(time.Duration(config.Agent.Interval) * time.Second)
	}
}

func coroutineWatch(ctx context.Context, wg *sync.WaitGroup, w *dto.Watcher) {
	defer wg.Done()
	// make request
	switch w.Method {
	case "http":
		// do request http
		rs := makeHttpRequest(ctx, w)
		syncServer(ctx, rs)
	default:
		fmt.Printf("Method %s not implemented\n", w.Method)
	}

}

func makeHttpRequest(ctx context.Context, w *dto.Watcher) *dto.Result {
	rs := &dto.Result{}
	// make get request
	startAt := time.Now()
	resp, err := http.Get(w.Addr)
	if err != nil {
		rs.Err = err
		return rs
	}
	// get result time
	endAt := time.Now()

	if resp.StatusCode >= 300 {
		rs.Err = fmt.Errorf("response dont return with successful status code: returns %d", resp.StatusCode)
	}

	rs.Duration = int(endAt.Sub(startAt).Microseconds())

	return rs
}

func syncServer(ctx context.Context, rs *dto.Result) {
	fmt.Print(rs)
}
