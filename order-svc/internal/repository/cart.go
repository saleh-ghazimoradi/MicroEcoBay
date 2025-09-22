package repository

import (
	"context"
	"errors"
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
	var cart domain.Cart
	if err := c.dbRead.WithContext(ctx).Preload("Items").Where("user_id = ?", userId).First(&cart).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			cart = domain.Cart{UserId: userId}
			if err := c.dbWrite.WithContext(ctx).Create(&cart).Error; err != nil {
				return nil, err
			}
			return &cart, nil
		}
		return nil, err
	}
	return &cart, nil
}

func (c *cartRepository) GetCart(ctx context.Context, userId uint) (*domain.Cart, error) {
	var cart domain.Cart
	if err := c.dbRead.WithContext(ctx).Preload("Items").Where("user_id = ?", userId).First(&cart).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

func (c *cartRepository) AddOrIncrement(ctx context.Context, userId uint, item *domain.CartItem) error {
	cart, err := c.GetCart(ctx, userId)
	if err != nil {
		return err
	}

	var existingCart domain.CartItem
	if err = c.dbRead.WithContext(ctx).Where("cart_id = ? AND product_id = ?", cart.ID, item.ProductID).First(&existingCart).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			item.CartID = cart.ID
			return c.dbWrite.WithContext(ctx).Create(&item).Error
		}
		return err
	}

	newQty := existingCart.Qty + item.Qty

	return c.dbWrite.WithContext(ctx).Model(&existingCart).Updates(map[string]any{
		"qty":          newQty,
		"price":        item.Price,
		"product_name": item.ProductName,
		"image_url":    item.ImageURL,
	}).Error

}

func (c *cartRepository) UpdateQty(ctx context.Context, userId, productId, qty uint) error {
	cart, err := c.GetOrCreateCart(ctx, userId)
	if err != nil {
		return err
	}

	if qty == 0 {
		return c.Remove(ctx, userId, productId)
	}

	return c.dbWrite.WithContext(ctx).Model(&domain.CartItem{}).Where("cart_id = ? AND product_id = ?", cart.ID, productId).Updates(map[string]any{
		"qty": qty,
	}).Error
}

func (c *cartRepository) Remove(ctx context.Context, userId, productId uint) error {
	cart, err := c.GetOrCreateCart(ctx, userId)
	if err != nil {
		return err
	}

	return c.dbWrite.WithContext(ctx).Where("cart_id = ? AND product_id = ?", cart.ID, productId).Delete(domain.CartItem{}).Error
}

func (c *cartRepository) Clear(ctx context.Context, userId uint) error {
	cart, err := c.GetOrCreateCart(ctx, userId)
	if err != nil {
		return err
	}

	return c.dbWrite.WithContext(ctx).Where("cart_id = ?", cart.ID).Delete(domain.CartItem{}).Error
}

func NewCartRepository(dbWrite, dbRead *gorm.DB) CartRepository {
	return &cartRepository{
		dbWrite: dbWrite,
		dbRead:  dbRead,
	}
}
