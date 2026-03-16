package http

import (
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/pkg/mylog"
	"encoding/json"
	"errors"
	"net/http"
)

type productAddRequest struct {
	Name          string  `json:"name"`
	BaseWeight    float64 `json:"base_weight"`
	BasePortion   float64 `json:"base_portion"`
	Fats          float64 `json:"fats"`
	Proteins      float64 `json:"proteins"`
	Carbohydrates float64 `json:"carbohydrates"`
}

func (a *App) ProductAdd(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := mylog.FromContext(ctx)

	if r.Header.Get("Content-Type") != "application/json" {
		ErrorResp(w, errInvalidHeader, http.StatusUnsupportedMediaType, logger)
		return
	}

	req := &productAddRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		ErrorResp(w, errInvalidRequestBody, http.StatusBadRequest, logger)
		return
	}

	user := getUserFromContext(ctx)
	product, err := a.service.AddProduct(ctx, domain.Product{
		Username:      user.Username,
		Name:          req.Name,
		BaseWeight:    req.BaseWeight,
		BasePortion:   req.BasePortion,
		Fats:          req.Fats,
		Proteins:      req.Proteins,
		Carbohydrates: req.Carbohydrates,
	})
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

	json.NewEncoder(w).Encode(productAddRequest{
		Name:          product.Name,
		BaseWeight:    product.BaseWeight,
		BasePortion:   product.BasePortion,
		Fats:          product.Fats,
		Proteins:      product.Proteins,
		Carbohydrates: product.Carbohydrates,
	})
}
