package useclient

import (
	"api/protos/userproto"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func UserClinet() (userproto.UserClient,context.Context, context.CancelFunc) {
	conn, err := grpc.NewClient("userservice:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	ctx,cancel:=context.WithTimeout(context.Background(),time.Minute*30)
	client := userproto.NewUserClient(conn)
	return client,ctx,cancel
}
