package service

import (
	"context"
	"database/sql"
	pbp "reservation_service/genproto/payment"
	pb "reservation_service/genproto/reservation"
	"reservation_service/storage/postgres"

	"github.com/pkg/errors"
)

type ReservationService struct {
	pb.UnimplementedReservationServer
	Repo          *postgres.ReservationRepo
	PaymentClient pbp.PaymentClient
}

func NewReservationService(db *sql.DB, paymentClient pbp.PaymentClient) *ReservationService {
	return &ReservationService{
		Repo:          postgres.NewReservationRepo(db),
		PaymentClient: paymentClient,
	}
}

func (r *ReservationService) CreateReservation(ctx context.Context, req *pb.ReservationDetails) (*pb.ID, error) {
	resp, err := r.Repo.CreateReservation(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create reservation")
	}

	_, err = r.PaymentClient.MakePayment(ctx, &pbp.PaymentDetails{
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

func (r *ReservationService) ValidateReservation(ctx context.Context, req *pb.ReservationDetails) (*pb.ID, error) {
	resp, err := r.Repo.ValidateReservation(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to validate reservation")
	}

	return resp, nil
}

func (r *ReservationService) Order(ctx context.Context, req *pb.ReservationOrders) (*pb.ID, error) {
	resp, err := r.Repo.Order(ctx, req)
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
