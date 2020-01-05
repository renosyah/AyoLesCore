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
	ClassRoomQualificationModule struct {
		db   *sql.DB
		Name string
	}
	OneClassRoomQualificationParam struct {
		ClassRoomID uuid.UUID `json:"classroom_id"`
	}
)

func NewClassRoomQualificationModule(db *sql.DB) *ClassRoomQualificationModule {
	return &ClassRoomQualificationModule{
		db:   db,
		Name: "module/classroom_qualification",
	}
}

func (m ClassRoomQualificationModule) One(ctx context.Context, param OneClassRoomQualificationParam) (model.ClassRoomQualificationResponse, *Error) {
	classRoomQualification := &model.ClassRoomQualification{
		ClassRoomID: param.ClassRoomID,
	}

	data, err := classRoomQualification.One(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one classRoom qualification"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no one classRoom qualification found"
		}

		return model.ClassRoomQualificationResponse{}, NewErrorWrap(err, m.Name, "one/course_progress_qualification",
			message, status)
	}

	return data.Response(), nil
}
