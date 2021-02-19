package auth

import (
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dsuhinin/suhinin-backend-1/core/errors"
)

// TokenModel represents `token` model.
type TokenModel struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	Token     string    `db:"token"`
	CreatedAt time.Time `db:"created_at"`
}

// NewTokenModel creates new instance of Token model.
func NewTokenModel(user *UserModel, token string) *TokenModel {
	return &TokenModel{
		UserID:    user.ID,
		Token:     token,
		CreatedAt: time.Now(),
	}
}

// UserModel represents `user` model.
type UserModel struct {
	ID        int       `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}

// NewUserModel creates new instance of User model.
func NewUserModel(email, password string) (*UserModel, error) {
	user := UserModel{
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
	}

	if err := user.hashPassword(); err != nil {
		return nil, err
	}

	return &user, nil
}

// hashPassword generates password hash and stores it internally.
func (m *UserModel) hashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.WithMessage(err, "impossible to generate password hash")
	}
	m.Password = string(hash)
	return nil
}

// IsPasswordValid makes password validation.
func (m *UserModel) IsPasswordValid(password string) bool {

	if err := bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(password)); err != nil {
		return false
	}
	return true
}
