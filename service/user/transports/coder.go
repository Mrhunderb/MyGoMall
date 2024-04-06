package transports

import (
	"context"
	"mygomall/service/user/endpoints"
	"mygomall/service/user/pb"
)

func decodeGRPCLoginRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.LoginRequest)
	return endpoints.LoginRequest{Username: req.Username, Password: req.Password}, nil
}

func encodeGRPCLoginResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.LoginResponse)
	return &pb.LoginResponse{Id: resp.ID, Token: resp.Token}, nil
}

func decodeGRPCRegisterRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.RegisterRequest)
	return endpoints.RegisterRequest{Username: req.Username, Password: req.Password}, nil
}

func encodeGRPCRegisterResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.RegisterResponse)
	return &pb.RegisterResponse{Id: resp.ID, Token: resp.Token}, nil
}

func decodeGRPCInfoRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.UserInfoRequest)
	return endpoints.InfoRequest{ID: req.Id}, nil
}

func encodeGRPCInfoResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.InfoResponse)
	return &pb.UserInfoResponse{
		Id:       resp.ID,
		Username: resp.Username,
		Gender:   resp.Gender,
		Phone:    resp.Phone,
	}, nil
}
