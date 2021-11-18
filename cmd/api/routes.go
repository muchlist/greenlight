package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	// convert notFoundResponse to http handler and set it as the custom error handler for 404 notfound
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	// convert methodNotAllowedResponse to http handler and set it as the custom error handler for 405 method not allowed
	router.NotFound = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/movies", app.listMoviesHandler)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthCheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/movies/:id", app.updateMovieHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", app.deleteMovieHandler)
	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)

	// Wrap the router with the rateLimit() middleware.
	// Wrap the router with the panic recovery middleware.
	return app.recoverPanic(app.rateLimit(router))
}
