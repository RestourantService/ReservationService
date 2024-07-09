package postgres

import (
	"context"
	"log"
	pb "reservation_service/genproto/reservation"
)

func (r *ReservationRepo)CreateReservation(ctx context.Context, reser *pb.ReservationDetails) (*pb.ID, error)  {

	id := pb.ID{}

	query := `
			INSERT INTO reservations(user_id, restaurant_id, reservation_time) 
			VALUES($1, $2, $3) 
			RETURNING id
			`
	err := r.DB.QueryRowContext(ctx, query, reser.UserId, reser.RestaurantId, reser.ReservationTime).Scan(&id)
	if err!= nil {
        log.Println("failed to insert reservation", err)
        return nil, err
    }

	return &id, nil
}

func (r *ReservationRepo) GetReservationById(ctx context.Context, id *pb.ID) (*pb.ReservationInfo, error) {
	reservation := &pb.ReservationInfo{}

    query := `
                SELECT id, user_id, restaurant_id, reservation_time
                FROM reservations
                WHERE id = $1
            `
    err := r.DB.QueryRowContext(ctx, query, id.Id).Scan(&reservation.Id, &reservation.UserId, &reservation.RestaurantId, &reservation.ReservationTime)
    if err!= nil {
        log.Println("reservation not found", err)
        return nil, err
    }
    return reservation, nil
}

func (r *ReservationRepo) UpdateReservation(ctx context.Context, reser *pb.ReservationInfo) error {
	query := `
	    	UPDATE reservations
			SET user_id = $1, reservation_id = $2, restaurant_time = $3, updated_at = NOW()
			WHERE id = $4
			`
	_, err := r.DB.ExecContext(ctx, query, &reser.UserId, &reser.RestaurantId, &reser.ReservationTime)
	if err!= nil {
        log.Println("failed to update reservation", err)
        return err
    }
	return nil
}

func (r *ReservationRepo) DeleteReservation(ctx context.Context, id *pb.ID) error {
	query := `
            UPDATE reservations
            SET deleted_at = NOW()
            WHERE id = $1
            `
    _, err := r.DB.ExecContext(ctx, query, id.Id)
    if err!= nil {
        log.Println("failed to delete reservation", err)
        return err
    }
    return nil
}

func (r *ReservationRepo) ValidateReservation(ctx context.Context, reser *pb.ReservationDetails) (*pb.ID, error) {

	id := pb.ID{}

	query := `
	        SELECT id
			FROM reservations
			WHERE user_id = $1 AND restaurant_id = $2 AND reservation_time = $3
			`
	err := r.DB.QueryRowContext(ctx, query, reser.UserId, reser.RestaurantId, reser.ReservationTime).Scan(&id.Id)
	if err!= nil {
        log.Println("reservation not found", err)
        return nil, err
    }
	return &id, nil
}

func (r *ReservationRepo) Order(ctx context.Context, reser *pb.ReservationOrders) (*pb.ID, error) {
	id := pb.ID{}

	query := `
			insert into reservation_orders (restaurant_id, menu_item_id, quantity)
			values ($1, $2, $3)
			RETURNING id
			`
	err := r.DB.QueryRowContext(ctx, query, reser.RestaurantId, reser.MenuItemId, reser.Quantity).Scan(&id)
	if err!= nil {
        log.Println("failed to insert order", err)
        return nil, err
    }
	return &id, nil
}

func (r *ReservationRepo) FetchReservations(ctx context.Context, filter *pb.Filter) (*pb.Reservations, error) {
	return nil, nil
}

