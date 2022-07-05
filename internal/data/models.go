package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Vessels    VesselModel
	Operations OperationModel
	Orders     OrderModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Vessels:    VesselModel{DB: db},
		Operations: OperationModel{DB: db},
		Orders:     OrderModel{DB: db},
	}
}
