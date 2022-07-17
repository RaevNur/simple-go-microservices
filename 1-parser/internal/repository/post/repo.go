package post

import (
	"context"
	"fmt"

	"github.com/RaevNur/simple-go-microservices-parser/internal/models"
)

type IPostRepo interface {
	Add(ctx context.Context, post *models.Post) error
}

func (r *PostRepo) Add(ctx context.Context, post *models.Post) error {
	query := `
		INSERT INTO "posts" (
			"user_id",
			"title",
			"body"
		) VALUES ($1, $2, $3) RETURNING "id";
	`

	row := r.dbHandler.DB.QueryRow(ctx, query, post.UserId, post.Title, post.Body)

	err := row.Scan(&post.Id)
	if err != nil {
		return fmt.Errorf("post adding error: %w", err)
	}

	return nil
}
