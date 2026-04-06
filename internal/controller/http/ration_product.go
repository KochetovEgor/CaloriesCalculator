package http

import (
	"CaloriesCalculator/internal/controller/http/models"
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/pkg/mylog"
	"encoding/json"
	"errors"
	"net/http"
)

func (a *App) RationsProductsPatch(w http.ResponseWriter, r *http.Request) {
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

	ration, err := a.service.AddProductsToRation(ctx, user, req.Date, productsEaten)
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

	json.NewEncoder(w).Encode(models.RationToModel(ration))
}
