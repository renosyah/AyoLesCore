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
	CourseMaterialModule struct {
		db   *sql.DB
		Name string
	}

	AddCourseMaterialParam struct {
		CourseID      uuid.UUID `json:"course_id"`
		MaterialIndex int32     `json:"material_index"`
		Title         string    `json:"title"`
	}

	OneCourseMaterialParam struct {
		ID uuid.UUID `json:"id"`
	}

	AllCourseMaterialParam struct {
		CourseID    uuid.UUID `json:"course_id"`
		SearchBy    string    `json:"search_by"`
		SearchValue string    `json:"search_value"`
		OrderBy     string    `json:"order_by"`
		OrderDir    string    `json:"order_dir"`
		Offset      int       `json:"offset"`
		Limit       int       `json:"limit"`
	}
)

func NewCourseMaterialModule(db *sql.DB) *CourseMaterialModule {
	return &CourseMaterialModule{
		db:   db,
		Name: "module/course_material",
	}
}

func (m CourseMaterialModule) All(ctx context.Context, param AllCourseMaterialParam) ([]model.CourseMaterialResponse, *Error) {
	var allResp []model.CourseMaterialResponse

	data, err := (&model.CourseMaterial{}).All(ctx, m.db, model.AllCourseMaterial{
		CourseID:    param.CourseID,
		SearchBy:    param.SearchBy,
		SearchValue: param.SearchValue,
		OrderBy:     param.OrderBy,
		OrderDir:    param.OrderDir,
		Offset:      param.Offset,
		Limit:       param.Limit,
	})
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query all course material"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no Course material found"
		}

		return []model.CourseMaterialResponse{}, NewErrorWrap(err, m.Name, "all/course_material",
			message, status)
	}

	for _, each := range data {
		allResp = append(allResp, each.Response())
	}

	return allResp, nil

}

func (m CourseMaterialModule) Add(ctx context.Context, param AddCourseMaterialParam) (model.CourseMaterialResponse, *Error) {
	courseMaterial := &model.CourseMaterial{
		CourseID:      param.CourseID,
		MaterialIndex: param.MaterialIndex,
		Title:         param.Title,
	}

	id, err := courseMaterial.Add(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add course material"

		return model.CourseMaterialResponse{}, NewErrorWrap(err, m.Name, "add/course_material",
			message, status)
	}

	courseMaterial.ID = id

	return courseMaterial.Response(), nil
}

func (m CourseMaterialModule) One(ctx context.Context, param OneCourseMaterialParam) (model.CourseMaterialResponse, *Error) {
	courseMaterial := &model.CourseMaterial{
		ID: param.ID,
	}

	data, err := courseMaterial.One(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one course material"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no course material found"
		}

		return model.CourseMaterialResponse{}, NewErrorWrap(err, m.Name, "one/course_material",
			message, status)
	}

	return data.Response(), nil
}

func (m CourseMaterialModule) Update(ctx context.Context, param AddCourseMaterialParam, id uuid.UUID) (model.CourseMaterialResponse, *Error) {
	var emptyUUID uuid.UUID

	courseMaterial := &model.CourseMaterial{
		ID:            id,
		CourseID:      param.CourseID,
		MaterialIndex: param.MaterialIndex,
		Title:         param.Title,
	}

	i, err := courseMaterial.Update(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on update course material"

		return model.CourseMaterialResponse{}, NewErrorWrap(err, m.Name, "update/course_material",
			message, status)
	}

	return courseMaterial.Response(), nil
}

func (m CourseMaterialModule) Delete(ctx context.Context, id uuid.UUID) (model.CourseMaterialResponse, *Error) {
	var emptyUUID uuid.UUID

	courseMaterial := &model.CourseMaterial{
		ID: id,
	}

	i, err := courseMaterial.Delete(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on delete course material"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no course material found"
		}

		return model.CourseMaterialResponse{}, NewErrorWrap(err, m.Name, "delete/course_material",
			message, status)
	}

	return courseMaterial.Response(), nil
}
