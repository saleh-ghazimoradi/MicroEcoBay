package repository

import (
	"context"
	"errors"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/customErr"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/domain"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/slg"
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
	if err := c.dbWrite.WithContext(ctx).Create(category).Error; err != nil {
		slg.Logger.Error("error on creating category", "error", err.Error())
		return customErr.ErrCatalogCreate
	}
	return nil
}

func (c *catalogRepository) GetAllCategories(ctx context.Context) ([]*domain.Category, error) {
	var categories []*domain.Category
	if err := c.dbRead.WithContext(ctx).Find(&categories).Error; err != nil {
		slg.Logger.Error("error on getting all categories", "error", err.Error())
		return nil, customErr.ErrCatalogGet
	}
	return categories, nil
}

func (c *catalogRepository) GetCategoryById(ctx context.Context, id uint) (*domain.Category, error) {
	var category domain.Category
	if err := c.dbRead.WithContext(ctx).First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slg.Logger.Error("category not found", "id", id)
			return nil, customErr.ErrNotFound
		}
		slg.Logger.Error("error on getting category", "error", err.Error())
		return nil, customErr.ErrCatalogGet
	}
	return &category, nil
}

func (c *catalogRepository) UpdateCategory(ctx context.Context, id uint, update *domain.Category) error {
	if err := c.dbWrite.WithContext(ctx).Model(&domain.Category{}).Where("id = ?", id).Updates(update).Error; err != nil {
		slg.Logger.Error("error on updating category", "error", err.Error())
		return customErr.ErrUpdateCatalog
	}
	return nil
}

func (c *catalogRepository) CreateProduct(ctx context.Context, product *domain.Product) error {
	if err := c.dbWrite.WithContext(ctx).Create(product).Error; err != nil {
		slg.Logger.Error("error on creating product", "error", err.Error())
		return customErr.ErrProductCreate
	}
	return nil
}

func (c *catalogRepository) GetAllProducts(ctx context.Context) ([]*domain.Product, error) {
	var products []*domain.Product
	if err := c.dbRead.WithContext(ctx).Find(&products).Error; err != nil {
		slg.Logger.Error("error on getting all products", "error", err.Error())
		return nil, customErr.ErrProductGet
	}
	return products, nil
}

func (c *catalogRepository) GetProductById(ctx context.Context, id uint) (*domain.Product, error) {
	var product domain.Product
	if err := c.dbRead.WithContext(ctx).First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slg.Logger.Error("product not found", "id", id)
			return nil, customErr.ErrProductGet
		}
		return nil, customErr.ErrProductGet
	}
	return &product, nil
}

func (c *catalogRepository) UpdateProduct(ctx context.Context, id uint, update *domain.Product) error {
	if err := c.dbWrite.WithContext(ctx).Model(&domain.Product{}).Where("id = ?", id).Updates(update).Error; err != nil {
		slg.Logger.Error("error on updating product", "error", err.Error())
		return customErr.ErrProductUpdate
	}
	return nil
}

func NewCatalogRepository(dbWrite, dbRead *gorm.DB) CatalogRepository {
	return &catalogRepository{
		dbWrite: dbWrite,
		dbRead:  dbRead,
	}
}
