package service

import "Finance/payload"

type LimitLoanService interface {
	InsertNewLoan(loan *payload.InsertLimitLoanRequest) (*payload.InsertLimitLoanResponse, error)
	GetLimitLoan(CustomerID int) (*payload.GetLimitLoanResponse, error)
	UpdateLimitLoan(limit *payload.UpdateLimitLoanRequest) (*payload.UpdateLimitLoanResponse, error)
}
