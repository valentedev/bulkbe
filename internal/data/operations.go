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

type Schedule struct {
	Creation string `json:"creation,omitempty"`
	Houston  string `json:"houston,omitempty"`
	Santos   string `json:"santos,omitempty"`
	Campana  string `json:"campana,omitempty"`
}

func (op OperationModel) SchedulesByVessel(id int64) ([]*Schedule, error) {

	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
	select distinct created_at::date
	from operations 
	where vessel=$1
	group by created_at
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
			&operation.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		operations = append(operations, &operation)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	schedules := []*Schedule{}

	for _, op := range operations {
		var schedule Schedule
		schedule.Creation = op.CreatedAt
		schedules = append(schedules, &schedule)
	}

	// /////////////////////

	query = `
	SELECT created_at::date, port, startop
	FROM operations
	WHERE vessel=$1
	ORDER BY created_at ASC;
	`

	rows, err = op.DB.QueryContext(ctx, query, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()
	operations = []*Operation{}

	for rows.Next() {
		var operation Operation
		err := rows.Scan(
			&operation.CreatedAt,
			&operation.Port,
			&operation.StartOp,
		)
		if err != nil {
			return nil, err
		}
		operations = append(operations, &operation)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	for _, op2 := range operations {
		for i := range schedules {
			if schedules[i].Creation == op2.CreatedAt {
				switch op2.Port {
				case "Houston":
					schedules[i].Houston = op2.StartOp
				case "Santos":
					schedules[i].Santos = op2.StartOp
				case "Campana":
					schedules[i].Campana = op2.StartOp
				}
			}
		}
	}

	return schedules, nil
}
