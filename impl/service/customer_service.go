package service

import (
	"Finance/api/repository"
	"Finance/entity"
	"Finance/payload"
	"Finance/utility"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

type customerServiceImpl struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerServiceImpl(customerRepository repository.CustomerRepository) *customerServiceImpl {
	return &customerServiceImpl{
		customerRepository: customerRepository,
	}
}

func (s *customerServiceImpl) Login(request payload.LoginRequest) (*payload.LoginResponse, error) {
	customer, err := s.customerRepository.Login(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	claims := payload.Claims{
		Username: customer.FullName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(3 * 60 * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	response := payload.LoginResponse{
		Username: customer.FullName,
		Token:    tokenString,
	}

	return &response, nil

}

func (s *customerServiceImpl) RegisterCustomer(request payload.CustomerRequest) (*payload.CustomerResponse, error) {
	customer, err := s.customerRepository.RegisterCustomer(&entity.Customer{
		NIK:          request.NIK,
		FullName:     request.FullName,
		LegalName:    request.LegalName,
		PlaceOfBirth: request.PlaceOfBirth,
		DateOfBirth:  request.DateOfBirth,
		Salary:       request.Salary,
		KTPImage:     request.KTPImage,
		SelfieImage:  request.SelfieImage,
	})
	if err != nil {
		return nil, err
	}

	// claims := Claims{
	// 	Email: siswa.Email,
	// 	Role:  siswa.KategoriUser,
	// 	StandardClaims: jwt.StandardClaims{
	// 		ExpiresAt: time.Now().Add(3 * 60 * time.Minute).Unix(),
	// 	},
	// }

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// tokenSting, err := token.SignedString(secretKey)
	// if err != nil {
	// 	return nil, err
	// }

	response := payload.CustomerResponse{
		Code:    200,
		Message: "Regitered Succesfully",
		Data: payload.Customer{
			CustomerID: customer.CustomerID,
			NIK:        customer.NIK,
			FullName:   customer.FullName,
			LegalName:  customer.LegalName,
		},
	}

	return &response, nil
}

func (s *customerServiceImpl) GetAllCustomer(request payload.ListCustomerRequest) (*payload.ListCustomerResponse, error) {
	log.Printf("Getting all customer")
	customerTotal, err := s.customerRepository.GetTotalCustomer(request.Nama)
	if err != nil {
		return nil, err
	}

	nextPage, prevPage, totalPage := utility.GetPaginateURL("/api/customer", &request.Page, &request.Limit, customerTotal)

	listCustomer, err := s.customerRepository.GetAllCustomer(request.Page, request.Limit, request.Nama)
	if err != nil {
		return nil, err
	}

	lenListCustomer := len(listCustomer)
	if lenListCustomer == 0 {
		return nil, fmt.Errorf("Data Not Found")
	}

	result := make([]payload.Customer, 0)
	for i := 0; i < lenListCustomer; i++ {
		customer := listCustomer[i]
		result = append(result, payload.Customer{
			CustomerID: customer.CustomerID,
			NIK:        customer.NIK,
			FullName:   customer.FullName,
			LegalName:  customer.LegalName,
		})
	}

	response := payload.ListCustomerResponse{
		Data: result,
		PaginateInfo: payload.PaginateInfo{
			NextPage:   nextPage,
			PrevPage:   prevPage,
			TotalPages: totalPage,
		},
	}

	return &response, nil
}

func (s *customerServiceImpl) UpdateCustomer(request payload.UpdateCustomer, customerID int) (*payload.CustomerResponse, error) {
	customer, err := s.customerRepository.UpdateCustomer(&entity.Customer{
		CustomerID:   request.CustomerID,
		NIK:          request.NIK,
		FullName:     request.FullName,
		LegalName:    request.LegalName,
		PlaceOfBirth: request.PlaceOfBirth,
		DateOfBirth:  request.DateOfBirth,
		Salary:       request.Salary,
		KTPImage:     request.KTPImage,
		SelfieImage:  request.SelfieImage,
	}, customerID)
	if err != nil {
		return nil, err
	}

	response := payload.CustomerResponse{
		Code:    200,
		Message: "Updated Succesfully",
		Data: payload.Customer{
			CustomerID: customer.CustomerID,
			NIK:        customer.NIK,
			FullName:   customer.FullName,
			LegalName:  customer.LegalName,
		},
	}

	return &response, nil

}
