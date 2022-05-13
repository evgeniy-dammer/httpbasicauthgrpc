package handlers

import (
	"context"
	"strings"

	productservice "github.com/evgeniy-dammer/httpbasicauthgrpc/productservice/proto"
)

type ProductServiceServer struct {
	productservice.UnimplementedProductServiceServer
}

var products []*productservice.Product

func init() {
	products = []*productservice.Product{
		{
			Id:       "p01",
			Name:     "tivi 1",
			Quantity: 27,
			Price:    3,
			Status:   true,
		},
		{
			Id:       "p02",
			Name:     "tivi 2",
			Quantity: 25,
			Price:    4,
			Status:   false,
		},
		{
			Id:       "p03",
			Name:     "laptop 3",
			Quantity: 23,
			Price:    2,
			Status:   true,
		},
	}
}

func (*ProductServiceServer) FindAll(ctx context.Context, in *productservice.FindAllRequest) (*productservice.FindAllResponse, error) {
	return &productservice.FindAllResponse{Products: products}, nil
}

func (*ProductServiceServer) Search(ctx context.Context, in *productservice.SearchRequest) (*productservice.SearchResponse, error) {
	var result []*productservice.Product

	for _, product := range products {
		if strings.Contains(product.Name, in.Keyword) {
			result = append(result, product)
		}
	}

	return &productservice.SearchResponse{Products: result}, nil
}
