package payload

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CustomerResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
