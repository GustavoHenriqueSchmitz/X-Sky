package models

type SignUp struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	AreaCode int    `json:"area_code"`
	Phone    int    `json:"phone"`
}
