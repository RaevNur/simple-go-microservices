package parser

import (
	"log"

	"github.com/RaevNur/simple-go-microservices-api-gateway/api/parser/pb"
	"github.com/RaevNur/simple-go-microservices-api-gateway/configs"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.ParserServiceClient
}

func InitServiceClinet(c *configs.Config) pb.ParserServiceClient {
	// WithInsecure because no ssl
	cc, err := grpc.Dial(c.ParserSrvUrl, grpc.WithInsecure())
	if err != nil {
		log.Println("Parser service is not working:", err)
	}

	return pb.NewParserServiceClient(cc)
}
