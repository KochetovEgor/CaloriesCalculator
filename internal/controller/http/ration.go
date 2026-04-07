package http

import (
	"CaloriesCalculator/internal/controller/http/models"
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/pkg/mylog"
	"encoding/json"
	"errors"
	"net/http"
)

// POST /rations
func (a *App) RationsPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := mylog.FromContext(ctx)

	if r.Header.Get("Content-Type") != "application/json" {
		ErrorResp(w, errInvalidHeader, http.StatusUnsupportedMediaType, logger)
		return
	}

	req := models.RationWithProducts{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorResp(w, errInvalidRequestBody, http.StatusBadRequest, logger)
		return
	}
	user := getUserFromContext(ctx)

	productsEaten := make([]domain.ProductEaten, len(req.Products))
	for i, p := range req.Products {
		productsEaten[i] = models.ProductEatenToDomain(p)
	}

	ration, err := a.service.AddRation(ctx, user, req.Date, productsEaten)
	if err != nil {
		var statusCode int
		if errors.Is(err, domain.ErrInternal) {
			statusCode = http.StatusInternalServerError
		} else {
			statusCode = http.StatusBadRequest
		}
		ErrorResp(w, err, statusCode, logger)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(models.Ration(ration))
}

type rationDeleteRequest struct {
	Date string `json:"date"`
}

// DELETE /rations
func (a *App) RationsDelete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := mylog.FromContext(ctx)

	if r.Header.Get("Content-Type") != "application/json" {
		ErrorResp(w, errInvalidHeader, http.StatusUnsupportedMediaType, logger)
		return
	}

	req := rationDeleteRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorResp(w, errInvalidRequestBody, http.StatusBadRequest, logger)
		return
	}
	user := getUserFromContext(ctx)

	if err := a.service.DeleteRation(ctx, user, req.Date); err != nil {
		var statusCode int
		if errors.Is(err, domain.ErrInternal) {
			statusCode = http.StatusInternalServerError
		} else {
			statusCode = http.StatusBadRequest
		}
		ErrorResp(w, err, statusCode, logger)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// PUT /rations
func (a *App) RationsPut(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := mylog.FromContext(ctx)

	if r.Header.Get("Content-Type") != "application/json" {
		ErrorResp(w, errInvalidHeader, http.StatusUnsupportedMediaType, logger)
		return
	}

	req := models.RationWithProducts{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorResp(w, errInvalidRequestBody, http.StatusBadRequest, logger)
		return
	}
	user := getUserFromContext(ctx)

	productsEaten := make([]domain.ProductEaten, len(req.Products))
	for i, p := range req.Products {
		productsEaten[i] = models.ProductEatenToDomain(p)
	}

	ration, err := a.service.UpdateRation(ctx, user, req.Date, productsEaten)
	if err != nil {
		var statusCode int
		if errors.Is(err, domain.ErrInternal) {
			statusCode = http.StatusInternalServerError
		} else {
			statusCode = http.StatusBadRequest
		}
		ErrorResp(w, err, statusCode, logger)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(models.Ration(ration))
}

// GET /rations
func (a *App) RationsGet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := mylog.FromContext(ctx)

	user := getUserFromContext(ctx)

	rations, err := a.service.SelectRationsByUser(ctx, user)
	if err != nil {
		var statusCode int
		if errors.Is(err, domain.ErrInternal) {
			statusCode = http.StatusInternalServerError
		} else {
			statusCode = http.StatusBadRequest
		}
		ErrorResp(w, err, statusCode, logger)
		return
	}

	resp := make([]models.Ration, len(rations))
	for i, r := range rations {
		resp[i] = models.RationToModel(r)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(resp)
}
