package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type (
	ClassRoomExamResult struct {
		CourseExamID    uuid.UUID           `json:"course_exam_id"`
		CourseID        uuid.UUID           `json:"course_id"`
		ClassRoomID     uuid.UUID           `json:"classroom_id"`
		StudentAnswerID uuid.UUID           `json:"student_answer_id"`
		ValidAnswerID   uuid.UUID           `json:"valid_answer_id"`
		TypeExam        int32               `json:"type_exam"`
		ExamIndex       int32               `json:"exam_index"`
		Text            string              `json:"text"`
		ImageURL        string              `json:"image_url"`
		Answers         []*CourseExamAnswer `json:"answers"`
	}

	ClassRoomExamResultResponse struct {
		CourseExamID    uuid.UUID                  `json:"course_exam_id"`
		CourseID        uuid.UUID                  `json:"course_id"`
		ClassRoomID     uuid.UUID                  `json:"classroom_id"`
		StudentAnswerID uuid.UUID                  `json:"student_answer_id"`
		ValidAnswerID   uuid.UUID                  `json:"valid_answer_id"`
		TypeExam        int32                      `json:"type_exam"`
		ExamIndex       int32                      `json:"exam_index"`
		Text            string                     `json:"text"`
		ImageURL        string                     `json:"image_url"`
		Answers         []CourseExamAnswerResponse `json:"answers"`
	}

	AllClassRoomExamResult struct {
		ClassRoomID uuid.UUID `json:"classroom_id"`
		SearchBy    string    `json:"search_by"`
		SearchValue string    `json:"search_value"`
		OrderBy     string    `json:"order_by"`
		OrderDir    string    `json:"order_dir"`
		Offset      int       `json:"offset"`
		Limit       int       `json:"limit"`
		LimitAnswer int       `json:"limit_answer"`
	}
)

func (c *ClassRoomExamResult) Response() ClassRoomExamResultResponse {
	answers := []CourseExamAnswerResponse{}
	for _, one := range c.Answers {
		answers = append(answers, one.Response())
	}
	return ClassRoomExamResultResponse{
		CourseExamID:    c.CourseExamID,
		CourseID:        c.CourseID,
		ClassRoomID:     c.ClassRoomID,
		StudentAnswerID: c.StudentAnswerID,
		ValidAnswerID:   c.ValidAnswerID,
		TypeExam:        c.TypeExam,
		ExamIndex:       c.ExamIndex,
		Text:            c.Text,
		ImageURL:        c.ImageURL,
		Answers:         answers,
	}
}

func (c *ClassRoomExamResult) One(ctx context.Context, db *sql.DB, LimitAnswer int) (*ClassRoomExamResult, error) {
	one := &ClassRoomExamResult{}
	query := `SELECT 
				course_exam.id,
				course_exam.course_id,
				classroom_exam_progress.classroom_id,
				classroom_exam_progress.course_exam_answer_id,
				course_exam_solution.course_exam_answer_id,
				course_exam.type_exam,
				course_exam.exam_index,
				course_exam.text,
				course_exam.image_url
			FROM 
				classroom_exam_progress
			INNER JOIN
				course_exam
			ON
				course_exam.id = classroom_exam_progress.course_exam_id
			INNER JOIN
				course_exam_solution
			ON
				course_exam_solution.course_exam_id = course_exam.id
			WHERE
				course_exam.id = $1
			AND
				classroom_exam_progress.classroom_id = $2
			LIMIT 1`

	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.CourseExamID, c.ClassRoomID).Scan(
		&one.CourseExamID, &one.CourseID, &one.ClassRoomID, &one.StudentAnswerID, &one.ValidAnswerID, &one.TypeExam, &one.ExamIndex, &one.Text, &one.ImageURL,
	)
	if err != nil {
		return one, errors.Wrap(err, "error at query one course exam progress")
	}

	one.Answers, err = (&CourseExamAnswer{CourseExamID: one.CourseExamID}).AllById(ctx, db, LimitAnswer)
	if err != nil {
		return one, errors.Wrap(err, "error at query one course exam for answers")
	}

	return one, nil
}

func (c *ClassRoomExamResult) All(ctx context.Context, db *sql.DB, param AllClassRoomExamResult) ([]*ClassRoomExamResult, error) {
	all := []*ClassRoomExamResult{}

	query := `SELECT 
				course_exam.id,
				course_exam.course_id,
				classroom_exam_progress.classroom_id,
				classroom_exam_progress.course_exam_answer_id,
				course_exam_solution.course_exam_answer_id,
				course_exam.type_exam,
				course_exam.exam_index,
				course_exam.text,
				course_exam.image_url
			FROM 
				classroom_exam_progress
			INNER JOIN
				course_exam
			ON
				course_exam.id = classroom_exam_progress.course_exam_id
			INNER JOIN
				course_exam_solution
			ON
				course_exam_solution.course_exam_id = course_exam.id
			WHERE
				%s LIKE $1
			AND
				classroom_exam_progress.classroom_id = $2
			ORDER BY 
				%s %s 
			OFFSET $3 
			LIMIT $4`

	rows, err := db.QueryContext(ctx, fmt.Sprintf(query, param.SearchBy, param.OrderBy, param.OrderDir), "%"+param.SearchValue+"%", param.ClassRoomID, param.Offset, param.Limit)
	if err != nil {
		return all, errors.Wrap(err, "error at query all course exam result")
	}

	defer rows.Close()

	for rows.Next() {
		one := &ClassRoomExamResult{}
		err = rows.Scan(
			&one.CourseExamID, &one.CourseID, &one.ClassRoomID, &one.StudentAnswerID, &one.ValidAnswerID, &one.TypeExam, &one.ExamIndex, &one.Text, &one.ImageURL,
		)
		if err != nil {
			return all, errors.Wrap(err, "error at query all and scan one of course examp result data")
		}

		one.Answers, err = (&CourseExamAnswer{CourseExamID: one.CourseExamID}).AllById(ctx, db, param.LimitAnswer)
		if err != nil {
			return all, errors.Wrap(err, "error at query all and scan one of course examp result  data for answers")
		}

		all = append(all, one)
	}

	return all, nil
}


// ITS DOESNOT HAVE TABLE
// THIS MODEL VALUE RESULT FROM
// QUERY JOIN
// NO UPDATE
// NO DELETE