package service

import (
	"context"
	"database/sql"
	pb "reservation_service/genproto/menu"
	"reservation_service/storage/postgres"

	"github.com/pkg/errors"
)

type MenuService struct {
	pb.UnimplementedMenuServer
	Repo *postgres.MenuRepo
}

func NewMenuService(db *sql.DB) *MenuService {
	return &MenuService{Repo: postgres.NewMenuRepo(db)}
}

func (m *MenuService) AddMeal(ctx context.Context, req *pb.MealDetails) (*pb.ID, error) {
	resp, err := m.Repo.AddMeal(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create meal")
	}

	return resp, nil
}

func (m *MenuService) GetMealByID(ctx context.Context, req *pb.ID) (*pb.MealInfo, error) {
	resp, err := m.Repo.GetMealByID(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read meal")
	}

	return resp, nil
}

func (m *MenuService) UpdateMeal(ctx context.Context, req *pb.MealInfo) (*pb.Void, error) {
	err := m.Repo.UpdateMeal(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update meal")
	}

	return &pb.Void{}, nil
}

func (m *MenuService) DeleteMeal(ctx context.Context, req *pb.ID) (*pb.Void, error) {
	err := m.Repo.DeleteMeal(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to delete meal")
	}

	return &pb.Void{}, nil
}

func (m *MenuService) FetchMeals(ctx context.Context, req *pb.Filter) (*pb.Meals, error) {
	resp, err := m.Repo.GetAllMeals(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch meals")
	}

	return resp, nil
}
