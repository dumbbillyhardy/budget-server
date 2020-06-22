package transaction

import (
	"fmt"

	o "github.com/dumbbillyhardy/budget-server/objects"
)

// InMemory version of Service.
type InMemory struct {
	Transactions []o.Transaction
}

// NotFoundError that will be mapped to 404
type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf(e.Message)
}

// GetByID the Transaction.
func (in_mem *InMemory) GetByID(id string) (*o.Transaction, error) {
	for _, Transaction := range in_mem.Transactions {
		if Transaction.ID == id {
			return &Transaction, nil
		}
	}
	return nil, &NotFoundError{"Not Found"}
}

// Save the Transaction.
func (in_mem *InMemory) Save(Transaction o.Transaction) (*o.Transaction, error) {
	in_mem.Transactions = append(in_mem.Transactions, Transaction)
	return &Transaction, nil
}

// Update the Transaction.
func (in_mem *InMemory) Update(Transaction o.Transaction) (*o.Transaction, error) {
	index := -1
	for i, b := range in_mem.Transactions {
		if b.ID == Transaction.ID {
			index = i
		}
	}
	if index == -1 {
		return nil, &NotFoundError{"Not found"}
	}
	in_mem.Transactions[index] = Transaction
	return &Transaction, nil
}

// Delete by id.
func (in_mem *InMemory) Delete(id string) error {
	index := -1
	for i, b := range in_mem.Transactions {
		if b.ID == id {
			index = i
		}
	}
	if index == -1 {
		return &NotFoundError{"Not found"}
	}
	end := len(in_mem.Transactions) - 1
	in_mem.Transactions[index] = in_mem.Transactions[end]
	in_mem.Transactions = in_mem.Transactions[0 : end-1]
	return nil
}

// GetAll the Transactions.
func (in_mem *InMemory) GetAll() ([]o.Transaction, error) {
	return in_mem.Transactions, nil
}

// SaveAll the Transactions.
func (in_mem *InMemory) SaveAll(Transactions []o.Transaction) ([]o.Transaction, error) {
	ret := make([]o.Transaction, 0, len(Transactions))
	for _, b := range Transactions {
		newAccount, err := in_mem.Save(b)
		if err != nil {
			return nil, err
		}
		ret = append(ret, *newAccount)
	}
	return ret, nil
}

// UpdateAll the Transactions.
func (in_mem *InMemory) UpdateAll(Transactions []o.Transaction) ([]o.Transaction, error) {
	ret := make([]o.Transaction, 0, len(Transactions))
	for _, b := range Transactions {
		newAccount, err := in_mem.Update(b)
		if err != nil {
			return nil, err
		}
		ret = append(ret, *newAccount)
	}
	return ret, nil
}

// DeleteAll the Transactions.
func (in_mem *InMemory) DeleteAll() error {
	in_mem.Transactions = make([]o.Transaction, 0, 100)
	return nil
}

// GetByBudgetID gets all with budget id.
func (in_mem *InMemory) GetByBudgetID(budgetID string) ([]o.Transaction, error) {
	ret := make([]o.Transaction, 0, len(in_mem.Transactions))
	for _, transaction := range in_mem.Transactions {
		if transaction.BudgetID == budgetID {
			ret = append(ret, transaction)
		}
	}
	return ret, nil
}

// SaveForBudget gets all with budget id.
func (in_mem *InMemory) SaveForBudget(budgetID string, transaction o.Transaction) (*o.Transaction, error) {
	transaction.BudgetID = budgetID
	in_mem.Transactions = append(in_mem.Transactions, transaction)
	return &transaction, nil
}
