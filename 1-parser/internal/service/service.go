package service

import (
	"context"
	"net/http"
	"sync"

	"github.com/RaevNur/simple-go-microservices-parser/internal/pb"
	"github.com/RaevNur/simple-go-microservices-parser/internal/repository"
	"github.com/RaevNur/simple-go-microservices-parser/internal/repository/post"
	"github.com/golang/protobuf/ptypes/empty"
)

type Service struct {
	pb.UnimplementedParserServiceServer
	repo post.IPostRepo
	mu   sync.RWMutex
	av   bool
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		repo: repo.Post,
		av:   true,
	}
}

func (s *Service) ParsePosts(context.Context, *empty.Empty) (*pb.ParsePostsResponce, error) {
	s.mu.Lock()
	if s.av {
		s.av = false
		s.mu.Unlock()
	} else {
		defer s.mu.Unlock()
		return &pb.ParsePostsResponce{
			Status:  http.StatusAccepted,
			Message: "Parser is working",
		}, nil
	}

	go s.startParse()

	return &pb.ParsePostsResponce{
		Status:  http.StatusAccepted,
		Message: "Parser is working",
	}, nil
}

func (s *Service) ParseStatus(context.Context, *empty.Empty) (*pb.ParseStatusResponce, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.av {
		return &pb.ParseStatusResponce{
			Status:  http.StatusOK,
			Message: "Parser is available",
		}, nil
	} else {
		return &pb.ParseStatusResponce{
			Status:  http.StatusAccepted,
			Message: "Parser is working",
		}, nil
	}
}
