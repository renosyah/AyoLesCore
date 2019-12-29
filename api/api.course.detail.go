package api

import (
	"context"
	"database/sql"
	"net/http"

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
