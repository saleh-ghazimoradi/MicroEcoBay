package repository

import (
	"context"
	"errors"
	"github.com/rs/zerolog"
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
	WithTx(tx *gorm.DB) CatalogRepository
}

type catalogRepository struct {
	dbWrite *gorm.DB
	dbRead  *gorm.DB
	tx      *gorm.DB
	logger  *zerolog.Logger
}

func (c *catalogRepository) CreateCategory(ctx context.Context, category *domain.Category) error {
	if err := exec(c.dbWrite, c.tx).WithContext(ctx).Create(category).Error; err != nil {
		c.logger.Error().Err(err).Msg("failed to create category")
		return ErrCatalogCreate
	}
	return nil
}

func (c *catalogRepository) GetAllCategories(ctx context.Context) ([]*domain.Category, error) {
	var categories []*domain.Category
	if err := exec(c.dbRead, c.tx).WithContext(ctx).Find(&categories).Error; err != nil {
		c.logger.Error().Err(err).Msg("failed to get all categories")
		return nil, ErrCatalogGet
	}
	return categories, nil
}

func (c *catalogRepository) GetCategoryById(ctx context.Context, id uint) (*domain.Category, error) {
	var category domain.Category
	if err := exec(c.dbRead, c.tx).WithContext(ctx).First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.logger.Warn().Uint("category_id", id).Msg("category not found")
			return nil, ErrNotFound
		}
		c.logger.Error().Err(err).Uint("category_id", id).Msg("failed to get category by id")
		return nil, ErrCatalogGet
	}
	return &category, nil
}

func (c *catalogRepository) UpdateCategory(ctx context.Context, id uint, update *domain.Category) error {
	if err := exec(c.dbWrite, c.tx).WithContext(ctx).Model(&domain.Category{}).Where("id = ?", id).Updates(update).Error; err != nil {
		c.logger.Error().Err(err).Uint("category_id", id).Msg("failed to update category")
		return ErrUpdateCatalog
	}
	return nil
}

func (c *catalogRepository) CreateProduct(ctx context.Context, product *domain.Product) error {
	if err := exec(c.dbWrite, c.tx).WithContext(ctx).Create(product).Error; err != nil {
		c.logger.Error().Err(err).Msg("failed to create product")
		return ErrProductCreate
	}
	return nil
}

func (c *catalogRepository) GetAllProducts(ctx context.Context) ([]*domain.Product, error) {
	var products []*domain.Product
	if err := exec(c.dbRead, c.tx).WithContext(ctx).Find(&products).Error; err != nil {
		c.logger.Error().Err(err).Msg("failed to get all products")
		return nil, ErrProductGet
	}
	return products, nil
}

func (c *catalogRepository) GetProductById(ctx context.Context, id uint) (*domain.Product, error) {
	var product domain.Product
	if err := exec(c.dbRead, c.tx).WithContext(ctx).First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.logger.Warn().Uint("product_id", id).Msg("product not found")
			return nil, ErrProductGet
		}
		c.logger.Error().Err(err).Uint("product_id", id).Msg("failed to get product by id")
		return nil, ErrProductGet
	}
	return &product, nil
}

func (c *catalogRepository) UpdateProduct(ctx context.Context, id uint, update *domain.Product) error {
	if err := exec(c.dbWrite, c.tx).WithContext(ctx).Model(&domain.Product{}).Where("id = ?", id).Updates(update).Error; err != nil {
		c.logger.Error().
			Err(err).Uint("product_id", id).Msg("failed to update product")
		return ErrProductUpdate
	}
	return nil
}

func (c *catalogRepository) WithTx(tx *gorm.DB) CatalogRepository {
	return &catalogRepository{
		dbWrite: c.dbWrite,
		dbRead:  c.dbRead,
		tx:      tx,
	}
}

func NewCatalogRepository(dbWrite, dbRead *gorm.DB, logger *zerolog.Logger) CatalogRepository {
	return &catalogRepository{
		dbWrite: dbWrite,
		dbRead:  dbRead,
		logger:  logger,
	}
}
