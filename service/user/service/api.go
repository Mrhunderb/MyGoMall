package service

import (
	"context"
	"fmt"
	"mygomall/common/cryptx"
	"mygomall/common/jwtx"
	"mygomall/service/user/model"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type service struct {
	logger zap.Logger
	db     gorm.DB
}

type UserService interface {
	Login(ctx context.Context, username, password string) (uint, string, error)
	Register(ctx context.Context, username, password string) (uint, string, error)
	Info(ctx context.Context, id uint64) (string, error)
}

func NewUserService(logger zap.Logger, db gorm.DB) UserService {
	return &service{
		logger: logger,
		db:     db,
	}
}

func (s *service) Login(ctx context.Context, username, password string) (uint, string, error) {
	return 0, "", nil
}

func (s *service) Register(ctx context.Context, username, password string) (id uint, token string, err error) {
	result := s.db.First(&model.User{}, "username = ?", username)
	if result.Error == nil {
		return 0, "", fmt.Errorf("username %s already exists", username)
	}
	cryptPass, err := cryptx.PasswordEncrypt("hello", password)
	if err != nil {
		return 0, "", fmt.Errorf("could not encrypt password")
	}
	user := model.User{
		Username: username,
		Password: cryptPass,
	}
	result = s.db.Create(&user)
	if result.Error == nil {
		return 0, "", result.Error
	}
	token, _ = jwtx.GetToken("hello", time.Now().Unix(), 60*60, int64(user.ID))
	return user.ID, token, nil
}

func (s *service) Info(ctx context.Context, id uint64) (string, error) {
	return "", nil
}
