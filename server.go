package main

import (
	"log"
	"net/http"

	"github.com/dumbbillyhardy/budget-server/budget"
	"github.com/dumbbillyhardy/budget-server/context"
	o "github.com/dumbbillyhardy/budget-server/objects"
	"github.com/dumbbillyhardy/budget-server/transaction"

	h "github.com/dumbbillyhardy/budget-server/handlers"
	"github.com/gorilla/mux"
)

func get(r *mux.Router, path string, f h.Handler) {
	r.HandleFunc(path, h.ErrorHandler(f)).Methods("GET")
}

func post(r *mux.Router, path string, f h.Handler) {
	r.HandleFunc(path, h.ErrorHandler(f)).Methods("POST")
}

func put(r *mux.Router, path string, f h.Handler) {
	r.HandleFunc(path, h.ErrorHandler(f)).Methods("PUT")
}

func delete(r *mux.Router, path string, f h.Handler) {
	r.HandleFunc(path, h.ErrorHandler(f)).Methods("DELETE")
}

func main() {
	ctx := context.Context{
		BudgetService:      &budget.InMemory{Budgets: make([]o.Budget, 0, 100)},
		TransactionService: &transaction.InMemory{Transactions: make([]o.Transaction, 0, 100)}}
	r := mux.NewRouter()
	get(r, "/", h.RootHandlerFactory(ctx))

	get(r, "/budget/{id}", h.GetBudgetHandlerFactory(ctx))
	post(r, "/budget", h.SaveBudgetHandlerFactory(ctx))
	put(r, "/budget", h.UpdateBudgetHandlerFactory(ctx))
	delete(r, "/budget/{id}", h.DeleteBudgetHandlerFactory(ctx))

	get(r, "/budget", h.GetAllBudgetsHandlerFactory(ctx))

	get(r, "/budget/{budgetid}/transaction/{id}", h.GetTransactionHandlerFactory(ctx))
	post(r, "/budget/{budgetid}/transaction", h.SaveTransactionForBudgetHandlerFactory(ctx))
	put(r, "/budget/{budgetid}/transaction", h.UpdateTransactionHandlerFactory(ctx))
	delete(r, "/budget/{budgetid}/transaction/{id}", h.DeleteTransactionHandlerFactory(ctx))

	get(r, "/budget/{budgetid}/transaction", h.GetAllTransactionsForBudgetHandlerFactory(ctx))

	log.Fatal(http.ListenAndServe(":8080", r))
}
