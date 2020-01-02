package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type (
	Course struct {
		ID            uuid.UUID      `json:"id"`
		CourseName    string         `json:"course_name"`
		ImageURL      string         `json:"image_url"`
		Teacher       *Teacher       `json:"teacher"`
		Category      *Category      `json:"category"`
		CourseDetails []CourseDetail `json:"course_details"`
	}

	CourseResponse struct {
		ID            uuid.UUID              `json:"id"`
		CourseName    string                 `json:"course_name"`
		ImageURL      string                 `json:"image_url"`
		Teacher       TeacherResponse        `json:"teacher"`
		Category      CategoryResponse       `json:"category"`
		CourseDetails []CourseDetailResponse `json:"course_details"`
	}

	AllCourse struct {
		CategoryID  uuid.UUID `json:"category_id"`
		TeacherID   uuid.UUID `json:"teacher_id"`
		SearchBy    string    `json:"search_by"`
		SearchValue string    `json:"search_value"`
		OrderBy     string    `json:"order_by"`
		OrderDir    string    `json:"order_dir"`
		Offset      int       `json:"offset"`
		Limit       int       `json:"limit"`
	}
)

func (c *Course) Response() CourseResponse {
	details := []CourseDetailResponse{}
	for _, v := range c.CourseDetails {
		details = append(details, v.Response())
	}
	return CourseResponse{
		ID:            c.ID,
		CourseName:    c.CourseName,
		ImageURL:      c.ImageURL,
		Teacher:       c.Teacher.Response(),
		Category:      c.Category.Response(),
		CourseDetails: details,
	}
}

func (a AllCourse) IsWithCategory() string {
	var emptyID uuid.UUID
	if a.CategoryID == emptyID {
		return ""
	}
	return fmt.Sprintf(`AND category_id = '%s'`, a.CategoryID)
}

func (a AllCourse) IsWithTeacher() string {
	var emptyID uuid.UUID
	if a.TeacherID == emptyID {
		return ""
	}
	return fmt.Sprintf(`AND teacher_id = '%s'`, a.TeacherID)
}

func (c *Course) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	query := `INSERT INTO course (course_name,teacher_id,category_id,image_url) VALUES ($1,$2,$3,$4) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.CourseName, c.Teacher.ID, c.Category.ID, c.ImageURL).Scan(
		&c.ID,
	)
	if err != nil {
		return c.ID, errors.Wrap(err, "error at insert new course")
	}

	return c.ID, nil
}

func (c *Course) One(ctx context.Context, db *sql.DB) (*Course, error) {
	one := &Course{
		Teacher:  &Teacher{},
		Category: &Category{},
	}
	query := `SELECT id,course_name,teacher_id,category_id,image_url FROM course WHERE id = $1 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.ID).Scan(
		&one.ID, &one.CourseName, &one.Teacher.ID, &one.Category.ID, &one.ImageURL,
	)
	if err != nil {
		return one, errors.Wrap(err, "error at query teacher with id")
	}

	one.Teacher, err = one.Teacher.One(ctx, db)
	if err != nil {
		return one, err
	}

	one.Category, err = one.Category.One(ctx, db)
	if err != nil {
		return one, err
	}

	one.CourseDetails, err = (&CourseDetail{CourseID: one.ID}).AllByCourseID(ctx, db)
	if err != nil {
		return one, err
	}

	return one, nil
}

func (c *Course) All(ctx context.Context, db *sql.DB, param AllCourse) ([]*Course, error) {
	all := []*Course{}
	query := `SELECT id,course_name,teacher_id,category_id,image_url FROM course WHERE %s LIKE $1 %s %s ORDER BY %s %s OFFSET $2 LIMIT $3 `
	rows, err := db.QueryContext(ctx, fmt.Sprintf(query, param.SearchBy, param.IsWithCategory(), param.IsWithTeacher(), param.OrderBy, param.OrderDir), "%"+param.SearchValue+"%", param.Offset, param.Limit)
	if err != nil {
		return all, errors.Wrap(err, "error at query all course")
	}

	defer rows.Close()

	for rows.Next() {
		one := &Course{
			Teacher:  &Teacher{},
			Category: &Category{},
		}
		err = rows.Scan(
			&one.ID, &one.CourseName, &one.Teacher.ID, &one.Category.ID, &one.ImageURL,
		)
		if err != nil {
			return all, errors.Wrap(err, "error at query all and scan one of course data")
		}

		one.Teacher, err = one.Teacher.One(ctx, db)
		if err != nil {
			return all, err
		}

		one.Category, err = one.Category.One(ctx, db)
		if err != nil {
			return all, err
		}

		one.CourseDetails, err = (&CourseDetail{CourseID: one.ID}).AllByCourseID(ctx, db)
		if err != nil {
			return all, err
		}

		all = append(all, one)
	}

	return all, nil
}
