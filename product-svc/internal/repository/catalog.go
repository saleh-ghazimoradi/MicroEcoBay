package repository

import (
	"context"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/domain"
	"gorm.io/gorm"
)

type CatalogRepository interface {
	CreateCategory(ctx context.Context, category *domain.Category) error
	GetAllCategories(ctx context.Context) ([]*domain.Category, error)
	GetCategoryById(ctx context.Context, id uint) (*domain.Category, error)
	UpdateCategory(ctx context.Context, update *domain.Category) error

	CreateProduct(ctx context.Context, product *domain.Product) error
	GetAllProducts(ctx context.Context) ([]*domain.Product, error)
	GetProductById(ctx context.Context, id uint) (*domain.Product, error)
	UpdateProduct(ctx context.Context, update *domain.Product) error
}

type catalogRepository struct {
	db *gorm.DB
}

func (c catalogRepository) CreateCategory(ctx context.Context, category *domain.Category) error {
	//TODO implement me
	panic("implement me")
}

func (c catalogRepository) GetAllCategories(ctx context.Context) ([]*domain.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c catalogRepository) GetCategoryById(ctx context.Context, id uint) (*domain.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c catalogRepository) UpdateCategory(ctx context.Context, update *domain.Category) error {
	//TODO implement me
	panic("implement me")
}

func (c catalogRepository) CreateProduct(ctx context.Context, product *domain.Product) error {
	//TODO implement me
	panic("implement me")
}

func (c catalogRepository) GetAllProducts(ctx context.Context) ([]*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (c catalogRepository) GetProductById(ctx context.Context, id uint) (*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (c catalogRepository) UpdateProduct(ctx context.Context, update *domain.Product) error {
	//TODO implement me
	panic("implement me")
}

func NewCatalogRepository(db *gorm.DB) CatalogRepository {
	return &catalogRepository{
		db: db,
	}
}
