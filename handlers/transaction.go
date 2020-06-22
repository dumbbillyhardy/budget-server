package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/dumbbillyhardy/budget-server/context"
	o "github.com/dumbbillyhardy/budget-server/objects"
	"github.com/gorilla/mux"
)

// GetBudget gets the budget from context and request.
func GetBudget(ctx context.Context, r *http.Request) (*o.Budget, error) {
	params := mux.Vars(r)
	return ctx.BudgetService.GetByID(params["budgetid"])
}

// GetTransactionHandlerFactory creates a handler that gets budgets
func GetTransactionHandlerFactory(ctx context.Context) Handler {
	return func(w http.ResponseWriter, r *http.Request) error {

		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		transaction, err := ctx.TransactionService.GetByID(params["id"])
		if err != nil {
			return err
		}
		bytes, err := json.Marshal(transaction)
		if err != nil {
			return err
		}

		return ServeContent(w, r, string(bytes), http.StatusOK)
	}
}

// SaveTransactionHandlerFactory creates a handler that saves the transaction
func SaveTransactionHandlerFactory(ctx context.Context) Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("Content-Type", "application/json")
		var transaction o.Transaction
		_ = json.NewDecoder(r.Body).Decode(&transaction)
		transaction.ID = strconv.Itoa(rand.Intn(1000000))
		b, err := ctx.TransactionService.Save(transaction)
		if err != nil {
			return err
		}

		return ServeContent(w, r, b.ID, http.StatusOK)
	}
}

// UpdateTransactionHandlerFactory creates a handler that updates the transaction
func UpdateTransactionHandlerFactory(ctx context.Context) Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("Content-Type", "application/json")
		var transaction o.Transaction
		_ = json.NewDecoder(r.Body).Decode(&transaction)
		b, err := ctx.TransactionService.Update(transaction)
		if err != nil {
			return err
		}

		return ServeContent(w, r, b.ID, http.StatusOK)
	}
}

// DeleteTransactionHandlerFactory creates a handler that updates the transaction
func DeleteTransactionHandlerFactory(ctx context.Context) Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		params := mux.Vars(r)
		err := ctx.TransactionService.Delete(params["id"])
		if err != nil {
			return err
		}

		return ServeContent(w, r, "Success", http.StatusOK)
	}
}

// GetAllTransactionsHandlerFactory creates a handler that gets budgets
func GetAllTransactionsHandlerFactory(ctx context.Context) Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("Content-Type", "application/json")
		transactions, err := ctx.TransactionService.GetAll()
		if err != nil {
			return err
		}
		bytes, err := json.Marshal(transactions)
		if err != nil {
			return err
		}

		return ServeContent(w, r, string(bytes), http.StatusOK)
	}
}

// GetAllTransactionsForBudgetHandlerFactory creates a handler that gets budgets
func GetAllTransactionsForBudgetHandlerFactory(ctx context.Context) Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		params := mux.Vars(r)
		w.Header().Set("Content-Type", "application/json")
		transactions, err := ctx.TransactionService.GetByBudgetID(params["budgetid"])
		if err != nil {
			return err
		}
		bytes, err := json.Marshal(transactions)
		if err != nil {
			return err
		}

		return ServeContent(w, r, string(bytes), http.StatusOK)
	}
}

// SaveTransactionForBudgetHandlerFactory creates a handler that gets budgets
func SaveTransactionForBudgetHandlerFactory(ctx context.Context) Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		params := mux.Vars(r)
		w.Header().Set("Content-Type", "application/json")
		var transaction o.Transaction
		_ = json.NewDecoder(r.Body).Decode(&transaction)
		transaction.ID = strconv.Itoa(rand.Intn(1000000))
		newTransaction, err := ctx.TransactionService.SaveForBudget(params["budgetid"], transaction)
		if err != nil {
			return err
		}
		bytes, err := json.Marshal(newTransaction)
		if err != nil {
			return err
		}

		return ServeContent(w, r, string(bytes), http.StatusOK)
	}
}
