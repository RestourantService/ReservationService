package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	pb "reservation_service/genproto/reservation"
	"reservation_service/storage/redis"
	"time"

	"github.com/google/uuid"
)

type ReservationRepo struct {
	DB *sql.DB
}

func NewReservationRepo(db *sql.DB) *ReservationRepo {
	return &ReservationRepo{DB: db}
}

func (r *ReservationRepo) CreateReservation(ctx context.Context, reser *pb.ReservationDetails) (*pb.ID, error) {
	query := `
			INSERT INTO reservations(user_id, restaurant_id, reservation_time) 
			VALUES($1, $2, $3) 
			RETURNING id
			`
	var id string
	err := r.DB.QueryRowContext(ctx, query,
		reser.UserId, reser.RestaurantId, reser.ReservationTime).Scan(&id)
	if err != nil {
		log.Println("failed to insert reservation", err)
		return nil, err
	}

	return &pb.ID{Id: id}, nil
}

func (r *ReservationRepo) GetReservationById(ctx context.Context, id *pb.ID) (*pb.ReservationInfo, error) {
	reser := pb.ReservationInfo{Id: id.Id}
	query := `
                SELECT user_id, restaurant_id, reservation_time, status
                FROM reservations
                WHERE deleted_at is null and id = $1
            `

	err := r.DB.QueryRowContext(ctx, query, id.Id).Scan(
		&reser.UserId, &reser.RestaurantId, &reser.ReservationTime, &reser.Status)
	if err != nil {
		log.Println("reservation not found", err)
		return nil, err
	}

	return &reser, nil
}

func (r *ReservationRepo) UpdateReservation(ctx context.Context, reser *pb.ReservationInfo) error {
	query := `
	    	UPDATE reservations
			SET user_id = $1, restaurant_id = $2, reservation_time = $3, status = $4, updated_at = NOW()
			WHERE deleted_at is null and id = $5
			`

	_, err := r.DB.ExecContext(ctx, query,
		reser.UserId, reser.RestaurantId, reser.ReservationTime, reser.Status, reser.Id)
	if err != nil {
		log.Println("failed to update reservation", err)
		return err
	}

	return nil
}

func (r *ReservationRepo) DeleteReservation(ctx context.Context, id *pb.ID) error {
	query := `
            UPDATE reservations
            SET deleted_at = NOW()
            WHERE deleted_at is null and id = $1
            `

	_, err := r.DB.ExecContext(ctx, query, id.Id)
	if err != nil {
		log.Println("failed to delete reservation", err)
		return err
	}

	return nil
}

func (r *ReservationRepo) ValidateReservation(ctx context.Context, id string) (*pb.Status, error) {
	query := `
	select
      	case 
        	when id = $1 then true
      	else
        	false
      	end
    from
		reservations
    where
        id = $1 and deleted_at is null
	`
	var status pb.Status
	err := r.DB.QueryRowContext(ctx, query, id).Scan(&status.Successful)
	if err != nil {
		log.Println("reservation not found", err)
		return nil, err
	}

	return &status, nil
}

func (r *ReservationRepo) Order(ctx context.Context, reser *pb.ReservationOrders, reserEndTime time.Time) (*pb.ID, error) {
	id := uuid.NewString()

	err := redis.StoreOrders(ctx, id, reser, reserEndTime)
	if err != nil {
		return nil, err
	}

	return &pb.ID{Id: id}, nil
}

func (r *ReservationRepo) ChangeStatus(ctx context.Context, id, status string) error {
	query := `
	update
		reservations
	set
		status = $1
	where deleted_at is null and id = $2`

	_, err := r.DB.ExecContext(ctx, query, status, id)
	return err
}

func (r *ReservationRepo) FetchReservations(ctx context.Context, f *pb.Filter) (*pb.Reservations, error) {
	query := `
	select
		id, user_id, restaurant_id, reservation_time, status
	from
		reservations
	where
		deleted_at is null
	`
	count := 1
	var params []interface{}
	if f.UserId != "" {
		query += fmt.Sprintf(" and user_id = $%d", count)
		params = append(params, f.UserId)
		count++
	}
	if f.RestaurantId != "" {
		query += fmt.Sprintf(" and restaurant_id = $%d", count)
		params = append(params, f.RestaurantId)
		count++
	}
	if f.Limit > 0 {
		query += fmt.Sprintf(" LIMIT $%d", count)
		params = append(params, f.Limit)
		count++
	}
	if f.Offset > 0 {
		query += fmt.Sprintf(" OFFSET $%d", count)
		params = append(params, f.Offset)
		count++
	}

	rows, err := r.DB.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservations []*pb.ReservationInfo
	for rows.Next() {
		var reser pb.ReservationInfo

		err := rows.Scan(&reser.Id, &reser.UserId, &reser.RestaurantId, &reser.ReservationTime, &reser.Status)
		if err != nil {
			return nil, err
		}

		reservations = append(reservations, &reser)
	}

	return &pb.Reservations{Reservations: reservations}, nil
}
