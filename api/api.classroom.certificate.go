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
	ClassRoomCertificateModule struct {
		db   *sql.DB
		Name string
	}

	AddClassRoomCertificateParam struct {
		ClassroomID uuid.UUID `json:"classroom_id"`
		HashID      string    `json:"hash_id"`
	}

	OneClassRoomCertificateParam struct {
		ClassroomID uuid.UUID `json:"classroom_id"`
		HashID      string    `json:"hash_id"`
	}

	AllClassRoomCertificateParam struct {
		StudentID uuid.UUID `json:"student_id"`
		OrderBy   string    `json:"order_by"`
		OrderDir  string    `json:"order_dir"`
		Offset    int       `json:"offset"`
		Limit     int       `json:"limit"`
	}
)

func NewClassRoomCertificateModule(db *sql.DB) *ClassRoomCertificateModule {
	return &ClassRoomCertificateModule{
		db:   db,
		Name: "module/classroom_certificate",
	}
}

func (m ClassRoomCertificateModule) All(ctx context.Context, param AllClassRoomCertificateParam) ([]model.ClassRoomCertificateResponse, *Error) {
	var allResp []model.ClassRoomCertificateResponse

	data, err := (&model.ClassRoomCertificate{}).All(ctx, m.db, model.AllClassRoomCertificate{
		StudentID: param.StudentID,
		OrderBy:   param.OrderBy,
		OrderDir:  param.OrderDir,
		Offset:    param.Offset,
		Limit:     param.Limit,
	})
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query all classRoom certificates"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no classRoom certificates found"
		}

		return []model.ClassRoomCertificateResponse{}, NewErrorWrap(err, m.Name, "all/classroom_certificate_module",
			message, status)
	}

	for _, each := range data {
		allResp = append(allResp, each.Response())
	}

	return allResp, nil
}

func (m ClassRoomCertificateModule) One(ctx context.Context, param OneClassRoomCertificateParam) (model.ClassRoomCertificateResponse, *Error) {
	classroomCert := &model.ClassRoomCertificate{
		ClassroomID: param.ClassroomID,
		HashID:      param.HashID,
	}

	data, err := classroomCert.One(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one all classRoom certificate"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no classRoom certificate found"
		}

		return model.ClassRoomCertificateResponse{}, NewErrorWrap(err, m.Name, "one/classroom_certificate_module",
			message, status)
	}

	return data.Response(), nil
}

func (m ClassRoomCertificateModule) Add(ctx context.Context, param AddClassRoomCertificateParam) (model.ClassRoomCertificateResponse, *Error) {
	classroomCert := &model.ClassRoomCertificate{
		ClassroomID: param.ClassroomID,
		HashID:      param.HashID,
	}

	id, err := classroomCert.Add(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add classRoom certificate"

		return model.ClassRoomCertificateResponse{}, NewErrorWrap(err, m.Name, "add/classroom_certificate_module",
			message, status)
	}

	classroomCert.ID = id

	return classroomCert.Response(), nil
}
