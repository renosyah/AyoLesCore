package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type (
	CategoryModel struct {
		ID       uuid.UUID `json:"id"`
		Name     string    `json:"name"`
		ImageURL string    `json:"image_url"`
	}

	CategoryModelResponse struct {
		ID       uuid.UUID `json:"id"`
		Name     string    `json:"name"`
		ImageURL string    `json:"image_url"`
	}

	AllCategory struct {
		SearchBy    string `json:"search_by"`
		SearchValue string `json:"search_value"`
		OrderBy     string `json:"order_by"`
		OrderDir    string `json:"order_dir"`
		Offset      int    `json:"offset"`
		Limit       int    `json:"limit"`
	}
)

func (c *CategoryModel) Response() CategoryModelResponse {
	return CategoryModelResponse{
		ID:       c.ID,
		Name:     c.Name,
		ImageURL: c.ImageURL,
	}
}

func (c *CategoryModel) Add(ctx context.Context, db *sql.DB) (*CategoryModel, error) {
	query := `INSERT INTO course_category (name,image_url) VALUES ($1,$2) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.Name, c.ImageURL).Scan(
		&c.ID,
	)
	if err != nil {
		return c, errors.Wrap(err, "error at insert new category")
	}

	return c, nil
}

func (c *CategoryModel) One(ctx context.Context, db *sql.DB) (*CategoryModel, error) {
	one := &CategoryModel{}
	query := `SELECT id,name,image_url FROM course_category WHERE id = $1 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.ID).Scan(
		&one.ID, &one.Name, &one.ImageURL,
	)
	if err != nil {
		return one, errors.Wrap(err, "error at query one category")
	}

	return one, nil
}

func (c *CategoryModel) All(ctx context.Context, db *sql.DB, param AllCategory) ([]*CategoryModel, error) {
	all := []*CategoryModel{}
	query := `SELECT id,name,image_url FROM course_category WHERE %s LIKE $1 ORDER BY %s %s OFFSET $2 LIMIT $3 `
	rows, err := db.QueryContext(ctx, fmt.Sprintf(query, param.SearchBy, param.OrderBy, param.OrderDir), "%"+param.SearchValue+"%", param.Offset, param.Limit)
	if err != nil {
		return all, errors.Wrap(err, "error at query all category")
	}

	defer rows.Close()

	for rows.Next() {
		one := &CategoryModel{}
		err = rows.Scan(
			&one.ID, &one.Name, &one.ImageURL,
		)
		if err != nil {
			return all, errors.Wrap(err, "error at scan one of category")
		}
		all = append(all, one)
	}

	return all, nil
}
