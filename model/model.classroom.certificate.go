package model

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type (
	ClassRoomCertificate struct {
		ID          uuid.UUID `json:"id"`
		ClassroomID uuid.UUID `json:"classroom_id"`
		HashID      string    `json:"hash_id"`
		CreateAt    time.Time `json:"-"`
	}

	ClassRoomCertificateResponse struct {
		ID          uuid.UUID `json:"id"`
		ClassroomID uuid.UUID `json:"classroom_id"`
		HashID      string    `json:"hash_id"`
		CreateAt    time.Time `json:"-"`
	}

	AllClassRoomCertificate struct {
		StudentID uuid.UUID `json:"student_id"`
		OrderBy   string    `json:"order_by"`
		OrderDir  string    `json:"order_dir"`
		Offset    int       `json:"offset"`
		Limit     int       `json:"limit"`
	}
)

func (c *ClassRoomCertificate) Response() ClassRoomCertificateResponse {
	return ClassRoomCertificateResponse{
		ID:          c.ID,
		ClassroomID: c.ClassroomID,
		HashID:      c.HashID,
		CreateAt:    c.CreateAt,
	}
}

func (c *ClassRoomCertificate) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var emptyUUID uuid.UUID

	one, _ := c.One(ctx, db)
	if one.ID != emptyUUID {
		return one.ID, nil
	}

	query := `INSERT INTO classroom_certificate (classroom_id,hash_id) VALUES ($1,$2) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.ClassroomID, c.HashID).Scan(
		&c.ID,
	)
	if err != nil {
		return c.ID, errors.Wrap(err, "error at add one classroom certificate")
	}
	return c.ID, nil
}

func (c *ClassRoomCertificate) One(ctx context.Context, db *sql.DB) (*ClassRoomCertificate, error) {
	one := &ClassRoomCertificate{}

	query := `SELECT id,classroom_id,hash_id,create_at FROM classroom_certificate WHERE classroom_id = $1 OR hash_id = $2 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), c.ClassroomID, c.HashID).Scan(
		&one.ID, &one.ClassroomID, &one.HashID, &one.CreateAt,
	)
	if err != nil {
		return one, errors.Wrap(err, "error at one classroom certificate")
	}
	return one, nil
}

func (c *ClassRoomCertificate) All(ctx context.Context, db *sql.DB, param AllClassRoomCertificate) ([]*ClassRoomCertificate, error) {
	all := []*ClassRoomCertificate{}

	query := `SELECT 
				classroom_certificate.id,
				classroom_certificate.classroom_id,
				classroom_certificate.hash_id,
				classroom_certificate.create_at 
			FROM 
				classroom_certificate 
			INNER JOIN
				classroom 
			ON
				classroom.id = classroom_certificate.classroom_id
			WHERE 
				classroom.student_id = $1
			ORDER BY %s %s 
			OFFSET $2 
			LIMIT $3`

	rows, err := db.QueryContext(ctx, fmt.Sprintf(query, param.OrderBy, param.OrderDir), param.StudentID, param.Offset, param.Limit)
	if err != nil {
		return all, errors.Wrap(err, "error at query all classroom certificate")
	}

	defer rows.Close()

	for rows.Next() {
		one := &ClassRoomCertificate{}
		err = rows.Scan(
			&one.ID, &one.ClassroomID, &one.HashID, &one.CreateAt,
		)
		if err != nil {
			return all, errors.Wrap(err, "error at query one of classroom certificate data")
		}
		all = append(all, one)
	}

	return all, nil
}
