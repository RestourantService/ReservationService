package test

import (
	"context"
	"testing"

	pb "reservation_service/genproto/menu"
	"reservation_service/storage/postgres"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestAddMeal(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	ctx := context.Background()
	repo := postgres.NewMenuRepo(db)

	addReq := &pb.MealDetails{
		RestaurantId: "550e8400-e29b-41d4-a716-446655440000",
		Name:         "Test Meal",
		Description:  "Test Description",
		Price:        10.0,
	}
	addResp, err := repo.AddMeal(ctx, addReq)
	assert.NoError(t, err)
	assert.NotNil(t, addResp)
	assert.NotEmpty(t, addResp.Id)

	// Clean up test data
	_, err = db.Exec(`DELETE FROM menu WHERE id = $1`, addResp.Id)
	if err != nil {
		t.Fatalf("Failed to clean up test data: %v", err)
	}
}

func TestGetMealByID(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	ctx := context.Background()
	repo := postgres.NewMenuRepo(db)

	addReq := &pb.MealDetails{
		RestaurantId: "550e8400-e29b-41d4-a716-446655440000",
		Name:         "Test Meal",
		Description:  "Test Description",
		Price:        10.0,
	}
	addResp, err := repo.AddMeal(ctx, addReq)
	assert.NoError(t, err)

	getReq := &pb.ID{Id: addResp.Id}
	getResp, err := repo.GetMealByID(ctx, getReq)
	assert.NoError(t, err)
	assert.NotNil(t, getResp)
	assert.Equal(t, "Test Meal", getResp.Name)

	_, err = db.Exec(`DELETE FROM menu WHERE id = $1`, addResp.Id)
	if err != nil {
		t.Fatalf("Failed to clean up test data: %v", err)
	}
}

func TestUpdateMeal(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	ctx := context.Background()
	repo := postgres.NewMenuRepo(db)

	addReq := &pb.MealDetails{
		RestaurantId: "550e8400-e29b-41d4-a716-446655440000",
		Name:         "Test Meal",
		Description:  "Test Description",
		Price:        10.0,
	}
	addResp, err := repo.AddMeal(ctx, addReq)
	assert.NoError(t, err)

	updateReq := &pb.MealInfo{
		Id:           addResp.Id,
		RestaurantId: "550e8400-e29b-41d4-a716-446655440000",
		Name:         "Updated Meal",
		Description:  "Updated Description",
		Price:        15.0,
	}
	err = repo.UpdateMeal(ctx, updateReq)
	assert.NoError(t, err)

	getResp, err := repo.GetMealByID(ctx, &pb.ID{Id: addResp.Id})
	assert.NoError(t, err)
	assert.Equal(t, "Updated Meal", getResp.Name)

	_, err = db.Exec(`DELETE FROM menu WHERE id = $1`, addResp.Id)
	if err != nil {
		t.Fatalf("Failed to clean up test data: %v", err)
	}
}

func TestDeleteMeal(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	ctx := context.Background()
	repo := postgres.NewMenuRepo(db)

	addReq := &pb.MealDetails{
		RestaurantId: "550e8400-e29b-41d4-a716-446655440000",
		Name:         "Test Meal",
		Description:  "Test Description",
		Price:        10.0,
	}
	addResp, err := repo.AddMeal(ctx, addReq)
	assert.NoError(t, err)

	deleteReq := &pb.ID{Id: addResp.Id}
	err = repo.DeleteMeal(ctx, deleteReq)
	assert.NoError(t, err)

	getResp, err := repo.GetMealByID(ctx, deleteReq)
	assert.Error(t, err)
	assert.Nil(t, getResp)
}

