package service

import (
	"context"
	"fmt"
	"mygomall/common/jwtx"
	"time"

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
	now := time.Now().Unix()
	expire := time.Second * 3600
	token, err := jwtx.GetToken("", now, int64(expire), id)
	if err != nil {
		return "", err
	}
	s.rdb.Set(ctx, fmt.Sprintf("%d", id), token, expire)
	return token, nil
}

func (s *service) VerifyToken(ctx context.Context, token string) (bool, error) {
	return jwtx.IsExpired(token, ""), nil
}
