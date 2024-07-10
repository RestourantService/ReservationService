package service

import (
	"context"
	"database/sql"
	pbp "reservation_service/genproto/payment"
	pb "reservation_service/genproto/reservation"
	pbu "reservation_service/genproto/user"
	"reservation_service/storage/postgres"
	"time"

	"github.com/pkg/errors"
)

type ReservationService struct {
	pb.UnimplementedReservationServer
	Repo          *postgres.ReservationRepo
	UserClient    pbu.UserClient
	PaymentClient pbp.PaymentClient
}

func NewReservationService(db *sql.DB, user pbu.UserClient, payment pbp.PaymentClient) *ReservationService {
	return &ReservationService{
		Repo:          postgres.NewReservationRepo(db),
		UserClient:    user,
		PaymentClient: payment,
	}
}

func (r *ReservationService) CreateReservation(ctx context.Context, req *pb.ReservationDetails) (*pb.ID, error) {
	status, err := r.UserClient.ValidateUser(ctx, &pbu.ID{Id: req.UserId})
	if !status.Successful || err != nil {
		return nil, errors.Wrap(err, "no such user")
	}

	resp, err := r.Repo.CreateReservation(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create reservation")
	}

	_, err = r.PaymentClient.CreatePayment(ctx, &pbp.PaymentDetails{
		ReservationId: resp.Id,
		Amount:        0,
		PaymentMethod: "cash",
	})

	if err != nil {
		return nil, errors.Wrap(err, "failed to create payment")
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

func (r *ReservationService) ValidateReservation(ctx context.Context, req *pb.ID) (*pb.Status, error) {
	resp, err := r.Repo.ValidateReservation(ctx, req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to validate reservation")
	}

	return resp, nil
}

func (r *ReservationService) Order(ctx context.Context, req *pb.ReservationOrders) (*pb.ID, error) {
	reserInfo, err := r.Repo.GetReservationById(ctx, &pb.ID{Id: req.Id})
	if err != nil {
		return nil, errors.Wrap(err, "failed to find reservation")
	}

	reserTime, err := time.Parse(time.RFC3339, reserInfo.ReservationTime)
	if err != nil {
		return nil, errors.Wrap(err, "invalid reservation time")
	}

	resp, err := r.Repo.Order(ctx, req, reserTime)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make order")
	}

	return resp, nil
}

func (r *ReservationService) Pay(ctx context.Context, req *pb.ID) (*pb.Status, error) {
	payment, err := r.PaymentClient.SearchByReservationID(ctx, &pbp.ID{Id: req.Id})
	if err != nil {
		return nil, errors.Wrap(err, "failed to find payment")
	}

	payment.PaymentStatus = "completed"
	_, err = r.PaymentClient.UpdatePayment(ctx, payment)
	if err != nil {
		return &pb.Status{Successful: false}, errors.Wrap(err, "failed to update payment")
	}

	err = r.Repo.ChangeStatus(ctx, req.Id, "confirmed")
	if err != nil {
		return &pb.Status{Successful: false}, errors.Wrap(err, "failed to confirm payment")
	}

	return &pb.Status{Successful: true}, nil
}

func (r *ReservationService) FetchReservations(ctx context.Context, req *pb.Filter) (*pb.Reservations, error) {
	resp, err := r.Repo.FetchReservations(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch reservations")
	}

	return resp, nil
}
