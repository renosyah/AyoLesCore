package api

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/renosyah/AyoLesCore/model"
	uuid "github.com/satori/go.uuid"
)

type (
	CourseModule struct {
		db   *sql.DB
		Name string
	}

	AddCourseParam struct {
		CourseName string          `json:"course_name"`
		ImageURL   string          `json:"image_url"`
		Teacher    *model.Teacher  `json:"teacher"`
		Category   *model.Category `json:"category"`
	}

	OneCourseParam struct {
		ID uuid.UUID `json:"id"`
	}

	AllCourseParam struct {
		CategoryID  uuid.UUID `json:"category_id"`
		SearchBy    string    `json:"search_by"`
		SearchValue string    `json:"search_value"`
		OrderBy     string    `json:"order_by"`
		OrderDir    string    `json:"order_dir"`
		Offset      int       `json:"offset"`
		Limit       int       `json:"limit"`
	}
)

func NewCourseModule(db *sql.DB) *CourseModule {
	return &CourseModule{
		db:   db,
		Name: "module/course",
	}
}

func (m CourseModule) All(ctx context.Context, param AllCourseParam) ([]model.CourseResponse, *Error) {
	var allResp []model.CourseResponse

	data, err := (&model.Course{}).All(ctx, m.db, model.AllCourse{
		CategoryID:  param.CategoryID,
		SearchBy:    param.SearchBy,
		SearchValue: param.SearchValue,
		OrderBy:     param.OrderBy,
		OrderDir:    param.OrderDir,
		Offset:      param.Offset,
		Limit:       param.Limit,
	})
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query all course"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no Course found"
		}

		return []model.CourseResponse{}, NewErrorWrap(err, m.Name, "all/course",
			message, status)
	}

	for _, each := range data {
		allResp = append(allResp, each.Response())
	}

	return allResp, nil

}
func (m CourseModule) Add(ctx context.Context, param AddCourseParam) (model.CourseResponse, *Error) {
	course := &model.Course{
		CourseName: param.CourseName,
		ImageURL:   param.ImageURL,
		Teacher:    param.Teacher,
		Category:   param.Category,
	}

	id, err := course.Add(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add course"

		return model.CourseResponse{}, NewErrorWrap(err, m.Name, "add/course",
			message, status)
	}

	course.ID = id

	return course.Response(), nil
}

func (m CourseModule) One(ctx context.Context, param OneCourseParam) (model.CourseResponse, *Error) {
	course := &model.Course{
		ID: param.ID,
	}

	data, err := course.One(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one course"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no course found"
		}

		return model.CourseResponse{}, NewErrorWrap(err, m.Name, "one/course",
			message, status)
	}

	return data.Response(), nil
}
