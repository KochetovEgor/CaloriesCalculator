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
	Name          string
	BaseWeight    float64
	BasePortion   float64
	Calories      float64
	Fats          float64
	Proteins      float64
	Carbohydrates float64
}

func (p Product) LogValue() slog.Value {
	return slog.StringValue(p.Name)
}

type ProductEaten struct {
	Name          string
	Weight        float64
	Portion       float64
	Calories      float64
	Fats          float64
	Proteins      float64
	Carbohydrates float64
}

func (pe ProductEaten) LogValue() slog.Value {
	return slog.StringValue(pe.Name)
}

type Ration struct {
	Date          string
	Calories      float64
	Fats          float64
	Proteins      float64
	Carbohydrates float64
}

func (r Ration) LogValue() slog.Value {
	return slog.StringValue(r.Date)
}
