package account

import (
	o "github.com/dumbbillyhardy/budget-server/objects"
)

// InMemory version of Service.
type InMemory struct {
	Accounts []o.Account
}

// GetByID the Account.
func (in_mem *InMemory) GetByID(id string) (*o.Account, error) {
	for _, Account := range in_mem.Accounts {
		if Account.ID == id {
			return &Account, nil
		}
	}
	return nil, &o.NotFoundError{Message: "Not Found"}
}

// Save the Account.
func (in_mem *InMemory) Save(Account o.Account) (*o.Account, error) {
	in_mem.Accounts = append(in_mem.Accounts, Account)
	return &Account, nil
}

// Update the Account.
func (in_mem *InMemory) Update(Account o.Account) (*o.Account, error) {
	index := -1
	for i, b := range in_mem.Accounts {
		if b.ID == Account.ID {
			index = i
		}
	}
	if index == -1 {
		return nil, &o.NotFoundError{Message: "Not found"}
	}
	in_mem.Accounts[index] = Account
	return &Account, nil
}

// Delete by id.
func (in_mem *InMemory) Delete(id string) error {
	index := -1
	for i, b := range in_mem.Accounts {
		if b.ID == id {
			index = i
		}
	}
	if index == -1 {
		return &o.NotFoundError{Message: "Not found"}
	}
	end := len(in_mem.Accounts) - 1
	in_mem.Accounts[index] = in_mem.Accounts[end]
	in_mem.Accounts = in_mem.Accounts[0 : end-1]
	return nil
}

// GetAll the Accounts.
func (in_mem *InMemory) GetAll() ([]o.Account, error) {
	return in_mem.Accounts, nil
}

// SaveAll the Accounts.
func (in_mem *InMemory) SaveAll(Accounts []o.Account) ([]o.Account, error) {
	ret := make([]o.Account, 0, len(Accounts))
	for _, b := range Accounts {
		newAccount, err := in_mem.Save(b)
		if err != nil {
			return nil, err
		}
		ret = append(ret, *newAccount)
	}
	return ret, nil
}

// UpdateAll the Accounts.
func (in_mem *InMemory) UpdateAll(Accounts []o.Account) ([]o.Account, error) {
	ret := make([]o.Account, 0, len(Accounts))
	for _, b := range Accounts {
		newAccount, err := in_mem.Update(b)
		if err != nil {
			return nil, err
		}
		ret = append(ret, *newAccount)
	}
	return ret, nil
}

// DeleteAll the Accounts.
func (in_mem *InMemory) DeleteAll() error {
	in_mem.Accounts = make([]o.Account, 0, 100)
	return nil
}
