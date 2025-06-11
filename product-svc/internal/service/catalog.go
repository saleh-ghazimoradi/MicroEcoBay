package service

import (
	"context"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/domain"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/dto"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/repository"
)

type CatalogService interface {
	CreateCategory(ctx context.Context, req dto.CreateCategoryReq) error
	GetAllCategories(ctx context.Context) ([]domain.Category, error)
	GetCategoryById(ctx context.Context, id uint) (*domain.Category, error)
	UpdateCategory(ctx context.Context, id uint, input dto.UpdateCategoryReq) error

	CreateProduct(ctx context.Context, input dto.CreateProductReq) error
	GetAllProducts(ctx context.Context) ([]domain.Product, error)
	GetProductById(ctx context.Context, id uint) (*domain.Product, error)
	UpdateProduct(ctx context.Context, id uint, input dto.UpdateProductReq) error
}

type catalogService struct {
	catalogRepository repository.CatalogRepository
}

func (c catalogService) CreateCategory(ctx context.Context, req dto.CreateCategoryReq) error {
	//TODO implement me
	panic("implement me")
}

func (c catalogService) GetAllCategories(ctx context.Context) ([]domain.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c catalogService) GetCategoryById(ctx context.Context, id uint) (*domain.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c catalogService) UpdateCategory(ctx context.Context, id uint, input dto.UpdateCategoryReq) error {
	//TODO implement me
	panic("implement me")
}

func (c catalogService) CreateProduct(ctx context.Context, input dto.CreateProductReq) error {
	//TODO implement me
	panic("implement me")
}

func (c catalogService) GetAllProducts(ctx context.Context) ([]domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (c catalogService) GetProductById(ctx context.Context, id uint) (*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (c catalogService) UpdateProduct(ctx context.Context, id uint, input dto.UpdateProductReq) error {
	//TODO implement me
	panic("implement me")
}

func NewCatalogService(catalogRepository repository.CatalogRepository) CatalogService {
	return &catalogService{
		catalogRepository: catalogRepository,
	}
}
