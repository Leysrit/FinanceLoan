package service

import "Finance/payload"

type TransactionService interface {
	AddTransaction(request *payload.AddTransactionRequest) (*payload.AddTransactionResponse, error)
}
