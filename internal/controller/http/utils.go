package http

import (
	"log/slog"
	"net/http"
)

func respErrWithLog(w http.ResponseWriter, err string, code int, logger *slog.Logger) {
	logger.Debug(err, "status code", code)
	http.Error(w, err, code)
}
