package user

import (
	"context"
	"fmt"

	pb "github.com/Buran_Company/Gateway/grpc"
)

type UserServer struct {
	pb.UnimplementedUserServer
}

func (u *UserServer) Profile(ctx context.Context, in *pb.Msg) (*pb.Response, error) {
	fmt.Println(in)
	return &pb.Response{Resp: "User success!"}, nil
}