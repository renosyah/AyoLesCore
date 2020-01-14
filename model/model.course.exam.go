package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type (
	CourseExam struct {
		ID        uuid.UUID           `json:"id"`
		CourseID  uuid.UUID           `json:"course_id"`
		TypeExam  int32               `json:"type_exam"`
		ExamIndex int32               `json:"exam_index"`
		Text      string              `json:"text"`
		ImageURL  string              `json:"image_url"`
		Answers   []*CourseExamAnswer `json:"answers"`
	}

	CourseExamResponse struct {
		ID        uuid.UUID                  `json:"id"`
		CourseID  uuid.UUID                  `json:"course_id"`
		TypeExam  int32                      `json:"type_exam"`
		ExamIndex int32                      `json:"exam_index"`
		Text      string                     `json:"text"`
		ImageURL  string                     `json:"image_url"`
		Answers   []CourseExamAnswerResponse `json:"answers"`
	}

	AllCourseExam struct {
		CourseID    uuid.UUID `json:"course_id"`
		SearchBy    string    `json:"search_by"`
		SearchValue string    `json:"search_value"`
		OrderBy     string    `json:"order_by"`
		OrderDir    string    `json:"order_dir"`
		Offset      int       `json:"offset"`
		Limit       int       `json:"limit"`
		LimitAnswer int       `json:"limit_answer"`
	}
)

func (c *CourseExam) Response() CourseExamResponse {
	answers := []CourseExamAnswerResponse{}
	for _, one := range c.Answers {
		answers = append(answers, one.Response())
	}
	return CourseExamResponse{
		ID:        c.ID,
		CourseID:  c.CourseID,
		TypeExam:  c.TypeExam,
		ExamIndex: c.ExamIndex,
		Text:      c.Text,
		ImageURL:  c.ImageURL,
		Answers:   answers,
	}
}

func (c *CourseExam) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	query := `INSERT INTO course_exam (course_id,type_exam,exam_index,text,image_url) VALUES ($1,$2,$3,$4,$5) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.CourseID, c.TypeExam, c.ExamIndex, c.Text, c.ImageURL).Scan(
		&c.ID,
	)
	if err != nil {
		return c.ID, errors.Wrap(err, "error at insert new course exam")
	}
	return c.ID, nil
}

func (c *CourseExam) One(ctx context.Context, db *sql.DB, LimitAnswer int) (*CourseExam, error) {
	one := &CourseExam{}
	query := `SELECT id,course_id,type_exam,exam_index,text,image_url FROM course_exam WHERE id = $1 AND flag_status=$2 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.ID, STATUS_AVAILABLE).Scan(
		&one.ID, &one.CourseID, &one.TypeExam, &one.ExamIndex, &one.Text, &one.ImageURL,
	)
	if err != nil {
		return one, errors.Wrap(err, "error at query one course exam")
	}

	one.Answers, err = (&CourseExamAnswer{CourseExamID: one.ID}).AllById(ctx, db, LimitAnswer)
	if err != nil {
		return one, errors.Wrap(err, "error at query one course exam for answers")
	}

	return one, nil
}

func (c *CourseExam) All(ctx context.Context, db *sql.DB, param AllCourseExam) ([]*CourseExam, error) {
	all := []*CourseExam{}

	query := `SELECT id,course_id,type_exam,exam_index,text,image_url FROM course_exam WHERE %s LIKE $1 AND course_id = $2 AND flag_status=$3 ORDER BY %s %s OFFSET $4 LIMIT $5`
	rows, err := db.QueryContext(ctx, fmt.Sprintf(query, param.SearchBy, param.OrderBy, param.OrderDir), "%"+param.SearchValue+"%", param.CourseID, STATUS_AVAILABLE, param.Offset, param.Limit)
	if err != nil {
		return all, errors.Wrap(err, "error at query all course exam")
	}

	defer rows.Close()

	for rows.Next() {
		one := &CourseExam{}
		err = rows.Scan(
			&one.ID, &one.CourseID, &one.TypeExam, &one.ExamIndex, &one.Text, &one.ImageURL,
		)
		if err != nil {
			return all, errors.Wrap(err, "error at query all and scan one of course examp data")
		}

		one.Answers, err = (&CourseExamAnswer{CourseExamID: one.ID}).AllById(ctx, db, param.LimitAnswer)
		if err != nil {
			return all, errors.Wrap(err, "error at query all and scan one of course examp data for answers")
		}

		all = append(all, one)
	}

	return all, nil
}

func (c *CourseExam) Update(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID
	query := `UPDATE course_exam SET course_id=$1,type_exam=$2,exam_index=$3,text=$4,image_url=$5 WHERE id = $6 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.CourseID, c.TypeExam, c.ExamIndex, c.Text, c.ImageURL, c.ID).Scan(
		&id,
	)
	if err != nil {
		return id, errors.Wrap(err, "error at update one of course exam data")
	}
	return id, nil
}

func (c *CourseExam) Delete(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID
	query := `UPDATE course_exam SET flag_status=$1 WHERE id=$2 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), STATUS_DELETE, c.ID).Scan(
		&id,
	)
	if err != nil {
		return id, errors.Wrap(err, "error at delete one of course exam data")
	}
	return id, nil
}
