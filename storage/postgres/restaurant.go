package postgres

import (
	"context"
	"database/sql"
	"log"
	pb "reservation_service/genproto/restaurant"
)

type ReservationRepo struct {
	DB *sql.DB
}

func NewReservationRepo(db *sql.DB) *ReservationRepo {
	return &ReservationRepo{DB: db}
}

func (r *ReservationRepo) CreateRestaurant(ctx context.Context, res *pb.RestaurantDetails) (*pb.ID, error) {
	query := `
				INSERT INTO restaurants (name, address, phone_number, description)
                VALUES ($1, $2, $3, $4)
                RETURNING id
            `
	id := pb.ID{}
	err := r.DB.QueryRowContext(ctx, query, res.Name, res.Address,
		res.PhoneNumber, res.Description).Scan(&id)
	if err != nil {
		log.Println("failed to insert restaurant", err)
		return nil, err
	}
	return &id, nil
}

func (r *ReservationRepo) GetRestaurant(ctx context.Context, id *pb.ID) (*pb.RestaurantInfo, error) {
	res := pb.RestaurantInfo{Id: id.Id}
	query := `
                SELECT name, address, phone_number, description
                FROM restaurants
                WHERE deleted_at is null and id = $1
            `
	err := r.DB.QueryRowContext(ctx, query, id.Id).Scan(
		&res.Name, &res.Address, &res.PhoneNumber, &res.Description)
	if err != nil {
		log.Println("restaurant not found", err)
		return nil, err
	}
	return &res, nil
}

func (r *ReservationRepo) UpdateRestaurant(ctx context.Context, res *pb.RestaurantInfo) error {
	query := `
                UPDATE restaurants
                SET name = $1, address = $2, phone_number = $3, description = $4 updated_at = NOW()
                WHERE deleted_at is null and id = $5
            `
	_, err := r.DB.ExecContext(ctx, query,
		res.Name, res.Address, res.PhoneNumber, res.Description, res.Id)
	if err != nil {
		log.Println("failed to update restaurant", err)
		return err
	}
	return nil
}

func (r *ReservationRepo) DeleteRestaurant(ctx context.Context, id *pb.ID) error {
	query := `
			UPDATE restaurants
			SET deleted_at = NOW()
			WHERE deleted_at is null and id = $1
            `
	_, err := r.DB.ExecContext(ctx, query, id.Id)
	if err != nil {
		log.Println("failed to delete", err)
		return err
	}
	return nil
}

func (r *ReservationRepo) FetchRestaurants(ctx context.Context, pag *pb.Pagination) ([]*pb.RestaurantInfo, error) {
	query := `
			SELECT id, name, address, phone_number, description
			FROM restaurants
			WHERE deleted_at IS NULL
			LIMIT $1
            OFFSET $2
            `
	rows, err := r.DB.QueryContext(ctx, query, pag.Limit, pag.Offset)
	if err != nil {
		log.Println("failed to fetch restaurants", err)
		return nil, err
	}
	defer rows.Close()

	var restaurants []*pb.RestaurantInfo
	for rows.Next() {
		res := &pb.RestaurantInfo{}

		err := rows.Scan(&res.Id, &res.Name, &res.Address, &res.PhoneNumber, &res.Description)
		if err != nil {
			log.Println("failed to scan row", err)
			return nil, err
		}

		restaurants = append(restaurants, res)
	}

	return restaurants, nil
}
