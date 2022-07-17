package post

import "github.com/RaevNur/simple-go-microservices-parser/internal/db"

type PostRepo struct {
	dbHandler db.DbHandler
}

func NewPostRepo(dbHandler db.DbHandler) *PostRepo {
	return &PostRepo{dbHandler}
}
