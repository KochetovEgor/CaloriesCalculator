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
	calories NUMERIC NOT NULL,
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
INSERT INTO products (user_id, name, base_weight, base_portion,
	calories, fats, proteins, carbohydrates)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
`

func (s *ProductStorage) Add(ctx context.Context,
	user domain.User, product domain.Product) error {
	attrs := []any{
		"table", tableProductsName,
	}
	logger := mylog.FromContext(ctx).With(attrs...)

	ct, err := s.pool.Exec(ctx, addProductToProducts,
		user.Id, product.Name, product.BaseWeight, product.BasePortion,
		product.Calories, product.Fats, product.Proteins, product.Carbohydrates)
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

	logger.Debug("product added in table")
	return nil
}

const deleteProductFromProducts = `
DELETE FROM products
WHERE user_id = $1 AND name = $2;
`

func (s *ProductStorage) Delete(ctx context.Context, user domain.User, productName string) error {
	attrs := []any{
		"table", tableProductsName,
	}
	logger := mylog.FromContext(ctx).With(attrs...)

	ct, err := s.pool.Exec(ctx, deleteProductFromProducts, user.Id, productName)
	if err != nil {
		err = mylog.WrapError(err, attrs...)
		return fmt.Errorf("error deleting product from table: %w", err)
	}
	if isNoAffectedRows(ct) {
		return domain.ErrProductNotExists
	}

	logger.Debug("product deleted from table")
	return nil
}

const updateProductFromProducts = `
UPDATE products SET
	base_weight = $3, base_portion = $4,
	calories = $5, fats = $6, proteins = $7, carbohydrates = $8
WHERE 
	user_id = $1 AND name = $2;
`

func (s *ProductStorage) Update(ctx context.Context,
	user domain.User, product domain.Product) error {
	attrs := []any{
		"table", tableProductsName,
	}
	logger := mylog.FromContext(ctx).With(attrs...)

	ct, err := s.pool.Exec(ctx, updateProductFromProducts,
		user.Id, product.Name, product.BaseWeight, product.BasePortion,
		product.Calories, product.Fats, product.Proteins, product.Carbohydrates)
	if err != nil {
		err = mylog.WrapError(err, attrs...)
		return fmt.Errorf("error updating product in table: %w", err)
	}
	if isNoAffectedRows(ct) {
		return domain.ErrProductNotExists
	}

	logger.Debug("product updated in table")
	return nil
}

const selectProductFromProducts = `
SELECT 
	name, base_weight, base_portion, calories, fats, proteins, carbohydrates
FROM 
	products
WHERE
	user_id = $1;
`

func (s *ProductStorage) SelectByUser(ctx context.Context,
	user domain.User) ([]domain.Product, error) {
	attrs := []any{
		"table", tableProductsName,
	}
	logger := mylog.FromContext(ctx).With(attrs...)

	rows, _ := s.pool.Query(ctx, selectProductFromProducts, user.Id)
	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Product])
	if err != nil {
		err = mylog.WrapError(err, attrs...)
		return nil, fmt.Errorf("error selecting products from table: %w", err)
	}

	logger.Debug("selected products from table")
	return products, nil
}
