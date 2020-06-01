package resolver

import (
	"context"
	"fmt"

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

// Products wil lreturn all product
func (r *Resolver) Products(ctx context.Context) ([]*ProductResolver, error) {
	products := make([]*ProductResolver, 0)

	// Query to DB or Upstream
	productList, err := r.productClient.GetProduct(ctx, &ppb.GetProductRequest{})
	if err != nil {
		return nil, err
	}

	fmt.Println(productList)

	for _, p := range productList.Products {
		s := &ProductResolver{
			m: &Product{
				ID:    p.Id,
				Name:  p.Name,
				Price: p.Price,
			},
		}

		products = append(products, s)
	}

	return products, nil
}
