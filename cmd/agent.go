package cmd

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/mandala-corps/abreaker/internal/entities"
)

func AgentExecute(ctx context.Context, config *entities.Config) {
	// TODO add defer function for recover panic
	for {
		var wg sync.WaitGroup
		// watcher all aims
		for _, w := range config.Watchers {
			wg.Add(1)
			go coroutineWatch(&wg, w)
		}
		// wait for all coroutines done
		wg.Wait()
		// wait for next request
		time.Sleep(time.Duration(config.Interval) * time.Second)
	}
}

func coroutineWatch(wg *sync.WaitGroup, w *entities.Watcher) {
	defer wg.Done()
	// make request
	switch w.Method {
	case "http":
		// do request http
	default:
		fmt.Printf("Method %s not implemented\n", w.Method)
	}

}
