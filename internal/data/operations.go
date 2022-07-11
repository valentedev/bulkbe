package data

import (
	"context"
	"database/sql"
	"time"
)

type Operation struct {
	ID        int64  `json:"id,omitempty"`
	CreatedBy string `json:"created_by,omitempty"`
	Type      string `json:"type,omitempty"`
	Port      string `json:"port,omitempty"`
	StartOp   string `json:"startop,omitempty"`
	EndOp     string `json:"endop,omitempty"`
	Vessel    int64  `json:"-"`
}

type OperationModel struct {
	DB *sql.DB
}

func (op OperationModel) Insert(operation *Operation) error {
	query := `
		INSERT INTO operations (created_by, type, port, startop, endop, vessel) 
		VALUES ($1,$2,$3,$4,$5,$6)
		RETURNING id
	`

	args := []interface{}{
		operation.CreatedBy,
		operation.Type,
		operation.Port,
		operation.StartOp,
		operation.EndOp,
		operation.Vessel,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return op.DB.QueryRowContext(ctx, query, args...).Scan(&operation.ID)
}
