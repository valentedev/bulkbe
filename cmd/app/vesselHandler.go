package main

import (
	"fmt"
	"github/valentedev/bulkbe/internal/data"
	"net/http"
)

func (app *application) insertVesselHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		ID           int64  `json:"id"`
		CreatedBy    string `json:"created_by"`
		Name         string `json:"name"`
		Voyage       string `json:"voyage"`
		Service      string `json:"service"`
		Status       string `json:"status"`
		Tolerance    string `json:"tolerance"`
		Booking      string `json:"booking"`
		InternalNote string `json:"internal_note"`
		ExternalNote string `json:"external_note"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	vessel := &data.Vessel{
		ID:           input.ID,
		CreatedBy:    input.CreatedBy,
		Name:         input.Name,
		Voyage:       input.Voyage,
		Service:      input.Service,
		Status:       input.Service,
		Tolerance:    input.Tolerance,
		Booking:      input.Booking,
		InternalNote: input.InternalNote,
		ExternalNote: input.ExternalNote,
	}

	err = app.models.Vessels.Insert(vessel)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/movies/%d", vessel.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"vessel": vessel}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
