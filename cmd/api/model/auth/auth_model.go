package authModel

import "time"

type CreateUserModel struct {
	Firstname   string    `json:"first_name" binding:"required"`
	Lastname    string    `json:"last_name" binding:"required"`
	PhoneNumber string    `json:"phone_number" binding:"required"`
	Birthdate   time.Time `json:"birthdate" binding:"required"`
	Email       string    `json:"email" binding:"required"`
}

type LoginModel struct {
	Otp       string `json:"otp" binding:"required"`
	Reference string `json:"reference" binding:"required"`
}
