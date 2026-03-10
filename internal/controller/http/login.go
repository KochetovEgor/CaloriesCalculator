package http

import (
	"CaloriesCalculator/internal/domain"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
)

func (a *App) Login(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")

	const basicAuthPrefix = "Basic "
	if authHeader == "" || !strings.HasPrefix(authHeader, basicAuthPrefix) {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	payload, err := base64.StdEncoding.DecodeString(authHeader[len(basicAuthPrefix):])
	if err != nil {
		http.Error(w, "Invalid Authorization header", http.StatusBadRequest)
		return
	}

	creds := strings.SplitN(string(payload), ":", 2)
	if len(creds) != 2 {
		http.Error(w, "Invalid Authorization header", http.StatusBadRequest)
		return
	}

	token, err := a.service.AuthUser(r.Context(), creds[0], creds[1])
	if err != nil {
		if err, ok := domain.ExtractErr(err); ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	type response struct {
		AccessToken string `json:"access_token"`
	}

	json.NewEncoder(w).Encode(response{AccessToken: token})
}
