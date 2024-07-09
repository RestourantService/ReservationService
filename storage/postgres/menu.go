package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	pb "reservation_service/genproto/menu"
)

type MenuRepo struct {
	DB *sql.DB
}

func NewMenuRepo(db *sql.DB) *MenuRepo {
	return &MenuRepo{DB: db}
}

func (r *MenuRepo) CreateMenu(ctx context.Context, req *pb.MealDetails) (*pb.ID, error) {
	query := `
			INSERT INTO menu(restaurant_id, name, description, price) 
			VALUES($1, $2, $3, $4) 
			RETURNING id
			`
	var id string
	err := r.DB.QueryRowContext(ctx, query,
		req.RestaurantId, req.Name, req.Description, req.Price).Scan(&id)
	if err != nil {
		log.Println("failed to insert reservation", err)
		return nil, err
	}
	return &pb.ID{Id: id}, nil
}

func (r *MenuRepo) GetMealByID(ctx context.Context, id *pb.ID) (*pb.MealInfo, error) {
	query := `
			SELECT restaurant_id, name, description, price
                FROM menu
                WHERE deleted_at is null and id = $1
			`
	res := pb.MealInfo{Id: id.Id}
	err := r.DB.QueryRowContext(ctx, query, id.Id).Scan(
		&res.RestaurantId, &res.Name, &res.Description, &res.Price)
	if err != nil {
		log.Println("failed to fetch meal", err)
		return nil, err
	}
	return &res, nil
}

func (r *MenuRepo) UpdateMeal(ctx context.Context, req *pb.MealInfo) error {
	query := `
	    	UPDATE menu
			SET restaurant_id = $1, name = $2, description = $3, price = $4, updated_at = NOW()
			WHERE deleted_at is null and id = $5
			`
	_, err := r.DB.ExecContext(ctx, query, req.RestaurantId, req.Name, req.Description, req.Price, req.Id)
	if err != nil {
		log.Println("failed to update meal", err)
		return err
	}
	return nil
}

func (r *MenuRepo) DeleteMeal(ctx context.Context, req *pb.ID) error {
	query := `
	    	UPDATE menu
			SET deleted_at= NOW()
			WHERE deleted_at is null and id = $1
			`
	_, err := r.DB.ExecContext(ctx, query, req.Id)
	if err != nil {
		log.Println("failed to delete meal", err)
		return err
	}
	return nil
}

func (r *MenuRepo) GetAllMeals(ctx context.Context, req *pb.Filter) (*pb.Meals, error) {
	query := `SELECT id,restaurant_id, name, description,price
			from menu
	    	where deleted_at is null `
	if req.RestaurantId != "" {
		query += fmt.Sprintf(" and restaurant_id = %s", req.RestaurantId)
	}
	if req.Limit != 0 {
		query += fmt.Sprintf(" LIMIT %d", req.Limit)
	}
	if req.Offset != 0 {
		query += fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println("failed to fetch meals", err)
		return nil, err
	}
	defer rows.Close()

	var res []*pb.MealInfo
	for rows.Next() {
		var in pb.MealInfo

		err := rows.Scan(&in.Id, &in.RestaurantId, &in.Name, &in.Description, &in.Price)
		if err != nil {
			log.Println("failed to fetch meals", err)
			return nil, err
		}

		res = append(res, &in)
	}
	return &pb.Meals{Meals: res}, nil
}
