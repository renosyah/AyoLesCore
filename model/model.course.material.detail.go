package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type (
	CourseMaterialDetail struct {
		ID               uuid.UUID `json:"id"`
		CourseMaterialID uuid.UUID `json:"course_material_id"`
		Position         int32     `json:"position"`
		Title            string    `json:"title"`
		TypeMaterial     int32     `json:"type_material"`
		Content          string    `json:"content"`
		ImageURL         string    `json:"image_url"`
	}

	CourseMaterialDetailResponse struct {
		ID               uuid.UUID `json:"id"`
		CourseMaterialID uuid.UUID `json:"course_material_id"`
		Position         int32     `json:"position"`
		Title            string    `json:"title"`
		TypeMaterial     int32     `json:"type_material"`
		Content          string    `json:"content"`
		ImageURL         string    `json:"image_url"`
	}

	AllCourseMaterialDetail struct {
		CourseMaterialID uuid.UUID `json:"course_material_id"`
		SearchBy         string    `json:"search_by"`
		SearchValue      string    `json:"search_value"`
		OrderBy          string    `json:"order_by"`
		OrderDir         string    `json:"order_dir"`
		Offset           int       `json:"offset"`
		Limit            int       `json:"limit"`
	}
)

func (c *CourseMaterialDetail) Response() CourseMaterialDetailResponse {
	return CourseMaterialDetailResponse{
		ID:               c.ID,
		CourseMaterialID: c.CourseMaterialID,
		Position:         c.Position,
		Title:            c.Title,
		TypeMaterial:     c.TypeMaterial,
		Content:          c.Content,
		ImageURL:         c.ImageURL,
	}
}

func (c *CourseMaterialDetail) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	query := `INSERT INTO course_material_detail (course_material_id,position,title,type_material,content,image_url) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.CourseMaterialID, c.Position, c.Title, c.TypeMaterial, c.Content, c.ImageURL).Scan(
		&c.ID,
	)
	if err != nil {
		return c.ID, errors.Wrap(err, "error at insert new course material detail")
	}

	return c.ID, nil
}

func (c *CourseMaterialDetail) All(ctx context.Context, db *sql.DB, param AllCourseMaterialDetail) ([]*CourseMaterialDetail, error) {
	all := []*CourseMaterialDetail{}
	query := `SELECT id,course_material_id,position,title,type_material,content,image_url FROM course_material_detail WHERE %s LIKE $1 AND course_material_id = $2 ORDER BY %s %s OFFSET $3 LIMIT $4 `
	rows, err := db.QueryContext(ctx, fmt.Sprintf(query, param.SearchBy, param.OrderBy, param.OrderDir), "%"+param.SearchValue+"%", param.CourseMaterialID, param.Offset, param.Limit)
	if err != nil {
		return all, errors.Wrap(err, "error at query all course material detail")
	}

	defer rows.Close()

	for rows.Next() {
		one := &CourseMaterialDetail{}
		err = rows.Scan(
			&one.ID, &one.CourseMaterialID, &one.Position, &one.Title, &one.TypeMaterial, &one.Content, &one.ImageURL,
		)
		if err != nil {
			return all, errors.Wrap(err, "error at query all and scan one of course material detail data")
		}

		all = append(all, one)
	}

	return all, nil
}
