package data

import "database/sql"

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

func (v VesselModel) Insert() error {
	return nil
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
