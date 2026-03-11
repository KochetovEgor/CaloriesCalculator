package http

import (
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/pkg/mylog"
	"encoding/json"
	"net/http"
)

func (a *App) Login(w http.ResponseWriter, r *http.Request) {
	logger := mylog.FromContext(r.Context())

	username, password, ok := r.BasicAuth()
	if !ok {
		respErrWithLog(w, "Unauthorized", http.StatusUnauthorized, logger)
	}

	token, err := a.service.AuthUser(r.Context(), username, password)
	if err != nil {
		if err, ok := domain.ExtractErr(err); ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			respErrWithLog(w, err.Error(), http.StatusUnauthorized, logger)
			return
		}
		respErrWithLog(w, "Internal error", http.StatusInternalServerError, logger)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	type response struct {
		AccessToken string `json:"access_token"`
	}

	json.NewEncoder(w).Encode(response{AccessToken: token})
}
