package post

import (
	"github.com/RaevNur/simple-go-microservices-crud/internal/db"
)

type PostRepo struct {
	dbHandler db.DbHandler
}

func NewPostRepo(dbHandler db.DbHandler) *PostRepo {
	return &PostRepo{dbHandler}
}
