package main

import "net/http"

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	if status != 200 {
		app.logger.Println(message)
		w.WriteHeader(status)
	}
}

func (app *application) rateLimitExceedResponse(w http.ResponseWriter, r *http.Request) {
	message := "rate limit exceeded"
	app.errorResponse(w, r, http.StatusTooManyRequests, message)
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}
