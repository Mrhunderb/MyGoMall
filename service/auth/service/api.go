package service

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type service struct {
	rdb redis.Client
}

type AuthService interface {
	GetToken(ctx context.Context, username, password string) (string, error)
	VerifyToken(ctx context.Context, token string) (bool, error)
}

func NewAuthService(rdb redis.Client) AuthService {
	return &service{
		rdb: rdb,
	}
}

func (s *service) GetToken(ctx context.Context, username, password string) (string, error) {
	panic("implement me")
}

func (s *service) VerifyToken(ctx context.Context, token string) (bool, error) {
	panic("implement me")
}
