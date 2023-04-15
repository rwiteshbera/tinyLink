package models

// User json request
type UserLogin struct {
	Email string `json:"email"`
}

// User model in MongoDB
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type OTP struct {
	Otp string `json:"otp"`
}
