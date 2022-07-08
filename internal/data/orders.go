package data

import (
	"context"
	"database/sql"
	"time"
)

type Order struct {
	ID               int64   `json:"id"`
	CreatedBy        string  `json:"created_by"`
	SalesNumber      string  `json:"sales_number"`
	PurchasingNumber string  `json:"purchasing_number"`
	Customer         string  `json:"customer"`
	LoadingBerth     string  `json:"loading_berth"`
	DestinationPort  string  `json:"destination_port"`
	DestinationBerth string  `json:"destination_berth"`
	Product          string  `json:"product"`
	Volume           float64 `json:"volume"`
	SalesRep         string  `json:"sales_rep"`
	CRP              string  `json:"crp"`
	Vessel           int64   `json:"-"`
}

type OrderModel struct {
	DB *sql.DB
}

func (om OrderModel) Insert(order *Order) error {
	query := `
		INSERT INTO orders (created_by, sales_number, purchasing_number, customer, loading_berth, destination_port, destination_berth, product, volume, sales_rep, crp, vessel) 
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
		RETURNING id
	`

	args := []interface{}{
		order.CreatedBy,
		order.SalesNumber,
		order.PurchasingNumber,
		order.Customer,
		order.LoadingBerth,
		order.DestinationPort,
		order.DestinationBerth,
		order.Product,
		order.Volume,
		order.SalesRep,
		order.CRP,
		order.Vessel,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return om.DB.QueryRowContext(ctx, query, args...).Scan(&order.ID)
}
