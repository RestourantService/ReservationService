package pkg

import (
	"errors"
	"log"
	"reservation_service/config"
	pbp "reservation_service/genproto/payment"
	pbu "reservation_service/genproto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateUserClient(cfg config.Config) pbu.UserClient {
	conn, err := grpc.NewClient(cfg.Server.USER_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(errors.New("failed to connect to the address: " + err.Error()))
		return nil
	}

	return pbu.NewUserClient(conn)
}

func CreatePaymentClient(cfg config.Config) pbp.PaymentClient {
	conn, err := grpc.NewClient(cfg.Server.PAYMENT_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(errors.New("failed to connect to the address: " + err.Error()))
		return nil
	}

	return pbp.NewPaymentClient(conn)
}
