package admin

import (
	"log"

	adminpb "github.com/shakezidin/pkg/admin/adminpb"
	"github.com/shakezidin/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Configure) (adminpb.AdminServiceClient, error) {
	grpc, err := grpc.Dial(":"+cfg.GRPCADMINPORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error Dialing to grpc client: %s, ", cfg.GRPCADMINPORT)
		return nil, err
	}
	log.Printf("succesfully Connected to Booking Client at port: %v", cfg.GRPCADMINPORT)
	return adminpb.NewAdminServiceClient(grpc), nil
}
