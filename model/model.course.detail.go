package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type (
	CourseDetail struct {
		ID              uuid.UUID `json:"id"`
		CourseID        uuid.UUID `json:"course_id"`
		OverviewText    string    `json:"overview_text"`
		DescriptionText string    `json:"description_text"`
		ImageURL        string    `json:"image_url"`
	}

	CourseDetailResponse struct {
		ID              uuid.UUID `json:"id"`
		CourseID        uuid.UUID `json:"course_id"`
		OverviewText    string    `json:"overview_text"`
		DescriptionText string    `json:"description_text"`
		ImageURL        string    `json:"image_url"`
	}

	AllCourseDetail struct {
		CourseID    uuid.UUID `json:"course_id"`
		SearchBy    string    `json:"search_by"`
		SearchValue string    `json:"search_value"`
		OrderBy     string    `json:"order_by"`
		OrderDir    string    `json:"order_dir"`
		Offset      int       `json:"offset"`
		Limit       int       `json:"limit"`
	}
)

func (a AllCourseDetail) IsWithCourseID() string {
	var emptyID uuid.UUID
	if a.CourseID == emptyID {
		return ""
	}
	return fmt.Sprintf(`AND course_id = '%s'`, a.CourseID)
}

func (c *CourseDetail) Response() CourseDetailResponse {
	return CourseDetailResponse{
		ID:              c.ID,
		CourseID:        c.CourseID,
		OverviewText:    c.OverviewText,
		DescriptionText: c.DescriptionText,
		ImageURL:        c.ImageURL,
	}
}

func (c *CourseDetail) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	query := `INSERT INTO course_detail (course_id,overview_text,description_text,image_url) VALUES ($1,$2,$3,$4) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.CourseID, c.OverviewText, c.DescriptionText, c.ImageURL).Scan(
		&c.ID,
	)
	if err != nil {
		return c.ID, errors.Wrap(err, "error at insert new detail course")
	}

	return c.ID, nil
}

func (c *CourseDetail) One(ctx context.Context, db *sql.DB) (*CourseDetail, error) {
	one := &CourseDetail{}
	query := `SELECT id,course_id,overview_text,description_text,image_url FROM course_detail WHERE id = $1 AND flag_status = $2 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.ID, STATUS_AVAILABLE).Scan(
		&one.ID, &one.CourseID, &one.OverviewText, &one.DescriptionText, &one.ImageURL,
	)
	if err != nil {
		return one, errors.Wrap(err, "error at query course detail with id")
	}

	return one, nil
}

func (c *CourseDetail) All(ctx context.Context, db *sql.DB, param AllCourseDetail) ([]*CourseDetail, error) {
	all := []*CourseDetail{}
	query := `SELECT id,course_id,overview_text,description_text,image_url FROM course_detail WHERE %s LIKE $1 AND flag_status = $2 %s ORDER BY %s %s OFFSET $3 LIMIT $4 `
	rows, err := db.QueryContext(ctx, fmt.Sprintf(query, param.SearchBy, param.IsWithCourseID(), param.OrderBy, param.OrderDir), "%"+param.SearchValue+"%", STATUS_AVAILABLE, param.Offset, param.Limit)
	if err != nil {
		return all, errors.Wrap(err, "error at query all course detail")
	}

	defer rows.Close()

	for rows.Next() {
		one := &CourseDetail{}
		err = rows.Scan(
			&one.ID, &one.CourseID, &one.OverviewText, &one.DescriptionText, &one.ImageURL,
		)
		if err != nil {
			return all, errors.Wrap(err, "error at query all and scan one of course detail data")
		}

		all = append(all, one)
	}

	return all, nil
}

func (c *CourseDetail) AllByCourseID(ctx context.Context, db *sql.DB) ([]CourseDetail, error) {
	all := []CourseDetail{}
	query := `SELECT id,course_id,overview_text,description_text,image_url FROM course_detail WHERE course_id = $1 AND flag_status = $2 LIMIT 3`
	rows, err := db.QueryContext(ctx, fmt.Sprintf(query), c.CourseID, STATUS_AVAILABLE)
	if err != nil {
		return all, errors.Wrap(err, "error at query all course detail")
	}

	defer rows.Close()

	for rows.Next() {
		one := CourseDetail{}
		err = rows.Scan(
			&one.ID, &one.CourseID, &one.OverviewText, &one.DescriptionText, &one.ImageURL,
		)
		if err != nil {
			return all, errors.Wrap(err, "error at query all and scan one of course detail data")
		}

		all = append(all, one)
	}

	return all, nil
}

func (c *CourseDetail) Update(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID
	query := `UPDATE course_detail SET course_id=$1,overview_text=$2,description_text=$3,image_url=$4 WHERE id=$5 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.CourseID, c.OverviewText, c.DescriptionText, c.ImageURL, c.ID).Scan(
		&id,
	)
	if err != nil {
		return id, errors.Wrap(err, "error at update one of course detail")
	}
	return id, nil
}

func (c *CourseDetail) Delete(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID
	query := `UPDATE course_detail SET flag_status=$1 WHERE id=$2 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), STATUS_DELETE, c.ID).Scan(
		&id,
	)
	if err != nil {
		return id, errors.Wrap(err, "error at delete one of course detail")
	}
	return id, nil
}
