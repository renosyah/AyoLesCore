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
	CourseExamAnswerModule struct {
		db   *sql.DB
		Name string
	}

	AddCourseExamAnswerParam struct {
		CourseExamID uuid.UUID `json:"course_exam_id"`
		TypeAnswer   int32     `json:"type_answer"`
		Label        string    `json:"label"`
		Text         string    `json:"text"`
		ImageURL     string    `json:"image_url"`
	}

	OneCourseExamAnswerParam struct {
		ID uuid.UUID `json:"id"`
	}

	AllCourseExamAnswerParam struct {
		CourseExamID uuid.UUID `json:"course_exam_id"`
		SearchBy     string    `json:"search_by"`
		SearchValue  string    `json:"search_value"`
		OrderBy      string    `json:"order_by"`
		OrderDir     string    `json:"order_dir"`
		Offset       int       `json:"offset"`
		Limit        int       `json:"limit"`
	}
)

func NewCourseExamAnswerModule(db *sql.DB) *CourseExamAnswerModule {
	return &CourseExamAnswerModule{
		db:   db,
		Name: "module/course_exam_answer",
	}
}

func (m CourseExamAnswerModule) All(ctx context.Context, param AllCourseExamAnswerParam) ([]model.CourseExamAnswerResponse, *Error) {
	var allResp []model.CourseExamAnswerResponse

	data, err := (&model.CourseExamAnswer{}).All(ctx, m.db, model.AllCourseExamAnswer{
		CourseExamID: param.CourseExamID,
		SearchBy:     param.SearchBy,
		SearchValue:  param.SearchValue,
		OrderBy:      param.OrderBy,
		OrderDir:     param.OrderDir,
		Offset:       param.Offset,
		Limit:        param.Limit,
	})
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query all course exam answer"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no Course exam answer found"
		}

		return []model.CourseExamAnswerResponse{}, NewErrorWrap(err, m.Name, "all/course_examp_answer",
			message, status)
	}

	for _, each := range data {
		allResp = append(allResp, each.Response())
	}

	return allResp, nil

}
func (m CourseExamAnswerModule) Add(ctx context.Context, param AddCourseExamAnswerParam) (model.CourseExamAnswerResponse, *Error) {
	courseExamAnswer := &model.CourseExamAnswer{
		CourseExamID: param.CourseExamID,
		TypeAnswer:   param.TypeAnswer,
		Label:        param.Label,
		Text:         param.Text,
		ImageURL:     param.ImageURL,
	}

	id, err := courseExamAnswer.Add(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add course exam answer"

		return model.CourseExamAnswerResponse{}, NewErrorWrap(err, m.Name, "add/course_exam_answer",
			message, status)
	}

	courseExamAnswer.ID = id

	return courseExamAnswer.Response(), nil
}

func (m CourseExamAnswerModule) One(ctx context.Context, param OneCourseExamAnswerParam) (model.CourseExamAnswerResponse, *Error) {
	courseExamAnswer := &model.CourseExamAnswer{
		ID: param.ID,
	}

	data, err := courseExamAnswer.One(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one course exam"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no course exam found"
		}

		return model.CourseExamAnswerResponse{}, NewErrorWrap(err, m.Name, "one/course_exam",
			message, status)
	}

	return data.Response(), nil
}
