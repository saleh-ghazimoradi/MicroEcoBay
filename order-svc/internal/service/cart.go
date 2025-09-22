package service

import (
	"context"
	"errors"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/internal/domain"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/internal/dto"
	"github.com/saleh-ghazimoradi/MicroEcoBay/order_service/internal/repository"
)

type CartService interface {
	Add(ctx context.Context, userId uint, input *dto.Cart) error
	UpdateQty(ctx context.Context, userId, productId, qty uint) error
	Remove(ctx context.Context, userId, productId uint) error
	Get(ctx context.Context, userId uint) (*dto.CartResponse, error)
}

type cartService struct {
	cartRepository repository.CartRepository
}

func (c *cartService) Add(ctx context.Context, userId uint, input *dto.Cart) error {
	if input.Qty == 0 {
		return errors.New("qty must be grater than zero")
	}

	return c.cartRepository.AddOrIncrement(ctx, userId, &domain.CartItem{
		ProductID:   input.ProductID,
		ProductName: input.ProductName,
		ImageURL:    input.ImageURL,
		Qty:         input.Qty,
		Price:       input.Price,
	})
}

func (c *cartService) UpdateQty(ctx context.Context, userId, productId, qty uint) error {
	return c.cartRepository.UpdateQty(ctx, userId, productId, qty)
}

func (c *cartService) Remove(ctx context.Context, userId, productId uint) error {
	return c.cartRepository.Remove(ctx, userId, productId)
}

func (c *cartService) Get(ctx context.Context, userId uint) (*dto.CartResponse, error) {
	cart, err := c.cartRepository.GetOrCreateCart(ctx, userId)
	if err != nil {
		return nil, err
	}

	resp := &dto.CartResponse{UserID: userId}
	var subtotal float64
	for _, it := range cart.Items {
		line := float64(it.Qty) * it.Price
		subtotal += line

		resp.Items = append(resp.Items, dto.CartItem{
			ProductID:   it.ProductID,
			ProductName: it.ProductName,
			ImageURL:    it.ImageURL,
			Qty:         it.Qty,
			Price:       it.Price,
			LineTotal:   line,
		})
	}

	resp.Count = len(cart.Items)
	resp.SubTotal = subtotal
	return resp, nil
}

func NewCartService(cartRepository repository.CartRepository) CartService {
	return &cartService{
		cartRepository: cartRepository,
	}
}
