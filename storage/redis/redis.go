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

func StoreOrders(ctx context.Context, id string, reser *pb.ReservationOrders, reservationEndTime time.Time) error {
	rdb := ConnectDB()

	orderID := "reservation_order:" + id
	order := map[string]interface{}{
		"reservation_id": reser.Id,
		"order": reser.Order,
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
