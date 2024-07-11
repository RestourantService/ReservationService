package test

import (
	"context"
	"testing"
	"time"

	pb "reservation_service/genproto/reservation"
	"reservation_service/storage/postgres"

	"github.com/stretchr/testify/assert"
)

func TestCreateReservation(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	ctx := context.Background()
	repo := postgres.NewReservationRepo(db)

	reser := &pb.ReservationDetails{
		UserId:          "550e8400-e29b-41d4-a716-446655440000",
		RestaurantId:    "550e8400-e29b-41d4-a716-446655440000",
		ReservationTime: "2024-07-10 18:00:00",
	}

	id, err := repo.CreateReservation(ctx, reser)
	assert.NoError(t, err)
	assert.NotEmpty(t, id.Id)

	_, err = db.Exec(`DELETE FROM reservations WHERE id = $1`, id.Id)
	if err != nil {
		t.Fatalf("Failed to clean up test data: %v", err)
	}
}

func TestGetReservationById(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	ctx := context.Background()
	repo := postgres.NewReservationRepo(db)

	reser := &pb.ReservationDetails{
		UserId:          "550e8400-e29b-41d4-a716-446655440000",
		RestaurantId:    "550e8400-e29b-41d4-a716-446655440000",
		ReservationTime: "2024-07-10 18:00:00",
	}

	id, err := repo.CreateReservation(ctx, reser)
	assert.NoError(t, err)

	fetchedReser, err := repo.GetReservationById(ctx, &pb.ID{Id: id.Id})
	assert.NoError(t, err)
	assert.NotNil(t, fetchedReser)
	assert.Equal(t, reser.UserId, fetchedReser.UserId)

	_, err = db.Exec(`DELETE FROM reservations WHERE id = $1`, id.Id)
	if err != nil {
		t.Fatalf("Failed to clean up test data: %v", err)
	}
}

func TestDeleteReservation(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	ctx := context.Background()
	repo := postgres.NewReservationRepo(db)

	reser := &pb.ReservationDetails{
		UserId:          "550e8400-e29b-41d4-a716-446655440000",
		RestaurantId:    "550e8400-e29b-41d4-a716-446655440000",
		ReservationTime: "2024-07-10 18:00:00",
	}

	id, err := repo.CreateReservation(ctx, reser)
	assert.NoError(t, err)

	err = repo.DeleteReservation(ctx, &pb.ID{Id: id.Id})
	assert.NoError(t, err)

	fetchedReser, err := repo.GetReservationById(ctx, &pb.ID{Id: id.Id})
	assert.Error(t, err)
	assert.Nil(t, fetchedReser)
}

func TestValidateReservation(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	ctx := context.Background()
	repo := postgres.NewReservationRepo(db)

	reser := &pb.ReservationDetails{
		UserId:          "550e8400-e29b-41d4-a716-446655440000",
		RestaurantId:    "550e8400-e29b-41d4-a716-446655440000",
		ReservationTime: "2024-07-10 18:00:00",
	}

	id, err := repo.CreateReservation(ctx, reser)
	assert.NoError(t, err)

	status, err := repo.ValidateReservation(ctx, id.Id)
	assert.NoError(t, err)
	assert.True(t, status.Successful)

	_, err = db.Exec(`DELETE FROM reservations WHERE id = $1`, id.Id)
	if err != nil {
		t.Fatalf("Failed to clean up test data: %v", err)
	}
}

func TestChangeStatus(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	ctx := context.Background()
	repo := postgres.NewReservationRepo(db)

	reser := &pb.ReservationDetails{
		UserId:          "550e8400-e29b-41d4-a716-446655440000",
		RestaurantId:    "550e8400-e29b-41d4-a716-446655440000",
		ReservationTime: "2024-07-10 18:00:00",
	}

	id, err := repo.CreateReservation(ctx, reser)
	assert.NoError(t, err)

	err = repo.ChangeStatus(ctx, id.Id, "confirmed")
	assert.NoError(t, err)

	fetchedReser, err := repo.GetReservationById(ctx, &pb.ID{Id: id.Id})
	assert.NoError(t, err)
	assert.Equal(t, "confirmed", fetchedReser.Status)

	_, err = db.Exec(`DELETE FROM reservations WHERE id = $1`, id.Id)
	if err != nil {
		t.Fatalf("Failed to clean up test data: %v", err)
	}
}

func TestFetchReservations(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	ctx := context.Background()
	repo := postgres.NewReservationRepo(db)

	reser := &pb.ReservationDetails{
		UserId:          "550e8400-e29b-41d4-a716-446655440000",
		RestaurantId:    "550e8400-e29b-41d4-a716-446655440000",
		ReservationTime: "2024-07-10 18:00:00",
	}

	id, err := repo.CreateReservation(ctx, reser)
	assert.NoError(t, err)

	filter := &pb.Filter{
		UserId:       "550e8400-e29b-41d4-a716-446655440000",
		RestaurantId: "550e8400-e29b-41d4-a716-446655440000",
		Limit:        10,
		Offset:       0,
	}

	reservations, err := repo.FetchReservations(ctx, filter)
	assert.NoError(t, err)
	assert.NotEmpty(t, reservations.Reservations)

	_, err = db.Exec(`DELETE FROM reservations WHERE id = $1`, id.Id)
	if err != nil {
		t.Fatalf("Failed to clean up test data: %v", err)
	}
}

// func TestOrder(t *testing.T) {
// 	db, err := ConnectDB()
// 	if err != nil {
// 		t.Fatalf("Failed to connect to database: %v", err)
// 	}
// 	defer db.Close()

// 	ctx := context.Background()
// 	repo := postgres.NewReservationRepo(db)

// 	reser := &pb.ReservationOrders{
// 		Id:         "550e8400-e29b-41d4-a716-446655440000",
// 		MenuItemId: "550e8400-e29b-41d4-a716-446655441001",
// 		Quantity:   2,
// 	}

// 	reserEndTime := time.Now().Add(1 * time.Hour)

// 	id, err := repo.Order(ctx, reser, reserEndTime)
// 	assert.NoError(t, err)
// 	assert.NotEmpty(t, id.Id)

// 	_, err = db.Exec(`DELETE FROM reservations WHERE id = $1`, id.Id)
// 	if err != nil {
// 		t.Fatalf("Failed to clean up test data: %v", err)
// 	}
// }

func TestUpdateReservation(t *testing.T) {
    db, err := ConnectDB()
    if err != nil {
        t.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()

    ctx := context.Background()
    repo := postgres.NewReservationRepo(db)

    reser := &pb.ReservationDetails{
        UserId:        "550e8400-e29b-41d4-a716-446655440000",
        RestaurantId:  "550e8400-e29b-41d4-a716-446655440000",
        ReservationTime: "2024-07-10 18:00:00",
    }

    id, err := repo.CreateReservation(ctx, reser)
    assert.NoError(t, err)

    updatedReser := &pb.ReservationInfo{
        Id:            id.Id,
        UserId:        "550e8400-e29b-41d4-a716-446655440000",
        RestaurantId:  "550e8400-e29b-41d4-a716-446655440000",
        ReservationTime: "2024-07-11 19:00:00",
        Status:        "confirmed",
    }

    err = repo.UpdateReservation(ctx, updatedReser)
    assert.NoError(t, err)

    fetchedReser, err := repo.GetReservationById(ctx, &pb.ID{Id: id.Id})
    assert.NoError(t, err)

    expectedTime, err := time.Parse("2006-01-02 15:04:05", "2024-07-11 19:00:00")
    assert.NoError(t, err)

    actualTime, err := time.Parse(time.RFC3339, fetchedReser.ReservationTime)
    assert.NoError(t, err)

    assert.True(t, expectedTime.Equal(actualTime), "expected: %v, actual: %v", expectedTime, actualTime)

    _, err = db.Exec(`DELETE FROM reservations WHERE id = $1`, id.Id)
    if err != nil {
        t.Fatalf("Failed to clean up test data: %v", err)
    }
}
