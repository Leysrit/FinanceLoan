package repository

import (
	"Finance/entity"
	"database/sql"
)

type customerLogin struct {
}

type customerRepositoryImpl struct {
	db *sql.DB
}

func NewCustomerRepositoryImpl(db *sql.DB) *customerRepositoryImpl {
	return &customerRepositoryImpl{
		db: db,
	}
}

func (s *customerRepositoryImpl) RegisterCustomer(customer *entity.Customer) (*entity.Customer, error) {
	query := `
	INSERT INTO Customers (CustomerID, NIK, FullName, LegalName, PlaceOfBirth, DateOfBirth, Salary, KTPImage, SelfieImage)
	VALUES
		(?, ?, ?,?,?,?,?,?,?)`

	result, err := s.db.Exec(query, customer.CustomerID, customer.NIK, customer.FullName, customer.LegalName, customer.PlaceOfBirth, customer.DateOfBirth, customer.Salary, customer.KTPImage, customer.SelfieImage)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	customer.CustomerID = int(id)

	return customer, nil
}
