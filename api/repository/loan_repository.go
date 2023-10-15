package repository

import "Finance/entity"

type LimitLoanRepository interface {
	GetLimitLoan(CustomerID int) ([]*entity.Loan, error)
	GetLimitLoanByID(CustomerID int) (*entity.Loan, error)
	UpdateLimitLoan(limit *entity.Loan) (*entity.Loan, error)
	InsertNewLoan(loan *entity.Loan) (*entity.Loan, error)
}
