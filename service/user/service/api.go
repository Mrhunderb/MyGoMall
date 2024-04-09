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
		s.logger.Error(
			"user login could not find user",
			zap.String("method", "Login"),
			zap.String("username", username),
			zap.Time("time", time.Now()),
		)
		return 0, "", fmt.Errorf("could not find user with username %s", username)
	}

	if !cryptx.PasswordVerify(salt, password, user.Password) {
		s.logger.Error(
			"user login password invalid",
			zap.String("method", "Login"),
			zap.String("username", username),
			zap.Time("time", time.Now()),
		)
		return 0, "", fmt.Errorf("password invalid")
	}

	s.logger.Info(
		"user login success",
		zap.String("method", "Login"),
		zap.String("username", username),
		zap.Time("time", time.Now()),
	)
	token, _ = jwtx.GetToken("hello", time.Now().Unix(), 60*60, uint64(user.ID))
	return user.ID, token, nil
}

func (s *service) Register(ctx context.Context, username, password string) (id uint, token string, err error) {
	result := s.db.First(&model.User{}, "username = ?", username)
	if result.Error == nil {
		s.logger.Error(
			"user register username already exists",
			zap.String("method", "Register"),
			zap.String("username", username),
			zap.Time("time", time.Now()),
		)
		return 0, "", fmt.Errorf("username %s already exists", username)
	}
	cryptPass, err := cryptx.PasswordEncrypt(salt, password)
	if err != nil {
		s.logger.Error(
			"user register could not encrypt password",
			zap.String("method", "Register"),
			zap.String("username", username),
			zap.Time("time", time.Now()),
		)
		return 0, "", fmt.Errorf("could not encrypt password")
	}
	user := model.User{
		Username: username,
		Password: cryptPass,
	}
	result = s.db.Create(&user)
	if result.Error != nil {
		s.logger.Error(
			"user register could not create user",
			zap.String("method", "Register"),
			zap.String("username", username),
			zap.Time("time", time.Now()),
		)
		return 0, "", result.Error
	}
	token, _ = jwtx.GetToken("hello", time.Now().Unix(), 60*60, uint64(user.ID))
	s.logger.Info(
		"user register success",
		zap.String("method", "Register"),
		zap.String("username", username),
		zap.Time("time", time.Now()),
	)
	return user.ID, token, nil
}

func (s *service) Info(ctx context.Context, id uint64) (*model.User, error) {
	var user model.User
	result := s.db.First(&user, id)
	if result.Error != nil {
		s.logger.Error(
			"user info could not find user",
			zap.String("method", "UserInfo"),
			zap.Uint64("id", id),
			zap.Time("time", time.Now()),
		)
		return nil, fmt.Errorf("could not find user with id %d", id)
	}
	s.logger.Info(
		"user info success",
		zap.String("method", "UserInfo"),
		zap.Uint64("id", id),
		zap.Time("time", time.Now()),
	)
	return &user, nil
}
