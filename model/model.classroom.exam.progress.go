package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type (
	ClassRoomExamProgress struct {
		ID                 uuid.UUID `json:"id"`
		ClassroomID        uuid.UUID `json:"classroom_id"`
		CourseExamID       uuid.UUID `json:"course_exam_id"`
		CourseExamAnswerID uuid.UUID `json:"course_exam_answer_id"`
	}
	ClassRoomExamProgressResponse struct {
		ID                 uuid.UUID `json:"id"`
		ClassroomID        uuid.UUID `json:"classroom_id"`
		CourseExamID       uuid.UUID `json:"course_exam_id"`
		CourseExamAnswerID uuid.UUID `json:"course_exam_answer_id"`
	}
	AllClassRoomExamProgress struct {
		ClassroomID uuid.UUID `json:"classroom_id"`
		OrderBy     string    `json:"order_by"`
		OrderDir    string    `json:"order_dir"`
		Offset      int       `json:"offset"`
		Limit       int       `json:"limit"`
	}
)

func (c *ClassRoomExamProgress) Response() ClassRoomExamProgressResponse {
	return ClassRoomExamProgressResponse{
		ID:                 c.ID,
		ClassroomID:        c.ClassroomID,
		CourseExamID:       c.CourseExamID,
		CourseExamAnswerID: c.CourseExamAnswerID,
	}
}

func (c *ClassRoomExamProgress) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var emptyUUID uuid.UUID

	query := `SELECT id FROM classroom_exam_progress WHERE classroom_id = $1 AND course_exam_id = $2 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.ClassroomID, c.CourseExamID).Scan(
		&c.ID,
	)
	if err != nil && err != sql.ErrNoRows {
		return c.ID, errors.Wrap(err, "error at query one course exam progress")
	}

	if c.ID != emptyUUID {
		return c.ID, nil
	}

	query = `INSERT INTO classroom_exam_progress (classroom_id,course_exam_id,course_exam_answer_id) VALUES ($1,$2,$3) RETURNING id`
	err = db.QueryRowContext(ctx, fmt.Sprintf(query), c.ClassroomID, c.CourseExamID, c.CourseExamAnswerID).Scan(
		&c.ID,
	)
	if err != nil {
		return c.ID, errors.Wrap(err, "error at insert new course exam progress")
	}
	return c.ID, nil
}

func (c *ClassRoomExamProgress) One(ctx context.Context, db *sql.DB) (*ClassRoomExamProgress, error) {
	one := &ClassRoomExamProgress{}

	query := `SELECT id,classroom_id,course_exam_id,course_exam_answer_id FROM classroom_exam_progress WHERE id = $1 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.ID).Scan(
		&one.ID, &one.ClassroomID, &one.CourseExamID, &one.CourseExamAnswerID,
	)
	if err != nil {
		return one, errors.Wrap(err, "error at query one course exam progress")
	}
	return one, nil
}

func (c *ClassRoomExamProgress) All(ctx context.Context, db *sql.DB, param AllClassRoomExamProgress) ([]*ClassRoomExamProgress, error) {
	all := []*ClassRoomExamProgress{}

	query := `SELECT id,classroom_id,course_exam_id,course_exam_answer_id FROM classroom_exam_progress WHERE classroom_id = $1 ORDER BY %s %s OFFSET $2 LIMIT $3`
	rows, err := db.QueryContext(ctx, fmt.Sprintf(query, param.OrderBy, param.OrderDir), param.ClassroomID, param.Offset, param.Limit)
	if err != nil {
		return all, errors.Wrap(err, "error at query all course exam progress")
	}

	defer rows.Close()

	for rows.Next() {
		one := &ClassRoomExamProgress{}
		err = rows.Scan(
			&one.ID, &one.ClassroomID, &one.CourseExamID, &one.CourseExamAnswerID,
		)
		if err != nil {
			return all, errors.Wrap(err, "error at query one of course exam progress data")
		}

		all = append(all, one)
	}

	return all, nil
}

func (c *ClassRoomExamProgress) Delete(ctx context.Context, db *sql.DB) error {
	query := `DELETE FROM classroom_exam_progress WHERE classroom_id = $1`
	_, err := db.ExecContext(ctx, fmt.Sprintf(query), c.ClassroomID)
	if err != nil {
		return errors.Wrap(err, "error at delete all course exam progress")
	}
	return nil
}
