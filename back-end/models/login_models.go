package models

type SignUp struct {
	Name     string `json:"name" db:"name"`
	LastName string `json:"last_name" db:"last_name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	AreaCode string `json:"area_code" db:"area_code"`
	Phone    string `json:"phone" db:"phone"`
}
