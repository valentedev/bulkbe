package data

import (
	"context"
	"database/sql"
	"time"
)

type Operation struct {
	ID        int64  `json:"id"`
	CreatedBy string `json:"created_by"`
	Type      string `json:"type"`
	Port      string `json:"port"`
	StartOp   string `json:"startop"`
	EndOp     string `json:"endop"`
	Vessel    int64  `json:"vessel"`
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
