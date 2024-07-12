package service

import (
	"context"
	"database/sql"
	"log/slog"
	pb "reservation_service/genproto/menu"
	"reservation_service/pkg/logger"
	"reservation_service/storage/postgres"

	"github.com/pkg/errors"
)

type MenuService struct {
	pb.UnimplementedMenuServer
	Repo   *postgres.MenuRepo
	Logger *slog.Logger
}

func NewMenuService(db *sql.DB) *MenuService {
	return &MenuService{
		Repo:   postgres.NewMenuRepo(db),
		Logger: logger.NewLogger(),
	}

}

func (m *MenuService) AddMeal(ctx context.Context, req *pb.MealDetails) (*pb.ID, error) {
	m.Logger.Info("AddMeal method is starting")

	resp, err := m.Repo.AddMeal(ctx, req)
	if err != nil {
		err := errors.Wrap(err, "failed to add meal")
		m.Logger.Error(err.Error())
		return nil, err
	}

	m.Logger.Info("AddMeal has successfully finished")
	return resp, nil
}

func (m *MenuService) GetMealByID(ctx context.Context, req *pb.ID) (*pb.MealInfo, error) {
	m.Logger.Info("GetMealByID method is starting")

	resp, err := m.Repo.GetMealByID(ctx, req)
	if err != nil {
		err := errors.Wrap(err, "failed to read meal")
		m.Logger.Error(err.Error())
		return nil, err
	}

	m.Logger.Info("GetMealById has successfully finished")
	return resp, nil
}

func (m *MenuService) UpdateMeal(ctx context.Context, req *pb.MealInfo) (*pb.Void, error) {
	m.Logger.Info("UpdateMeal method is starting")

	err := m.Repo.UpdateMeal(ctx, req)
	if err != nil {
		err := errors.Wrap(err, "failed to update meal")
		m.Logger.Error(err.Error())
		return nil, err
	}

	m.Logger.Info("UpdateMeal has successfully finished")
	return &pb.Void{}, nil
}

func (m *MenuService) DeleteMeal(ctx context.Context, req *pb.ID) (*pb.Void, error) {
	m.Logger.Info("DeleteMeal method is starting")

	err := m.Repo.DeleteMeal(ctx, req)
	if err != nil {
		err := errors.Wrap(err, "failed to delete meal")
		m.Logger.Error(err.Error())
		return nil, err
	}

	m.Logger.Info("DeleteMeal has successfully finished")
	return &pb.Void{}, nil
}

func (m *MenuService) FetchMeals(ctx context.Context, req *pb.Filter) (*pb.Meals, error) {
	m.Logger.Info("FetchMeals method is starting")

	resp, err := m.Repo.GetAllMeals(ctx, req)
	if err != nil {
		err := errors.Wrap(err, "failed to fetch meals")
		m.Logger.Error(err.Error())
		return nil, err
	}

	m.Logger.Info("FetchMeals has successfully finished")
	return resp, nil
}
