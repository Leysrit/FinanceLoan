package entity

type Transaction struct {
	ContractNumber    int     `json:"contract_number"`
	CustomerID        int     `json:"customer_id"`
	OTR               float64 `json:"otr"`
	AdminFee          float64 `json:"admin_fee"`
	InstallmentAmount float64 `json:"installment_amount"`
	InterestAmount    float64 `json:"interest_amount"`
	AssetName         string  `json:"asset_name"`
}
