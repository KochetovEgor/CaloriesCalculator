package domain

import "log/slog"

type User struct {
	Username     string
	HashPassword []byte
}

func (u User) LogValue() slog.Value {
	return slog.StringValue(u.Username)
}

type Product struct {
	Username     string
	Name         string
	BaseWeight   float64
	BasePortion  float64
	Fat          float64
	Protein      float64
	Carbohydrate float64
}

func (p Product) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Any("username", p.Username),
		slog.String("name", p.Name),
	)
}
