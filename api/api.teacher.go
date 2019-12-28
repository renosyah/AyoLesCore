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
	TeacherModule struct {
		db   *sql.DB
		Name string
	}

	TeacherLoginParam struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	AddTeacherParam struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	OneTeacherParam struct {
		ID uuid.UUID `json:"id"`
	}

	AllTeacherParam struct {
		SearchBy    string `json:"search_by"`
		SearchValue string `json:"search_value"`
		OrderBy     string `json:"order_by"`
		OrderDir    string `json:"order_dir"`
		Offset      int    `json:"offset"`
		Limit       int    `json:"limit"`
	}
)

func NewTeacherModule(db *sql.DB) *TeacherModule {
	return &TeacherModule{
		db:   db,
		Name: "module/teacher",
	}
}

func (m TeacherModule) Add(ctx context.Context, param AddTeacherParam) (model.TeacherResponse, *Error) {

	teacher := &model.Teacher{
		Name:     param.Name,
		Email:    param.Email,
		Password: param.Password,
	}

	check, err := teacher.OneByEmail(ctx, m.db)
	if err != nil && errors.Cause(err) != sql.ErrNoRows {
		status := http.StatusInternalServerError
		message := "error on check existing teacher"

		return model.TeacherResponse{}, NewErrorWrap(err, m.Name, "add/teacher",
			message, status)
	}

	if check.Email != "" && check.Email == teacher.Email {
		status := http.StatusOK
		message := "teacher with this email is exist"

		return model.TeacherResponse{}, NewErrorWrap(err, m.Name, "add/teacher",
			message, status)
	}

	id, err := teacher.Add(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add teacher"

		return model.TeacherResponse{}, NewErrorWrap(err, m.Name, "add/teacher",
			message, status)
	}

	teacher.ID = id

	return teacher.Response(), nil
}

func (m TeacherModule) All(ctx context.Context, param AllTeacherParam) ([]model.TeacherResponse, *Error) {
	var allResp []model.TeacherResponse

	data, err := (&model.Teacher{}).All(ctx, m.db, model.AllTeacher{
		SearchBy:    param.SearchBy,
		SearchValue: param.SearchValue,
		OrderBy:     param.OrderBy,
		OrderDir:    param.OrderDir,
		Offset:      param.Offset,
		Limit:       param.Limit,
	})
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query all teacher"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no teacher found"
		}

		return []model.TeacherResponse{}, NewErrorWrap(err, m.Name, "all/teacher",
			message, status)
	}

	for _, each := range data {
		allResp = append(allResp, each.Response())
	}

	return allResp, nil

}

func (m TeacherModule) One(ctx context.Context, param OneTeacherParam) (model.TeacherResponse, *Error) {
	var resp model.TeacherResponse

	teacher, err := (&model.Teacher{ID: param.ID}).One(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query one teacher"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no teacher found"
		}

		return resp, NewErrorWrap(err, m.Name, "one/teacher",
			message, status)
	}

	resp = teacher.Response()

	return resp, nil
}

func (m TeacherModule) Login(ctx context.Context, param TeacherLoginParam) (model.TeacherResponse, *Error) {
	var resp model.TeacherResponse

	teacher, err := (&model.Teacher{Email: param.Email}).OneByEmail(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on login teacher"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusUnauthorized
			message = "no teacher found"
		}

		return resp, NewErrorWrap(err, m.Name, "login/teacher",
			message, status)
	}

	if teacher.Password != param.Password {
		status := http.StatusOK
		message := "password is invalid"

		return resp, NewErrorWrap(err, m.Name, "login/teacher",
			message, status)
	}

	resp = teacher.Response()

	return resp, nil
}
