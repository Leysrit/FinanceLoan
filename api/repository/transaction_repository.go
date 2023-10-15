package repository

import "Finance/entity"

type TransactionRepository interface {
	AddTransaction(transaction *entity.Transaction) (*entity.Transaction, error)
}
