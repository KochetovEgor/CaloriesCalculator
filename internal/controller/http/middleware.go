package http

import (
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/internal/pkg/auth"
	"CaloriesCalculator/pkg/mylog"
	"context"
	"log/slog"
	"net/http"
	"strings"
)

type statusCodeWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *statusCodeWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

// logMiddleware is for logging incoming requests.
func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		scw := &statusCodeWriter{ResponseWriter: w, statusCode: 200}

		logger := slog.Default().With(
			"method", r.Method,
			"url", r.URL.Path,
			"ip", r.RemoteAddr,
		)

		ctx := mylog.NewContext(r.Context(), logger)

		logger.Info("request received")
		next.ServeHTTP(scw, r.WithContext(ctx))
		logger.Info("request finished", "status code", scw.statusCode)
	})
}

type contextKey string

const userContextKey contextKey = "user"

func putUserToContext(ctx context.Context, user domain.User) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}

func getUserFromContext(ctx context.Context) domain.User {
	if user, ok := ctx.Value(userContextKey).(domain.User); ok {
		return user
	}
	return domain.User{}
}

func bearerAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := mylog.FromContext(r.Context())

		authHeader := r.Header.Get("Authorization")

		const bearerPrefix = "Bearer "
		if authHeader == "" || !strings.HasPrefix(authHeader, bearerPrefix) {
			ErrorResp(w, errMissingAccessToken, http.StatusUnauthorized, logger)
			return
		}

		rawToken := strings.TrimPrefix(authHeader, bearerPrefix)

		user, err := auth.GetUserFromToken(rawToken)
		if err != nil {
			ErrorResp(w, errInvalidOrExpiredToken, http.StatusUnauthorized, logger)
			return
		}

		ctx := putUserToContext(r.Context(), user)
		next(w, r.WithContext(ctx))
	}
}
