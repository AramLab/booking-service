package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         int       `json:"id"`
	Username   string    `json:"username" validate:"required,min=3,max=50,alphanum"`
	Password   string    `json:"password" validate:"required,min=8"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

// HashPassword
// Хеширует пароль и возвращает ошибку в случае её возникновения или же nil.
func (u *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword
// Проверяет корректность введённого пароля и возвращает ошибку в случае её возникновения или же nil.
func (u *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
