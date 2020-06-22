package account

import o "github.com/dumbbillyhardy/budget-server/objects"

// Service to get your Accounts.
type Service interface {
	GetByID(string) (*o.Account, error)
	Save(o.Account) (*o.Account, error)
	Update(o.Account) (*o.Account, error)
	Delete(string) error

	GetAll() ([]o.Account, error)
	SaveAll([]o.Account) ([]o.Account, error)
	UpdateAll([]o.Account) ([]o.Account, error)
	DeleteAll() error
}
