package http

import (
	"CaloriesCalculator/pkg/myerrors"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

var (
	errInvalidRequestBody    = errors.New("invalid request body")
	errInvalidHeader         = errors.New("invalid header")
	errUnauthorized          = errors.New("unauthorized")
	errMissingAccessToken    = errors.New("missing access token")
	errInvalidOrExpiredToken = errors.New("invalid or expired token")
)

type errorResponse struct {
	Errors []string `json:"errors"`
}

func ErrorResp(w http.ResponseWriter, err error, code int, logger *slog.Logger) {
	w.WriteHeader(code)
	logger.Debug(err.Error())
	json.NewEncoder(w).Encode(errorResponse{Errors: myerrors.ExtractWrapped(err)})
}
