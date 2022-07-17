package service

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/RaevNur/simple-go-microservices-crud/api/parser"
	"github.com/RaevNur/simple-go-microservices-crud/internal/helper"
	"github.com/RaevNur/simple-go-microservices-crud/internal/models"
	"github.com/RaevNur/simple-go-microservices-crud/internal/pb"
	"github.com/RaevNur/simple-go-microservices-crud/internal/repository"
	"github.com/RaevNur/simple-go-microservices-crud/internal/repository/post"
)

type Service struct {
	pb.UnimplementedCrudServiceServer
	repo         post.IPostRepo
	parserClient parser.ParserServiceClient
}

func NewService(repo *repository.Repository, parserClient parser.ParserServiceClient) *Service {
	return &Service{
		repo:         repo.Post,
		parserClient: parserClient,
	}
}

func (s *Service) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	parserStatus, err := s.parserClient.ParseStatus()
	if err != nil {
		log.Println(err.Error())
		return &pb.GetResponse{Status: http.StatusBadGateway, Error: "Can't get response from parser server"}, nil
	} else if parserStatus.Status != http.StatusOK {
		return &pb.GetResponse{Status: parserStatus.Status, Error: parserStatus.Message}, nil
	}

	var dbError *helper.DbError
	post, err := s.repo.Get(ctx, req.Id)
	if err != nil {
		if errors.As(err, &dbError) {
			return &pb.GetResponse{Status: http.StatusBadRequest, Error: dbError.Error()}, nil
		}

		log.Println(err.Error())
		return &pb.GetResponse{Status: http.StatusInternalServerError, Error: http.StatusText(http.StatusInternalServerError)}, nil
	}

	return &pb.GetResponse{
		Status: http.StatusOK,
		Post: &pb.Post{
			Id:     post.Id,
			UserId: post.UserId,
			Title:  post.Title,
			Body:   post.Body,
		},
	}, nil
}

func (s *Service) GetPosts(ctx context.Context, req *pb.GetPostsRequest) (*pb.GetPostsResponse, error) {
	parserStatus, err := s.parserClient.ParseStatus()
	if err != nil {
		log.Println(err.Error())
		return &pb.GetPostsResponse{Status: http.StatusBadGateway, Message: "Can't get response from parser server"}, nil
	} else if parserStatus.Status != http.StatusOK {
		return &pb.GetPostsResponse{Status: parserStatus.Status, Message: parserStatus.Message}, nil
	}

	posts, err := s.repo.GetPosts(ctx, req.Id)
	if err != nil {
		log.Println(err.Error())
		return &pb.GetPostsResponse{Status: http.StatusInternalServerError, Message: http.StatusText(http.StatusInternalServerError)}, nil
	}

	respPosts := make([]*pb.Post, 0, len(posts))
	for i := range posts {
		respPosts = append(respPosts, &pb.Post{
			Id:     posts[i].Id,
			UserId: posts[i].UserId,
			Title:  posts[i].Title,
			Body:   posts[i].Body,
		})
	}

	return &pb.GetPostsResponse{
		Status: http.StatusOK,
		Post:   respPosts,
	}, nil
}

func (s *Service) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	parserStatus, err := s.parserClient.ParseStatus()
	if err != nil {
		log.Println(err.Error())
		return &pb.DeleteResponse{Status: http.StatusBadGateway, Message: "Can't get response from parser server"}, nil
	} else if parserStatus.Status != http.StatusOK {
		return &pb.DeleteResponse{Status: parserStatus.Status, Message: parserStatus.Message}, nil
	}

	var dbError *helper.DbError
	err = s.repo.Delete(ctx, req.Id)
	if err != nil {
		if errors.As(err, &dbError) {
			return &pb.DeleteResponse{Status: http.StatusBadRequest, Message: dbError.Error()}, nil
		}

		log.Println(err.Error())
		return &pb.DeleteResponse{Status: http.StatusInternalServerError, Message: http.StatusText(http.StatusInternalServerError)}, nil
	}

	return &pb.DeleteResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *Service) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	parserStatus, err := s.parserClient.ParseStatus()
	if err != nil {
		log.Println(err.Error())
		return &pb.UpdateResponse{Status: http.StatusBadGateway, Message: "Can't get response from parser server"}, nil
	} else if parserStatus.Status != http.StatusOK {
		return &pb.UpdateResponse{Status: parserStatus.Status, Message: parserStatus.Message}, nil
	}

	var dbError *helper.DbError
	err = s.repo.Update(ctx, &models.Post{
		Id:     req.Id,
		UserId: req.UserId,
		Title:  req.Title,
		Body:   req.Body,
	})
	if err != nil {
		if errors.As(err, &dbError) {
			return &pb.UpdateResponse{Status: http.StatusBadRequest, Message: dbError.Error()}, nil
		}

		log.Println(err.Error())
		return &pb.UpdateResponse{Status: http.StatusInternalServerError, Message: http.StatusText(http.StatusInternalServerError)}, nil
	}

	return &pb.UpdateResponse{
		Status: http.StatusOK,
	}, nil
}
