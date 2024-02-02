package auth

import (
	"context"
	"fmt"

	pb "github.com/Buran_Company/Gateway/grpc"
)

type AuthServer struct {
	pb.UnimplementedAuthServer
}

func (a *AuthServer) Register(ctx context.Context, in *pb.Msg) (*pb.Response, error) {
	fmt.Println(in)
	return &pb.Response{Resp: "Auth success!"}, nil
}