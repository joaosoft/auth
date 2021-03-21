package auth

import (
	"github.com/joaosoft/web"

	"time"
)

type getSessionRequest struct {
	Email    string `json:"email" validate:"not-empty, error={{ErrorInvalidBodyParameter}}"`
	Password string `json:"password" validate:"not-empty, error={{ErrorInvalidBodyParameter}}"`
}

type refreshSessionRequest struct {
	Authorization string `json:"authorization" validate:"not-empty"`
}

type updateUserStatusRequest struct {
	IdUser string `json:"id_user" db:"id_user" validate:"not-empty"`
}

type signUpRequest struct {
	FirstName       string `json:"first_name" db:"first_name" validate:"not-empty, error={{ErrorInvalidBodyParameter}}"`
	LastName        string `json:"last_name" db:"last_name" validate:"not-empty, error={{ErrorInvalidBodyParameter}}"`
	Email           string `json:"email" db:"email" validate:"not-empty, email, error={{ErrorInvalidBodyParameter}}" `
	Password        string `json:"password" validate:"id=password, not-empty, error={{ErrorInvalidBodyParameter}}"`
	PasswordConfirm string `json:"password_confirm" validate:"not-empty, value={password}, error={{ErrorInvalidBodyParameter}}"`
}

type ErrorResponse struct {
	Code    web.Status `json:"code,omitempty"`
	Message string     `json:"message,omitempty"`
	Cause   string     `json:"cause,omitempty"`
}

type SessionResponse struct {
	TokenType    string `json:"token_type"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type SignUpResponse struct {
	IdUser string `json:"id_user" db:"id_user"`
}

type User struct {
	IdUser       string    `json:"id_user" db:"id_user"`
	FirstName    string    `json:"first_name" db:"first_name"`
	LastName     string    `json:"last_name" db:"last_name"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"-" db.write:"password_hash"`
	RefreshToken string    `json:"refresh_token" db:"refresh_token"`
	Active       bool      `json:"active" db:"active"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
