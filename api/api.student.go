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
	StudentModule struct {
		db   *sql.DB
		Name string
	}

	StudentLoginParam struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	AddStudentParam struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	UpdateStudentParam struct {
		Name     string    `json:"name"`
		Email    string    `json:"email"`
		Password string    `json:"password"`
		ID       uuid.UUID `json:"id"`
	}

	OneStudentParam struct {
		ID uuid.UUID `json:"id"`
	}

	AllStudentParam struct {
		SearchBy    string `json:"search_by"`
		SearchValue string `json:"search_value"`
		OrderBy     string `json:"order_by"`
		OrderDir    string `json:"order_dir"`
		Offset      int    `json:"offset"`
		Limit       int    `json:"limit"`
	}
)

func NewStudentModule(db *sql.DB) *StudentModule {
	return &StudentModule{
		db:   db,
		Name: "module/student",
	}
}

func (m StudentModule) Update(ctx context.Context, param UpdateStudentParam) (model.StudentResponse, *Error) {

	student := &model.Student{
		ID:       param.ID,
		Name:     param.Name,
		Email:    param.Email,
		Password: param.Password,
	}

	id, err := student.Update(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on update one student"

		return model.StudentResponse{}, NewErrorWrap(err, m.Name, "add/student",
			message, status)
	}

	student.ID = id

	return student.Response(), nil
}
func (m StudentModule) Add(ctx context.Context, param AddStudentParam) (model.StudentResponse, *Error) {

	student := &model.Student{
		Name:     param.Name,
		Email:    param.Email,
		Password: param.Password,
	}

	check, err := student.OneByEmail(ctx, m.db)
	if err != nil && errors.Cause(err) != sql.ErrNoRows {
		status := http.StatusInternalServerError
		message := "error on check existing student"

		return model.StudentResponse{}, NewErrorWrap(err, m.Name, "add/student",
			message, status)
	}

	if check.Email != "" && check.Email == student.Email {
		status := http.StatusOK
		message := "student with this email is exist"

		return model.StudentResponse{}, NewErrorWrap(err, m.Name, "add/student",
			message, status)
	}

	id, err := student.Add(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add student"

		return model.StudentResponse{}, NewErrorWrap(err, m.Name, "add/student",
			message, status)
	}

	student.ID = id

	return student.Response(), nil
}

func (m StudentModule) All(ctx context.Context, param AllStudentParam) ([]model.StudentResponse, *Error) {
	var allResp []model.StudentResponse

	data, err := (&model.Student{}).All(ctx, m.db, model.AllStudent{
		SearchBy:    param.SearchBy,
		SearchValue: param.SearchValue,
		OrderBy:     param.OrderBy,
		OrderDir:    param.OrderDir,
		Offset:      param.Offset,
		Limit:       param.Limit,
	})
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query all student"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no student found"
		}

		return []model.StudentResponse{}, NewErrorWrap(err, m.Name, "all/student",
			message, status)
	}

	for _, each := range data {
		allResp = append(allResp, each.Response())
	}

	return allResp, nil

}

func (m StudentModule) One(ctx context.Context, param OneStudentParam) (model.StudentResponse, *Error) {
	var resp model.StudentResponse

	student, err := (&model.Student{ID: param.ID}).One(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query one student"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no student found"
		}

		return resp, NewErrorWrap(err, m.Name, "one/student",
			message, status)
	}

	resp = student.Response()

	return resp, nil
}

func (m StudentModule) Login(ctx context.Context, param StudentLoginParam) (model.StudentResponse, *Error) {
	var resp model.StudentResponse

	student, err := (&model.Student{Email: param.Email}).OneByEmail(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on login student"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusUnauthorized
			message = "no student found"
		}

		return resp, NewErrorWrap(err, m.Name, "login/student",
			message, status)
	}

	if student.Password != param.Password {
		status := http.StatusOK
		message := "password is invalid"

		return resp, NewErrorWrap(err, m.Name, "login/student",
			message, status)
	}

	resp = student.Response()

	return resp, nil
}
