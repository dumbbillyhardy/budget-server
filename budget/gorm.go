package budget

import (
	o "github.com/dumbbillyhardy/budget-server/objects"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Gorm struct{}

// GetByID the budget.
func (budget_gorm *Gorm) GetByID(id string) (*o.Budget, error) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var budget o.Budget
	db.First(&budget, id)
	return &budget, nil
}

// Save the budget.
func (budget_gorm *Gorm) Save(budget o.Budget) (*o.Budget, error) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Create
	db.Create(&budget)
	return &budget, nil
}

// Delete the budget.
func (budget_gorm *Gorm) Delete(id string) error {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		return err
	}

	var budget o.Budget
	budget.ID = id

	// Delete - delete product
	db.Delete(&budget)
	return nil
}
