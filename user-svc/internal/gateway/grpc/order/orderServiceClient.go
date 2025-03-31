package order

import (
	"context"
	"github.com/saleh-ghazimoradi/MicroEcoBay/user_service/slg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcOrderServiceClient struct {
	client OrderServiceClient
}

func (g *GrpcOrderServiceClient) GetUserOrders(ctx context.Context, userId int64) (*GetOrderResponse, error) {
	res, err := g.client.GetOrders(ctx, &GetOrderRequest{
		UserId: userId,
	})
	if err != nil {
		slg.Logger.Error("Failed to get orders for user", "userId", userId, "error", err)
		return nil, err
	}
	return res, nil
}

func NewGrpcOrderServiceClient(orderServiceAddress string) (*GrpcOrderServiceClient, error) {
	client, err := grpc.NewClient(orderServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	orderServiceClient := NewOrderServiceClient(client)
	return &GrpcOrderServiceClient{client: orderServiceClient}, nil
}
