package user

import (
	"encoding/json"
	"net/http"

	"github.com/AramLab/booking-service/models"
	"github.com/AramLab/booking-service/service"
	"github.com/AramLab/booking-service/validation"
	"github.com/gorilla/mux"
)

// UserHandlers содержит ссылки на службы для обработки пользователей.
type UserHandlers struct {
	UserService service.UserService
}

// NewUserHandlers создает новый объект UserHandlers с заданной службой UserService.
func NewUserHandlers(userService service.UserService) *UserHandlers {
	return &UserHandlers{UserService: userService}
}

// CreateUser обрабатывает запрос на создание нового пользователя.
// Он декодирует тело запроса, валидирует данные пользователя, хеширует пароль
// и создает нового пользователя в базе данных. В случае успеха возвращает созданного пользователя без пароля.

// @Summary Create a new user
// @Description Create a new user with the provided details.
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User details"
// @Success 201 {object} models.User "Successfully created user"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user [post]
func (h *UserHandlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	// Декодируем JSON из тела запроса в структуру пользователя.
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Валидируем пользователя.
	err = validation.ValidateUser(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Хешируем пароль пользователя.
	err = newUser.HashPassword(newUser.Password)
	if err != nil {
		http.Error(w, `{"error":"error hashing password"}`, http.StatusInternalServerError)
		return
	}

	//newUser.Created_at = time.Now()
	//newUser.Updated_at = time.Now()

	// Создаем нового пользователя с помощью UserService.
	err = h.UserService.Create(&newUser)
	if err != nil {
		http.Error(w, `{"error":"error creating user"}`, http.StatusInternalServerError)
		return
	}

	// Убираем пароль из ответа для безопасности.
	newUser.Password = ""

	// Устанавливаем заголовки ответа и отправляем данные пользователя.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Кодируем и отправляем данные пользователя как JSON.
	err = json.NewEncoder(w).Encode(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// DeleteUser обрабатывает запрос на удаление пользователя по его ID.
// Сначала проверяет наличие пользователя в базе данных, затем удаляет его.
// В случае успеха возвращает статус OK.

// @Summary Delete a user by ID
// @Description Delete an existing user using their ID.
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {string} string "Successfully deleted"
// @Failure 400 {object} map[string]string "Invalid ID supplied"
// @Failure 404 {object} map[string]string "User not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/{id} [delete]
func (h *UserHandlers) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Получаем ID пользователя из параметров URL.
	params := mux.Vars(r)
	id := params["id"]

	// Проверяем, что ID не пустой.
	if id == "" {
		http.Error(w, `{"error":"id is not found"}`, http.StatusBadRequest)
		return
	}

	// Ищем пользователя в базе данных по его ID.
	_, err := h.UserService.Get(id)
	if err != nil {
		http.Error(w, `{"error":"user is not found"}`, http.StatusNotFound)
		return
	}

	// Удаляем пользователя с заданным ID.
	err = h.UserService.Delete(id)
	if err != nil {
		http.Error(w, `{"error":"error to delete user"}`, http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовки ответа и отправляем статус OK.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
