package budget

import (
	o "github.com/dumbbillyhardy/budget-server/objects"
)

// InMemory version of Service.
type InMemory struct {
	Budgets []o.Budget
}

// GetByID the budget.
func (in_mem *InMemory) GetByID(id string) (*o.Budget, error) {
	for _, budget := range in_mem.Budgets {
		if budget.ID == id {
			return &budget, nil
		}
	}
	return nil, &o.NotFoundError{Message: "Not Found"}
}

// Save the budget.
func (in_mem *InMemory) Save(budget o.Budget) (*o.Budget, error) {
	in_mem.Budgets = append(in_mem.Budgets, budget)
	return &budget, nil
}

// Update the budget.
func (in_mem *InMemory) Update(budget o.Budget) (*o.Budget, error) {
	index := -1
	for i, b := range in_mem.Budgets {
		if b.ID == budget.ID {
			index = i
		}
	}
	if index == -1 {
		return nil, &o.NotFoundError{Message: "Not found"}
	}
	in_mem.Budgets[index] = budget
	return &budget, nil
}

// Delete by id.
func (in_mem *InMemory) Delete(id string) error {
	index := -1
	for i, b := range in_mem.Budgets {
		if b.ID == id {
			index = i
		}
	}
	if index == -1 {
		return &o.NotFoundError{Message: "Not found"}
	}
	end := len(in_mem.Budgets) - 1
	in_mem.Budgets[index] = in_mem.Budgets[end]
	in_mem.Budgets = in_mem.Budgets[0 : end-1]
	return nil
}

// GetAll the budgets.
func (in_mem *InMemory) GetAll() ([]o.Budget, error) {
	return in_mem.Budgets, nil
}

// SaveAll the budgets.
func (in_mem *InMemory) SaveAll(budgets []o.Budget) ([]o.Budget, error) {
	ret := make([]o.Budget, 0, len(budgets))
	for _, b := range budgets {
		newBudget, err := in_mem.Save(b)
		if err != nil {
			return nil, err
		}
		ret = append(ret, *newBudget)
	}
	return ret, nil
}

// UpdateAll the budgets.
func (in_mem *InMemory) UpdateAll(budgets []o.Budget) ([]o.Budget, error) {
	ret := make([]o.Budget, 0, len(budgets))
	for _, b := range budgets {
		newBudget, err := in_mem.Update(b)
		if err != nil {
			return nil, err
		}
		ret = append(ret, *newBudget)
	}
	return ret, nil
}

// DeleteAll the budgets.
func (in_mem *InMemory) DeleteAll() error {
	in_mem.Budgets = make([]o.Budget, 0, 100)
	return nil
}
