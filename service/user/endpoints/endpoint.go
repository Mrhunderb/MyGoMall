package endpoints

import (
	"context"
	"mygomall/service/user/service"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	Login    endpoint.Endpoint
	Register endpoint.Endpoint
	Info     endpoint.Endpoint
}

func MakeEndpoints(s service.UserService) Endpoints {
	return Endpoints{
		Login:    makeLoginEndpoint(s),
		Register: makeRegisterEndpoint(s),
		Info:     makeInfoEndpoint(s),
	}
}

func makeLoginEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(LoginRequest)
		id, token, err := s.Login(ctx, req.Username, req.Password)
		if err != nil {
			return LoginResponse{ID: uint64(id), Token: token}, err
		}
		return LoginResponse{ID: uint64(id), Token: token}, nil
	}
}

func makeRegisterEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(RegisterRequest)
		id, token, err := s.Register(ctx, req.Username, req.Password)
		if err != nil {
			return RegisterResponse{ID: uint64(id), Token: token}, err
		}
		return RegisterResponse{ID: uint64(id), Token: token}, nil
	}
}

func makeInfoEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(InfoRequest)
		user, err := s.Info(ctx, req.ID)
		if err != nil {
			return InfoResponse{
				ID:       0,
				Username: "",
				Gender:   0,
				Phone:    "",
			}, err
		}
		return InfoResponse{
			ID:       uint64(user.ID),
			Username: user.Username,
			Gender:   int32(user.Gender),
			Phone:    user.Phone,
		}, nil
	}
}
