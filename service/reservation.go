package service

import (
	"context"
	"database/sql"
	pb "reservation_service/genproto/reservation"
	"reservation_service/storage/postgres"

	"github.com/pkg/errors"
)

type ReservationService struct {
	pb.UnimplementedReservationServer
	Repo *postgres.ReservationRepo
}

func NewReservationService(db *sql.DB) *ReservationService {
	return &ReservationService{Repo: postgres.NewReservationRepo(db)}
}

func (s *ReservationService)CreateReservation(ctx context.Context, reser *pb.ReservationDetails) (*pb.ID, error){
	resp, err := s.Repo.CreateReservation(ctx, reser)
    if err!= nil {
        return nil, errors.Wrap(err, "failed to create reservation")
    }
    return resp, nil
}

func (s *ReservationService) GetReservationByID(ctx context.Context, id *pb.ID) (*pb.ReservationInfo, error) {
	resp, err := s.Repo.GetReservationById(ctx, id)
    if err!= nil {
        return nil, errors.Wrap(err, "failed to get reservation by ID")
    }
    return resp, nil
}

func (s *ReservationService) UpdateReservation(ctx context.Context, reser *pb.ReservationInfo) (*pb.Void, error) {
	err := s.Repo.UpdateReservation(ctx, reser)
    if err!= nil {
        return nil, errors.Wrap(err, "failed to update reservation")
    }
    return &pb.Void{}, nil
}

func (s *ReservationService) DeleteReservation(ctx context.Context, id *pb.ID) (*pb.Void, error) {
	err := s.Repo.DeleteReservation(ctx, id)
    if err!= nil {
        return nil, errors.Wrap(err, "failed to delete reservation")
    }
    return &pb.Void{}, nil
}

func (s *ReservationService) ValidateReservation(ctx context.Context, reser *pb.ReservationDetails) (*pb.ID, error) {
	resp, err := s.Repo.ValidateReservation(ctx, reser)
    if err!= nil {
        return nil, errors.Wrap(err, "failed to validate reservation")
    }
    return resp, nil
}

func (s *ReservationService) Order(ctx context.Context, reser *pb.ReservationOrders) (*pb.ID, error) {
	resp, err := s.Repo.Order(ctx, reser)
    if err!= nil {
        return nil, errors.Wrap(err, "failed to order")
    }
    return resp, nil
}

func (s *ReservationService) ChangeStatus(ctx context.Context, id, status string) error {
	err := s.Repo.ChangeStatus(ctx, id, status)
    if err!= nil {
        return errors.Wrap(err, "failed to change reservation status")
    }
    return nil
}

func (s *ReservationService) FetchReservations(ctx context.Context, filter *pb.Filter) (*pb.Reservations, error) {
	resp, err := s.Repo.FetchReservations(ctx, filter)
    if err!= nil {
        return nil, errors.Wrap(err, "failed to fetch reservations")
    }
    return resp, nil
}
