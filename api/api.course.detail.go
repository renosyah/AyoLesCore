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
	CourseDetailModule struct {
		db   *sql.DB
		Name string
	}

	AddCourseDetailParam struct {
		CourseID        uuid.UUID `json:"course_id"`
		OverviewText    string    `json:"overview_text"`
		DescriptionText string    `json:"description_text"`
		ImageURL        string    `json:"image_url"`
	}

	AllCourseDetailParam struct {
		CourseID    uuid.UUID `json:"course_id"`
		SearchBy    string    `json:"search_by"`
		SearchValue string    `json:"search_value"`
		OrderBy     string    `json:"order_by"`
		OrderDir    string    `json:"order_dir"`
		Offset      int       `json:"offset"`
		Limit       int       `json:"limit"`
	}
)

func NewCourseDetailModule(db *sql.DB) *CourseDetailModule {
	return &CourseDetailModule{
		db:   db,
		Name: "module/course_detail",
	}
}

func (m CourseDetailModule) Add(ctx context.Context, param AddCourseDetailParam) (model.CourseDetailResponse, *Error) {
	courseDetail := &model.CourseDetail{
		CourseID:        param.CourseID,
		OverviewText:    param.OverviewText,
		DescriptionText: param.DescriptionText,
		ImageURL:        param.ImageURL,
	}

	id, err := courseDetail.Add(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add course detail"

		return model.CourseDetailResponse{}, NewErrorWrap(err, m.Name, "add/course_detail",
			message, status)
	}

	courseDetail.ID = id

	return courseDetail.Response(), nil
}

func (m CourseDetailModule) All(ctx context.Context, param AllCourseDetailParam) ([]model.CourseDetailResponse, *Error) {
	var allResp []model.CourseDetailResponse

	data, err := (&model.CourseDetail{}).All(ctx, m.db, model.AllCourseDetail{
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
		message := "error on query all course detail"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no Course detal found"
		}

		return []model.CourseDetailResponse{}, NewErrorWrap(err, m.Name, "all/course_detail",
			message, status)
	}

	for _, each := range data {
		allResp = append(allResp, each.Response())
	}

	return allResp, nil
}

func (m CourseDetailModule) Update(ctx context.Context, param AddCourseDetailParam, id uuid.UUID) (model.CourseDetailResponse, *Error) {
	var emptyUUID uuid.UUID

	courseDetail := &model.CourseDetail{
		ID:              id,
		CourseID:        param.CourseID,
		OverviewText:    param.OverviewText,
		DescriptionText: param.DescriptionText,
		ImageURL:        param.ImageURL,
	}

	i, err := courseDetail.Update(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on update course detail"

		return model.CourseDetailResponse{}, NewErrorWrap(err, m.Name, "update/course_detail",
			message, status)
	}

	return courseDetail.Response(), nil
}

func (m CourseDetailModule) Delete(ctx context.Context, id uuid.UUID) (model.CourseDetailResponse, *Error) {
	var emptyUUID uuid.UUID

	courseDetail := &model.CourseDetail{
		ID: id,
	}

	i, err := courseDetail.Delete(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on delete course detail"

		return model.CourseDetailResponse{}, NewErrorWrap(err, m.Name, "delete/course_detail",
			message, status)
	}

	return courseDetail.Response(), nil
}
