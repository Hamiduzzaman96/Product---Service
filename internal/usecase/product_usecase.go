package usecase

import (
	"context"

	"github.com/Hamiduzzaman96/Product---Service/internal/domain"
	"github.com/Hamiduzzaman96/Product---Service/internal/repository"
)

type ProductUsecase struct {
	repo *repository.ProductRepository
}

func NewProductUsecase(r *repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{repo: r}
}

func (u *ProductUsecase) GetProductByID(ctx context.Context, id int64) (*domain.Product, error) {
	return u.repo.GetByID(ctx, id)
}
