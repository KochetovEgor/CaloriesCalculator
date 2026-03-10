package http

import (
	"CaloriesCalculator/pkg/mylog"
	"log/slog"
	"net/http"
)

type statusCodeWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *statusCodeWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		scw := &statusCodeWriter{ResponseWriter: w}

		logger := slog.Default().With(
			"method", r.Method,
			"url", r.URL.String(),
			"ip", r.RemoteAddr,
		)

		ctx := mylog.NewContext(r.Context(), logger)

		logger.Info("request received")
		next.ServeHTTP(scw, r.WithContext(ctx))
		logger.Info("request finished", "status code", scw.statusCode)
	})
}
