package service

import (
	"database/sql"
	pb "reservation_service/genproto/reservation"
	"reservation_service/storage/postgres"
)

type ReservationService struct {
	pb.UnimplementedReservationServer
	Repo *postgres.ReservationRepo
}

func NewReservationService(db *sql.DB) *ReservationService {
	return &ReservationService{Repo: postgres.NewReservationRepo(db)}
}
