package repository

import (
	"context"
	"github.com/saleh-ghazimoradi/MicroEcoBay/product_service/internal/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func setupTestDB(t *testing.T) (*gorm.DB, *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	err = db.AutoMigrate(&domain.Category{}, &domain.Product{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	return db, db
}

func TestCatalogRepository_CreateCategory(t *testing.T) {
	dbWrite, dbRead := setupTestDB(t)
	repo := NewCatalogRepository(dbWrite, dbRead)
	ctx := context.Background()

	category := &domain.Category{
		Name:        "Electronics",
		Description: "Electronic Product",
		ImageURL:    "electronics.jpg",
		Status:      "active",
	}
	if err := repo.CreateCategory(ctx, category); err != nil {
		t.Errorf("CreateCategory failed: %v", err)
	}

	var createdCategory domain.Category
	if err := dbRead.First(&createdCategory, category.Id).Error; err != nil {
		t.Errorf("CreateCategory failed: %v", err)
	}

	if createdCategory.Name != category.Name {
		t.Errorf("CreateCategory failed: %v %v", createdCategory.Name, category.Name)
	}
}

func TestCatalogRepository_GetAllCategories(t *testing.T) {
	dbWrite, dbRead := setupTestDB(t)
	repo := NewCatalogRepository(dbWrite, dbRead)
	ctx := context.Background()

	categories := []*domain.Category{
		{Name: "Books", Description: "Book Category", Status: "active"},
		{Name: "Clothing", Description: "Clothing Category", Status: "active"},
	}

	for _, cat := range categories {
		if err := repo.CreateCategory(ctx, cat); err != nil {
			t.Errorf("CreateCategory failed: %v", err)
		}
	}

	result, err := repo.GetAllCategories(ctx)
	if err != nil {
		t.Errorf("GetAllCategories failed: %v", err)
	}

	if len(result) < 2 {
		t.Errorf("GetAllCategories failed: %v %v", len(result), result)
	}
}

func TestCatalogRepository_GetCategoryById(t *testing.T) {
	dbWrite, dbRead := setupTestDB(t)
	repo := NewCatalogRepository(dbWrite, dbRead)
	ctx := context.Background()

	category := &domain.Category{
		Name:        "Toys",
		Description: "Toys Category",
		Status:      "active",
	}

	if err := repo.CreateCategory(ctx, category); err != nil {
		t.Errorf("CreateCategory failed: %v", err)
	}

	result, err := repo.GetCategoryById(ctx, uint(category.Id))
	if err != nil {
		t.Errorf("GetCategoryById failed: %v", err)
	}
	if result.Name != category.Name {
		t.Errorf("GetCategoryById failed: %v %v", result.Name, category.Name)
	}
}

func TestCatalogRepository_UpdateCategory(t *testing.T) {
	dbWrite, dbRead := setupTestDB(t)
	repo := NewCatalogRepository(dbWrite, dbRead)
	ctx := context.Background()
	category := &domain.Category{
		Name:        "Sports",
		Description: "Sports equipment",
		Status:      "active",
	}
	if err := repo.CreateCategory(ctx, category); err != nil {
		t.Fatalf("Setup failed: %v", err)
	}

	update := &domain.Category{
		Name:        "Sports Updated",
		Description: "Updated sports equipment",
		Status:      "inactive",
	}
	err := repo.UpdateCategory(ctx, uint(category.Id), update)
	if err != nil {
		t.Errorf("UpdateCategory failed: %v", err)
	}

	updated, err := repo.GetCategoryById(ctx, uint(category.Id))
	if err != nil {
		t.Errorf("Failed to get updated category: %v", err)
	}
	if updated.Name != update.Name {
		t.Errorf("Expected name %s, got %s", update.Name, updated.Name)
	}
}

func TestCatalogRepository_CreateProduct(t *testing.T) {
	dbWrite, dbRead := setupTestDB(t)
	repo := NewCatalogRepository(dbWrite, dbRead)
	ctx := context.Background()

	category := &domain.Category{Name: "Tech", Status: "active"}
	if err := repo.CreateCategory(ctx, category); err != nil {
		t.Fatalf("Setup failed: %v", err)
	}

	product := &domain.Product{
		Name:        "Laptop",
		Description: "High-end laptop",
		CategoryId:  uint(category.Id),
		Price:       999.99,
		Stock:       10,
		ImageURL:    "laptop.jpg",
		Status:      "active",
	}

	if err := repo.CreateProduct(ctx, product); err != nil {
		t.Errorf("CreateProduct failed: %v", err)
	}

	var createdProduct domain.Product
	if err := dbRead.First(&createdProduct, product.Id).Error; err != nil {
		t.Errorf("CreateProduct failed: %v", err)
	}
	if createdProduct.Name != product.Name {
		t.Errorf("CreateProduct failed: %v %v", createdProduct.Name, product.Name)
	}
}

func TestCatalogRepository_GetProductById(t *testing.T) {
	dbWrite, dbRead := setupTestDB(t)
	repo := NewCatalogRepository(dbWrite, dbRead)
	ctx := context.Background()

	category := &domain.Category{Name: "Appliances", Status: "active"}
	if err := repo.CreateCategory(ctx, category); err != nil {
		t.Fatalf("Setup failed: %v", err)
	}

	product := &domain.Product{
		Name:       "Fridge",
		CategoryId: uint(category.Id),
		Price:      799.99,
		Stock:      5,
		Status:     "active",
	}
	if err := repo.CreateProduct(ctx, product); err != nil {
		t.Fatalf("Setup failed: %v", err)
	}

	result, err := repo.GetProductById(ctx, uint(product.Id))
	if err != nil {
		t.Errorf("GetProductById failed: %v", err)
	}
	if result.Name != product.Name {
		t.Errorf("Expected name %s, got %s", product.Name, result.Name)
	}

}

func TestCatalogRepository_GetAllProducts(t *testing.T) {
	dbWrite, dbRead := setupTestDB(t)
	repo := NewCatalogRepository(dbWrite, dbRead)

	ctx := context.Background()

	category := &domain.Category{Name: "Gadgets", Status: "active"}
	if err := repo.CreateCategory(ctx, category); err != nil {
		t.Fatalf("Setup failed: %v", err)
	}

	products := []*domain.Product{
		{Name: "Phone", CategoryId: uint(category.Id), Price: 599.99, Stock: 20, Status: "active"},
		{Name: "Tablet", CategoryId: uint(category.Id), Price: 399.99, Stock: 15, Status: "active"},
	}
	for _, prod := range products {
		if err := repo.CreateProduct(ctx, prod); err != nil {
			t.Fatalf("Setup failed: %v", err)
		}
	}

	result, err := repo.GetAllProducts(ctx)
	if err != nil {
		t.Errorf("GetAllProducts failed: %v", err)
	}
	if len(result) < 2 {
		t.Errorf("Expected at least 2 products, got %d", len(result))
	}
}

func TestCatalogRepository_UpdateProduct(t *testing.T) {
	dbWrite, dbRead := setupTestDB(t)
	repo := NewCatalogRepository(dbWrite, dbRead)
	ctx := context.Background()

	category := &domain.Category{Name: "Furniture", Status: "active"}
	if err := repo.CreateCategory(ctx, category); err != nil {
		t.Fatalf("Setup failed: %v", err)
	}

	product := &domain.Product{
		Name:       "Chair",
		CategoryId: uint(category.Id),
		Price:      99.99,
		Stock:      25,
		Status:     "active",
	}
	if err := repo.CreateProduct(ctx, product); err != nil {
		t.Fatalf("Setup failed: %v", err)
	}

	update := &domain.Product{
		Name:        "Chair Updated",
		Description: "Updated comfy chair",
		Price:       129.99,
		Stock:       20,
		Status:      "inactive",
	}
	err := repo.UpdateProduct(ctx, uint(product.Id), update)
	if err != nil {
		t.Errorf("UpdateProduct failed: %v", err)
	}

	updated, err := repo.GetProductById(ctx, uint(product.Id))
	if err != nil {
		t.Errorf("Failed to get updated product: %v", err)
	}
	if updated.Name != update.Name {
		t.Errorf("Expected name %s, got %s", update.Name, updated.Name)
	}
	if !reflect.DeepEqual(updated.Price, update.Price) {
		t.Errorf("Expected price %v, got %v", update.Price, updated.Price)
	}
}
