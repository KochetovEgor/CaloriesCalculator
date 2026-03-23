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
	Calories      float64 `json:"calories"`
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

	req := productAddRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorResp(w, errInvalidRequestBody, http.StatusBadRequest, logger)
		return
	}
	user := getUserFromContext(ctx)

	product, err := a.service.AddProduct(ctx, user, domain.Product(req))
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

	json.NewEncoder(w).Encode(productAddRequest(product))
}

type productDeleteRequest struct {
	Name string `json:"name"`
}

func (a *App) ProductDelete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := mylog.FromContext(ctx)

	if r.Header.Get("Content-Type") != "application/json" {
		ErrorResp(w, errInvalidHeader, http.StatusUnsupportedMediaType, logger)
		return
	}

	req := &productDeleteRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		ErrorResp(w, errInvalidRequestBody, http.StatusBadRequest, logger)
		return
	}
	user := getUserFromContext(ctx)

	if err := a.service.DeleteProduct(ctx, user, req.Name); err != nil {
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

type productUpdateRequest struct {
	Name          string  `json:"name"`
	BaseWeight    float64 `json:"base_weight"`
	BasePortion   float64 `json:"base_portion"`
	Calories      float64 `json:"calories"`
	Fats          float64 `json:"fats"`
	Proteins      float64 `json:"proteins"`
	Carbohydrates float64 `json:"carbohydrates"`
}

func (a *App) ProductUpdate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := mylog.FromContext(ctx)

	if r.Header.Get("Content-Type") != "application/json" {
		ErrorResp(w, errInvalidHeader, http.StatusUnsupportedMediaType, logger)
		return
	}

	req := productUpdateRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorResp(w, errInvalidRequestBody, http.StatusBadRequest, logger)
		return
	}
	user := getUserFromContext(ctx)

	product, err := a.service.UpdateProduct(ctx, user, domain.Product(req))
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

	json.NewEncoder(w).Encode(productUpdateRequest(product))
}

type productResponse struct {
	Name          string  `json:"name"`
	BaseWeight    float64 `json:"base_weight"`
	BasePortion   float64 `json:"base_portion"`
	Calories      float64 `json:"calories"`
	Fats          float64 `json:"fats"`
	Proteins      float64 `json:"proteins"`
	Carbohydrates float64 `json:"carbohydrates"`
}

func (a *App) Product(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := mylog.FromContext(ctx)

	user := getUserFromContext(ctx)

	products, err := a.service.SelectProductsByUser(ctx, user)
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

	resp := make([]productResponse, len(products))
	for i, p := range products {
		resp[i] = productResponse{
			Name:          p.Name,
			BaseWeight:    p.BaseWeight,
			BasePortion:   p.BasePortion,
			Calories:      p.Calories,
			Fats:          p.Fats,
			Proteins:      p.Proteins,
			Carbohydrates: p.Carbohydrates,
		}
	}

	json.NewEncoder(w).Encode(resp)
}
