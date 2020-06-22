package transaction

import o "github.com/dumbbillyhardy/budget-server/objects"

// Service to get your Accounts.
type Service interface {
	GetByID(string) (*o.Transaction, error)
	Save(o.Transaction) (*o.Transaction, error)
	Update(o.Transaction) (*o.Transaction, error)
	Delete(string) error

	GetAll() ([]o.Transaction, error)
	SaveAll([]o.Transaction) ([]o.Transaction, error)
	UpdateAll([]o.Transaction) ([]o.Transaction, error)
	DeleteAll() error

	GetByBudgetID(string) ([]o.Transaction, error)
	SaveForBudget(budgetID string, transaction o.Transaction) (*o.Transaction, error)
}
