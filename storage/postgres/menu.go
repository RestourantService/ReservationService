package postgres

import (
	"context"
	"log"
	pb "reservation_service/genproto/menu"
)

func (m ReservationRepo) AddMeal(ctx context.Context, menu *pb.MealDetails) (*pb.ID, error) {
	
	id := pb.ID{}

	query := `
			INSERT INTO menu_items(restaurant_id, name, price) 
            VALUES($1, $2, $3) 
            RETURNING id
            `
    err := m.DB.QueryRowContext(ctx, query, menu.RestaurantId, menu.Name, menu.Price).Scan(&id)
	if err!= nil {
        log.Println("failed to insert meal", err)
        return nil, err
    }
	return &id, nil
}