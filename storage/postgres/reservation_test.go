package postgres

import (
	"context"
	"reflect"
	"testing"

	pb "reservation_service/genproto/reservation"
)

func TestCreateReservation(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}

	repo := NewReservationRepo(db)

	reser := &pb.ReservationDetails{
		UserId:          "789abcde-f012-3456-789a-bcdef0123456",
		RestaurantId:    "d3d29d70-1c75-40a7-a2e6-d1c2f3e3f3b3",
		ReservationTime: "2024-07-12 20:00:00",
	}

	id, err := repo.CreateReservation(context.Background(), reser)
	if err != nil {
		t.Error("failed to create reservation", id)
	}
}

func TestGetReservationByID(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}

	repo := NewReservationRepo(db)

	id := &pb.ID{Id: "550e8400-e29b-41d4-a716-446655440002"}

	reser, err := repo.GetReservationById(context.Background(), id)
	if err != nil {
		t.Fatal(err)
	}

	exp := &pb.ReservationInfo{
		Id:              "550e8400-e29b-41d4-a716-446655440002",
		UserId:          "123e4567-e89b-12d3-a456-426614174000",
		RestaurantId:    "710b962e-041c-11e1-9234-0123456789ab",
		ReservationTime: "2024-07-11 19:30:00",
		Status:          "confirmed",
	}

	if reflect.DeepEqual(reser, exp) {
		t.Error("expected reservation to be", exp, "but got", reser)
	}
}

func TestUpdateReservation(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}

	repo := NewReservationRepo(db)

	reser := &pb.ReservationInfo{
		Id:              "550e8400-e29b-41d4-a716-446655440006",
		UserId:          "a577b9de-0b98-43ce-9133-4f08a0f966c3",
		RestaurantId:    "a0a3d5ef-8e81-47be-b7f5-f23dfb4e9e44",
		ReservationTime: "2024-07-15 17:00:00",
		Status:          "cancelled",
	}
	err = repo.UpdateReservation(context.Background(), reser)
	if err!= nil {
        t.Error("failed to update reservation", err)
    }
}

func TestDeleteReservation(t *testing.T) {
	db, err := ConnectDB()
    if err!= nil {
        t.Fatal(err)
    }

    repo := NewReservationRepo(db)

    id := &pb.ID{Id: "550e8400-e29b-41d4-a716-446655440006"}

    err = repo.DeleteReservation(context.Background(), id)
    if err!= nil {
        t.Error("failed to delete reservation", err)
    }
}

func TestValidateReservation(t *testing.T) {
	db, err := ConnectDB()
    if err!= nil {
        t.Fatal(err)
    }

    repo := NewReservationRepo(db)

    id := &pb.ID{Id: "550e8400-e29b-41d4-a716-446655440004"}

    status, err := repo.ValidateReservation(context.Background(), id.Id)
    if err!= nil {
        t.Fatal(err)
    }

    exp := &pb.Status{
        Successful: true,
    }

    if!reflect.DeepEqual(status, exp) {
        t.Error("expected status to be", exp, "but got", status)
    }
}

func TestOrder(t *testing.T) {
	db, err := ConnectDB()
    if err!= nil {
        t.Fatal(err)
    }

    repo := NewReservationRepo(db)

    reser := &pb.ReservationOrders{
		Id:          "550e8400-e29b-41d4-a716-446655440007",
        MenuItemId:   "c9c98556-1627-4979-8075-175106816597",
        Quantity:     2,
    }

    id, err := repo.Order(context.Background(), reser)
    if err!= nil {
        t.Error("failed to order reservation", err)
    }
	exp := &pb.ID{
		Id: "550e8400-e29b-41d4-a716-446655440002",
	}
	if!reflect.DeepEqual(id, exp) {
        t.Error("expected id to be", exp, "but got", id)
    }
}

