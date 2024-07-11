package test

import (
    "context"
    "database/sql"
    "fmt"
    "testing"
    "time"

    pb "reservation_service/genproto/restaurant"
    "reservation_service/storage/postgres"

    "github.com/stretchr/testify/assert"
)

func cleanupRestaurant(db *sql.DB, id string) {
    _, err := db.Exec(`DELETE FROM restaurants WHERE id = $1`, id)
    if err != nil {
        fmt.Printf("Failed to clean up test data: %v\n", err)
    }
}

func TestCreateRestaurant(t *testing.T) {
    db, err := ConnectDB()
    if err != nil {
        t.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()

    ctx := context.Background()
    repo := postgres.NewRestaurantRepo(db)

    uniqueName := fmt.Sprintf("Test Restaurant %d", time.Now().UnixNano())
    res := &pb.RestaurantDetails{
        Name:        uniqueName,
        Address:     "123 Test St",
        PhoneNumber: "123-456-7890",
        Description: "A test restaurant",
    }

    id, err := repo.CreateRestaurant(ctx, res)
    assert.NoError(t, err)
    assert.NotEmpty(t, id.Id)

    // Clean up test data
    cleanupRestaurant(db, id.Id)
}

func TestGetRestaurant(t *testing.T) {
    db, err := ConnectDB()
    if err != nil {
        t.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()

    ctx := context.Background()
    repo := postgres.NewRestaurantRepo(db)

    uniqueName := fmt.Sprintf("Test Restaurant %d", time.Now().UnixNano())
    res := &pb.RestaurantDetails{
        Name:        uniqueName,
        Address:     "123 Test St",
        PhoneNumber: "123-456-7890",
        Description: "A test restaurant",
    }

    id, err := repo.CreateRestaurant(ctx, res)
    assert.NoError(t, err)

    fetchedRes, err := repo.GetRestaurant(ctx, id)
    assert.NoError(t, err)
    assert.NotNil(t, fetchedRes)
    assert.Equal(t, res.Name, fetchedRes.Name)

    // Clean up test data
    cleanupRestaurant(db, id.Id)
}

func TestUpdateRestaurant(t *testing.T) {
    db, err := ConnectDB()
    if err != nil {
        t.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()

    ctx := context.Background()
    repo := postgres.NewRestaurantRepo(db)

    uniqueName := fmt.Sprintf("Test Restaurant %d", time.Now().UnixNano())
    res := &pb.RestaurantDetails{
        Name:        uniqueName,
        Address:     "123 Test St",
        PhoneNumber: "123-456-7890",
        Description: "A test restaurant",
    }

    id, err := repo.CreateRestaurant(ctx, res)
    assert.NoError(t, err)

    updatedRes := &pb.RestaurantInfo{
        Id:          id.Id,
        Name:        "Updated Restaurant",
        Address:     "456 Updated St",
        PhoneNumber: "987-654-3210",
        Description: "An updated test restaurant",
    }

    err = repo.UpdateRestaurant(ctx, updatedRes)
    assert.NoError(t, err)

    fetchedRes, err := repo.GetRestaurant(ctx, id)
    assert.NoError(t, err)
    assert.Equal(t, "Updated Restaurant", fetchedRes.Name)

    // Clean up test data
    cleanupRestaurant(db, id.Id)
}

func TestDeleteRestaurant(t *testing.T) {
    db, err := ConnectDB()
    if err != nil {
        t.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()

    ctx := context.Background()
    repo := postgres.NewRestaurantRepo(db)

    uniqueName := fmt.Sprintf("Test Restaurant %d", time.Now().UnixNano())
    res := &pb.RestaurantDetails{
        Name:        uniqueName,
        Address:     "123 Test St",
        PhoneNumber: "123-456-7890",
        Description: "A test restaurant",
    }

    id, err := repo.CreateRestaurant(ctx, res)
    assert.NoError(t, err)

    err = repo.DeleteRestaurant(ctx, id)
    assert.NoError(t, err)

    fetchedRes, err := repo.GetRestaurant(ctx, id)
    assert.Error(t, err)
    assert.Nil(t, fetchedRes)
}

func TestFetchRestaurants(t *testing.T) {
    db, err := ConnectDB()
    if err != nil {
        t.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()

    ctx := context.Background()
    repo := postgres.NewRestaurantRepo(db)

    uniqueName := fmt.Sprintf("Test Restaurant %d", time.Now().UnixNano())
    res := &pb.RestaurantDetails{
        Name:        uniqueName,
        Address:     "123 Test St",
        PhoneNumber: "123-456-7890",
        Description: "A test restaurant",
    }

    id, err := repo.CreateRestaurant(ctx, res)
    assert.NoError(t, err)

    pag := &pb.Pagination{
        Limit:  10,
        Offset: 0,
    }

    restaurants, err := repo.FetchRestaurants(ctx, pag)
    assert.NoError(t, err)
    assert.NotEmpty(t, restaurants)

    // Clean up test data
    cleanupRestaurant(db, id.Id)
}
