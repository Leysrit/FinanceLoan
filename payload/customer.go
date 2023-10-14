package payload

type CustomerRequest struct {
	CustomerID   int     `json:"customer_id"`
	NIK          string  `json:"nik"`
	FullName     string  `json:"full_name"`
	LegalName    string  `json:"legal_name"`
	PlaceOfBirth string  `json:"place_of_birth"`
	DateOfBirth  string  `json:"date_of_birth"`
	Salary       float64 `json:"salary"`
	KTPImage     string  `json:"ktp_image"`
	SelfieImage  string  `json:"selfie_image"`
}

type Customer struct {
	CustomerID int    `json:"customer_id"`
	NIK        string `json:"nik"`
	FullName   string `json:"full_name"`
	LegalName  string `json:"legal_name"`
}

type CustomerResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    Customer `json:"data"`
}

type ListCustomerRequest struct {
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	Nama  string `json:"nama"`
}

type ListCustomerResponse struct {
	Data         []Customer   `json:"data"`
	PaginateInfo PaginateInfo `json:"paginate_info"`
}

type UpdateCustomer struct {
	CustomerID   int     `json:"customer_id"`
	NIK          string  `json:"nik"`
	FullName     string  `json:"full_name"`
	LegalName    string  `json:"legal_name"`
	PlaceOfBirth string  `json:"place_of_birth"`
	DateOfBirth  string  `json:"date_of_birth"`
	Salary       float64 `json:"salary"`
	KTPImage     string  `json:"ktp_image"`
	SelfieImage  string  `json:"selfie_image"`
}
