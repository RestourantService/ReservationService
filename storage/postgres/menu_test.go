package postgres

import (
	"context"
	"reflect"
	pb "reservation_service/genproto/menu"
	"testing"
)

func TestAddMeal(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := NewMenuRepo(db)
	test := pb.MealDetails{
		RestaurantId: "550e8400-e29b-41d4-a716-446655440000",
		Name:         "Margherita Pizza",
		Description:  "Classic Italian pizza with tomato sauce, mozzarella, and basil",
		Price:        12.99,
	}

	id, err := repo.AddMeal(context.Background(), &test)
	if err != nil {
		t.Error("failed to addMeal", id)
	}
}

func TestGetMealByID(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := NewMenuRepo(db)
	id := &pb.ID{Id: "1ed38daf-3f38-4406-b365-fc56433dffd0"}

	meal, err := repo.GetMealByID(context.Background(), id)
	if err != nil {
		t.Fatal(err)
	}

	exp := &pb.MealInfo{
		Id:           "1ed38daf-3f38-4406-b365-fc56433dffd0",
		RestaurantId: "550e8400-e29b-41d4-a716-446655440000",
		Name:         "Margherita Pizza",
		Description:  "Classic Italian pizza with tomato sauce, mozzarella, and basil",
		Price:        12.99,
	}

	if !reflect.DeepEqual(exp, meal) {
		t.Error("expected meal to be", exp, "but got", meal)
	}
}

func TestUpdateMeal(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := NewMenuRepo(db)
	test := &pb.MealInfo{
		Id:           "550e8400-e29b-41d4-a716-446655441001",
		RestaurantId: "550e8400-e29b-41d4-a716-446655440000",
		Name:         "Updated Margherita Pizza",
		Description:  "Classic Italian pizza with tomato sauce, mozzarella, and basil",
		Price:        12.99,
	}

	err = repo.UpdateMeal(context.Background(), test)
	if err != nil {
		t.Error("failed to update meal", err)
	}
}

func TestDeleteMeal(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := NewMenuRepo(db)
	id := &pb.ID{Id: "550e8400-e29b-41d4-a716-446655441001"}

	err = repo.DeleteMeal(context.Background(), id)
	if err != nil {
		t.Error("failed to delete meal", err)
	}
}

func TestGetAllMeals(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := NewMenuRepo(db)
	filter := &pb.Filter{
		RestaurantId: "550e8400-e29b-41d4-a716-446655440000",
		Limit:        1,
		Offset:       0,
	}

	meals, err := repo.GetAllMeals(context.Background(), filter)
	if err != nil {
		t.Fatal(err)
	}

	meal := pb.MealInfo{
		Id:           "550e8400-e29b-41d4-a716-446655441010",
		RestaurantId: "550e8400-e29b-41d4-a716-446655440000",
		Name:         "Tiramisu",
		Description:  "Classic Italian dessert made with layers of coffee-soaked ladyfingers and mascarpone cheese",
		Price:        8.99,
	}
	exp := pb.Meals{
		Meals: []*pb.MealInfo{
			&meal,
		},
	}

	if !reflect.DeepEqual(&exp, meals) {
		t.Error("expected filter to be", &exp, "but got", meals)
	}

}
