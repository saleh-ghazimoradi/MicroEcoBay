package repository

import "errors"

var (
	ErrCatalogCreate = errors.New("error on creating catalog")
	ErrProductCreate = errors.New("error on creating product")
	ErrCatalogGet    = errors.New("error on getting catalog")
	ErrProductGet    = errors.New("error on getting product")
	ErrNotFound      = errors.New("catalog not found")
	ErrUpdateCatalog = errors.New("error on updating catalog")
	ErrProductUpdate = errors.New("error on updating product")
)
