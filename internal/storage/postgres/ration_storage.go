package postgres

import (
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/pkg/mylog"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
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
VALUES ($1, $2, ROUND($3, 2), ROUND($4, 2), ROUND($5, 2), ROUND($6, 2))
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

	logger.Debug("ration added in table")
	return id, nil
}

const deleteRationFromRations = `
DELETE FROM rations
WHERE user_id = $1 AND date = $2;
`

func (s *RationStorage) DeleteRation(ctx context.Context, user domain.User, date string) error {
	attrs := []any{
		"table", tableRationName,
	}
	logger := mylog.FromContext(ctx).With(attrs...)

	ct, err := s.pool.Exec(ctx, deleteRationFromRations, user.Id, date)
	if err != nil {
		err = mylog.WrapError(err, attrs...)
		return fmt.Errorf("error deleting ration from table: %w", err)
	}
	if isNoAffectedRows(ct) {
		return domain.ErrRationNotExists
	}

	logger.Debug("ration deleted from table")
	return nil
}

const updateRationFromRations = `
UPDATE rations SET
	calories = ROUND($3 , 2), fats = ROUND($4, 2), 
	proteins = ROUND($5, 2), carbohydrates = ROUND($6, 2)
WHERE
	user_id = $1 AND date = $2
RETURNING id;
`

func (s *RationStorage) UpdateRation(ctx context.Context,
	user domain.User, ration domain.Ration) (int, error) {
	attrs := []any{
		"table", tableRationName,
	}
	logger := mylog.FromContext(ctx).With(attrs...)

	row := s.pool.QueryRow(ctx, updateRationFromRations, user.Id,
		ration.Date, ration.Calories, ration.Fats, ration.Proteins, ration.Carbohydrates)

	var id int
	if err := row.Scan(&id); err != nil {
		if isNoRows(err) {
			return 0, domain.ErrRationNotExists
		}
		err = mylog.WrapError(err, attrs...)
		return 0, fmt.Errorf("error updating ration in table: %w", err)
	}

	logger.Debug("ration updated in table")
	return id, nil
}

const selectRationsByUser = `
SELECT
	date::text, calories, fats, proteins, carbohydrates
FROM
	rations
WHERE
	user_id = $1
ORDER BY date;
`

func (s *RationStorage) SelectRationsByUser(ctx context.Context,
	user domain.User) ([]domain.Ration, error) {
	attrs := []any{
		"table", tableRationName,
	}
	logger := mylog.FromContext(ctx).With(attrs...)

	rows, _ := s.pool.Query(ctx, selectRationsByUser, user.Id)
	rations, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Ration])
	if err != nil {
		err = mylog.WrapError(err, attrs...)
		return nil, fmt.Errorf("error selecting rations from table: %w", err)
	}

	logger.Debug("rations selected from table")
	return rations, nil
}

const addRationToRation = `
UPDATE rations SET
	calories = ROUND(calories + $3, 2), fats = ROUND(fats + $4, 2),
	proteins = ROUND(proteins + $5, 2), carbohydrates = ROUND(carbohydrates + $6, 2)
WHERE
	user_id = $1 AND date = $2
RETURNING id;
`

func (s *RationStorage) AddRationToRation(ctx context.Context,
	user domain.User, ration domain.Ration) (int, error) {
	attrs := []any{
		"table", tableRationName,
	}
	logger := mylog.FromContext(ctx).With(attrs...)

	row := s.pool.QueryRow(ctx, addRationToRation, user.Id, ration.Date,
		ration.Calories, ration.Fats, ration.Proteins, ration.Carbohydrates)

	var id int
	if err := row.Scan(&id); err != nil {
		if isNoRows(err) {
			return 0, domain.ErrRationNotExists
		}
		err = mylog.WrapError(err, attrs...)
		return 0, fmt.Errorf("error updating ration in table: %w", err)
	}

	logger.Debug("ration updated in table")
	return id, nil
}
