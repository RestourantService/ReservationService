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

func (r *ReservationService) CreateReservation(ctx context.Context, req *pb.ReservationDetails) (*pb.ID, error) {
	resp, err := r.Repo.CreateReservation(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create reservation")
	}

	return resp, nil
}

func (r *ReservationService) GetReservationByID(ctx context.Context, req *pb.ID) (*pb.ReservationInfo, error) {
	resp, err := r.Repo.GetReservationById(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read reservation")
	}

	return resp, nil
}

func (r *ReservationService) UpdateReservation(ctx context.Context, req *pb.ReservationInfo) (*pb.Void, error) {
	err := r.Repo.UpdateReservation(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update reservation")
	}

	return &pb.Void{}, nil
}

func (r *ReservationService) DeleteReservation(ctx context.Context, req *pb.ID) (*pb.Void, error) {
	err := r.Repo.DeleteReservation(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update reservation")
	}

	return &pb.Void{}, nil
}

func (r *ReservationService) ValidateReservation(ctx context.Context, req *pb.ReservationDetails) (*pb.ID, error) {
	resp, err := r.Repo.ValidateReservation(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to validate reservation")
	}

	return resp, nil
}

func (r *ReservationService) Order(ctx context.Context, req *pb.ReservationOrders) (*pb.ID, error)