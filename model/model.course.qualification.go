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
	query := `SELECT id,course_id,course_level,min_score,course_material_total,course_exam_total FROM course_qualification WHERE (id = $1 OR course_id = $2) AND flag_status=$3 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.ID, c.CourseID, STATUS_AVAILABLE).Scan(
		&one.ID, &one.CourseID, &one.CourseLevel, &one.MinScore, &one.CourseMaterialTotal, &one.CourseExamTotal,
	)
	if err != nil {
		return one, errors.Wrap(err, "error at query one course qualification")
	}
	return one, nil
}

func (c *CourseQualification) Update(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID
	query := `UPDATE course_qualification SET course_id=$1,course_level=$2,min_score=$3,course_material_total=$4,course_exam_total=$5 WHERE id = $6 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.CourseID, c.CourseLevel, c.MinScore, c.CourseMaterialTotal, c.CourseExamTotal, c.ID).Scan(
		&id,
	)
	if err != nil {
		return id, errors.Wrap(err, "error at update one of course qualification data")
	}
	return id, nil
}

func (c *CourseQualification) Delete(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID
	query := `UPDATE course_qualification SET flag_status=$1 WHERE id=$2 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), STATUS_DELETE, c.ID).Scan(
		&id,
	)
	if err != nil {
		return id, errors.Wrap(err, "error at delete one of course qualification data")
	}
	return id, nil
}
