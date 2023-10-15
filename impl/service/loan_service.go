package service

import (
	"Finance/api/repository"
	"Finance/entity"
	"Finance/payload"
	"fmt"
	"log"
)

type limitLoanServiceImpl struct {
	loanRepository repository.LimitLoanRepository
}

func NewLimitLoanServiceImpl(loanRepository repository.LimitLoanRepository) *limitLoanServiceImpl {
	return &limitLoanServiceImpl{
		loanRepository: loanRepository,
	}
}

func (s *limitLoanServiceImpl) InsertNewLoan(request *payload.InsertLimitLoanRequest) (*payload.InsertLimitLoanResponse, error) {
	log.Printf("Inserting new limit loan for CustomerID: %d", request.CustomerID)
	loan, err := s.loanRepository.InsertNewLoan(&entity.Loan{
		CustomerID:  request.CustomerID,
		TenorMonths: request.TenorMonths,
		LimitAmount: request.LimitAmount,
	})

	if err != nil {
		return nil, err
	}

	response := payload.InsertLimitLoanResponse{
		Code:    200,
		Message: "Success",
		Data:    *loan,
	}

	log.Printf("Successfully inserted new limit loan for CustomerID: %d", request.CustomerID)

	return &response, nil
}

func (s *limitLoanServiceImpl) GetLimitLoan(CustomerID int) (*payload.GetLimitLoanResponse, error) {
	log.Printf("Getting limit loan for CustomerID: %d", CustomerID)

	listLoan, err := s.loanRepository.GetLimitLoan(CustomerID)
	if err != nil {
		log.Printf("Failed to get limit loan: %v", err)
		return nil, err
	}

	if len(listLoan) == 0 {
		return nil, fmt.Errorf("Data Not Found")
	}

	result := make([]payload.GetLimitLoan, len(listLoan))
	for i, loan := range listLoan {
		result[i] = payload.GetLimitLoan{
			LimitID:     loan.LimitID,
			TenorMonths: loan.TenorMonths,
			LimitAmount: loan.LimitAmount,
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("Data Not Found")
	}

	response := payload.GetLimitLoanResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	}

	log.Printf("Successfully retrieved limit loan for CustomerID: %d", CustomerID)

	return &response, nil
}

func (s *limitLoanServiceImpl) UpdateLimitLoan(request *payload.UpdateLimitLoanRequest) (*payload.UpdateLimitLoanResponse, error) {
	log.Printf("Updating limit loan for CustomerID: %d", request.CustomerID)
	response := payload.UpdateLimitLoanResponse{}
	loan, err := s.loanRepository.GetLimitLoanByID(request.CustomerID)
	if err != nil {
		if loan == nil || err == fmt.Errorf("Data Not Found") {
			insertedLoan, err := s.loanRepository.InsertNewLoan(&entity.Loan{
				CustomerID:  request.CustomerID,
				TenorMonths: request.TenorMonths,
				LimitAmount: request.LoanAmount,
			})

			if err != nil {
				log.Printf("Failed to insert new limit loan: %v", err)
				return nil, err
			}

			response = payload.UpdateLimitLoanResponse{
				Code:    200,
				Message: "Success",
				Data: payload.GetLimitLoan{
					LimitID:     insertedLoan.LimitID,
					TenorMonths: insertedLoan.TenorMonths,
					LimitAmount: insertedLoan.LimitAmount,
				},
			}

			log.Printf("Successfully inserted new limit loan for CustomerID: %d", request.CustomerID)

			return &response, nil
		} else {
			log.Printf("Failed to get limit loan: %v", err)
			return nil, err
		}
	}

	if loan.CustomerID != request.CustomerID {
		loan.CustomerID = request.CustomerID
	}

	Message := ""
	if request.Action == "pinjam" {
		loan.LimitAmount = loan.LimitAmount - request.LoanAmount
		Message = "Success Pinjam"
	} else if request.Action == "bayar" {
		loan.LimitAmount = loan.LimitAmount + request.LoanAmount
		Message = "Success bayar"
	} else {
		return nil, fmt.Errorf("Invalid Action")
	}

	updatedLimitLoan, err := s.loanRepository.UpdateLimitLoan(&entity.Loan{
		CustomerID:  request.CustomerID,
		TenorMonths: request.TenorMonths,
		LimitAmount: loan.LimitAmount,
	})
	if err != nil {
		log.Printf("Failed to update limit loan: %v", err)
		return nil, err
	}

	response = payload.UpdateLimitLoanResponse{
		Code:    200,
		Message: Message,
		Data: payload.GetLimitLoan{
			LimitID:     updatedLimitLoan.LimitID,
			TenorMonths: updatedLimitLoan.TenorMonths,
			LimitAmount: updatedLimitLoan.LimitAmount,
		},
	}

	log.Printf("Successfully updated limit loan for CustomerID: %d", request.CustomerID)

	return &response, nil
}
