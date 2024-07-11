package service

import (
	"context"
	"database/sql"
	"log/slog"
	pb "reservation_service/genproto/restaurant"
	"reservation_service/pkg/logger"
	"reservation_service/storage/postgres"

	"github.com/pkg/errors"
)

type RestaurantService struct {
	pb.UnimplementedRestaurantServer
	Repo   *postgres.RestaurantRepo
	Logger *slog.Logger
}

func NewRestaurantService(db *sql.DB) *RestaurantService {
	return &RestaurantService{
		Repo:   postgres.NewRestaurantRepo(db),
		Logger: logger.NewLogger(),
	}
}

func (r *RestaurantService) CreateRestaurant(ctx context.Context, req *pb.RestaurantDetails) (*pb.ID, error) {
	r.Logger.Info("CreateRestaurant method is starting")

	resp, err := r.Repo.CreateRestaurant(ctx, req)
	if err != nil {
		err := errors.Wrap(err, "failed to create restaurant")
		r.Logger.Error(err.Error())
		return nil, err
	}

	r.Logger.Info("CreateRestaurant has successfully finished")
	return resp, nil
}

func (r *RestaurantService) GetRestaurantByID(ctx context.Context, req *pb.ID) (*pb.RestaurantInfo, error) {
	r.Logger.Info("GetRestaurantByID method is starting")

	resp, err := r.Repo.GetRestaurant(ctx, req)
	if err != nil {
		err := errors.Wrap(err, "failed to read restaurant")
		r.Logger.Error(err.Error())
		return nil, err
	}

	r.Logger.Info("GetRestaurantByID has successfully finished")
	return resp, nil
}

func (r *RestaurantService) UpdateRestaurant(ctx context.Context, req *pb.RestaurantInfo) (*pb.Void, error) {
	r.Logger.Info("UpdateRestaurant method is starting")

	err := r.Repo.UpdateRestaurant(ctx, req)
	if err != nil {
		err := errors.Wrap(err, "failed to update restaurant")
		r.Logger.Error(err.Error())
		return nil, err
	}

	r.Logger.Info("UpdateRestaurant has successfully finished")
	return &pb.Void{}, nil
}

func (r *RestaurantService) DeleteRestaurant(ctx context.Context, req *pb.ID) (*pb.Void, error) {
	r.Logger.Info("DeleteRestaurant method is starting")

	err := r.Repo.DeleteRestaurant(ctx, req)
	if err != nil {
		err := errors.Wrap(err, "failed to update restaurant")
		r.Logger.Error(err.Error())
		return nil, err
	}

	r.Logger.Info("DeleteRestaurant has successfully finished")
	return &pb.Void{}, nil
}

func (r *RestaurantService) FetchRestaurants(ctx context.Context, req *pb.Pagination) (*pb.Restaurants, error) {
	r.Logger.Info("FetchRestaurants method is starting")

	resp, err := r.Repo.FetchRestaurants(ctx, req)
	if err != nil {
		err := errors.Wrap(err, "failed to fetch restaurants")
		r.Logger.Error(err.Error())
		return nil, err
	}

	r.Logger.Info("FetchRestaurants has successfully finished")
	return &pb.Restaurants{Restaurants: resp}, nil
}
