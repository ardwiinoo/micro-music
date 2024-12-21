package request

type RegisterRequestPayload struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
}

type LoginRequestPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}