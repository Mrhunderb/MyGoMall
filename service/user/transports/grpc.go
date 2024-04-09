package transports

import (
	"context"
	"mygomall/service/user/endpoints"
	"mygomall/service/user/pb"

	"github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	pb.UnimplementedUserServer
	login    grpc.Handler
	register grpc.Handler
	info     grpc.Handler
}

func NewGRPCServer(endpoints endpoints.UserEndpoint) pb.UserServer {
	return &gRPCServer{
		login: grpc.NewServer(
			endpoints.Login,
			decodeGRPCLoginRequest,
			encodeGRPCLoginResponse,
		),
		register: grpc.NewServer(
			endpoints.Register,
			decodeGRPCRegisterRequest,
			encodeGRPCRegisterResponse,
		),
		info: grpc.NewServer(
			endpoints.Info,
			decodeGRPCInfoRequest,
			encodeGRPCInfoResponse,
		),
	}
}

func (s *gRPCServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	_, resp, err := s.login.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.LoginResponse), nil
}

func (s *gRPCServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	_, resp, err := s.register.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.RegisterResponse), nil
}

func (s *gRPCServer) UserInfo(ctx context.Context, req *pb.UserInfoRequest) (*pb.UserInfoResponse, error) {
	_, resp, err := s.info.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.UserInfoResponse), nil
}
