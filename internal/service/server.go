package service

import (
	"context"

	"github.com/mandala-corps/abreaker/internal/dto"
)

type ServerSyncService struct{}

func (*ServerSyncService) SendResult(ctx context.Context, rs *dto.Result) error {
	return nil
}

func NewServerService() *ServerSyncService {
	return &ServerSyncService{}
}
