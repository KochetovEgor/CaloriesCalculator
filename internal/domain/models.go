package domain

import "log/slog"

type User struct {
	Username     string
	HashPassword []byte
}

func (u User) LogValue() slog.Value {
	return slog.StringValue(u.Username)
}
