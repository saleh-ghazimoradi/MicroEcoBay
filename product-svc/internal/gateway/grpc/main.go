package main

import (
	"context"
	"fmt"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/gateway/grpc/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := product.NewProductServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.GetProductById(ctx, &product.GetProductRequest{Id: 123})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
