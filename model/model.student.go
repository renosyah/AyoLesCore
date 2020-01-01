package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type (
	Student struct {
		ID       uuid.UUID `json:"id"`
		Name     string    `json:"name"`
		Email    string    `json:"email"`
		Password string    `json:"password"`
	}

	StudentResponse struct {
		ID       uuid.UUID `json:"id"`
		Name     string    `json:"name"`
		Email    string    `json:"email"`
		Password string    `json:"-"`
	}

	AllStudent struct {
		SearchBy    string `json:"search_by"`
		SearchValue string `json:"search_value"`
		OrderBy     string `json:"order_by"`
		OrderDir    string `json:"order_dir"`
		Offset      int    `json:"offset"`
		Limit       int    `json:"limit"`
	}
)

func (s *Student) Response() StudentResponse {
	return StudentResponse{
		ID:       s.ID,
		Name:     s.Name,
		Email:    s.Email,
		Password: s.Password,
	}
}

func (s *Student) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	query := `INSERT INTO student (name,email,password) VALUES ($1,$2,$3) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), s.Name, s.Email, s.Password).Scan(
		&s.ID,
	)
	if err != nil {
		return s.ID, errors.Wrap(err, "error at insert new student")
	}

	return s.ID, nil
}

func (s *Student) Update(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var emptyId uuid.UUID
	query := `UPDATE student SET name = $1,email = $2 WHERE id = $3 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), s.Name, s.Email, s.ID).Scan(
		&s.ID,
	)

	if s.Password != "" {
		query = `UPDATE student SET password = $1 WHERE id = $2 RETURNING id`
		err = db.QueryRowContext(ctx, fmt.Sprintf(query), s.Password, s.ID).Scan(
			&s.ID,
		)
	}

	if err != nil || s.ID == emptyId {
		return s.ID, errors.Wrap(err, "error at update student")
	}

	return s.ID, nil
}

func (s *Student) One(ctx context.Context, db *sql.DB) (*Student, error) {
	one := &Student{}
	query := `SELECT id,name,email,password FROM student WHERE id = $1 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), s.ID).Scan(
		&one.ID, &one.Name, &one.Email, &one.Password,
	)
	if err != nil {
		return one, errors.Wrap(err, "error at query student with id")
	}

	return one, nil
}

func (s *Student) All(ctx context.Context, db *sql.DB, param AllStudent) ([]*Student, error) {
	all := []*Student{}
	query := `SELECT id,name,email,password FROM student WHERE %s LIKE $1 ORDER BY %s %s OFFSET $2 LIMIT $3 `
	rows, err := db.QueryContext(ctx, fmt.Sprintf(query, param.SearchBy, param.OrderBy, param.OrderDir), "%"+param.SearchValue+"%", param.Offset, param.Limit)
	if err != nil {
		return all, errors.Wrap(err, "error at query all student")
	}

	defer rows.Close()

	for rows.Next() {
		one := &Student{}
		err = rows.Scan(
			&one.ID, &one.Name, &one.Email, &one.Password,
		)
		if err != nil {
			return all, errors.Wrap(err, "error at query all and scan one of student data")
		}
		all = append(all, one)
	}

	return all, nil
}

func (s *Student) OneByEmail(ctx context.Context, db *sql.DB) (*Student, error) {
	one := &Student{}
	query := `SELECT id,name,email,password FROM student WHERE email = $1 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), s.Email).Scan(
		&one.ID, &one.Name, &one.Email, &one.Password,
	)
	if err != nil {
		return one, errors.Wrap(err, "error at query student with email")
	}
	return one, nil
}
