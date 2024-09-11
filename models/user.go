package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User представляет пользователя в системе.
type User struct {
	ID         int       `json:"id"`
	Username   string    `json:"username" validate:"required,min=3,max=50,alphanum"`
	Password   string    `json:"password" validate:"required,min=8"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

// HashPassword
// Хеширует пароль и устанавливает его в поле Password.
// Возвращает ошибку, если произошла ошибка при хешировании.
func (u *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword
// Проверяет введённый пароль, сравнивая его с хешированным паролем.
// Возвращает ошибку, если пароли не совпадают.
func (u *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
