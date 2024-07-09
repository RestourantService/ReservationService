package service

import (
	"database/sql"
	pb "reservation_service/genproto/menu"
	"reservation_service/storage/postgres"
)

type MenuService struct {
	pb.UnimplementedMenuServer
	Repo *postgres.MenuRepo
}

func NewMenuService(db *sql.DB) *MenuService {
	return &MenuService{Repo: postgres.NewMenuRepo(db)}
}
