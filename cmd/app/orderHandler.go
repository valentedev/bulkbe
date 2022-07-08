package main

import (
	"fmt"
	"github/valentedev/bulkbe/internal/data"
	"net/http"
)

func (app *application) insertOrderHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
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
		Vessel           int64   `json:"vessel"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	order := &data.Order{
		CreatedBy:        input.CreatedBy,
		SalesNumber:      input.SalesNumber,
		PurchasingNumber: input.PurchasingNumber,
		Customer:         input.Customer,
		LoadingBerth:     input.LoadingBerth,
		DestinationPort:  input.DestinationPort,
		DestinationBerth: input.DestinationBerth,
		Product:          input.Product,
		Volume:           input.Volume,
		SalesRep:         input.SalesRep,
		CRP:              input.CRP,
		Vessel:           input.Vessel,
	}

	err = app.models.Orders.Insert(order)
	if err != nil {
		fmt.Println(err)
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/order/%d", order.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"order": order}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
