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
		Nis      string    `json:"nis"`
		Password string    `json:"password"`
	}

	StudentResponse struct {
		ID       uuid.UUID `json:"id"`
		Name     string    `json:"name"`
		Nis      string    `json:"nis"`
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
		Nis:      s.Nis,
		Password: s.Password,
	}
}

func (s *Student) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	query := `INSERT INTO student (name,nis,password) VALUES ($1,$2,$3) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), s.Name, s.Nis, s.Password).Scan(
		&s.ID,
	)
	if err != nil {
		return s.ID, errors.Wrap(err, "error at insert new student")
	}

	return s.ID, nil
}

func (s *Student) One(ctx context.Context, db *sql.DB) (*Student, error) {
	one := &Student{}
	query := `SELECT id,name,nis,password FROM student WHERE id = $1 AND flag_status = $2 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), s.ID, STATUS_AVAILABLE).Scan(
		&one.ID, &one.Name, &one.Nis, &one.Password,
	)
	if err != nil {
		return one, errors.Wrap(err, "error at query student with id")
	}

	return one, nil
}

func (s *Student) All(ctx context.Context, db *sql.DB, param AllStudent) ([]*Student, error) {
	all := []*Student{}
	query := `SELECT id,name,nis,password FROM student WHERE %s LIKE $1 AND flag_status = $2 ORDER BY %s %s OFFSET $3 LIMIT $4 `
	rows, err := db.QueryContext(ctx, fmt.Sprintf(query, param.SearchBy, param.OrderBy, param.OrderDir), "%"+param.SearchValue+"%", STATUS_AVAILABLE, param.Offset, param.Limit)
	if err != nil {
		return all, errors.Wrap(err, "error at query all student")
	}

	defer rows.Close()

	for rows.Next() {
		one := &Student{}
		err = rows.Scan(
			&one.ID, &one.Name, &one.Nis, &one.Password,
		)
		if err != nil {
			return all, errors.Wrap(err, "error at query all and scan one of student data")
		}
		all = append(all, one)
	}

	return all, nil
}

func (s *Student) OneByNis(ctx context.Context, db *sql.DB) (*Student, error) {
	one := &Student{}
	query := `SELECT id,name,nis,password FROM student WHERE nis = $1 AND flag_status = $2 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), s.Nis, STATUS_AVAILABLE).Scan(
		&one.ID, &one.Name, &one.Nis, &one.Password,
	)
	if err != nil {
		return one, errors.Wrap(err, "error at query student with nis")
	}
	return one, nil
}

func (s *Student) Update(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var emptyId uuid.UUID
	query := `UPDATE student SET name = $1,nis = $2 WHERE id = $3 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), s.Name, s.Nis, s.ID).Scan(
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

func (s *Student) Delete(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID
	query := `UPDATE student SET flag_status=$1 WHERE id=$2 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), STATUS_DELETE, s.ID).Scan(
		&id,
	)
	if err != nil {
		return id, errors.Wrap(err, "error at delete one of student data")
	}
	return id, nil
}
