package user

import (
	"log"

	"github.com/shakezidin/pkg/config"
	userpb "github.com/shakezidin/pkg/user/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Configure) (userpb.UserServiceClient, error) {
	grpc, err := grpc.Dial(":"+cfg.GRPCUSERPORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error Dialing to grpc client: %s, ", cfg.GRPCUSERPORT)
		return nil, err
	}
	log.Printf("succesfully Connected to Booking Client at port: %v", cfg.GRPCUSERPORT)
	return userpb.NewUserServiceClient(grpc), nil
}
