package main

import (
	"errors"
	"fmt"
	"github/valentedev/bulkbe/internal/data"
	"net/http"
)

func (app *application) insertOperationHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		CreatedBy string `json:"created_by"`
		Type      string `json:"type"`
		Port      string `json:"port"`
		StartOp   string `json:"startop"`
		EndOp     string `json:"endop"`
		Vessel    int64  `json:"vessel"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	operation := &data.Operation{
		CreatedBy: input.CreatedBy,
		Type:      input.Type,
		Port:      input.Port,
		StartOp:   input.StartOp,
		EndOp:     input.EndOp,
		Vessel:    input.Vessel,
	}

	err = app.models.Operations.Insert(operation)
	if err != nil {
		fmt.Println(err)
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/operations/%d", operation.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"operation": operation}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) GetLoadByVesselHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	operations, err := app.models.Operations.GetLoadByVessel(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"operations": operations}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) GetOpsByVesselHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	operations, err := app.models.Operations.SchedulesByVessel(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"operations": operations}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
