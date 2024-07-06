package services

import (
	pb "reservationService/genproto/ReservationService"
	"reservationService/storage/postgres"

	"github.com/jmoiron/sqlx"
)

type reservationService struct {
	reservationRepo *postgres.ReservationRepo
	pb.UnimplementedReservationServiceServer
}

func NewuserManagementService(db *sqlx.DB) *reservationService {
	return &reservationService{reservationRepo: postgres.NewReservationRepo(db)}
}
