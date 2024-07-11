package service

import (
	"context"
	"database/sql"
	"log/slog"
	pbp "reservation_service/genproto/payment"
	pb "reservation_service/genproto/reservation"
	pbu "reservation_service/genproto/user"
	"reservation_service/pkg/logger"
	"reservation_service/storage/postgres"
	"time"

	"github.com/pkg/errors"
)

type ReservationService struct {
	pb.UnimplementedReservationServer
	Repo          *postgres.ReservationRepo
	UserClient    pbu.UserClient
	PaymentClient pbp.PaymentClient
	Logger *slog.Logger
}

func NewReservationService(db *sql.DB, user pbu.UserClient, payment pbp.PaymentClient) *ReservationService {
	return &ReservationService{
		Repo:          postgres.NewReservationRepo(db),
		UserClient:    user,
		PaymentClient: payment,
		Logger: logger.NewLogger(),
	}
}

func (r *ReservationService) CreateReservation(ctx context.Context, req *pb.ReservationDetails) (*pb.ID, error) {
	r.Logger.Info("CreateReservation method is starting")

	status, err := r.UserClient.ValidateUser(ctx, &pbu.ID{Id: req.UserId})
	if err != nil {
		err := errors.Wrap(err, "failed to validate user")
		r.Logger.Error(err.Error())
		return nil, err
	}
	if !status.Successful {
		err := errors.Wrap(err, "user validation failed")
		r.Logger.Error(err.Error())
		return nil, err
	}

	r.Logger.Info("User has been validated")

	resp, err := r.Repo.CreateReservation(ctx, req)
	if err != nil {
		err := errors.Wrap(err, "failed to create reservation")
		r.Logger.Error(err.Error())
		return nil, err
	}

	r.Logger.Info("Reservation has been created")

	_, err = r.PaymentClient.CreatePayment(ctx, &pbp.PaymentDetails{
		ReservationId: resp.Id,
		Amount:        0,
		PaymentMethod: "cash",
	})

	if err != nil {
		err := errors.Wrap(err, "failed to create payment")
		r.Logger.Error(err.Error())
		return nil, err
	}

	r.Logger.Info("CreateReservation has successfully finished")
	return resp, nil
}

func (r *ReservationService) GetReservationByID(ctx context.Context, req *pb.ID) (*pb.ReservationInfo, error) {
	r.Logger.Info("GetReservationByID method is starting")
	
	resp, err := r.Repo.GetReservationById(ctx, req)
	if err != nil {
		err := errors.Wrap(err, "failed to read reservation")
		r.Logger.Error(err.Error())
		return nil, err
	}

	r.Logger.Info("GetReservationByID has successfully finished")
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
	status, err := r.Repo.ValidateReservation(ctx, req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to validate reservation")
	}
	if !status.Successful {
		return nil, errors.Wrap(err, "reservation validation failed")
	}

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
