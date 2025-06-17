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
	UpdateCategory(ctx context.Context, id uint, update *domain.Category) error

	CreateProduct(ctx context.Context, product *domain.Product) error
	GetAllProducts(ctx context.Context) ([]*domain.Product, error)
	GetProductById(ctx context.Context, id uint) (*domain.Product, error)
	UpdateProduct(ctx context.Context, id uint, update *domain.Product) error
}

type catalogRepository struct {
	dbWrite *gorm.DB
	dbRead  *gorm.DB
}

func (c *catalogRepository) CreateCategory(ctx context.Context, category *domain.Category) error {
	return c.dbWrite.WithContext(ctx).Create(category).Error
}

func (c *catalogRepository) GetAllCategories(ctx context.Context) ([]*domain.Category, error) {
	var categories []*domain.Category
	if err := c.dbRead.WithContext(ctx).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *catalogRepository) GetCategoryById(ctx context.Context, id uint) (*domain.Category, error) {
	var category domain.Category
	if err := c.dbRead.WithContext(ctx).First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c *catalogRepository) UpdateCategory(ctx context.Context, id uint, update *domain.Category) error {
	return c.dbWrite.WithContext(ctx).Model(&domain.Category{}).Where("id = ?", id).Updates(update).Error
}

func (c *catalogRepository) CreateProduct(ctx context.Context, product *domain.Product) error {
	return c.dbWrite.WithContext(ctx).Create(product).Error
}

func (c *catalogRepository) GetAllProducts(ctx context.Context) ([]*domain.Product, error) {
	var products []*domain.Product
	if err := c.dbRead.WithContext(ctx).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (c *catalogRepository) GetProductById(ctx context.Context, id uint) (*domain.Product, error) {
	var product domain.Product
	if err := c.dbRead.WithContext(ctx).First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (c *catalogRepository) UpdateProduct(ctx context.Context, id uint, update *domain.Product) error {
	return c.dbWrite.WithContext(ctx).Model(&domain.Product{}).Where("id = ?", id).Updates(update).Error
}

func NewCatalogRepository(dbWrite, dbRead *gorm.DB) CatalogRepository {
	return &catalogRepository{
		dbWrite: dbWrite,
		dbRead:  dbRead,
	}
}
