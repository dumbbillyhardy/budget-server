package budget

import (
	o "github.com/dumbbillyhardy/budget-server/objects"
)

// Service to get your budgets.
type Service interface {
	GetByID(id string) (*o.Budget, error)
	Save(budget o.Budget) (*o.Budget, error)
	Update(budget o.Budget) (*o.Budget, error)
	Delete(id string) error

	GetAll() ([]o.Budget, error)
	SaveAll(budgets []o.Budget) ([]o.Budget, error)
	UpdateAll(budgets []o.Budget) ([]o.Budget, error)
	DeleteAll() error
}
