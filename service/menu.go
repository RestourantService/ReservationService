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

func (m *MenuService) CreateMenu(ctx context.Context, req *pb.MealDetails) (*pb.ID, error) {
	resp, err := m.Repo.CreateMenu(ctx, req)
    if err!= nil {
        return nil, errors.Wrap(err, "failed to create menu")
    }
    return resp, nil
}

func (m *MenuService) GetMealByID(ctx context.Context, req *pb.ID) (*pb.MealInfo, error) {
	resp, err := m.Repo.GetMealByID(ctx, req)
    if err!= nil {
        return nil, errors.Wrap(err, "failed to get menu")
    }
    return resp, nil
}

func (m *MenuService) UpdateMeal(ctx context.Context, req *pb.MealInfo) (*pb.Void, error) {
	err := m.Repo.UpdateMeal(ctx, req)
    if err!= nil {	
        return nil, errors.Wrap(err, "failed to update menu")
    }
    return &pb.Void{}, nil
}

func (m *MenuService) DeleteMeal(ctx context.Context, req *pb.ID) (*pb.Void, error) {
	err := m.Repo.DeleteMeal(ctx, req)
    if err!= nil {
        return nil, errors.Wrap(err, "failed to delete menu")
    }
    return &pb.Void{}, nil
}

func (m *MenuService) GetAllMeals(ctx context.Context, req *pb.Filter) (*pb.Meals, error) {
	resp, err := m.Repo.GetAllMeals(ctx, req)
    if err!= nil {
        return nil, errors.Wrap(err, "failed to get all meals")
    }
    return resp, nil
}