package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Vessel struct {
	ID           int64       `json:"id,omitempty"`
	CreatedBy    string      `json:"created_by,omitempty"`
	Name         string      `json:"name,omitempty"`
	Voyage       string      `json:"voyage,omitempty"`
	Service      string      `json:"service,omitempty"`
	Status       string      `json:"status,omitempty"`
	Tolerance    string      `json:"tolerance,omitempty"`
	Booking      string      `json:"booking,omitempty"`
	InternalNote string      `json:"internal_note,omitempty"`
	ExternalNote string      `json:"external_note,omitempty"`
	Operations   []Operation `json:"operations,omitempty"`
	Orders       []Order     `json:"orders,omitempty"`
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

func (v VesselModel) Get(id int64) (*Vessel, Metadata, error) {
	if id < 1 {
		return nil, Metadata{}, ErrRecordNotFound
	}
	// Vessel Query ////////////////////////
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
			return nil, Metadata{}, ErrRecordNotFound
		default:
			return nil, Metadata{}, err
		}
	}

	// Operation Query ////////////////////////
	query = `
		SELECT DISTINCT ON (port) id, created_by, type, port, startop, endop FROM operations
		WHERE vessel = $1
		ORDER BY port, id DESC;
	`

	var operations []Operation

	rows, err := v.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var operation Operation
		err := rows.Scan(
			&operation.ID,
			&operation.CreatedBy,
			&operation.Type,
			&operation.Port,
			&operation.StartOp,
			&operation.EndOp,
		)
		if err != nil {
			fmt.Println(err)
			return nil, Metadata{}, err
		}
		operations = append(operations, operation)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return nil, Metadata{}, err
	}

	vessel.Operations = operations

	// Orders Query ////////////////////////
	query = `
		SELECT id, created_by, sales_number, purchasing_number, customer, loading_berth, destination_port, destination_berth, product, volume, sales_rep, crp FROM orders
        WHERE vessel = $1;
	`

	var orders []Order

	rows, err = v.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var order Order
		err := rows.Scan(
			&order.ID,
			&order.CreatedBy,
			&order.SalesNumber,
			&order.PurchasingNumber,
			&order.Customer,
			&order.LoadingBerth,
			&order.DestinationPort,
			&order.DestinationBerth,
			&order.Product,
			&order.Volume,
			&order.SalesRep,
			&order.CRP,
		)
		if err != nil {
			fmt.Println(err)
			return nil, Metadata{}, err
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return nil, Metadata{}, err
	}

	vessel.Orders = orders

	metadata := Metadata{
		TotalVolume:    calcVesselVolume(vessel.Orders),
		TotalByProduct: totalByProduct(vessel.Orders),
	}

	return &vessel, metadata, nil
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

func (v VesselModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
	DELETE FROM vessels
	WHERE id = $1`

	result, err := v.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}

type Metadata struct {
	TotalVolume    float64            `json:"total_volume"`
	TotalByProduct map[string]float64 `json:"total_by_product"`
}

func calcVesselVolume(orders []Order) float64 {
	var total float64
	for i := 0; i < len(orders); i++ {
		total = total + orders[i].Volume
	}
	return total
}

func totalByProduct(orders []Order) map[string]float64 {
	productMap := map[string]float64{}
	for i := 0; i < len(orders); i++ {
		product := orders[i].Product
		volume := orders[i].Volume
		if v, ok := productMap[product]; ok {
			productMap[product] = volume + v
		} else {
			productMap[product] = volume
		}
	}
	return productMap
}
