package repository

import (
	"context"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/internal/domain"
	"gorm.io/gorm"
)

type CartRepository interface {
	GetOrCreateCart(ctx context.Context, userId uint) (*domain.Cart, error)
	GetCart(ctx context.Context, userId uint) (*domain.Cart, error)
	AddOrIncrement(ctx context.Context, userId uint, item *domain.CartItem) error
	UpdateQty(ctx context.Context, userId, productId, qty uint) error
	Remove(ctx context.Context, userId, productId uint) error
	Clear(ctx context.Context, userId uint) error
}

type cartRepository struct {
	dbWrite *gorm.DB
	dbRead  *gorm.DB
}

func (c *cartRepository) GetOrCreateCart(ctx context.Context, userId uint) (*domain.Cart, error) {
	return nil, nil
}

func (c *cartRepository) GetCart(ctx context.Context, userId uint) (*domain.Cart, error) {
	return nil, nil
}

func (c *cartRepository) AddOrIncrement(ctx context.Context, userId uint, item *domain.CartItem) error {
	return nil
}

func (c *cartRepository) UpdateQty(ctx context.Context, userId, productId, qty uint) error {
	return nil
}

func (c *cartRepository) Remove(ctx context.Context, userId, productId uint) error {
	return nil
}

func (c *cartRepository) Clear(ctx context.Context, userId uint) error {
	return nil
}

func NewCartRepository(dbWrite, dbRead *gorm.DB) CartRepository {
	return &cartRepository{
		dbWrite: dbWrite,
		dbRead:  dbRead,
	}
}
