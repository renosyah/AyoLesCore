package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type (
	CourseQualification struct {
		ID                  uuid.UUID `json:"id"`
		CourseID            uuid.UUID `json:"course_id"`
		CourseLevel         string    `json:"course_level"`
		MinScore            int32     `json:"min_score"`
		CourseMaterialTotal int32     `json:"course_material_total"`
		CourseExamTotal     int32     `json:"course_exam_total"`
	}
	CourseQualificationResponse struct {
		ID                  uuid.UUID `json:"id"`
		CourseID            uuid.UUID `json:"course_id"`
		CourseLevel         string    `json:"course_level"`
		MinScore            int32     `json:"min_score"`
		CourseMaterialTotal int32     `json:"course_material_total"`
		CourseExamTotal     int32     `json:"course_exam_total"`
	}
)

func (c *CourseQualification) Response() CourseQualificationResponse {
	return CourseQualificationResponse{
		ID:                  c.ID,
		CourseID:            c.CourseID,
		CourseLevel:         c.CourseLevel,
		MinScore:            c.MinScore,
		CourseMaterialTotal: c.CourseMaterialTotal,
		CourseExamTotal:     c.CourseExamTotal,
	}
}

func (c *CourseQualification) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	query := `INSERT INTO course_qualification (course_id,course_level,min_score,course_material_total,course_exam_total) VALUES ($1,$2,$3,$4,$5) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.CourseID, c.CourseLevel, c.MinScore, c.CourseMaterialTotal, c.CourseExamTotal).Scan(
		&c.ID,
	)
	if err != nil {
		return c.ID, errors.Wrap(err, "error at insert new course qualification")
	}
	return c.ID, nil
}

func (c *CourseQualification) One(ctx context.Context, db *sql.DB) (*CourseQualification, error) {
	one := &CourseQualification{}
	query := `SELECT id,course_id,course_level,min_score,course_material_total,course_exam_total FROM course_qualification WHERE id = $1 OR course_id = $2 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.ID, c.CourseID).Scan(
		&one.ID, &one.CourseID, &one.CourseLevel, &one.MinScore, &one.CourseMaterialTotal, &one.CourseExamTotal,
	)
	if err != nil {
		return one, errors.Wrap(err, "error at query one course qualification")
	}
	return one, nil
}
