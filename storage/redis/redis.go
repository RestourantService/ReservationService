package redis

import (
	"context"
	pb "reservation_service/genproto/reservation"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

func ConnectDB() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return rdb
}

func StoreOrders(ctx context.Context, id string, reser *pb.ReservationOrder, reservationEndTime time.Time) error {
	rdb := ConnectDB()

	orderID := "reservation_order:" + id
	order := map[string]interface{}{
		"reservation_id": reser.Id,
		"menu_item_id":   reser.MenuItemId,
		"quantity":       reser.Quantity,
	}

	for k, v := range order {
		err := rdb.HSet(ctx, orderID, k, v).Err()
		if err != nil {
			return errors.Wrap(err, "failed to store order")
		}
	}

	expiration := time.Until(reservationEndTime)
	if expiration > 0 {
		err := rdb.Expire(ctx, orderID, expiration).Err()
		if err != nil {
			return errors.Wrap(err, "failed to set expiration time")
		}
	}

	return nil
}

func GetOrders(ctx context.Context, id string) (map[string]string, error) {
	rdb := ConnectDB()

	orderID := "reservation_order:" + id

	order, err := rdb.HGetAll(ctx, orderID).Result()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get order")
	}

	return order, nil
}
