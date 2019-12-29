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
	ClassRoomModule struct {
		db   *sql.DB
		Name string
	}
	AddClassRoomParam struct {
		ID        uuid.UUID     `json:"id"`
		Course    *model.Course `json:"course"`
		StudentID uuid.UUID     `json:"student_id"`
	}

	OneClassRoomParam struct {
		ID uuid.UUID `json:"id"`
	}

	AllClassRoomParam struct {
		StudentID   uuid.UUID `json:"student_id"`
		SearchBy    string    `json:"search_by"`
		SearchValue string    `json:"search_value"`
		OrderBy     string    `json:"order_by"`
		OrderDir    string    `json:"order_dir"`
		Offset      int       `json:"offset"`
		Limit       int       `json:"limit"`
	}
)

func NewClassRoomModule(db *sql.DB) *ClassRoomModule {
	return &ClassRoomModule{
		db:   db,
		Name: "module/classroom",
	}
}

func (m ClassRoomModule) All(ctx context.Context, param AllClassRoomParam) ([]model.ClassRoomResponse, *Error) {
	var allResp []model.ClassRoomResponse

	data, err := (&model.ClassRoom{}).All(ctx, m.db, model.AllClassRoom{
		StudentID:   param.StudentID,
		SearchBy:    param.SearchBy,
		SearchValue: param.SearchValue,
		OrderBy:     param.OrderBy,
		OrderDir:    param.OrderDir,
		Offset:      param.Offset,
		Limit:       param.Limit,
	})
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query all classroom"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no classroom found"
		}

		return []model.ClassRoomResponse{}, NewErrorWrap(err, m.Name, "all/classroom",
			message, status)
	}

	for _, each := range data {
		allResp = append(allResp, each.Response())
	}

	return allResp, nil

}
func (m ClassRoomModule) Add(ctx context.Context, param AddClassRoomParam) (model.ClassRoomResponse, *Error) {
	classroom := &model.ClassRoom{
		Course: &model.Course{
			ID: param.Course.ID,
		},
		StudentID: param.StudentID,
	}

	id, err := classroom.Add(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add classroom"

		return model.ClassRoomResponse{}, NewErrorWrap(err, m.Name, "add/classroom",
			message, status)
	}

	classroom.ID = id

	classroom.Course, err = classroom.Course.One(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get classroom data after added"

		return model.ClassRoomResponse{}, NewErrorWrap(err, m.Name, "add/classroom",
			message, status)
	}

	return classroom.Response(), nil
}

func (m ClassRoomModule) One(ctx context.Context, param OneClassRoomParam) (model.ClassRoomResponse, *Error) {
	classroom := &model.ClassRoom{
		ID: param.ID,
	}

	data, err := classroom.One(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one classroom"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no classroom found"
		}

		return model.ClassRoomResponse{}, NewErrorWrap(err, m.Name, "one/classroom",
			message, status)
	}

	return data.Response(), nil
}
