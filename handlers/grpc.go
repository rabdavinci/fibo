package handlers

import (
	"context"
	"strconv"
	"strings"

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
	var out []string
	for _, v := range list {
		out = append(out, strconv.FormatInt(v, 10))
	}
	return &pb.FiboResponse{List: strings.Join(out, ",")}, nil
}
