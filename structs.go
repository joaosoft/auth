package auth

import (
	"github.com/joaosoft/web"

	"time"
)

type ErrorResponse struct {
	Code    web.Status `json:"code,omitempty"`
	Message string     `json:"message,omitempty"`
	Cause   string     `json:"cause,omitempty"`
}

type GetSessionRequest struct {
	Email    string `json:"email" validate:"notzero, error={InvalidBodyParameter}"`
	Password string `json:"password" validate:"notzero, error={InvalidBodyParameter}"`
}

type RefreshSessionRequest struct {
	Authorization string `json:"authorization" validate:"notzero"`
}

type SessionResponse struct {
	TokenType    string `json:"token_type"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type SignUpRequest struct {
	FirstName       string `json:"first_name" db:"first_name" validate:"notzero, error={InvalidBodyParameter}"`
	LastName        string `json:"last_name" db:"last_name" validate:"notzero, error={InvalidBodyParameter}"`
	Email           string `json:"email" db:"email" validate:"notzero, email, error={InvalidBodyParameter}" `
	Password        string `json:"password" validate:"id=password, notzero, error={InvalidBodyParameter}"`
	PasswordConfirm string `json:"password_confirm" validate:"notzero, value={password}, error={InvalidBodyParameter}"`
}

type SignUpResponse struct {
	IdUser string `json:"id_user" db:"id_user"`
}

type ChangeUserStatusRequest struct {
	IdUser string `json:"id_user" db:"id_user" validate:"notzero"`
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
