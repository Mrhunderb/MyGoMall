package service

import (
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type service struct {
	logger zap.Logger
	db     gorm.DB
}

type UserService interface {
	Login(ctx context.Context, username, password string) (string, error)
	Register(ctx context.Context, username, password string) (string, error)
	Info(ctx context.Context, id string) (string, error)
}

func NewUserService(logger zap.Logger, db gorm.DB) UserService {
	return &service{
		logger: logger,
		db:     db,
	}
}

func (s *service) Login(ctx context.Context, username, password string) (string, error) {
	return "", nil
}

func (s *service) Register(ctx context.Context, username, password string) (string, error) {
	return "", nil
}

func (s *service) Info(ctx context.Context, id string) (string, error) {
	return "", nil
}
