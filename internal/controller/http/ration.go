package http

import (
	"CaloriesCalculator/internal/domain"
	"CaloriesCalculator/pkg/mylog"
	"encoding/json"
	"errors"
	"net/http"
)

type rationAddRequest struct {
	Date     string `json:"date"`
	Products []struct {
		Name    string  `json:"name"`
		Weight  float64 `json:"weight"`
		Portion float64 `json:"portion"`
	} `json:"products"`
}

type rationAddResponse struct {
	Date          string  `json:"date"`
	Calories      float64 `json:"calories"`
	Fats          float64 `json:"fats"`
	Proteins      float64 `json:"proteins"`
	Carbohydrates float64 `json:"carbohydrates"`
}

func (a *App) RationAdd(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := mylog.FromContext(ctx)

	if r.Header.Get("Content-Type") != "application/json" {
		ErrorResp(w, errInvalidHeader, http.StatusUnsupportedMediaType, logger)
		return
	}

	req := rationAddRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorResp(w, errInvalidRequestBody, http.StatusBadRequest, logger)
		return
	}
	user := getUserFromContext(ctx)

	productsEaten := make([]domain.ProductEaten, len(req.Products))
	for i, p := range req.Products {
		productsEaten[i] = domain.ProductEaten{
			Name:    p.Name,
			Weight:  p.Weight,
			Portion: p.Portion,
		}
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

	json.NewEncoder(w).Encode(rationAddResponse(ration))
}

type rationDeleteRequest struct {
	Date string `json:"date"`
}

func (a *App) RationDelete(w http.ResponseWriter, r *http.Request) {
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

type rationUpdateRequest struct {
	Date     string `json:"date"`
	Products []struct {
		Name    string  `json:"name"`
		Weight  float64 `json:"weight"`
		Portion float64 `json:"portion"`
	} `json:"products"`
}

type rationUpdateResponse struct {
	Date          string  `json:"date"`
	Calories      float64 `json:"calories"`
	Fats          float64 `json:"fats"`
	Proteins      float64 `json:"proteins"`
	Carbohydrates float64 `json:"carbohydrates"`
}

func (a *App) RationUpdate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := mylog.FromContext(ctx)

	if r.Header.Get("Content-Type") != "application/json" {
		ErrorResp(w, errInvalidHeader, http.StatusUnsupportedMediaType, logger)
		return
	}

	req := rationUpdateRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorResp(w, errInvalidRequestBody, http.StatusBadRequest, logger)
		return
	}
	user := getUserFromContext(ctx)

	productsEaten := make([]domain.ProductEaten, len(req.Products))
	for i, p := range req.Products {
		productsEaten[i] = domain.ProductEaten{
			Name:    p.Name,
			Weight:  p.Weight,
			Portion: p.Portion,
		}
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

	json.NewEncoder(w).Encode(rationUpdateResponse(ration))
}
