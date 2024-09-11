package user

import (
	"errors"

	"github.com/AramLab/booking-service/models"
	"github.com/AramLab/booking-service/storage"
)

type UserService struct {
	UserRepository storage.UserRepository
}

func NewUserService(UserRepository storage.UserRepository) *UserService {
	return &UserService{UserRepository: UserRepository}
}

func (us *UserService) Create(user *models.User) error {
	err := us.UserRepository.Save(user)
	if err != nil {
		return errors.New(`{"error":"error of creating user"}`)
	}
	return nil
}

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

func (us *UserService) Get(id string) (models.User, error) {
	user, err := us.UserRepository.FindById(id)
	if err != nil {
		return models.User{}, errors.New(`{"error":"error to get user"}`)
	}
	return user, nil
}

func (us *UserService) GetAll() ([]models.User, error) {
	var users []models.User
	users, err := us.UserRepository.FindAll()
	if err != nil {
		return []models.User{}, errors.New(`{"error":"error get all users"}`)
	}
	return users, nil
}
