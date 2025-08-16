package dtos

type GenerateOTPRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ValidateOTPRequest struct {
	Email string `json:"email" validate:"required,email"`
	OTP   string `json:"otp" validate:"required,len=6"`
}
