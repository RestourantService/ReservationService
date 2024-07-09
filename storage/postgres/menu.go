package postgres

import "database/sql"

type MenuRepo struct {
	DB *sql.DB
}

func NewMenuRepo(db *sql.DB) *MenuRepo {
	return &MenuRepo{DB: db}
}
