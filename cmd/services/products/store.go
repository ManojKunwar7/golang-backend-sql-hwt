package products

import (
	"database/sql"
	"test-project/types"
)

type ProductStore struct {
	db *sql.DB
}

func NewProductStore(db *sql.DB) *ProductStore {
	return &ProductStore{db: db}
}

func (s *ProductStore) GetProducts() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	var products []types.Product = make([]types.Product, 0)
	for rows.Next() {
		product, err := ScanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *product)
	}
	return products, nil
}

func ScanRowIntoProduct(row *sql.Rows) (*types.Product, error) {
	product := new(types.Product)
	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return product, nil
}

//! 1:10:29
