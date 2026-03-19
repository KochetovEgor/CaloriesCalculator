package domain

import "log/slog"

type User struct {
	Id           int
	Username     string
	HashPassword []byte
}

func (u User) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Int("id", u.Id),
		slog.String("username", u.Username),
	)
}

type Product struct {
	Username      string
	Name          string
	BaseWeight    float64
	BasePortion   float64
	Fats          float64
	Proteins      float64
	Carbohydrates float64
}

func (p Product) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Any("username", p.Username),
		slog.String("name", p.Name),
	)
}
