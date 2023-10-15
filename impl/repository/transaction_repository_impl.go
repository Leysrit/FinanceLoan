package repository

import (
	"Finance/entity"
	"database/sql"
	"sync"
)

type transactionRepositoryImpl struct {
	db *sql.DB
}

func NewTransactionRepositoryImpl(db *sql.DB) *transactionRepositoryImpl {
	return &transactionRepositoryImpl{
		db: db,
	}
}

func (r *transactionRepositoryImpl) AddTransaction(transaction *entity.Transaction) (*entity.Transaction, error) {
	var dbMutex sync.Mutex

	dbMutex.Lock()
	defer dbMutex.Unlock()

	trx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	query := `INSERT INTO transactions (ContractNumber, CustomerID, OTR, AdminFee, InstallmentAmount, InterestAmount, AssetName) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err = trx.Exec(query, transaction.ContractNumber, transaction.CustomerID, transaction.OTR, transaction.AdminFee, transaction.InstallmentAmount, transaction.InterestAmount, transaction.AssetName)
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	err = trx.Commit()
	if err != nil {
		trx.Rollback()
		return nil, err
	}

	return transaction, nil

	// return transaction, nil
}
