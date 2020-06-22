package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	context "github.com/dumbbillyhardy/budget-server/context"
	o "github.com/dumbbillyhardy/budget-server/objects"
	"github.com/gorilla/mux"
)

// GetBudgetHandlerFactory creates a handler that gets budgets
func GetBudgetHandlerFactory(ctx context.Context) Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		budget, err := ctx.BudgetService.GetByID(params["id"])
		if err != nil {
			return err
		}
		bytes, err := json.Marshal(budget)
		if err != nil {
			return err
		}

		return ServeContent(w, r, string(bytes), http.StatusOK)
	}
}

// SaveBudgetHandlerFactory creates a handler that saves the budget
func SaveBudgetHandlerFactory(ctx context.Context) Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("Content-Type", "application/json")
		var budget o.Budget
		_ = json.NewDecoder(r.Body).Decode(&budget)
		budget.ID = strconv.Itoa(rand.Intn(1000000))
		b, err := ctx.BudgetService.Save(budget)
		if err != nil {
			return err
		}

		return ServeContent(w, r, b.ID, http.StatusOK)
	}
}

// UpdateBudgetHandlerFactory creates a handler that updates the budget
func UpdateBudgetHandlerFactory(ctx context.Context) Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("Content-Type", "application/json")
		var budget o.Budget
		_ = json.NewDecoder(r.Body).Decode(&budget)
		b, err := ctx.BudgetService.Update(budget)
		if err != nil {
			return err
		}

		return ServeContent(w, r, b.ID, http.StatusOK)
	}
}

// DeleteBudgetHandlerFactory creates a handler that updates the budget
func DeleteBudgetHandlerFactory(ctx context.Context) Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		params := mux.Vars(r)
		err := ctx.BudgetService.Delete(params["id"])
		if err != nil {
			return err
		}

		return ServeContent(w, r, "Success", http.StatusOK)
	}
}

// GetAllBudgetsHandlerFactory creates a handler that gets budgets
func GetAllBudgetsHandlerFactory(ctx context.Context) Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("Content-Type", "application/json")
		budgets, err := ctx.BudgetService.GetAll()
		if err != nil {
			return err
		}
		bytes, err := json.Marshal(budgets)
		if err != nil {
			return err
		}

		return ServeContent(w, r, string(bytes), http.StatusOK)
	}
}
