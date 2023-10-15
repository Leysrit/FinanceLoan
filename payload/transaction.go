package payload

type AddTransactionRequest struct {
	CustomerID        int     `json:"customer_id"`
	OTR               float64 `json:"otr"`
	AdminFee          float64 `json:"admin_fee"`
	InstallmentAmount float64 `json:"installment_amount"`
	InterestAmount    float64 `json:"interest_amount"`
	AssetName         string  `json:"asset_name"`
}

type AddTransactionResponse struct {
	Code    int                   `json:"code"`
	Message string                `json:"message"`
	Data    AddTransactionRequest `json:"data"`
}
