package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Vessel struct {
	ID           int64  `json:"id"`
	CreatedBy    string `json:"created_by"`
	Name         string `json:"name"`
	Voyage       string `json:"voyage"`
	Service      string `json:"service"`
	Status       string `json:"status"`
	Tolerance    string `json:"tolerance"`
	Booking      string `json:"booking"`
	InternalNote string `json:"internal_note"`
	ExternalNote string `json:"external_note"`
	Operations   []Operation
	Orders       []Order
}

type VesselModel struct {
	DB *sql.DB
}

func (v VesselModel) Insert(vessel *Vessel) error {
	query := `
		INSERT INTO vessels (created_by, name, voyage, service, status, tolerance, booking, internal_note, external_note)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
	`
	args := []interface{}{
		vessel.CreatedBy,
		vessel.Name,
		vessel.Voyage,
		vessel.Service,
		vessel.Status,
		vessel.Tolerance,
		vessel.Booking,
		vessel.InternalNote,
		vessel.ExternalNote,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return v.DB.QueryRowContext(ctx, query, args...).Scan(&vessel.ID)
}

func (v VesselModel) Get(id int64) (*Vessel, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
        SELECT id, created_by, name, voyage, service, status, tolerance, booking, internal_note, external_note
        FROM vessels
        WHERE id = $1`

	var vessel Vessel

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := v.DB.QueryRowContext(ctx, query, id).Scan(
		&vessel.ID,
		&vessel.CreatedBy,
		&vessel.Name,
		&vessel.Voyage,
		&vessel.Service,
		&vessel.Status,
		&vessel.Tolerance,
		&vessel.Booking,
		&vessel.InternalNote,
		&vessel.ExternalNote,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &vessel, nil
}

func (v VesselModel) GetAll() ([]*Vessel, error) {
	query := `
		SELECT id, name, status, voyage
		FROM vessels;
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := v.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	vessels := []*Vessel{}

	for rows.Next() {
		var vessel Vessel
		err := rows.Scan(
			&vessel.ID,
			&vessel.Name,
			&vessel.Status,
			&vessel.Voyage,
		)
		if err != nil {
			return nil, err
		}
		vessels = append(vessels, &vessel)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return vessels, nil
}

func (v VesselModel) Update(vessel *Vessel) error {
	query := `
		UPDATE vessels
		SET created_by = $1, name = $2, voyage = $3, service = $4, status = $5, tolerance = $6, booking = $7, internal_note = $8, external_note = $9, version = version + 1
		WHERE id = $10
		RETURNING id
	`
	args := []interface{}{
		vessel.CreatedBy,
		vessel.Name,
		vessel.Voyage,
		vessel.Service,
		vessel.Status,
		vessel.Tolerance,
		vessel.Booking,
		vessel.InternalNote,
		vessel.ExternalNote,
		vessel.ID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := v.DB.QueryRowContext(ctx, query, args...).Scan(&vessel.ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil
}

func (v VesselModel) Delete() error {
	return nil
}
