package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mandala-corps/abreaker/internal/dto"
	"github.com/mandala-corps/abreaker/internal/http/handle"
)

func ExecuteServer(ctx context.Context, cfg *dto.Config) {
	router := httprouter.New()

	router.POST("/recive", handle.Recive)

	addr := fmt.Sprintf("%s:%s", cfg.Server.Addr, cfg.Server.Port)

	log.Fatal(http.ListenAndServe(addr, router))
}
