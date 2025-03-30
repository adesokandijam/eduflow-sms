package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	// router.NotFound = http.HandlerFunc(app.notFoundResponse)
	// router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/healthz", app.HealthHandler)

	router.HandlerFunc(http.MethodPost, "/api/v1/students", app.CreateStudentHandler)
	router.HandlerFunc(http.MethodGet, "/api/v1/students/:id", app.GetStudentByIDHandler)
	router.HandlerFunc(http.MethodGet, "/api/v1/students", app.GetAllStudentsHandler)

	return router
}
