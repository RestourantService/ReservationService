package main

import (
	"log"
	"net"
	"reservation_service/config"
	pbM "reservation_service/genproto/menu"
	pbReser "reservation_service/genproto/reservation"
	pbRest "reservation_service/genproto/restaurant"
	"reservation_service/service"
	"reservation_service/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()
	lis, err := net.Listen("tcp", cfg.Server.RESERVATION_PORT)
	if err != nil {
		log.Fatalf("error while listening: %v", err)
	}
	defer lis.Close()

	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatalf("error while connecting to database: %v", err)
	}
	defer db.Close()

	restaurantService := service.NewRestaurantService(db)
	reservationService := service.NewReservationService(db)
	menuService := service.NewMenuService(db)
	server := grpc.NewServer()
	pbRest.RegisterRestaurantServer(server, restaurantService)
	pbReser.RegisterReservationServer(server, reservationService)
	pbM.RegisterMenuServer(server, menuService)

	log.Printf("server listening at %v", lis.Addr())
	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("error while serving: %v", err)
	}
}
