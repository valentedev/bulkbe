package data

import (
	"context"
	"database/sql"
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

	return v.DB.QueryRowContext(ctx, query, args...).Scan(
		&vessel.ID,
	)
}

func (v VesselModel) Get() error {
	return nil
}

func (v VesselModel) GetAll() error {
	return nil
}

func (v VesselModel) Update() error {
	return nil
}

func (v VesselModel) Delete() error {
	return nil
}