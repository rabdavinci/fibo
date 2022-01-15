package handlers

import (
	"context"

	"github.com/rabdavinci/fibo/data"
	pb "github.com/rabdavinci/fibo/gen/proto"
)

type FiboApiServer struct {
	pb.UnimplementedFiboApiServer
}

func (s *FiboApiServer) CalculateList(ctx context.Context, req *pb.FiboRequest) (*pb.FiboResponse, error) {
	fi := data.FiboInterval{}
	fi.X = req.X

	list := data.CalculateSlice(&fi)

	return &pb.FiboResponse{List: list}, nil
}
