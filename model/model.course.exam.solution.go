package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type (
	CourseExamSolution struct {
		ID                 uuid.UUID `json:"id"`
		CourseExamID       uuid.UUID `json:"course_exam_id"`
		CourseExamAnswerID uuid.UUID `json:"course_exam_answer_id"`
	}

	CourseExamSolutionResponse struct {
		ID                 uuid.UUID `json:"id"`
		CourseExamID       uuid.UUID `json:"course_exam_id"`
		CourseExamAnswerID uuid.UUID `json:"course_exam_answer_id"`
	}
	AllCourseExamSolution struct {
		CourseExamID uuid.UUID `json:"course_exam_id"`
		OrderBy      string    `json:"order_by"`
		OrderDir     string    `json:"order_dir"`
		Offset       int       `json:"offset"`
		Limit        int       `json:"limit"`
	}
)

func (c *CourseExamSolution) Response() CourseExamSolutionResponse {
	return CourseExamSolutionResponse{
		ID:                 c.ID,
		CourseExamID:       c.CourseExamID,
		CourseExamAnswerID: c.CourseExamAnswerID,
	}
}

func (c *CourseExamSolution) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	query := `INSERT INTO course_exam_solution (course_exam_id,course_exam_answer_id) VALUES ($1,$2) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.CourseExamID, c.CourseExamAnswerID).Scan(
		&c.ID,
	)
	if err != nil {
		return c.ID, errors.Wrap(err, "error at insert new exam solution")
	}
	return c.ID, nil
}

func (c *CourseExamSolution) One(ctx context.Context, db *sql.DB) (*CourseExamSolution, error) {
	one := &CourseExamSolution{}
	query := `SELECT id,course_exam_id,course_exam_answer_id FROM WHERE course_exam_solution id=$1 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.ID).Scan(
		&one.ID, &one.CourseExamID, &one.CourseExamAnswerID,
	)
	if err != nil {
		return one, errors.Wrap(err, "error at query one exam solution")
	}
	return one, nil
}

func (c *CourseExamSolution) All(ctx context.Context, db *sql.DB, param AllCourseExamSolution) ([]*CourseExamSolution, error) {
	all := []*CourseExamSolution{}
	query := `SELECT id,course_exam_id,course_exam_answer_id FROM course_exam_solution WHERE course_exam_id = $1 ORDER BY %s %s OFFSET $2 LIMIT $3`
	rows, err := db.QueryContext(ctx, fmt.Sprintf(query, param.OrderBy, param.OrderDir), param.CourseExamID, param.Offset, param.Limit)
	if err != nil {
		return all, errors.Wrap(err, "error at query all exam solution")
	}

	defer rows.Close()

	for rows.Next() {
		one := &CourseExamSolution{}
		err = rows.Scan(
			&one.ID, &one.CourseExamID, &one.CourseExamAnswerID,
		)
		if err != nil {
			return all, errors.Wrap(err, "error at query one of exam solution data")
		}
		all = append(all, one)
	}

	return all, nil
}
