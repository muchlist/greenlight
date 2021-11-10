package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, _ *http.Request) {

	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and cloud not process your request", http.StatusInternalServerError)
	}
}
