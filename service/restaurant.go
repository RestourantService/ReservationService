package service

import (
	"context"
	"database/sql"
	pb "reservation_service/genproto/restaurant"
	"reservation_service/storage/postgres"

	"github.com/pkg/errors"
)

type RestaurantService struct {
	pb.UnimplementedRestaurantServer
	Repo *postgres.RestaurantRepo
}

func NewRestaurantService(db *sql.DB) *RestaurantService {
	return &RestaurantService{Repo: postgres.NewRestaurantRepo(db)}
}

func (r *RestaurantService) CreateRestaurant(ctx context.Context, req *pb.RestaurantDetails) (*pb.ID, error) {
	resp, err := r.Repo.CreateRestaurant(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create restaurant")
	}

	return resp, nil
}

func (r *RestaurantService) GetRestaurantByID(ctx context.Context, req *pb.ID) (*pb.RestaurantInfo, error) {
	resp, err := r.Repo.GetRestaurant(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read restaurant")
	}

	return resp, nil
}

func (r *RestaurantService) UpdateRestaurant(ctx context.Context, req *pb.RestaurantInfo) (*pb.Void, error) {
	err := r.Repo.UpdateRestaurant(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update restaurant")
	}

	return &pb.Void{}, nil
}

func (r *RestaurantService) DeleteRestaurant(ctx context.Context, req *pb.ID) (*pb.Void, error) {
	err := r.Repo.DeleteRestaurant(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update restaurant")
	}

	return &pb.Void{}, nil
}

func (r *RestaurantService) FetchRestaurants(ctx context.Context, req *pb.Pagination) (*pb.Restaurants, error) {
	resp, err := r.Repo.FetchRestaurants(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch restaurants")
	}

	return &pb.Restaurants{Restaurants: resp}, nil
}
