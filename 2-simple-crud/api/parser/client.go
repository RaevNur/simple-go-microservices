package parser

import (
	"context"
	"log"

	"github.com/RaevNur/simple-go-microservices-crud/api/parser/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ParserServiceClient struct {
	Client pb.ParserServiceClient
}

func InitParserServiceClient(url string) ParserServiceClient {
	// WithInsecure because no ssl
	cc, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Println("Parser service is not working:", err)
	}

	client := ParserServiceClient{
		Client: pb.NewParserServiceClient(cc),
	}

	return client
}

func (c *ParserServiceClient) ParseStatus() (*pb.ParseStatusResponce, error) {
	return c.Client.ParseStatus(context.Background(), &emptypb.Empty{})
}
