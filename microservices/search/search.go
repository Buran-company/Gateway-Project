package search

import (
	"context"
	"fmt"

	pb "github.com/Buran_Company/Gateway/grpc"
)

type SearchServer struct {
	pb.UnimplementedSearchServer
}

func (s *SearchServer) AddressSearch(ctx context.Context, in *pb.Msg) (*pb.Response, error) {
	fmt.Println(in)
	return &pb.Response{Resp: "Search success!"}, nil
}