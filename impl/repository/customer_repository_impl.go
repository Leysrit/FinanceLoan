package repository

import (
	"Finance/entity"
	"database/sql"
	"fmt"
	"strings"
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

func (s *customerRepositoryImpl) GetTotalCustomer(nama string) (int, error) {
	count := 0
	query := `SELECT COUNT(*) FROM Customers`

	if len(strings.Trim(nama, " ")) != 0 {
		query = fmt.Sprintf(`SELECT COUNT(*) FROM Customers WHERE FullName LIKE '%s'`, nama)
	}
	row := s.db.QueryRow(query)
	if err := row.Scan(
		&count,
	); err != nil {
		return -1, err
	}
	fmt.Println("count is", count)
	return count, nil

}

func (s *customerRepositoryImpl) GetAllCustomer(page int, limit int, nama string) ([]*entity.Customer, error) {
	offset := (page - 1) * limit

	query := `
	SELECT CustomerID, NIK, FullName, LegalName, PlaceOfBirth, DateOfBirth, Salary
	FROM Customers
	WHERE FullName LIKE ?
	LIMIT ? OFFSET ?`

	if len(strings.Trim(nama, " ")) != 0 {
		query = fmt.Sprintf(`SELECT CustomerID, NIK, FullName, LegalName, PlaceOfBirth, DateOfBirth, Salary
	FROM Customers
	WHERE FullName LIKE '%s'
	LIMIT ? OFFSET ?`, nama)
	}

	rows, err := s.db.Query(query, "%"+nama+"%", limit, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	listCustomer := make([]*entity.Customer, 0)

	for rows.Next() {
		customer := &entity.Customer{}
		err := rows.Scan(
			&customer.CustomerID,
			&customer.NIK,
			&customer.FullName,
			&customer.LegalName,
			&customer.PlaceOfBirth,
			&customer.DateOfBirth,
			&customer.Salary,
		)
		if err != nil {
			return listCustomer, err
		}
		listCustomer = append(listCustomer, customer)
	}

	return listCustomer, nil

}

func (s *customerRepositoryImpl) GetCustomerByID(id int) (*entity.Customer, error) {
	query := `
	SELECT CustomerID, NIK, FullName, LegalName, PlaceOfBirth, DateOfBirth, Salary
	FROM Customers
	WHERE CustomerID = ?`

	row := s.db.QueryRow(query, id)
	customer := &entity.Customer{}
	err := row.Scan(
		&customer.CustomerID,
		&customer.NIK,
		&customer.FullName,
		&customer.LegalName,
		&customer.PlaceOfBirth,
		&customer.DateOfBirth,
		&customer.Salary,
	)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *customerRepositoryImpl) UpdateCustomer(customer *entity.Customer, customerID int) (*entity.Customer, error) {
	query := `
	UPDATE Customers
	SET
		CustomerID = ?,
		NIK = ?,
		FullName = ?,
		LegalName = ?,
		PlaceOfBirth = ?,
		DateOfBirth = ?,
		Salary = ?,
		KTPImage = ?,
		SelfieImage = ?
	WHERE CustomerID = ?`

	_, err := s.db.Exec(
		query,
		customer.CustomerID,
		customer.NIK,
		customer.FullName,
		customer.LegalName,
		customer.PlaceOfBirth,
		customer.DateOfBirth,
		customer.Salary,
		customer.KTPImage,
		customer.SelfieImage,
		customerID,
	)
	if err != nil {
		return nil, err
	}

	updatedCustomer, err := s.GetCustomerByID(customerID)
	if err != nil {
		return nil, err
	}

	return updatedCustomer, nil

}
