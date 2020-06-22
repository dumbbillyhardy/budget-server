package objects

// Budget you know
type Budget struct {
	ID         string
	Name       string
	FirstMonth string
	LastMonth  string
	Accounts   []Account
}

// Account you know
type Account struct {
	BudgetID         string
	ID               string
	Name             string
	Balance          int64
	ClearedBalance   int64
	UnclearedBalance int64
	Closed           bool
}

// CategoryGroup groups categories
type CategoryGroup struct {
	BudgetID   string
	ID         string
	Name       string
	Categories []Category
}

// Category for transactions
type Category struct {
	BudgetID string
	ID       string
	Name     string
	GroupID  string
	Budgeted int64
	Activity int64
	Balance  int64
}

// Transaction you know
type Transaction struct {
	BudgetID     string
	ID           string
	AccountID    string
	AccountName  string
	Date         string
	Amount       int64
	PayeeID      string
	PayeeName    string
	CategoryID   string
	CategoryName string
	Memo         string
	Cleared      bool
	Approved     bool
}
