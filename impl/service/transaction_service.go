package service

import (
	"Finance/api/repository"
	"Finance/entity"
	"Finance/payload"
	"fmt"
	"sync"
)

type transactionServiceImpl struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionServiceImpl(transactionRepository repository.TransactionRepository) *transactionServiceImpl {
	return &transactionServiceImpl{
		transactionRepository: transactionRepository,
	}
}

func (s *transactionServiceImpl) AddTransaction(request *payload.AddTransactionRequest) (*payload.AddTransactionResponse, error) {
	wg := sync.WaitGroup{}
	errCh := make(chan error, 1)
	transactionCh := make(chan *entity.Transaction, 1)

	wg.Add(1)
	go func() {
		defer wg.Done()

		transaction, err := s.transactionRepository.AddTransaction(&entity.Transaction{
			ContractNumber:    request.ContractNumber,
			CustomerID:        request.CustomerID,
			OTR:               request.OTR,
			AdminFee:          request.AdminFee,
			InstallmentAmount: request.InstallmentAmount,
			InterestAmount:    request.InterestAmount,
			AssetName:         request.AssetName,
		})
		if err != nil {
			errCh <- err
			return
		}
		transactionCh <- transaction
	}()

	go func() {
		wg.Wait()
		close(errCh)
		close(transactionCh)
	}()

	var errors []error
	for err := range errCh {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return nil, fmt.Errorf("Errors while adding transactions: %v", errors)
	}

	trx := make([]*entity.Transaction, 0, 1)
	for transaction := range transactionCh {
		trx = append(trx, transaction)
	}

	response := payload.AddTransactionResponse{
		Code:    200,
		Message: "Success",
		Data: payload.AddTransactionRequest{
			ContractNumber:    trx[0].ContractNumber,
			CustomerID:        trx[0].CustomerID,
			OTR:               trx[0].OTR,
			AdminFee:          trx[0].AdminFee,
			InstallmentAmount: trx[0].InstallmentAmount,
			InterestAmount:    trx[0].InterestAmount,
			AssetName:         trx[0].AssetName,
		},
	}

	return &response, nil
}
