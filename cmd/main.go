package main

import (
	"context"
	"log"
	"net/http"
	"os"

	_ "github.com/AramLab/booking-service/docs" // импорт сгенерированных Swagger файлов
	"github.com/AramLab/booking-service/server/routes"
	"github.com/AramLab/booking-service/service/domain"
	"github.com/AramLab/booking-service/storage/postgres"
	"github.com/AramLab/booking-service/storage/setUpDB"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Booking Service API
// @version 1.0
// @description This is a sample server for a booking service API.
// @host localhost:8080
// @BasePath /
func main() {
	// Загружаем переменные окружения из файла `.env`.
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// Конфигурация базы данных, основанная на переменных окружения.
	config := &setUpDB.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	ctx := context.Background()

	// Подключаемся к базе данных с помощью функции `setUpDB.ConnectDB`.
	db, err := setUpDB.ConnectDB(ctx, config)
	if err != nil {
		log.Fatal("could not connect to the database")
	}
	defer db.Close()

	// Создаем таблицы в базе данных (если они не созданы).
	err = setUpDB.CreateTables(db)
	if err != nil {
		log.Fatalf("error creating tables: %v", err)
	}

	// Создаем репозиторий для работы с базой данных.
	repository := postgres.NewRepository(db)

	// Инициализируем сервисы для работы с пользователями и бронированиями.
	service := domain.NewService(repository)

	// Регистрируем маршруты для API, передавая сервисы пользователей и бронирований.
	r := routes.RegisterRoutes(service.UserServ, service.BookingServ)

	// Swagger UI
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // URL pointing to API definition
		httpSwagger.DeepLinking(true),        // Deep linking for documentation
	))

	port := os.Getenv("DB_SERVER_PORT")

	log.Printf("Starting server on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
