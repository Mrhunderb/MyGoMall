package service

import (
	"context"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type service struct {
	logger zap.Logger
	rdb    redis.Client
}

type AuthService interface {
	GetToken(ctx context.Context, id uint64) (string, error)
	VerifyToken(ctx context.Context, token string) (bool, error)
}

func NewAuthService(logger zap.Logger, rdb redis.Client) AuthService {
	return &service{
		logger: logger,
		rdb:    rdb,
	}
}

func (s *service) GetToken(ctx context.Context, id uint64) (string, error) {
	s.logger.Info("get token")
	return "token", nil
}

func (s *service) VerifyToken(ctx context.Context, token string) (bool, error) {
	s.logger.Info("verify token")
	return true, nil
}
