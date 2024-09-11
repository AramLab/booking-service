package user

import (
	"context"
	"errors"

	"github.com/AramLab/booking-service/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Save(user *models.User) error {
	err := ur.db.QueryRow(context.Background(),
		`INSERT INTO "user" (username, password, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id`,
		user.Username, user.Password, user.Created_at, user.Updated_at).Scan(&user.ID)
	if err != nil {
		return errors.New(`{"error":"failed to add user"}`)
	}
	return nil
}

func (ur *UserRepository) DeleteById(id string) error {
	_, err := ur.db.Exec(context.Background(), "DELETE FROM booking WHERE user_id=$1", id)
	if err != nil {
		return errors.New(`{"error":"failed to delete bookings"}`)
	}
	_, err = ur.db.Exec(context.Background(), `DELETE FROM "user" WHERE id=$1`, id)
	if err != nil {
		return errors.New(`{"error":"failed to delete user"}`)
	}
	return nil
}

func (ur *UserRepository) FindAll() ([]models.User, error) {
	rows, err := ur.db.Query(context.Background(), `SELECT id, username, created_at, updated_at FROM "user"`)
	if err != nil {
		return []models.User{}, errors.New(`{"error":"failed to find all users"}`)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Created_at, &user.Updated_at)
		if err != nil {
			return []models.User{}, errors.New(`{"error":"error scanning users"}`)
		}
		users = append(users, user)
	}
	return users, nil
}

func (ur *UserRepository) FindById(id string) (models.User, error) {
	var user models.User
	row := ur.db.QueryRow(context.Background(), `SELECT id, username, created_at, updated_at FROM "user" WHERE id=$1`, id)
	err := row.Scan(&user.ID, &user.Username, &user.Created_at, &user.Updated_at)
	if err != nil {
		return models.User{}, errors.New(`{"error":"error scanning user"}`)
	}
	return user, nil
}
