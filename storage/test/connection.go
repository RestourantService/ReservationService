package test

import (
	"database/sql"
	"reservation_service/storage/postgres"
)

func ConnectDB() (*sql.DB, error) {
	return postgres.ConnectDB()
}
