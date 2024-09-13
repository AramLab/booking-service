package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/AramLab/booking-service/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Create(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserService) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserService) GetAll() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *MockUserService) Get(id string) (*models.User, error) {
	args := m.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}

func TestCreateUser(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewUserHandlers(mockService)

	user := models.User{
		ID:         1,
		Username:   "mayor12",
		Password:   "123456Mayor12",
		Created_at: time.Now().UTC(),
		Updated_at: time.Now().UTC(),
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	hashedUser := user
	hashedUser.Password = string(hashedPassword)

	body, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("error marshaling the user: %v", err)
	}
	req := httptest.NewRequest("POST", "/user", bytes.NewReader(body))
	w := httptest.NewRecorder()

	mockService.On("Create", mock.MatchedBy(func(u *models.User) bool {
		err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte("123456Mayor12"))
		return err == nil && u.Username == hashedUser.Username
	})).Return(nil)

	handler.CreateUser(w, req)

	res := w.Result()
	defer res.Body.Close()
	require.Equal(t, http.StatusCreated, res.StatusCode)

	var createdUser models.User
	err = json.NewDecoder(res.Body).Decode(&createdUser)
	require.NoError(t, err)

	require.Equal(t, user.ID, createdUser.ID)
	require.Empty(t, createdUser.Password)
}

func TestCreateUser_InvalidInput(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewUserHandlers(mockService)

	user := models.User{
		ID:       1,
		Username: "mayor12",
	}

	body, err := json.Marshal(user)
	if err != nil {
		t.Errorf("error of creating the user")
	}
	req := httptest.NewRequest("POST", "/user", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler.CreateUser(w, req)

	res := w.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestCreateUser_CreateUserError(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewUserHandlers(mockService)

	newUser := models.User{
		ID:         12,
		Username:   "peta",
		Password:   "password123",
		Created_at: time.Now().UTC(),
		Updated_at: time.Now().UTC(),
	}

	mockService.On("Create", mock.MatchedBy(func(u *models.User) bool {
		return u.Username == newUser.Username
	})).Return(errors.New("database error"))

	body, _ := json.Marshal(newUser)

	req := httptest.NewRequest("POST", "/user", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	handler.CreateUser(w, req)

	res := w.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusInternalServerError, res.StatusCode)
}

func TestDeleteUser_Success(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewUserHandlers(mockService)

	user_id := "1"

	userID, err := strconv.Atoi(user_id)
	if err != nil {
		t.Errorf("error of converting string to integer")
	}
	mockService.On("Get", user_id).Return(&models.User{ID: userID}, nil)

	mockService.On("Delete", user_id).Return(nil)

	req := httptest.NewRequest("DELETE", "/user/"+user_id, nil)
	req = mux.SetURLVars(req, map[string]string{"id": user_id})
	w := httptest.NewRecorder()

	handler.DeleteUser(w, req)

	res := w.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusOK, res.StatusCode)
}

func TestDeleteUser_UserNotFound(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewUserHandlers(mockService)

	userID := "1"

	// Мокаем метод Get, чтобы вернуть (nil, error)
	mockService.On("Get", userID).Return((*models.User)(nil), errors.New("user not found"))

	// Создаем запрос DELETE с userID
	req := httptest.NewRequest("DELETE", "/user/"+userID, nil)
	req = mux.SetURLVars(req, map[string]string{"id": userID})
	w := httptest.NewRecorder()

	// Вызываем DeleteUser
	handler.DeleteUser(w, req)

	// Проверяем, что статус ответа - 404
	res := w.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusNotFound, res.StatusCode)
}

func TestDeleteUser_DeleteUserError(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewUserHandlers(mockService)

	user_id := "1"

	userID, err := strconv.Atoi(user_id)
	if err != nil {
		t.Errorf("error of converting string to integer")
	}
	mockService.On("Get", user_id).Return(&models.User{ID: userID}, nil)

	mockService.On("Delete", user_id).Return(errors.New("deletion error"))

	req := httptest.NewRequest("DELETE", "/user/"+user_id, nil)
	req = mux.SetURLVars(req, map[string]string{"id": user_id})
	w := httptest.NewRecorder()

	handler.DeleteUser(w, req)

	res := w.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusInternalServerError, res.StatusCode)
}
