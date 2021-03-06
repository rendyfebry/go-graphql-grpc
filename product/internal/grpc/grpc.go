package product

import (
	"context"

	pb "github.com/rendyfebry/go-graphql-grpc/product/api"
)

type grpcServer struct {
}

// NewServer creates a grpc server with required dependencies.
func NewServer() *grpcServer {
	return &grpcServer{}
}

func (gs *grpcServer) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	res := &pb.GetProductResponse{
		Products: []*pb.Product{
			&pb.Product{
				Id:    1,
				Name:  "Product A",
				Price: 1000,
			},
			&pb.Product{
				Id:    2,
				Name:  "Product B",
				Price: 2500,
			},
			&pb.Product{
				Id:    3,
				Name:  "Product C",
				Price: 5000,
			},
		},
	}

	return res, nil
}
