package postgres

import (
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/pkg/mylog"
	"context"
	"fmt"
)

const tableProductsEatenName = "products_eaten"

const createTableProductsEaten = `
CREATE TABLE IF NOT EXISTS products_eaten (
	id SERIAL PRIMARY KEY,
	ration_id INT NOT NULL,
	product_id INT,
	name TEXT NOT NULL,
	weight NUMERIC NOT NULL,
	portion NUMERIC NOT NULL DEFAULT 0.00,
	calories NUMERIC NOT NULL,
	fats NUMERIC NOT NULL,
	proteins NUMERIC NOT NULL,
	carbohydrates NUMERIC NOT NULL,
	CONSTRAINT unique_product_eaten UNIQUE (ration_id, product_id),
	CONSTRAINT foreign_key_ration_id FOREIGN KEY (ration_id) REFERENCES rations(id) ON DELETE CASCADE,
	CONSTRAINT foreign_key_product_id FOREIGN KEY (product_id) REFERENCES products(id)
);
`

type argSliceProductEaten struct {
	name          []string
	weight        []float64
	portion       []float64
	calories      []float64
	fats          []float64
	proteins      []float64
	carbohydrates []float64
}

func newArgSliceFromProductsEaten(productsEaten []domain.ProductEaten) argSliceProductEaten {
	n := len(productsEaten)
	argSlice := argSliceProductEaten{
		name: make([]string, n), weight: make([]float64, n), portion: make([]float64, n),
		calories: make([]float64, n), fats: make([]float64, n), proteins: make([]float64, n),
		carbohydrates: make([]float64, n),
	}
	for i, p := range productsEaten {
		argSlice.name[i] = p.Name
		argSlice.weight[i] = p.Weight
		argSlice.portion[i] = p.Portion
		argSlice.calories[i] = p.Calories
		argSlice.fats[i] = p.Fats
		argSlice.proteins[i] = p.Proteins
		argSlice.carbohydrates[i] = p.Carbohydrates
	}

	return argSlice
}

const addProductsToProductsEaten = `
INSERT INTO products_eaten (ration_id, product_id, name, weight,
		portion, calories, fats, proteins, carbohydrates)
SELECT 
	$2, products.id, arr.name, arr.weight, arr.portion,
	arr.calories, arr.fats, arr.proteins, arr.carbohydrates
FROM
	UNNEST($3::text[], $4::numeric[], $5::numeric[], $6::numeric[], $7::numeric[],
	$8::numeric[], $9::numeric[]) as arr(name, weight, portion, calories, 
	fats, proteins, carbohydrates)
JOIN
	products ON (products.user_id = $1 AND products.name = arr.name);
`

func (s *RationStorage) AddProductsEaten(ctx context.Context,
	user domain.User, rationId int, productsEaten []domain.ProductEaten) error {
	attrs := []any{
		"table", tableProductsEatenName,
	}
	logger := mylog.FromContext(ctx).With(attrs...)

	argSlice := newArgSliceFromProductsEaten(productsEaten)

	_, err := s.pool.Exec(ctx, addProductsToProductsEaten,
		user.Id, rationId, argSlice.name, argSlice.weight, argSlice.portion,
		argSlice.calories, argSlice.fats, argSlice.proteins, argSlice.carbohydrates)
	if err != nil {
		err = mylog.WrapError(err, attrs...)
		return fmt.Errorf("error adding products eaten to table: %w", err)
	}
	logger.Debug("products eaten added")

	return nil
}
