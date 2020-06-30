package objects

// Budget you know
type Budget struct {
	ID         string
	Name       string
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

// Category for transactions
type Category struct {
	BudgetID string
	ID       string
	Name     string
	Group    string
}

// Transaction you know
type Transaction struct {
	BudgetID    string
	ID          string
	Date        string
	Payee       string
	Memo        string
	Cleared     bool
	Approved    bool
}

type SubTransaction interface {
	ID() string
	AccountID   string
	Amount() int64
	CategoryID  string
}

type CategoryTransfer struct {
    TransactionID string
	AccountID   string

    FromCategoryID string
}

type AccountTransfer struct {
    TransactionID string
    FromAccountID string
}

type PaymentTransaction struct {
    TransactionID string
    FromCategoryID string
}
