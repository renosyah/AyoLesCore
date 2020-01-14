package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type (
	CourseExamAnswer struct {
		ID           uuid.UUID `json:"id"`
		CourseExamID uuid.UUID `json:"course_exam_id"`
		TypeAnswer   int32     `json:"type_answer"`
		Label        string    `json:"label"`
		Text         string    `json:"text"`
		ImageURL     string    `json:"image_url"`
	}

	CourseExamAnswerResponse struct {
		ID           uuid.UUID `json:"id"`
		CourseExamID uuid.UUID `json:"course_exam_id"`
		TypeAnswer   int32     `json:"type_answer"`
		Label        string    `json:"label"`
		Text         string    `json:"text"`
		ImageURL     string    `json:"image_url"`
	}

	AllCourseExamAnswer struct {
		CourseExamID uuid.UUID `json:"course_exam_id"`
		SearchBy     string    `json:"search_by"`
		SearchValue  string    `json:"search_value"`
		OrderBy      string    `json:"order_by"`
		OrderDir     string    `json:"order_dir"`
		Offset       int       `json:"offset"`
		Limit        int       `json:"limit"`
	}
)

func (c *CourseExamAnswer) Response() CourseExamAnswerResponse {
	return CourseExamAnswerResponse{
		ID:           c.ID,
		CourseExamID: c.CourseExamID,
		TypeAnswer:   c.TypeAnswer,
		Label:        c.Label,
		Text:         c.Text,
		ImageURL:     c.ImageURL,
	}
}

func (c *CourseExamAnswer) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	query := `INSERT INTO course_exam_answer (course_exam_id,type_answer,label,text,image_url) VALUES ($1,$2,$3,$4,$5) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.CourseExamID, c.TypeAnswer, c.Label, c.Text, c.ImageURL).Scan(
		&c.ID,
	)
	if err != nil {
		return c.ID, errors.Wrap(err, "error at insert new course exam answer")
	}
	return c.ID, nil
}

func (c *CourseExamAnswer) One(ctx context.Context, db *sql.DB) (*CourseExamAnswer, error) {
	one := &CourseExamAnswer{}
	query := `SELECT id,course_exam_id,type_answer,label,text,image_url FROM course_exam_answer WHERE id = $1 AND flag_status=$2 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.ID, STATUS_AVAILABLE).Scan(
		&one.ID, &one.CourseExamID, &one.TypeAnswer, &one.Label, &one.Text, &one.ImageURL,
	)
	if err != nil {
		return one, errors.Wrap(err, "error at query one course exam answer")
	}
	return one, nil
}

func (c *CourseExamAnswer) All(ctx context.Context, db *sql.DB, param AllCourseExamAnswer) ([]*CourseExamAnswer, error) {
	all := []*CourseExamAnswer{}
	query := `SELECT id,course_exam_id,type_answer,label,text,image_url FROM course_exam_answer WHERE %s LIKE $1 AND course_exam_id = $2 AND flag_status = $3 ORDER BY %s %s OFFSET $4 LIMIT $5`
	rows, err := db.QueryContext(ctx, fmt.Sprintf(query, param.SearchBy, param.OrderBy, param.OrderDir), "%"+param.SearchValue+"%", param.CourseExamID, STATUS_AVAILABLE, param.Offset, param.Limit)
	if err != nil {
		return all, errors.Wrap(err, "error at query one course exam answer")
	}

	defer rows.Close()

	for rows.Next() {
		one := &CourseExamAnswer{}
		err = rows.Scan(
			&one.ID, &one.CourseExamID, &one.TypeAnswer, &one.Label, &one.Text, &one.ImageURL,
		)
		if err != nil {
			return all, errors.Wrap(err, "error at query all and scan one of course exam data")
		}

		all = append(all, one)
	}

	return all, nil
}

func (c *CourseExamAnswer) AllById(ctx context.Context, db *sql.DB, LimitAnswer int) ([]*CourseExamAnswer, error) {
	all := []*CourseExamAnswer{}
	query := `SELECT id,course_exam_id,type_answer,label,text,image_url FROM course_exam_answer WHERE course_exam_id = $1 AND flag_status = $2 ORDER BY create_at ASC LIMIT $3`
	rows, err := db.QueryContext(ctx, fmt.Sprintf(query), c.CourseExamID, STATUS_AVAILABLE, LimitAnswer)
	if err != nil {
		return all, errors.Wrap(err, "error at query one course exam answer")
	}

	defer rows.Close()

	for rows.Next() {
		one := &CourseExamAnswer{}
		err = rows.Scan(
			&one.ID, &one.CourseExamID, &one.TypeAnswer, &one.Label, &one.Text, &one.ImageURL,
		)
		if err != nil {
			return all, errors.Wrap(err, "error at query all and scan one of course exam data")
		}

		all = append(all, one)
	}

	return all, nil
}

func (c *CourseExamAnswer) Update(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID
	query := `UPDATE course_exam_answer SET course_exam_id=$1,type_answer=$2,label=$3,text=$4,image_url=$5 WHERE id = $6 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.CourseExamID, c.TypeAnswer, c.Label, c.Text, c.ImageURL, c.ID).Scan(
		&id,
	)
	if err != nil {
		return id, errors.Wrap(err, "error at update one of course exam answer data")
	}
	return id, nil
}

func (c *CourseExamAnswer) Delete(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID
	query := `UPDATE course_exam_answer SET flag_status=$1 WHERE id=$2 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), STATUS_DELETE, c.ID).Scan(
		&id,
	)
	if err != nil {
		return id, errors.Wrap(err, "error at delete one of course exam answer data")
	}
	return id, nil
}
