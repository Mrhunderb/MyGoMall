package endpoints

import (
	"context"
	"mygomall/service/auth/service"

	"github.com/go-kit/kit/endpoint"
)

type AuthEndpoint struct {
	GetToken    endpoint.Endpoint
	VerifyToken endpoint.Endpoint
}

func MakeEndpoints(s service.AuthService) AuthEndpoint {
	return AuthEndpoint{
		GetToken:    makeGetTokenEndpoint(s),
		VerifyToken: makeVerifyTokenEndpoint(s),
	}
}

func makeGetTokenEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetTokenRequset)
		token, err := s.GetToken(ctx, req.UserId)
		if err != nil {
			return Token{Token: token}, err
		}
		return Token{Token: token}, nil
	}
}

func makeVerifyTokenEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(Token)
		valid, err := s.VerifyToken(ctx, req.Token)
		if err != nil {
			return VerifyTokenResponse{Valid: false}, err
		}
		return VerifyTokenResponse{Valid: valid}, nil
	}
}
