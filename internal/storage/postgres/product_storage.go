package postgres

import (
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/pkg/mylog"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductStorage struct {
	pool *pgxpool.Pool
}

func NewProductStorage(pool *pgxpool.Pool) *ProductStorage {
	return &ProductStorage{pool: pool}
}

func (s *ProductStorage) Close() error {
	if s.pool != nil {
		s.pool.Close()
	}
	return nil
}

const tableProductsName = "products"

const createTableProducts = `
CREATE TABLE IF NOT EXISTS products (
	id SERIAL PRIMARY KEY,
	user_id INT NOT NULL,
	name TEXT NOT NULL,
	base_weight NUMERIC NOT NULL,
	base_portion NUMERIC NOT NULL DEFAULT 0.00,
	fats NUMERIC NOT NULL,
	proteins NUMERIC NOT NULL,
	carbohydrates NUMERIC NOT NULL,
	CONSTRAINT foreign_key_user_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
	CONSTRAINT unique_food UNIQUE (user_id, name)
);
`

func (s *ProductStorage) Init(ctx context.Context) error {
	attrs := []any{
		"table", tableProductsName,
	}
	logger := mylog.FromContext(ctx).With(attrs...)

	_, err := s.pool.Exec(ctx, createTableProducts)
	if err != nil {
		err = mylog.WrapError(err, attrs...)
		return fmt.Errorf("error creating table: %w", err)
	}

	logger.Debug("table created")
	return nil
}

const addProductToProducts = `
INSERT INTO products (user_id, name, base_weight, base_portion, fats, proteins, carbohydrates)
	SELECT id, $2, $3, $4, $5, $6, $7 FROM users
		WHERE username = $1;
`

func (s *ProductStorage) Add(ctx context.Context, product domain.Product) error {
	attrs := []any{
		"table", tableProductsName,
		"product", product,
	}
	logger := mylog.FromContext(ctx).With(attrs...)

	ct, err := s.pool.Exec(ctx, addProductToProducts,
		product.Username, product.Name, product.BaseWeight,
		product.BasePortion, product.Fats, product.Proteins, product.Carbohydrates)
	if err != nil {
		if isUniqueViolation(err) {
			return domain.ErrProductAlreadyExists
		}
		err = mylog.WrapError(err, attrs...)
		return fmt.Errorf("error adding product to table: %w", err)
	}
	if isNoAffectedRows(ct) {
		return domain.ErrUserNotExists
	}

	logger.Debug("product added")
	return nil
}

const deleteProductFromProducts = `
DELETE FROM products
WHERE user_id = (SELECT id FROM users WHERE username = $1) AND name = $2;
`

func (s *ProductStorage) Delete(ctx context.Context, username, productName string) error {
	attrs := []any{
		"table", tableProductsName,
		"product name", productName,
		"username", username,
	}
	logger := mylog.FromContext(ctx).With(attrs...)

	_, err := s.pool.Exec(ctx, deleteProductFromProducts, username, productName)
	if err != nil {
		err = mylog.WrapError(err, attrs...)
		return fmt.Errorf("error deleting product from table: %w", err)
	}

	logger.Debug("product deleted")
	return nil
}

const updateProductFromProducts = `
UPDATE products SET
	base_weight = $3, base_portion = $4,
	fats = $5, proteins = $6, carbohydrates = $7
WHERE 
	user_id = (SELECT id FROM users WHERE username = $1) AND name = $2;
`

func (s *ProductStorage) Update(ctx context.Context, product domain.Product) error {
	attrs := []any{
		"table", tableProductsName,
	}
	logger := mylog.FromContext(ctx).With(attrs...)

	ct, err := s.pool.Exec(ctx, updateProductFromProducts,
		product.Username, product.Name, product.BaseWeight, product.BasePortion,
		product.Fats, product.Proteins, product.Carbohydrates)
	if err != nil {
		err = mylog.WrapError(err, attrs...)
		return fmt.Errorf("error updating product in table: %w", err)
	}
	if isNoAffectedRows(ct) {
		return domain.ErrProductNotExists
	}

	logger.Debug("product updated")
	return nil
}

const selectProductFromProducts = `
SELECT 
	username, name, base_weight, base_portion, fats, proteins, carbohydrates
FROM 
	products
	JOIN users ON user_id = users.id
WHERE 
	username = $1;
`

func (s *ProductStorage) SelectByUser(ctx context.Context,
	username string) ([]domain.Product, error) {
	attrs := []any{
		"table", tableProductsName,
		"username", username,
	}
	logger := mylog.FromContext(ctx).With(attrs...)

	rows, _ := s.pool.Query(ctx, selectProductFromProducts, username)
	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Product])
	if err != nil {
		err = mylog.WrapError(err, attrs...)
		return nil, fmt.Errorf("error selecting products from table: %w", err)
	}

	logger.Debug("selected products from table")
	return products, nil
}
