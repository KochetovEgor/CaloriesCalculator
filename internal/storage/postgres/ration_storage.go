package postgres

import (
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/pkg/mylog"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RationStorage struct {
	pool *pgxpool.Pool
}

func NewRationStorage(pool *pgxpool.Pool) *RationStorage {
	return &RationStorage{pool: pool}
}

func (s *RationStorage) Close() error {
	if s.pool != nil {
		s.pool.Close()
	}
	return nil
}

const tableRationName = "rations"

const createTableRations = `
CREATE TABLE IF NOT EXISTS rations (
	id SERIAL PRIMARY KEY,
	user_id INT NOT NULL,
	date DATE NOT NULL,
	calories NUMERIC NOT NULL,
	fats NUMERIC NOT NULL,
	proteins NUMERIC NOT NULL,
	carbohydrates NUMERIC NOT NULL,
	CONSTRAINT unique_ration UNIQUE (user_id, date),
	CONSTRAINT foreign_key_user_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
`

func (s *RationStorage) Init(ctx context.Context) error {
	logger := mylog.FromContext(ctx)

	_, err := s.pool.Exec(ctx, createTableRations)
	if err != nil {
		err = mylog.WrapError(err, "table", tableRationName)
		return fmt.Errorf("error creating table: %w", err)
	}
	logger.Debug("table created", "table", tableRationName)

	_, err = s.pool.Exec(ctx, createTableProductsEaten)
	if err != nil {
		err = mylog.WrapError(err, "table", tableProductsEatenName)
		return fmt.Errorf("error creating table: %w", err)
	}
	logger.Debug("table created", "table", tableProductsEatenName)

	return nil
}

const addNewRationToRations = `
INSERT INTO rations (user_id, date, calories, fats, proteins, carbohydrates)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;
`

func (s *RationStorage) AddNewRation(ctx context.Context,
	user domain.User, ration domain.Ration) (int, error) {
	attrs := []any{
		"table", tableRationName,
	}
	logger := mylog.FromContext(ctx).With(attrs...)

	row := s.pool.QueryRow(ctx, addNewRationToRations, user.Id,
		ration.Date, ration.Calories, ration.Fats, ration.Proteins, ration.Carbohydrates)

	var id int
	if err := row.Scan(&id); err != nil {
		if isUniqueViolation(err) {
			return 0, domain.ErrRationAlreadyExists
		}
		if isNoRows(err) {
			return 0, domain.ErrUserNotExists
		}
		err = mylog.WrapError(err, attrs...)
		return 0, fmt.Errorf("error adding ration to table: %w", err)
	}

	logger.Debug("ration added")
	return id, nil
}
