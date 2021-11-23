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

	query := `SELECT id FROM classroom_exam_progress WHERE classroom_id = $1 AND course_exam_id = $2 AND flag_status = $3 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.ClassroomID, c.CourseExamID, STATUS_AVAILABLE).Scan(
		&c.ID,
	)
	if err != nil && err != sql.ErrNoRows {
		return c.ID, errors.Wrap(err, "error at query one course exam progress")
	}

	if c.ID != emptyUUID {
		return c.ID, errors.Wrap(err, "course exam progress already inserted")
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

	query := `SELECT id,classroom_id,course_exam_id,course_exam_answer_id FROM classroom_exam_progress WHERE id = $1 AND flag_status = $2 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.ID, STATUS_AVAILABLE).Scan(
		&one.ID, &one.ClassroomID, &one.CourseExamID, &one.CourseExamAnswerID,
	)
	if err != nil {
		return one, errors.Wrap(err, "error at query one course exam progress")
	}
	return one, nil
}

func (c *ClassRoomExamProgress) All(ctx context.Context, db *sql.DB, param AllClassRoomExamProgress) ([]*ClassRoomExamProgress, error) {
	all := []*ClassRoomExamProgress{}

	query := `SELECT id,classroom_id,course_exam_id,course_exam_answer_id FROM classroom_exam_progress WHERE classroom_id = $1 AND flag_status = $2 ORDER BY %s %s OFFSET $3 LIMIT $4`
	rows, err := db.QueryContext(ctx, fmt.Sprintf(query, param.OrderBy, param.OrderDir), param.ClassroomID, STATUS_AVAILABLE, param.Offset, param.Limit)
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

func (c *ClassRoomExamProgress) Update(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID
	query := `UPDATE classroom_exam_progress SET classroom_id=$1,course_exam_id=$2,course_exam_answer_id=$3 WHERE id = $4 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.ClassroomID, c.CourseExamID, c.CourseExamAnswerID, c.ID).Scan(
		&id,
	)
	if err != nil {
		return id, errors.Wrap(err, "error at update one of classroom exam progress data")
	}
	return id, nil
}

func (c *ClassRoomExamProgress) Delete(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID
	query := `UPDATE classroom_exam_progress SET flag_status = $1 WHERE id = $2 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), STATUS_DELETE, c.ID).Scan(
		&id,
	)
	if err != nil {
		return id, errors.Wrap(err, "error at delete one of classroom exam progress data")
	}
	return id, nil
}
