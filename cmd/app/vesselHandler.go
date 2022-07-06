package main

import (
	"errors"
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
		Status:       input.Status,
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

func (app *application) getVesselHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	vessel, err := app.models.Vessels.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"vessel": vessel}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getVesselsHandler(w http.ResponseWriter, r *http.Request) {
	vessels, err := app.models.Vessels.GetAll()
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"vessels": vessels}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateVesselHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	vessel, err := app.models.Vessels.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
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

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	vessel.CreatedBy = input.CreatedBy
	vessel.Name = input.Name
	vessel.Voyage = input.Voyage
	vessel.Service = input.Service
	vessel.Status = input.Status
	vessel.Tolerance = input.Tolerance
	vessel.Booking = input.Booking
	vessel.InternalNote = input.InternalNote
	vessel.ExternalNote = input.ExternalNote

	err = app.models.Vessels.Update(vessel)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"vessel": vessel}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
