package grpc

import (
	"context"
	"database/sql"

	pb "github.com/Hamiduzzaman96/Product---Service/Product---Service/proto/productpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Hamiduzzaman96/Product---Service/internal/usecase"
)

type ProductHandler struct {
	pb.UnimplementedProductServiceServer
	usecase *usecase.ProductUsecase
}

func NewProductHandler(u *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{usecase: u}
}

func (h *ProductHandler) GetProductByID(
	ctx context.Context,
	req *pb.GetProductRequest,
) (*pb.ProductResponse, error) {
	product, err := h.usecase.GetProductByID(ctx, req.Id)

	if err == sql.ErrNoRows {
		return nil, status.Error(codes.NotFound, "Product Not Found")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.ProductResponse{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       int32(product.Stock),
		CreatedAt:   product.CreatedAt.String(),
	}, nil
}
