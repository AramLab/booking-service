package setUpDB

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func CreateTables(db *pgxpool.Pool) error {
	ctx := context.Background()

	userTable := `
	CREATE TABLE IF NOT EXISTS "user" (
		id SERIAL PRIMARY KEY,
		username VARCHAR(100) NOT NULL,
		password VARCHAR(250) NOT NULL,
		created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
	);`

	bookingTable := `CREATE TABLE IF NOT EXISTS booking (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL REFERENCES "user"(id) ON DELETE CASCADE,
		start_time TIMESTAMPTZ NOT NULL,
		end_time TIMESTAMPTZ NOT NULL
	);`

	_, err := db.Exec(ctx, userTable)
	if err != nil {
		return fmt.Errorf("error creating user table: %v", err)
	}

	_, err = db.Exec(ctx, bookingTable)
	if err != nil {
		return fmt.Errorf("error creating booking table: %v", err)
	}

	fmt.Println("Tables created successfully")
	return nil
}
