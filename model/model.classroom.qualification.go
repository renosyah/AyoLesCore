package model

import (
	"context"
	"database/sql"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

const (
	STATUS_NO_PROGRESS   = 0
	STATUS_PASS_EXAM     = 1
	STATUS_NOT_PASS_EXAM = 2
)

type (
	ClassRoomQualification struct {
		ClassRoomID         uuid.UUID            `json:"classroom_id"`
		CourseQualification *CourseQualification `json:"course_qualification"`
		TotalScore          int32                `json:"total_score"`
		Status              int32                `json:"status"`
	}

	ClassRoomQualificationResponse struct {
		ClassRoomID         uuid.UUID                   `json:"classroom_id"`
		CourseQualification CourseQualificationResponse `json:"course_qualification"`
		TotalScore          int32                       `json:"total_score"`
		Status              int32                       `json:"status"`
	}
)

func (c *ClassRoomQualification) Response() ClassRoomQualificationResponse {
	return ClassRoomQualificationResponse{
		ClassRoomID:         c.ClassRoomID,
		CourseQualification: c.CourseQualification.Response(),
		TotalScore:          c.TotalScore,
		Status:              c.Status,
	}
}

func (c *ClassRoomQualification) One(ctx context.Context, db *sql.DB) (*ClassRoomQualification, error) {
	var totalExamAnswered int32
	one := &ClassRoomQualification{
		ClassRoomID: c.ClassRoomID,
	}
	classroom, err := (&ClassRoom{ID: c.ClassRoomID}).One(ctx, db)
	if err != nil {
		return one, err
	}
	courseQ, err := (&CourseQualification{CourseID: classroom.Course.ID}).One(ctx, db)
	if err != nil {
		return one, err
	}
	one.CourseQualification = courseQ

	query := `SELECT
					SUM(1),
					SUM(CASE classroom_exam_progress.course_exam_answer_id
						WHEN course_exam_solution.course_exam_answer_id THEN 1
						ELSE 0
					END)
				FROM
					classroom_exam_progress
				INNER JOIN
					course_exam_solution
				ON
					course_exam_solution.course_exam_id = classroom_exam_progress.course_exam_id
				WHERE
					classroom_exam_progress.classroom_id = $1`

	err = db.QueryRowContext(ctx, fmt.Sprintf(query), c.ClassRoomID).Scan(
		&totalExamAnswered, &one.TotalScore,
	)
	if err != nil {
		return one, nil
	}

	if totalExamAnswered == one.CourseQualification.CourseExamTotal {
		if one.TotalScore >= one.CourseQualification.MinScore {
			one.Status = STATUS_PASS_EXAM
		} else {
			one.Status = STATUS_NOT_PASS_EXAM
		}
	} else {
		one.Status = STATUS_NO_PROGRESS
	}

	return one, nil
}
