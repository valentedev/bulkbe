package data

import "database/sql"

type Order struct {
	ID               int64  `json:"id"`
	CreatedBy        string `json:"created_by"`
	SalesNumber      string `json:"sales_number"`
	PurchasingNumber string `json:"purchasing_number"`
	Customer         string `json:"customer"`
	LoadingBerth     string `json:"loading_berth"`
	DestinationPort  string `json:"destination_port"`
	DestinationBerth string `json:"destination_berth"`
	Product          string `json:"product"`
	Volume           string `json:"volume"`
	SalesRep         string `json:"sales_rep"`
	CRP              string `json:"crp"`
	Vessel           int64  `json:"vessel"`
}

type OrderModel struct {
	DB *sql.DB
}
