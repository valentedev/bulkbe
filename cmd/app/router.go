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

	return app.recoverPanic(app.enableCORS(app.rateLimit(router)))
}
