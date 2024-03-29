package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	// Healthcheck
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	// Vessel Handler
	router.HandlerFunc(http.MethodPost, "/v1/vessels", app.insertVesselHandler)
	router.HandlerFunc(http.MethodGet, "/v1/vessels", app.getVesselsHandler)
	router.HandlerFunc(http.MethodGet, "/v1/vessels/:id", app.getVesselHandler)
	router.HandlerFunc(http.MethodPut, "/v1/vessels/:id", app.updateVesselHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/vessels/:id", app.deleteVesselHandler)

	// Operation Handlers
	router.HandlerFunc(http.MethodPost, "/v1/operations", app.insertOperationHandler)
	//router.HandlerFunc(http.MethodGet, "/v1/operations/:id", app.GetLoadByVesselHandler)
	router.HandlerFunc(http.MethodGet, "/v1/operations/:id", app.GetOpsByVesselHandler)

	// Order Handlers
	router.HandlerFunc(http.MethodPost, "/v1/orders", app.insertOrderHandler)

	// Calendar Handlers
	router.HandlerFunc(http.MethodGet, "/v1/calendar", app.vesselsForCalendarHandler)

	return app.recoverPanic(app.enableCORS(app.rateLimit(router)))
}
