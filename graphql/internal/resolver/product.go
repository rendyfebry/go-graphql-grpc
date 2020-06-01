package resolver

import (
	"context"
)

// Product ...
type Product struct {
	ID    int32
	Name  string
	Price int32
}

// ProductResolver ...
type ProductResolver struct {
	m *Product
}

// ID ...
func (u *ProductResolver) ID(ctx context.Context) *int32 {
	return &u.m.ID
}

// Name ...
func (u *ProductResolver) Name(ctx context.Context) *string {
	return &u.m.Name
}

// Price ...
func (u *ProductResolver) Price(ctx context.Context) *int32 {
	return &u.m.Price
}
