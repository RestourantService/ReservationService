package postgres

import (
	"database/sql"
)

type ReservationRepo struct {
	DB *sql.DB
}

func NewReservationRepo(db *sql.DB) *ReservationRepo {
	return &ReservationRepo{DB: db}
}
