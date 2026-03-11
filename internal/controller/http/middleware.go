package http

import (
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/internal/pkg/auth"
	"CaloriesCalculator/pkg/mylog"
	"context"
	"log/slog"
	"net/http"
	"slices"
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

const userContextKey = "user"

func putUserToContext(ctx context.Context, user domain.User) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}

func getUserFromContext(ctx context.Context) (domain.User, bool) {
	user, ok := ctx.Value(userContextKey).(domain.User)
	return user, ok
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if methods, ok := publicPaths[r.URL.Path]; ok {
			if methods == nil || slices.Contains(methods, r.Method) {
				next.ServeHTTP(w, r)
				return
			}
		}

		logger := mylog.FromContext(r.Context())

		authHeader := r.Header.Get("Authorization")

		const bearerPrefix = "Bearer "
		if authHeader == "" || !strings.HasPrefix(authHeader, bearerPrefix) {
			errorWithLog(w, "missing access token", http.StatusUnauthorized, logger)
			return
		}

		rawToken := strings.TrimPrefix(authHeader, bearerPrefix)

		user, err := auth.GetUserFromToken(rawToken)
		if err != nil {
			errorWithLog(w, "Invalid or expired token", http.StatusUnauthorized, logger)
			return
		}

		ctx := putUserToContext(r.Context(), user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
