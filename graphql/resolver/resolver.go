package resolver

import (
	"context"

	"google.golang.org/grpc"

	ppb "github.com/rendyfebry/go-graphql-grpc/product/api"
)

// Resolver ...
type Resolver struct {
	productClient ppb.ProductSvcClient
}

// NewResolver ...
func NewResolver(conn *grpc.ClientConn) *Resolver {
	return &Resolver{
		productClient: ppb.NewProductSvcClient(conn),
	}
}

// Hello ...
func (r *Resolver) Hello(ctx context.Context) (string, error) {
	return "Hello, world!", nil
}
