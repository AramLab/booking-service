package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/AramLab/booking-service/server/routes"
	"github.com/AramLab/booking-service/service/domain"
	"github.com/AramLab/booking-service/storage/postgres"
	"github.com/AramLab/booking-service/storage/setUpDB"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &setUpDB.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
	ctx := context.Background()
	db, err := setUpDB.ConnectDB(ctx, config)
	if err != nil {
		log.Fatal("could not connect to the database")
	}
	defer db.Close()

	repository := postgres.NewRepository(db)
	service := domain.NewService(repository)

	r := routes.RegisterRoutes(service.UserServ, service.BookingServ)

	port := os.Getenv("DB_SERVER_PORT")

	log.Printf("Starting server on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))

}
