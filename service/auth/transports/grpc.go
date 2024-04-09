package transports

import (
	"context"
	"mygomall/service/auth/endpoints"
	"mygomall/service/auth/pb"

	"github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	pb.UnimplementedAuthServiceServer
	getToken    grpc.Handler
	verifyToken grpc.Handler
}

func NewAuthGRPCServer(endpoints endpoints.AuthEndpoint) pb.AuthServiceServer {
	return &gRPCServer{
		getToken: grpc.NewServer(
			endpoints.GetToken,
			decodeGetTokenRequest,
			encodeToken,
		),
		verifyToken: grpc.NewServer(
			endpoints.VerifyToken,
			decodeToken,
			encodeVerifyTokenResponse,
		),
	}
}

func (s *gRPCServer) GetToken(ctx context.Context, req *pb.GetTokenRequest) (*pb.Token, error) {
	_, resp, err := s.getToken.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.Token), nil
}

func (s *gRPCServer) VerifyToken(ctx context.Context, req *pb.Token) (*pb.VerifyTokenResponse, error) {
	_, resp, err := s.verifyToken.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.VerifyTokenResponse), nil
}
