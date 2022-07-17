package main

import (
	"log"
	"net"
	"os"

	"github.com/RaevNur/simple-go-microservices-crud/api/parser"
	"github.com/RaevNur/simple-go-microservices-crud/configs"
	"github.com/RaevNur/simple-go-microservices-crud/internal/db"
	"github.com/RaevNur/simple-go-microservices-crud/internal/pb"
	"github.com/RaevNur/simple-go-microservices-crud/internal/repository"
	"github.com/RaevNur/simple-go-microservices-crud/internal/service"
	"google.golang.org/grpc"
)

func main() {
	c, err := configs.LoadConfig()
	if err != nil {
		log.Println("Can't load configs:", err)
		os.Exit(1)
	}

	dbHandler, err := db.InitPostgresDB(c.DbUrl)
	if err != nil {
		log.Println("Can't connect DB:", err)
		os.Exit(2)
	}

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Println("Failed to listening:", err)
		os.Exit(3)
	}

	parserServ := parser.InitParserServiceClient(c.ParserSrvUrl)
	repo := repository.NewRepo(dbHandler)
	serv := service.NewService(repo, parserServ)

	log.Println("CRUD service on", c.Port)

	grpcServer := grpc.NewServer()

	pb.RegisterCrudServiceServer(grpcServer, serv)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Println("Failed to serve:", err)
		os.Exit(4)
	}
}
