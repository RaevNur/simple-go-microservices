package post

import (
	"context"
	"fmt"
	"strings"

	"github.com/RaevNur/simple-go-microservices-crud/internal/helper"
	"github.com/RaevNur/simple-go-microservices-crud/internal/models"
	"github.com/jackc/pgx/v4"
)

type IPostRepo interface {
	Get(ctx context.Context, id int64) (*models.Post, error)
	GetPosts(ctx context.Context, ids []int64) ([]*models.Post, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, post *models.Post) error
}

func (r *PostRepo) Get(ctx context.Context, id int64) (*models.Post, error) {
	query := `
		SELECT *
		FROM "posts"
		WHERE "id" = $1;
	`

	post := models.Post{}
	row := r.dbHandler.DB.QueryRow(ctx, query, id)

	err := row.Scan(&post.Id, &post.UserId, &post.Title, &post.Body)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, &helper.DbError{
				Title:       "post not found",
				Description: fmt.Sprintf("post with id %d not exists", id),
			}
		}

		return nil, fmt.Errorf("cant scan row: %w", err)
	}

	return &post, nil
}

func (r *PostRepo) GetPosts(ctx context.Context, ids []int64) ([]*models.Post, error) {
	params := make([]string, len(ids))
	args := make([]interface{}, len(ids))

	for i := 0; i < len(ids); i++ {
		args[i] = ids[i]
		params[i] = fmt.Sprintf("$%d", i+1)
	}

	query := fmt.Sprintf(`
		SELECT *
		FROM "posts"
		WHERE "id" IN (%s);
	`, strings.Join(params, ","))

	rows, err := r.dbHandler.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("cant execute query: %w", err)
	}

	defer rows.Close()

	posts := make([]*models.Post, 0, len(ids))
	for rows.Next() {
		post := models.Post{}

		err := rows.Scan(&post.Id, &post.UserId, &post.Title, &post.Body)
		if err != nil {
			return nil, fmt.Errorf("cant scan row: %w", err)
		}

		posts = append(posts, &post)
	}

	return posts, nil
}

func (r *PostRepo) Delete(ctx context.Context, id int64) error {
	query := `
		DELETE FROM "posts" 
		WHERE "id" = $1;
	`

	_, err := r.dbHandler.DB.Exec(ctx, query, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return &helper.DbError{
				Title:       "post not found",
				Description: fmt.Sprintf("post with id %d not exists", id),
			}
		}

		return fmt.Errorf("cant execute: %w", err)
	}

	return nil
}

func (r *PostRepo) Update(ctx context.Context, post *models.Post) error {
	query := `
		UPDATE "posts" 
		SET "user_id" = $1, "title" = $2, "body" = $3 
		WHERE "id" = $4;
	`

	_, err := r.dbHandler.DB.Exec(ctx, query, post.UserId, post.Title, post.Body, post.Id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return &helper.DbError{
				Title:       "post not found",
				Description: fmt.Sprintf("post with id %d not exists", post.Id),
			}
		}

		return fmt.Errorf("cant execute: %w", err)
	}

	return nil
}
