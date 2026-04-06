package http

import (
	"CaloriesCalculator/internal/controller/http/models"
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/pkg/mylog"
	"encoding/json"
	"errors"
	"net/http"
)

func (a *App) ProductsPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := mylog.FromContext(ctx)

	if r.Header.Get("Content-Type") != "application/json" {
		ErrorResp(w, errInvalidHeader, http.StatusUnsupportedMediaType, logger)
		return
	}

	req := models.Product{}
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

	json.NewEncoder(w).Encode(models.Product(product))
}

type productDeleteRequest struct {
	Name string `json:"name"`
}

func (a *App) ProductsDelete(w http.ResponseWriter, r *http.Request) {
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

func (a *App) ProductsPut(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := mylog.FromContext(ctx)

	if r.Header.Get("Content-Type") != "application/json" {
		ErrorResp(w, errInvalidHeader, http.StatusUnsupportedMediaType, logger)
		return
	}

	req := models.Product{}
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

	json.NewEncoder(w).Encode(models.Product(product))
}

func (a *App) ProductsGet(w http.ResponseWriter, r *http.Request) {
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

	resp := make([]models.Product, len(products))
	for i, p := range products {
		resp[i] = models.ProductToModel(p)
	}

	json.NewEncoder(w).Encode(resp)
}
