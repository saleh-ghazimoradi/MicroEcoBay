package service

import (
	"context"
	"errors"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/domain"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/dto"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/repository"
)

type CatalogService interface {
	CreateCategory(ctx context.Context, input *dto.CreateCategoryReq) error
	GetAllCategories(ctx context.Context) ([]*domain.Category, error)
	GetCategoryById(ctx context.Context, id uint) (*domain.Category, error)
	UpdateCategory(ctx context.Context, id uint, input *dto.UpdateCategoryReq) error

	CreateProduct(ctx context.Context, input *dto.CreateProductReq) error
	GetAllProducts(ctx context.Context) ([]*domain.Product, error)
	GetProductById(ctx context.Context, id uint) (*domain.Product, error)
	UpdateProduct(ctx context.Context, id uint, input *dto.UpdateProductReq) error
}

type catalogService struct {
	catalogRepository repository.CatalogRepository
}

func (c *catalogService) CreateCategory(ctx context.Context, input *dto.CreateCategoryReq) error {
	return c.catalogRepository.CreateCategory(ctx, &domain.Category{
		Name:        input.Name,
		Description: input.Description,
		ImageURL:    input.ImageURL,
		Status:      "publish",
	})
}

func (c *catalogService) GetAllCategories(ctx context.Context) ([]*domain.Category, error) {
	return c.catalogRepository.GetAllCategories(ctx)
}

func (c *catalogService) GetCategoryById(ctx context.Context, id uint) (*domain.Category, error) {
	return c.catalogRepository.GetCategoryById(ctx, id)
}

func (c *catalogService) UpdateCategory(ctx context.Context, id uint, input *dto.UpdateCategoryReq) error {
	existingCategory, err := c.catalogRepository.GetCategoryById(ctx, id)
	if err != nil {
		return errors.New("category not found")
	}

	if input.Name != nil {
		existingCategory.Name = *input.Name
	}
	if input.Description != nil {
		existingCategory.Description = *input.Description
	}
	if input.Status != nil {
		existingCategory.Status = *input.Status
	}

	return c.catalogRepository.UpdateCategory(ctx, id, existingCategory)
}

func (c *catalogService) CreateProduct(ctx context.Context, input *dto.CreateProductReq) error {
	return c.catalogRepository.CreateProduct(ctx, &domain.Product{
		Name:        input.Name,
		Description: input.Description,
		CategoryId:  input.CategoryID,
		Price:       input.Price,
		Stock:       input.Stock,
		ImageURL:    input.ImageURL,
		Status:      input.Status,
	})
}

func (c *catalogService) GetAllProducts(ctx context.Context) ([]*domain.Product, error) {
	return c.catalogRepository.GetAllProducts(ctx)
}

func (c *catalogService) GetProductById(ctx context.Context, id uint) (*domain.Product, error) {
	return c.catalogRepository.GetProductById(ctx, id)
}

func (c *catalogService) UpdateProduct(ctx context.Context, id uint, input *dto.UpdateProductReq) error {
	existingProduct, err := c.catalogRepository.GetProductById(ctx, id)
	if err != nil {
		return errors.New("product not found")
	}

	if input.Name != nil {
		existingProduct.Name = *input.Name
	}

	if input.Description != nil {
		existingProduct.Description = *input.Description
	}

	if input.Price != nil {
		existingProduct.Price = *input.Price
	}

	if input.ImageURL != nil {
		existingProduct.ImageURL = *input.ImageURL
	}

	if input.Status != nil {
		existingProduct.Status = *input.Status
	}

	if input.Stock != nil {
		existingProduct.Stock = *input.Stock
	}

	return c.catalogRepository.UpdateProduct(ctx, id, existingProduct)
}

func NewCatalogService(catalogRepository repository.CatalogRepository) CatalogService {
	return &catalogService{
		catalogRepository: catalogRepository,
	}
}
