package payload

import "Finance/entity"

type GetLimitLoanRequest struct {
	CustomerID int `json:"customer_id"`
}

type GetLimitLoan struct {
	LimitID     int     `json:"limit_id"`
	TenorMonths int     `json:"tenor_months"`
	LimitAmount float64 `json:"limit_amount"`
}

type GetLimitLoanResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    []GetLimitLoan `json:"data"`
}

type UpdateLimitLoanRequest struct {
	CustomerID  int     `json:"customer_id"`
	Action      string  `json:"action"`
	TenorMonths int     `json:"tenor_months"`
	LoanAmount  float64 `json:"loan_amount"`
}

type UpdateLimitLoanResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    GetLimitLoan `json:"data"`
}

type InsertLimitLoanRequest struct {
	CustomerID  int     `json:"customer_id"`
	Action      string  `json:"action"`
	TenorMonths int     `json:"tenor_months"`
	LimitAmount float64 `json:"loan_amount"`
}

type InsertLimitLoanResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    entity.Loan `json:"data"`
}
