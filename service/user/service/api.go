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

const (
	salt = "HWVOFkGgPTryzICwd7qnJaZR9KQ2i8xe"
)

type service struct {
	logger zap.Logger
	db     gorm.DB
}

type UserService interface {
	Login(ctx context.Context, username, password string) (uint, string, error)
	Register(ctx context.Context, username, password string) (uint, string, error)
	Info(ctx context.Context, id uint64) (*model.User, error)
}

func NewUserService(logger zap.Logger, db gorm.DB) UserService {
	return &service{
		logger: logger,
		db:     db,
	}
}

func (s *service) Login(ctx context.Context, username, password string) (id uint, token string, err error) {
	var user model.User
	result := s.db.First(&user, "username = ?", username)
	if result.Error != nil {
		return 0, "", fmt.Errorf("could not find user with username %s", username)
	}

	if !cryptx.PasswordVerify(salt, password, user.Password) {
		return 0, "", fmt.Errorf("password invalid")
	}

	token, _ = jwtx.GetToken("hello", time.Now().Unix(), 60*60, int64(user.ID))
	return user.ID, token, nil
}

func (s *service) Register(ctx context.Context, username, password string) (id uint, token string, err error) {
	result := s.db.First(&model.User{}, "username = ?", username)
	if result.Error == nil {
		return 0, "", fmt.Errorf("username %s already exists", username)
	}
	cryptPass, err := cryptx.PasswordEncrypt(salt, password)
	if err != nil {
		return 0, "", fmt.Errorf("could not encrypt password")
	}
	user := model.User{
		Username: username,
		Password: cryptPass,
	}
	result = s.db.Create(&user)
	if result.Error != nil {
		return 0, "", result.Error
	}
	token, _ = jwtx.GetToken("hello", time.Now().Unix(), 60*60, int64(user.ID))
	return user.ID, token, nil
}

func (s *service) Info(ctx context.Context, id uint64) (*model.User, error) {
	var user model.User
	result := s.db.First(&user, id)
	if result.Error != nil {
		return nil, fmt.Errorf("could not find user with id %d", id)
	}
	return &user, nil
}
