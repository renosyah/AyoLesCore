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
	CourseExamSolutionModule struct {
		db   *sql.DB
		Name string
	}
	AddCourseExamSolutionParam struct {
		CourseExamID       uuid.UUID `json:"course_exam_id"`
		CourseExamAnswerID uuid.UUID `json:"course_exam_answer_id"`
	}
	OneCourseExamSolutionParam struct {
		ID uuid.UUID `json:"id"`
	}
	AllCourseExamSolutionParam struct {
		CourseExamID uuid.UUID `json:"course_exam_id"`
		OrderBy      string    `json:"order_by"`
		OrderDir     string    `json:"order_dir"`
		Offset       int       `json:"offset"`
		Limit        int       `json:"limit"`
	}
)

func NewCourseExamSolutionModule(db *sql.DB) *CourseExamSolutionModule {
	return &CourseExamSolutionModule{
		db:   db,
		Name: "module/course_exam_solution",
	}
}

func (m CourseExamSolutionModule) Add(ctx context.Context, param AddCourseExamSolutionParam) (model.CourseExamSolutionResponse, *Error) {
	courseExamSolution := &model.CourseExamSolution{
		CourseExamID:       param.CourseExamID,
		CourseExamAnswerID: param.CourseExamAnswerID,
	}

	id, err := courseExamSolution.Add(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add exam solution"

		return model.CourseExamSolutionResponse{}, NewErrorWrap(err, m.Name, "add/course_exam_solution",
			message, status)
	}

	courseExamSolution.ID = id

	return courseExamSolution.Response(), nil
}

func (m CourseExamSolutionModule) All(ctx context.Context, param AllCourseExamSolutionParam) ([]model.CourseExamSolutionResponse, *Error) {
	var allResp []model.CourseExamSolutionResponse

	data, err := (&model.CourseExamSolution{}).All(ctx, m.db, model.AllCourseExamSolution{
		CourseExamID: param.CourseExamID,
		OrderBy:      param.OrderBy,
		OrderDir:     param.OrderDir,
		Offset:       param.Offset,
		Limit:        param.Limit,
	})
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query all course exam solution"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no course exam solution found"
		}

		return []model.CourseExamSolutionResponse{}, NewErrorWrap(err, m.Name, "all/course_exam_solution",
			message, status)
	}

	for _, each := range data {
		allResp = append(allResp, each.Response())
	}

	return allResp, nil
}

func (m CourseExamSolutionModule) One(ctx context.Context, param OneCourseExamSolutionParam) (model.CourseExamSolutionResponse, *Error) {
	courseExamSolution := &model.CourseExamSolution{
		ID: param.ID,
	}

	data, err := courseExamSolution.One(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one course exam"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no course exam found"
		}

		return model.CourseExamSolutionResponse{}, NewErrorWrap(err, m.Name, "one/course_exam_solution",
			message, status)
	}

	return data.Response(), nil
}
