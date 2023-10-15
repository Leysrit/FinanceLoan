package entity

type (
	Loan struct {
		LimitID     int     `json:"limit_id"`
		CustomerID  int     `json:"customer_id"`
		TenorMonths int     `json:"tenor_months"`
		LimitAmount float64 `json:"limit_amount"`
	}
)
