package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type (
	CourseMaterial struct {
		ID            uuid.UUID `json:"id"`
		CourseID      uuid.UUID `json:"course_id"`
		MaterialIndex int32     `json:"material_index"`
		Title         string    `json:"title"`
	}

	CourseMaterialResponse struct {
		ID            uuid.UUID `json:"id"`
		CourseID      uuid.UUID `json:"course_id"`
		MaterialIndex int32     `json:"material_index"`
		Title         string    `json:"title"`
	}

	AllCourseMaterial struct {
		CourseID    uuid.UUID `json:"course_id"`
		SearchBy    string    `json:"search_by"`
		SearchValue string    `json:"search_value"`
		OrderBy     string    `json:"order_by"`
		OrderDir    string    `json:"order_dir"`
		Offset      int       `json:"offset"`
		Limit       int       `json:"limit"`
	}
)

func (c *CourseMaterial) Response() CourseMaterialResponse {
	return CourseMaterialResponse{
		ID:            c.ID,
		CourseID:      c.CourseID,
		MaterialIndex: c.MaterialIndex,
		Title:         c.Title,
	}
}

func (c *CourseMaterial) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	query := `INSERT INTO course_material (course_id,material_index,title) VALUES ($1,$2,$3) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.CourseID, c.MaterialIndex, c.Title).Scan(
		&c.ID,
	)
	if err != nil {
		return c.ID, errors.Wrap(err, "error at insert new course material")
	}

	return c.ID, nil
}

func (c *CourseMaterial) One(ctx context.Context, db *sql.DB) (*CourseMaterial, error) {
	one := &CourseMaterial{}
	query := `SELECT id,course_id,material_index,title FROM course_material WHERE id = $1 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.ID).Scan(
		&one.ID, &one.CourseID, &one.MaterialIndex, &one.Title,
	)
	if err != nil {
		return one, errors.Wrap(err, "error at query course material with id")
	}

	return one, nil
}

func (c *CourseMaterial) All(ctx context.Context, db *sql.DB, param AllCourseMaterial) ([]*CourseMaterial, error) {
	all := []*CourseMaterial{}
	query := `SELECT id,course_id,material_index,title FROM course_material WHERE %s LIKE $1 AND course_id = $2 ORDER BY %s %s OFFSET $3 LIMIT $4 `
	rows, err := db.QueryContext(ctx, fmt.Sprintf(query, param.SearchBy, param.OrderBy, param.OrderDir), "%"+param.SearchValue+"%", param.CourseID, param.Offset, param.Limit)
	if err != nil {
		return all, errors.Wrap(err, "error at query all course material")
	}

	defer rows.Close()

	for rows.Next() {
		one := &CourseMaterial{}
		err = rows.Scan(
			&one.ID, &one.CourseID, &one.MaterialIndex, &one.Title,
		)
		if err != nil {
			return all, errors.Wrap(err, "error at query all and scan one of course material data")
		}

		all = append(all, one)
	}

	return all, nil
}
