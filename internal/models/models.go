package models

type SuccessResponce struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Status  int               `json:"status"`           // HTTP status code
	Message string            `json:"message"`          // A brief message explaining the error
	Error   string            `json:"error,omitempty"`  // Detailed error message (optional)
	Fields  map[string]string `json:"fields,omitempty"` // Field-specific errors (optional)
}

type CheckUser struct {
	Username    string `json:"username"`
	PhoneNumber string `json:"phone_number"`
}

type UserRegisterRequst struct {
	Username    string `json:"username"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type UserRegisterResponce struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserProfile struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
}

type UpdateUserProfile struct {
	Username     string `json:"username"`
	FullName     string `json:"full_name"`
	PhoneNumber  string `json:"phone_number"`
	PasswordHash string `json:"password_hash"`
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
