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
	ClassRoomProgressModule struct {
		db   *sql.DB
		Name string
	}

	AddClassRoomProgressParam struct {
		ClassRoomID      uuid.UUID `json:"classroom_id"`
		CourseMaterialID uuid.UUID `json:"course_material_id"`
	}

	OneClassRoomProgressParam struct {
		ID uuid.UUID `json:"id"`
	}

	AllClassRoomProgressParam struct {
		ClassRoomID uuid.UUID `json:"classroom_id"`
		Offset      int       `json:"offset"`
		Limit       int       `json:"limit"`
	}
)

func NewClassRoomProgressModule(db *sql.DB) *ClassRoomProgressModule {
	return &ClassRoomProgressModule{
		db:   db,
		Name: "module/course_progress_module",
	}
}

func (m ClassRoomProgressModule) All(ctx context.Context, param AllClassRoomProgressParam) ([]model.ClassRoomProgressResponse, *Error) {
	var allResp []model.ClassRoomProgressResponse

	data, err := (&model.ClassRoomProgress{}).All(ctx, m.db, model.AllClassRoomProgress{
		ClassRoomID: param.ClassRoomID,
		Offset:      param.Offset,
		Limit:       param.Limit,
	})
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query all classRoom progress"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no Course found"
		}

		return []model.ClassRoomProgressResponse{}, NewErrorWrap(err, m.Name, "all/course_progress_module",
			message, status)
	}

	for _, each := range data {
		allResp = append(allResp, each.Response())
	}

	return allResp, nil

}
func (m ClassRoomProgressModule) Add(ctx context.Context, param AddClassRoomProgressParam) (model.ClassRoomProgressResponse, *Error) {
	courseProgress := &model.ClassRoomProgress{
		ClassRoomID:      param.ClassRoomID,
		CourseMaterialID: param.CourseMaterialID,
	}

	id, err := courseProgress.Add(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add classRoom progress"

		return model.ClassRoomProgressResponse{}, NewErrorWrap(err, m.Name, "add/course_progress_module",
			message, status)
	}

	courseProgress.ID = id

	return courseProgress.Response(), nil
}

func (m ClassRoomProgressModule) One(ctx context.Context, param OneClassRoomProgressParam) (model.ClassRoomProgressResponse, *Error) {
	courseProgress := &model.ClassRoomProgress{
		ID: param.ID,
	}

	data, err := courseProgress.One(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one all classRoom progress"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no all classRoom progress found"
		}

		return model.ClassRoomProgressResponse{}, NewErrorWrap(err, m.Name, "one/course_progress_module",
			message, status)
	}

	return data.Response(), nil
}
