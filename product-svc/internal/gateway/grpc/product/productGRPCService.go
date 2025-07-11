package product

import (
	"context"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/service"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/slg"
	"strconv"
)

type productGRPCService struct {
	UnimplementedProductServiceServer
	catalogService service.CatalogService
}

func (p *productGRPCService) GetProductById(ctx context.Context, req *GetProductRequest) (*ProductResponse, error) {

	product, err := p.catalogService.GetProductById(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}

	slg.Logger.Info("GetProductById", "grpc request", req)
	return &ProductResponse{
		Id:         int32(product.ID),
		Name:       product.Name,
		CategoryId: strconv.Itoa(int(product.CategoryId)),
		Price:      product.Price,
		Stock:      int32(product.Stock),
	}, nil
}

func NewProductGRPCService(catalogService service.CatalogService) *productGRPCService {
	return &productGRPCService{
		catalogService: catalogService,
	}
}
