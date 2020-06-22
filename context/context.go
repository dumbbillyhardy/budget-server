package context

import (
	"github.com/dumbbillyhardy/budget-server/account"
	"github.com/dumbbillyhardy/budget-server/budget"
	"github.com/dumbbillyhardy/budget-server/transaction"
)

// Context to pass around.
type Context struct {
	BudgetService      budget.Service
	AccountService     account.Service
	TransactionService transaction.Service
}
