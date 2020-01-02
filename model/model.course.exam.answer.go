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

func (c CourseExamAnswer) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	query := `INSERT INTO course_exam_answer (course_exam_id,type_answer,label,text,image_url) VALUES ($1,$2,$3,$4,$5) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.CourseExamID, c.TypeAnswer, c.Label, c.Text, c.ImageURL).Scan(
		&c.ID,
	)
	if err != nil {
		return c.ID, errors.Wrap(err, "error at insert new course exam answer")
	}
	return c.ID, nil
}

func (c CourseExamAnswer) One(ctx context.Context, db *sql.DB) (*CourseExamAnswer, error) {
	one := &CourseExamAnswer{}
	query := `SELECT id,course_exam_id,type_answer,label,text,image_url FROM course_exam_answer WHERE id = $1 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.ID).Scan(
		&one.ID, &one.CourseExamID, &one.TypeAnswer, &one.Label, &one.Text, &one.ImageURL,
	)
	if err != nil {
		return one, errors.Wrap(err, "error at query one course exam answer")
	}
	return one, nil
}

func (c CourseExamAnswer) All(ctx context.Context, db *sql.DB, param AllCourseExamAnswer) ([]*CourseExamAnswer, error) {
	all := []*CourseExamAnswer{}
	query := `SELECT id,course_exam_id,type_answer,label,text,image_url FROM course_exam_answer WHERE %s LIKE $1 AND course_exam_id = $2 ORDER BY %s %s OFFSET $3 LIMIT $4`
	rows, err := db.QueryContext(ctx, fmt.Sprintf(query, param.SearchBy, param.OrderBy, param.OrderDir), "%"+param.SearchValue+"%", param.CourseExamID, param.Offset, param.Limit)
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

func (c CourseExamAnswer) AllById(ctx context.Context, db *sql.DB, LimitAnswer int) ([]*CourseExamAnswer, error) {
	all := []*CourseExamAnswer{}
	query := `SELECT id,course_exam_id,type_answer,label,text,image_url FROM course_exam_answer WHERE course_exam_id = $1 ORDER BY random() LIMIT $2`
	rows, err := db.QueryContext(ctx, fmt.Sprintf(query), c.CourseExamID, LimitAnswer)
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
