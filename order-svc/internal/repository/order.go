package repository

import (
	"context"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/internal/domain"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *domain.Order) error
	GetOrderByUser(ctx context.Context, userId uint) ([]*domain.Order, error)
	GetOrderById(ctx context.Context, orderId, userId uint) (*domain.Order, error)
	WithTx(tx *gorm.DB) OrderRepository
}

type orderRepository struct {
	dbWrite *gorm.DB
	dbRead  *gorm.DB
	tx      *gorm.DB
}

func (o *orderRepository) CreateOrder(ctx context.Context, order *domain.Order) error {
	return exec(o.dbWrite, o.tx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return err
		}
		return nil
	})
}

func (o *orderRepository) GetOrderByUser(ctx context.Context, userId uint) ([]*domain.Order, error) {
	var orders []*domain.Order
	if err := exec(o.dbRead, o.tx).Preload("OrderItems").Where("user_id = ?", userId).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (o *orderRepository) GetOrderById(ctx context.Context, orderId, userId uint) (*domain.Order, error) {
	var order *domain.Order
	if err := exec(o.dbRead, o.tx).Preload("OrderItems").Where("order_id = ? AND user_id = ?", orderId, userId).First(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (o *orderRepository) WithTx(tx *gorm.DB) OrderRepository {
	return &orderRepository{
		dbWrite: o.dbWrite,
		dbRead:  o.dbRead,
		tx:      tx,
	}
}

func NewOrderRepository(dbWrite, dbRead *gorm.DB) OrderRepository {
	return &orderRepository{
		dbWrite: dbWrite,
		dbRead:  dbRead,
	}
}
