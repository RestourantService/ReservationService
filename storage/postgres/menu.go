package postgres

import "database/sql"

type Menu struct {
	Db *sql.DB
}

func NewMenuRepo(db *sql.DB) *Menu {
	return &Menu{Db: db}
}
