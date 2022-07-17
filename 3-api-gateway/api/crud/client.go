package crud

import (
	"log"

	"github.com/RaevNur/simple-go-microservices-api-gateway/api/crud/pb"
	"github.com/RaevNur/simple-go-microservices-api-gateway/configs"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.CrudServiceClient
}

func InitServiceClinet(c *configs.Config) pb.CrudServiceClient {
	// WithInsecure because no ssl
	cc, err := grpc.Dial(c.CrudSrvUrl, grpc.WithInsecure())
	if err != nil {
		log.Println("Crud service is not working:", err)
	}

	return pb.NewCrudServiceClient(cc)
}
