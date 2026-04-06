package http

import (
	"CaloriesCalculator/internal/controller/http/models"
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/pkg/mylog"
	"encoding/json"
	"errors"
	"net/http"
)

func (a *App) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := mylog.FromContext(ctx)

	username, password, ok := r.BasicAuth()
	if !ok {
		ErrorResp(w, errUnauthorized, http.StatusUnauthorized, logger)
		return
	}

	token, err := a.service.AuthUser(ctx, username, password)
	if err != nil {
		var statusCcode int
		if errors.Is(err, domain.ErrInternal) {
			statusCcode = http.StatusInternalServerError
		} else {
			statusCcode = http.StatusUnauthorized
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		}
		ErrorResp(w, err, statusCcode, logger)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(models.AccessToken{AccessToken: token})
}

func (a *App) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := mylog.FromContext(ctx)

	userReq := &models.UserPWD{}
	err := json.NewDecoder(r.Body).Decode(userReq)
	if err != nil {
		ErrorResp(w, errInvalidRequestBody, http.StatusBadRequest, logger)
		return
	}

	user, err := a.service.RegisterUser(ctx, userReq.Username, userReq.Password)
	if err != nil {
		var statusCode int
		if errors.Is(err, domain.ErrInternal) {
			statusCode = http.StatusInternalServerError
		} else {
			statusCode = http.StatusUnauthorized
		}
		ErrorResp(w, err, statusCode, logger)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(models.User{Username: user.Username})
}
