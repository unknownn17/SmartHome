package deviceclient

import (
	"api/protos/deviceproto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func UserClinet() deviceproto.DevicesClient {
	conn, err := grpc.NewClient("deviceservice:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := deviceproto.NewDevicesClient(conn)
	return client
}
