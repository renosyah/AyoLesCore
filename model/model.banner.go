package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type (
	Banner struct {
		ID       uuid.UUID `json:"id"`
		Title    string    `json:"title"`
		Content  string    `json:"content"`
		ImageURL string    `json:"image_url"`
	}

	BannerResponse struct {
		ID       uuid.UUID `json:"id"`
		Title    string    `json:"title"`
		Content  string    `json:"content"`
		ImageURL string    `json:"image_url"`
	}

	AllBanner struct {
		SearchBy    string `json:"search_by"`
		SearchValue string `json:"search_value"`
		OrderBy     string `json:"order_by"`
		OrderDir    string `json:"order_dir"`
		Offset      int    `json:"offset"`
		Limit       int    `json:"limit"`
	}
)

func (c *Banner) Response() BannerResponse {
	return BannerResponse{
		ID:       c.ID,
		Title:    c.Title,
		Content:  c.Content,
		ImageURL: c.ImageURL,
	}
}

func (b *Banner) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	query := `INSERT INTO banner (title,content,image_url) VALUES ($1,$2,$3) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), b.Title, b.Content, b.ImageURL).Scan(
		&b.ID,
	)
	if err != nil {
		return b.ID, errors.Wrap(err, "error at add banner")
	}

	return b.ID, nil
}

func (b *Banner) One(ctx context.Context, db *sql.DB) (*Banner, error) {
	one := &Banner{}
	query := `SELECT id,title,content,image_url FROM banner WHERE id = $1 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), b.ID).Scan(
		&one.ID, &one.Title, &one.Content, &one.ImageURL,
	)
	if err != nil {
		return b, errors.Wrap(err, "error at query one banner")
	}

	return one, nil
}

func (b *Banner) All(ctx context.Context, db *sql.DB, param AllBanner) ([]*Banner, error) {
	all := []*Banner{}
	query := `SELECT id,title,content,image_url FROM banner WHERE %s LIKE $1 ORDER BY %s %s OFFSET $2 LIMIT $3 `
	rows, err := db.QueryContext(ctx, fmt.Sprintf(query, param.SearchBy, param.OrderBy, param.OrderDir), "%"+param.SearchValue+"%", param.Offset, param.Limit)
	if err != nil {
		return all, errors.Wrap(err, "error at query all banner")
	}

	defer rows.Close()

	for rows.Next() {
		one := &Banner{}
		err = rows.Scan(
			&one.ID, &one.Title, &one.Content, &one.ImageURL,
		)
		if err != nil {
			return all, errors.Wrap(err, "error at scan one of banner")
		}

		all = append(all, one)
	}

	return all, nil
}
