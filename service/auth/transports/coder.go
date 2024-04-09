package transports

import (
	"context"
	"mygomall/service/auth/endpoints"
	"mygomall/service/auth/pb"
)

func decodeGetTokenRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GetTokenRequest)
	return endpoints.GetTokenRequset{UserId: req.UserId}, nil
}

func encodeToken(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoints.Token)
	return &pb.Token{Token: resp.Token}, nil
}

func decodeToken(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(*pb.Token)
	return endpoints.Token{Token: resp.Token}, nil
}

func encodeVerifyTokenResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoints.VerifyTokenResponse)
	return &pb.VerifyTokenResponse{Valid: resp.Valid}, nil
}
