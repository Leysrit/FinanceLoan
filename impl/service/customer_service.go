package service

import (
	"Finance/api/repository"
	"Finance/entity"
	"Finance/payload"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte("Final Project Beasiswa")

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

type customerServiceImpl struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerServiceImpl(customerRepository repository.CustomerRepository) *customerServiceImpl {
	return &customerServiceImpl{
		customerRepository: customerRepository,
	}
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

	response := &payload.CustomerResponse{
		Code:    200,
		Message: "Regitered Succesfully",
		Data:    map[string]interface{}{"customer": customer},
	}

	return response, nil
}
