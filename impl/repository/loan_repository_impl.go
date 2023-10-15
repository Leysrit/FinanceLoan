package repository

import (
	"Finance/entity"
	"database/sql"
	"fmt"
)

type loanRepositoryImpl struct {
	db *sql.DB
}

func NewLimitLoanRepositoryImpl(db *sql.DB) *loanRepositoryImpl {
	return &loanRepositoryImpl{
		db: db,
	}
}

func (r *loanRepositoryImpl) GetLimitLoan(CustomerID int) ([]*entity.Loan, error) {
	query := `SELECT LimitID, TenorMonths, LimitAmount FROM limit_loan WHERE CustomerID = ?`

	rows, err := r.db.Query(query, CustomerID)
	if err != nil {
		return nil, err
	}

	loans := make([]*entity.Loan, 0)
	for rows.Next() {
		loan := &entity.Loan{}
		err := rows.Scan(&loan.LimitID, &loan.TenorMonths, &loan.LimitAmount)
		if err != nil {
			return nil, err
		}
		loans = append(loans, loan)
	}

	return loans, nil

}

func (r *loanRepositoryImpl) GetLimitLoanByID(CustomerID int) (*entity.Loan, error) {
	getLoan, err := r.GetLimitLoan(CustomerID)
	if err != nil {
		return nil, err
	}

	if len(getLoan) == 0 {
		return nil, fmt.Errorf("Data Not Found")
	}

	loanLimitID := 0
	for _, loan := range getLoan {
		loanLimitID = loan.LimitID
	}

	query := `SELECT LimitID, TenorMonths, LimitAmount FROM limit_loan WHERE LimitID = ?`

	rows, err := r.db.Query(query, loanLimitID)
	if err != nil {
		return nil, err
	}

	loan := &entity.Loan{}
	for rows.Next() {
		err := rows.Scan(&loan.LimitID, &loan.TenorMonths, &loan.LimitAmount)
		if err != nil {
			return nil, err
		}
	}

	return loan, nil
}

func (r *loanRepositoryImpl) UpdateLimitLoan(limit *entity.Loan) (*entity.Loan, error) {
	query := `INSERT INTO limit_loan (LimitID, TenorMonths, LimitAmount, CustomerID) VALUES (?, ?, ?, ?)`

	_, err := r.db.Exec(query, limit.LimitID, limit.TenorMonths, limit.LimitAmount, limit.CustomerID)
	if err != nil {
		return nil, err
	}

	updatedLoan, err := r.GetLimitLoanByID(limit.CustomerID)

	return updatedLoan, err

}

func (r *loanRepositoryImpl) InsertNewLoan(loan *entity.Loan) (*entity.Loan, error) {
	query := `INSERT INTO limit_loan (LimitID, TenorMonths, LimitAmount, CustomerID) VALUES (?, ?, ?, ?)`

	result, err := r.db.Exec(query, loan.LimitID, loan.TenorMonths, loan.LimitAmount, loan.CustomerID)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	loan.LimitID = int(id)

	return loan, nil
}
