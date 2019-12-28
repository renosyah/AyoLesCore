package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type (
	Teacher struct {
		ID       uuid.UUID `json:"id"`
		Name     string    `json:"name"`
		Email    string    `json:"email"`
		Password string    `json:"password"`
	}

	TeacherResponse struct {
		ID       uuid.UUID `json:"id"`
		Name     string    `json:"name"`
		Email    string    `json:"email"`
		Password string    `json:"-"`
	}
	AllTeacher struct {
		SearchBy    string `json:"search_by"`
		SearchValue string `json:"search_value"`
		OrderBy     string `json:"order_by"`
		OrderDir    string `json:"order_dir"`
		Offset      int    `json:"offset"`
		Limit       int    `json:"limit"`
	}
)

func (t *Teacher) Response() TeacherResponse {
	return TeacherResponse{
		ID:       t.ID,
		Name:     t.Name,
		Email:    t.Email,
		Password: t.Password,
	}
}

func (t *Teacher) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	query := `INSERT INTO teacher (name,email,password) VALUES ($1,$2,$3) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), t.Name, t.Email, t.Password).Scan(
		&t.ID,
	)
	if err != nil {
		return t.ID, errors.Wrap(err, "error at insert new teacher")
	}

	return t.ID, nil
}

func (t *Teacher) One(ctx context.Context, db *sql.DB) (*Teacher, error) {
	one := &Teacher{}
	query := `SELECT id,name,email,password FROM teacher WHERE id = $1 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), t.ID).Scan(
		&one.ID, &one.Name, &one.Email, &one.Password,
	)
	if err != nil {
		return one, errors.Wrap(err, "error at query teacher with id")
	}

	return one, nil
}

func (t *Teacher) All(ctx context.Context, db *sql.DB, param AllTeacher) ([]*Teacher, error) {
	all := []*Teacher{}
	query := `SELECT id,name,email,password FROM teacher WHERE %s LIKE $1 ORDER BY %s %s OFFSET $2 LIMIT $3 `
	rows, err := db.QueryContext(ctx, fmt.Sprintf(query, param.SearchBy, param.OrderBy, param.OrderDir), "%"+param.SearchValue+"%", param.Offset, param.Limit)
	if err != nil {
		return all, errors.Wrap(err, "error at query all teacher")
	}

	defer rows.Close()

	for rows.Next() {
		one := &Teacher{}
		err = rows.Scan(
			&one.ID, &one.Name, &one.Email, &one.Password,
		)
		if err != nil {
			return all, errors.Wrap(err, "error at query all and scan one of teacher data")
		}
		all = append(all, one)
	}

	return all, nil
}

func (t *Teacher) OneByEmail(ctx context.Context, db *sql.DB) (*Teacher, error) {
	one := &Teacher{}
	query := `SELECT id,name,email,password FROM teacher WHERE email = $1 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), t.Email).Scan(
		&one.ID, &one.Name, &one.Email, &one.Password,
	)
	if err != nil {
		return one, errors.Wrap(err, "error at query teacher with email")
	}
	return one, nil
}
