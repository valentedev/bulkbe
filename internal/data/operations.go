package data

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Operation struct {
	ID        int64  `json:"id,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
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

func (op OperationModel) GetLoadByVessel(id int64) ([]*Operation, error) {

	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, created_at, port, startop, vessel 
		FROM operations
		WHERE type='load'
		AND vessel=$1
		ORDER BY created_at ASC;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := op.DB.QueryContext(ctx, query, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()
	operations := []*Operation{}

	for rows.Next() {
		var operation Operation
		err := rows.Scan(
			&operation.ID,
			&operation.CreatedAt,
			&operation.Port,
			&operation.StartOp,
			&operation.Vessel,
		)
		if err != nil {
			return nil, err
		}
		operations = append(operations, &operation)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return operations, nil
}
