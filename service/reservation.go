package service

import (
	"context"
	"database/sql"
	"log/slog"
	"reservation_service/genproto/menu"
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
	Logger        *slog.Logger
}

func NewReservationService(db *sql.DB, user pbu.UserClient, payment pbp.PaymentClient) *ReservationService {
	return &ReservationService{
		Repo:          postgres.NewReservationRepo(db),
		UserClient:    user,
		PaymentClient: payment,
		Logger:        logger.NewLogger(),
	}
}

func (r *ReservationService) CreateReservation(ctx context.Context, req *pb.ReservationDetails) (*pb.ID, error) {
	r.Logger.Info("CreateReservation method is starting")

	status, err := r.UserClient.ValidateUser(ctx, &pbu.ID{Id: req.UserId})
	if err != nil {
		err := errors.Wrap(err, "user validation failed")
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
	r.Logger.Info("UpdateReservation method is starting")

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

	err = r.Repo.UpdateReservation(ctx, req)
	if err != nil {
		err := errors.Wrap(err, "failed to update reservation")
		r.Logger.Error(err.Error())
		return nil, err
	}

	r.Logger.Info("UpdateReservation has successfully finished")
	return &pb.Void{}, nil
}

func (r *ReservationService) DeleteReservation(ctx context.Context, req *pb.ID) (*pb.Void, error) {
	r.Logger.Info("DeleteReservation method is starting")

	err := r.Repo.DeleteReservation(ctx, req)
	if err != nil {
		err := errors.Wrap(err, "failed to delete reservation")
		r.Logger.Error(err.Error())
		return nil, err
	}

	r.Logger.Info("DeleteReservation has successfully finished")
	return &pb.Void{}, nil
}

func (r *ReservationService) ValidateReservation(ctx context.Context, req *pb.ID) (*pb.Status, error) {
	r.Logger.Info("ValidateReservation method is starting")

	resp, err := r.Repo.ValidateReservation(ctx, req.Id)
	if err != nil {
		err := errors.Wrap(err, "failed to validate reservation")
		r.Logger.Error(err.Error())
		return nil, err
	}

	r.Logger.Info("ValidateReservation has successfully finished")
	return resp, nil
}

func (r *ReservationService) Order(ctx context.Context, req *pb.ReservationOrder) (*pb.ID, error) {
	r.Logger.Info("Order method is starting")

	reserInfo, err := r.Repo.GetReservationById(ctx, &pb.ID{Id: req.Id})
	if err != nil {
		err := errors.Wrap(err, "failed to find reservation")
		r.Logger.Error(err.Error())
		return nil, err
	}

	r.Logger.Info("Reservation info has been retrieved")

	if reserInfo.Status != "confirmed" {
		err := errors.New("failed: incomplete reservation")
		r.Logger.Error(err.Error())
		return nil, err
	}

	reserTime, err := time.Parse(time.RFC3339, reserInfo.ReservationTime)
	if err != nil {
		err := errors.Wrap(err, "invalid reservation time")
		r.Logger.Error(err.Error())
		return nil, err
	}

	r.Logger.Info("Reservation end time has been parsed")

	resp, err := r.Repo.Order(ctx, req, reserTime)
	if err != nil {
		err := errors.Wrap(err, "failed to make order")
		r.Logger.Error(err.Error())
		return nil, err
	}

	m := MenuService{}
	meal, err := m.Repo.GetMealByID(ctx, &menu.ID{Id: req.MenuItemId})
	if err != nil {
		err := errors.Wrap(err, "failed to find meal")
		r.Logger.Error(err.Error())
		return nil, err
	}

	_, err = r.PaymentClient.CreatePayment(ctx, &pbp.PaymentDetails{
		ReservationId: req.Id,
		Amount:        meal.Price,
		PaymentMethod: "cash",
	})

	if err != nil {
		err := errors.Wrap(err, "failed to create payment")
		r.Logger.Error(err.Error())
		return nil, err
	}

	r.Logger.Info("Order has successfully finished")
	return resp, nil
}

func (r *ReservationService) Pay(ctx context.Context, req *pb.ID) (*pb.Status, error) {
	r.Logger.Info("Pay method is starting")

	status, err := r.Repo.ValidateReservation(ctx, req.Id)
	if err != nil {
		err := errors.Wrap(err, "failed to validate reservation")
		r.Logger.Error(err.Error())
		return nil, err
	}
	if !status.Successful {
		err := errors.Wrap(err, "reservation validation failed")
		r.Logger.Error(err.Error())
		return nil, err
	}

	r.Logger.Info("Reservation has been validated")

	payment, err := r.PaymentClient.SearchByReservationID(ctx, &pbp.ID{Id: req.Id})
	if err != nil {
		err := errors.Wrap(err, "failed to find payment")
		r.Logger.Error(err.Error())
		return nil, err
	}

	r.Logger.Info("Payment has been found")

	payment.PaymentStatus = "completed"
	_, err = r.PaymentClient.UpdatePayment(ctx, payment)
	if err != nil {
		err := errors.Wrap(err, "failed to update payment")
		r.Logger.Error(err.Error())
		return &pb.Status{Successful: false}, err
	}

	err = r.Repo.ChangeStatus(ctx, req.Id, "confirmed")
	if err != nil {
		err := errors.Wrap(err, "failed to confirm payment")
		r.Logger.Error(err.Error())
		return &pb.Status{Successful: false}, err
	}

	r.Logger.Info("Pay has successfully finished")
	return &pb.Status{Successful: true}, nil
}

func (r *ReservationService) FetchReservations(ctx context.Context, req *pb.Filter) (*pb.Reservations, error) {
	r.Logger.Info("FetchReservations method is starting")

	resp, err := r.Repo.FetchReservations(ctx, req)
	if err != nil {
		err := errors.Wrap(err, "failed to fetch reservations")
		r.Logger.Error(err.Error())
		return nil, err
	}

	r.Logger.Info("FetchReservations has successfully finished")
	return resp, nil
}

func (r *ReservationService) DeleteReservationByUserID(ctx context.Context, req *pb.ID) (*pb.Status, error) {
	r.Logger.Info("DeleteReservationByUserID method is starting")

	err := r.Repo.DeleteReservationByUserID(ctx, req)
	if err != nil {
		err := errors.Wrap(err, "failed to delete reservations by user ID")
		r.Logger.Error(err.Error())
		return nil, err
	}

	r.Logger.Info("DeleteReservationByUserID has successfully finished")
	return &pb.Status{Successful: true}, nil
}
