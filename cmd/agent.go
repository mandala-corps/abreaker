package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/mandala-corps/abreaker/internal/dependency"
	"github.com/mandala-corps/abreaker/internal/dto"
	"github.com/mandala-corps/abreaker/internal/service"
)

func AgentExecute(ctx context.Context, config *dto.Config) {
	// TODO add defer function for recover panic
	for {
		var wg sync.WaitGroup
		// watcher all aims
		for n, w := range config.Agent.Watchers {
			wg.Add(1)
			go coroutineWatch(ctx, &wg, w, n)
		}
		// wait for all coroutines done
		wg.Wait()
		// wait for next request
		time.Sleep(time.Duration(config.Agent.Interval) * time.Second)
	}
}

func coroutineWatch(ctx context.Context, wg *sync.WaitGroup, w *dto.Watcher, n string) {
	defer wg.Done()
	// make request
	switch w.Method {
	case "http":
		// do request http
		rs := makeHttpRequest(ctx, w, n)
		// TODO check if returns error
		syncServer(ctx, rs)
	default:
		fmt.Printf("Method %s not implemented\n", w.Method)
	}

}

func makeHttpRequest(ctx context.Context, w *dto.Watcher, name string) *dto.Result {
	rs := &dto.Result{
		Name: name,
	}
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

	rs.Duration = int(endAt.Sub(startAt).Milliseconds())

	return rs
}

func syncServer(ctx context.Context, rs *dto.Result) error {
	log.Printf("Request for %s replied in %d ms", rs.Name, rs.Duration)

	s, ok := ctx.Value(dependency.ServerServiceKey).(service.ServerSyncService)
	if !ok {
		return errors.New("server service not return ServerSyncService")
	}

	return s.SendResult(ctx, rs)
}
