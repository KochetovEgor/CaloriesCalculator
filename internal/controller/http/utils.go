package http

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func errorWithLog(w http.ResponseWriter, err string, code int, logger *slog.Logger) {
	logger.Debug(err, "status code", code)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(errorResponse{Error: err})
}
