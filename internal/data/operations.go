package data

import "database/sql"

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
