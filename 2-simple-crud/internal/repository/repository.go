package repository

import (
	"github.com/RaevNur/simple-go-microservices-crud/internal/db"
	"github.com/RaevNur/simple-go-microservices-crud/internal/repository/post"
)

type Repository struct {
	Post post.IPostRepo
}

func NewRepo(dbHandler db.DbHandler) *Repository {
	return &Repository{
		Post: post.NewPostRepo(dbHandler),
	}
}
