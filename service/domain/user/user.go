package user

import (
	"errors"

	"github.com/AramLab/booking-service/models"
	"github.com/AramLab/booking-service/storage"
)

// UserService представляет собой слой бизнес-логики для управления пользователями.
// Он взаимодействует с UserRepository для выполнения операций над сущностями пользователя.
type UserService struct {
	UserRepository storage.UserRepository
}

// NewUserService создает и возвращает новый экземпляр UserService с предоставленным репозиторием UserRepository.
func NewUserService(UserRepository storage.UserRepository) *UserService {
	return &UserService{UserRepository: UserRepository}
}

// Create создает нового пользователя, вызывая метод Save репозитория.
// Если происходит ошибка во время сохранения пользователя, возвращает соответствующую ошибку.
func (us *UserService) Create(user *models.User) error {
	err := us.UserRepository.Save(user)
	if err != nil {
		return errors.New(`{"error":"error of creating user"}`)
	}
	return nil
}

// Delete удаляет пользователя по указанному идентификатору.
// Сначала проверяет, существует ли пользователь, затем вызывает метод DeleteById репозитория.
// Если пользователь не найден или возникает ошибка при удалении, возвращает соответствующую ошибку.
func (us *UserService) Delete(id string) error {
	if id == "" {
		return errors.New(`{"error":"id is not specified"}`)
	}
	_, err := us.Get(id)
	if err != nil {
		return errors.New(`{"error":"user is not found"}`)
	}

	err = us.UserRepository.DeleteById(id)
	if err != nil {
		return errors.New(`{"error":"error to delete user"}`)
	}
	return nil
}

// Get возвращает пользователя по указанному идентификатору, используя метод FindById репозитория.
// В случае ошибки возвращает пустого пользователя и ошибку.
func (us *UserService) Get(id string) (models.User, error) {
	user, err := us.UserRepository.FindById(id)
	if err != nil {
		return models.User{}, errors.New(`{"error":"error to get user"}`)
	}
	return user, nil
}

// GetAll возвращает всех существующих пользователей, используя метод FindAll репозитория.
// В случае ошибки возвращает пустой список пользователей и соответствующую ошибку.
func (us *UserService) GetAll() ([]models.User, error) {
	var users []models.User
	users, err := us.UserRepository.FindAll()
	if err != nil {
		return []models.User{}, errors.New(`{"error":"error get all users"}`)
	}
	return users, nil
}
