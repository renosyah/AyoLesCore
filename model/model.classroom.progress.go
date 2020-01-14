package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type (
	ClassRoomProgress struct {
		ID               uuid.UUID `json:"id"`
		ClassRoomID      uuid.UUID `json:"classroom_id"`
		CourseMaterialID uuid.UUID `json:"course_material_id"`
	}
	ClassRoomProgressResponse struct {
		ID               uuid.UUID `json:"id"`
		ClassRoomID      uuid.UUID `json:"classroom_id"`
		CourseMaterialID uuid.UUID `json:"course_material_id"`
	}

	AllClassRoomProgress struct {
		ClassRoomID uuid.UUID `json:"classroom_id"`
		Offset      int       `json:"offset"`
		Limit       int       `json:"limit"`
	}
)

func (c *ClassRoomProgress) Response() ClassRoomProgressResponse {
	return ClassRoomProgressResponse{
		ID:               c.ID,
		ClassRoomID:      c.ClassRoomID,
		CourseMaterialID: c.CourseMaterialID,
	}
}

func (c *ClassRoomProgress) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var emptyUUID uuid.UUID

	query := `SELECT id FROM classroom_progress WHERE classroom_id = $1 AND course_material_id = $2 AND flag_status = $3 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.ClassRoomID, c.CourseMaterialID, STATUS_AVAILABLE).Scan(
		&c.ID,
	)
	if err != nil && err != sql.ErrNoRows {
		return c.ID, errors.Wrap(err, "error at query one classroom progress")
	}

	if c.ID != emptyUUID {
		return c.ID, nil
	}

	query = `INSERT INTO classroom_progress (classroom_id,course_material_id) VALUES ($1,$2) RETURNING id`
	err = db.QueryRowContext(ctx, fmt.Sprintf(query), c.ClassRoomID, c.CourseMaterialID).Scan(
		&c.ID,
	)
	if err != nil {
		return c.ID, errors.Wrap(err, "error at insert new classroom progress")
	}

	return c.ID, nil
}

func (c *ClassRoomProgress) One(ctx context.Context, db *sql.DB) (*ClassRoomProgress, error) {
	one := &ClassRoomProgress{}
	query := `SELECT id,classroom_id,course_material_id FROM classroom_progress WHERE id = $1 AND flag_status = $2 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.ID, STATUS_AVAILABLE).Scan(
		&one.ID, &one.ClassRoomID, &one.CourseMaterialID,
	)
	if err != nil {
		return one, errors.Wrap(err, "error at query one classroom progress")
	}

	return one, nil
}

func (c *ClassRoomProgress) All(ctx context.Context, db *sql.DB, param AllClassRoomProgress) ([]*ClassRoomProgress, error) {
	all := []*ClassRoomProgress{}

	query := `SELECT id,classroom_id,course_material_id FROM classroom_progress WHERE classroom_id = $1 AND flag_status = $2 OFFSET $3 LIMIT $4`
	rows, err := db.QueryContext(ctx, fmt.Sprintf(query), param.ClassRoomID, STATUS_AVAILABLE, param.Offset, param.Limit)
	if err != nil {
		return all, errors.Wrap(err, "error at query all classroom progress")
	}

	defer rows.Close()

	for rows.Next() {
		one := &ClassRoomProgress{}
		err = rows.Scan(
			&one.ID, &one.ClassRoomID, &one.CourseMaterialID,
		)
		if err != nil {
			return all, errors.Wrap(err, "error at query one of classroom progress data")
		}
		all = append(all, one)
	}

	return all, nil
}

func (c *ClassRoomProgress) Update(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID
	query := `UPDATE classroom_progress SET classroom_id=$1,course_material_id=$2 WHERE id = $3 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.ClassRoomID, c.CourseMaterialID, c.ID).Scan(
		&id,
	)
	if err != nil {
		return id, errors.Wrap(err, "error at update one of classroom progress data")
	}
	return id, nil
}

func (c *ClassRoomProgress) Delete(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID
	query := `UPDATE classroom_progress SET flag_status=$1 WHERE id=$2 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), STATUS_DELETE, c.ID).Scan(
		&id,
	)
	if err != nil {
		return id, errors.Wrap(err, "error at delete one of classroom progress data")
	}
	return id, nil
}
