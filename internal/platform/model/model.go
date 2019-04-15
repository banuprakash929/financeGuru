package model

import (
	"time"
)

//User struct to add or update user
type User struct {
	ID               string
	FirstName        string
	LastName         string
	Age              int
	Contact          string
	Email            string
	CreatedTimestamp time.Time
	Loans            []Loan
}

//Loan struct to describe loan details
type Loan struct {
	UserID                 string
	ID                     string
	LoanType               string
	Principle              float64
	InterestRate           float64
	Payments               []Payment
	LendTimestamp          time.Time
	BalancePrincipleAmount float64
	BalanceInterestAmount  float64
}

//Payment struct to track payments
type Payment struct {
	LoanID        string
	ID            string
	PaymentType   string
	Amount        float64
	PaidTimestamp time.Time
	PaidBy        string
}
