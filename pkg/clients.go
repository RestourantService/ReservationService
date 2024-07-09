package pkg

import (
	"errors"
	"log"
	"reservation_service/config"
	pb "reservation_service/genproto/payment"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreatePaymentClient(cfg config.Config) pb.PaymentClient {
	conn, err := grpc.NewClient(cfg.Server.PAYMENT_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(errors.New("failed to connect to the address: " + err.Error()))
		return nil
	}

	return pb.NewPaymentClient(conn)
}
