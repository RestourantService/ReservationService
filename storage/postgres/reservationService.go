package postgres

import (
	"github.com/jmoiron/sqlx"
)

type ReservationRepo struct {
	db *sqlx.DB
}

func NewReservationRepo(db *sqlx.DB) *ReservationRepo {
	return &ReservationRepo{db: db}
}
