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

func (r *ReservationRepo) CreateRestaurant(ctx context.Context, restaurant *pb.RestaurantDetails) (*pb.ID, error) {
	id := pb.ID{}

	query := `
				INSERT INTO restaurants (name, address, phone_number)
                VALUES ($1, $2, $3, $4)
                RETURNING id
            `
	err := r.DB.QueryRowContext(ctx, query, restaurant.Name, restaurant.Address, restaurant.PhoneNumber).Scan(&id)
	if err != nil {
		log.Println("failed to insert", err)
		return nil, err
	}
	return &id, nil
}

func (r *ReservationRepo) GetRestaurant(ctx context.Context, id *pb.ID) (*pb.RestaurantInfo, error) {
	restaurant := &pb.RestaurantInfo{}

    query := `
                SELECT id, name, address, phone_number
                FROM restaurants
                WHERE id = $1
            `
    err := r.DB.QueryRowContext(ctx, query, id.Id).Scan(&restaurant.Id, &restaurant.Name, &restaurant.Address, &restaurant.PhoneNumber)
    if err!= nil {
        log.Println("restaurant not found", err)
        return nil, err
    }
    return restaurant, nil
}

func (r *ReservationRepo) UpdateRestaurant(ctx context.Context, restaurant *pb.RestaurantInfo) error {
	query := `
                UPDATE restaurants
                SET name = $1, address = $2, phone_number = $3, updated_at = NOW()
                WHERE id = $4
            `
    _, err := r.DB.ExecContext(ctx, query, restaurant.Name, restaurant.Address, restaurant.PhoneNumber, restaurant.Id)
    if err!= nil {
        log.Println("failed to update", err)
        return err
    }
    return nil
}

func (r *ReservationRepo) DeleteRestaurant(ctx context.Context, id *pb.ID) error {
	query := `
			UPDATE restaurants
			SET deleted_at = NOW()
			WHERE id = $1
            `
    _, err := r.DB.ExecContext(ctx, query, id.Id)
    if err!= nil {
        log.Println("failed to delete", err)
        return err
    }
    return nil
}

func (r *ReservationRepo) FetchRestaurants(ctx context.Context, pagination *pb.Pagination) ([]*pb.RestaurantInfo, error) {
	restaurants := []*pb.RestaurantInfo{}

    query := `
			SELECT id, name, address, phone_number
			FROM restaurants
			WHERE deleted_at IS NULL
			LIMIT 10 OFFSET 20
            `
    rows, err := r.DB.QueryContext(ctx, query, pagination.Limit, pagination.Offset)
    if err!= nil {
        log.Println("failed to fetch restaurants", err)
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        restaurant := &pb.RestaurantInfo{}
        err := rows.Scan(&restaurant.Id, &restaurant.Name, &restaurant.Address, &restaurant.PhoneNumber)
        if err!= nil {
            log.Println("failed to scan row", err)
            return nil, err
        }
        restaurants = append(restaurants, restaurant)
    }

    return restaurants, nil
}