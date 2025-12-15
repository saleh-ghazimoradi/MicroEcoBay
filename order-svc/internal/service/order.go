package service

import (
	"context"
	"errors"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/internal/domain"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/internal/dto"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/internal/repository"
)

type OrderService interface {
	CreateOrder(ctx context.Context, input *dto.CreateOrderRequest) (*domain.Order, error)
	GetOrderByUser(ctx context.Context, userId uint) ([]*domain.Order, error)
	GetOrderById(ctx context.Context, orderId, userId uint) (*domain.Order, error)
}

type orderService struct {
	orderRepository repository.OrderRepository
}

func (o *orderService) CreateOrder(ctx context.Context, input *dto.CreateOrderRequest) (*domain.Order, error) {
	if len(input.Items) == 0 {
		return nil, errors.New("order must have at least one item")
	}

	var (
		items []domain.OrderItem
		total float64
	)

	for _, item := range input.Items {
		if item.Quantity == 0 {
			continue
		}
		items = append(items, domain.OrderItem{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
			Price:     item.Price,
		})
		total += float64(item.Quantity) * item.Price
	}

	if len(items) == 0 {
		return nil, errors.New("order must have at least one item")
	}

	order := &domain.Order{
		UserId:     input.UserId,
		Total:      total,
		Status:     "pending",
		OrderRef:   generateOrderRef(),
		OrderItems: items,
	}

	if err := o.orderRepository.CreateOrder(ctx, order); err != nil {
		return nil, err
	}

	return order, nil
}

func (o *orderService) GetOrderByUser(ctx context.Context, userId uint) ([]*domain.Order, error) {
	orders, err := o.orderRepository.GetOrderByUser(ctx, userId)
	if err != nil {
		return nil, errors.New("failed to retrieve orders")
	}
	return orders, nil
}

func (o *orderService) GetOrderById(ctx context.Context, orderId, userId uint) (*domain.Order, error) {
	order, err := o.orderRepository.GetOrderById(ctx, orderId, userId)
	if err != nil {
		return nil, errors.New("order does not exist")
	}
	return order, nil
}

func generateOrderRef() string {
	return "ORD_123456"
}

func NewOrderService(orderRepository repository.OrderRepository) OrderService {
	return &orderService{
		orderRepository: orderRepository,
	}
}
