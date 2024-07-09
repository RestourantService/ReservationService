package service

import (
	"database/sql"
	pb "reservation_service/genproto/restaurant"
	"reservation_service/storage/postgres"
)

type RestaurantService struct {
	pb.UnimplementedRestaurantServer
	Repo *postgres.RestaurantRepo
}

func NewRestaurantService(db *sql.DB) *RestaurantService {
	return &RestaurantService{Repo: postgres.NewRestaurantRepo(db)}
}
