package user

import (
	"device/protos/userproto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func UserClinet() userproto.UserClient {
	conn, err := grpc.NewClient("userservice:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := userproto.NewUserClient(conn)
	return client
}
