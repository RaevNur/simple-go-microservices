package repository

import (
	"github.com/RaevNur/simple-go-microservices-parser/internal/db"
	"github.com/RaevNur/simple-go-microservices-parser/internal/repository/post"
)

type Repository struct {
	Post post.IPostRepo
}

func NewRepo(dbHandler db.DbHandler) *Repository {
	return &Repository{
		Post: post.NewPostRepo(dbHandler),
	}
}
