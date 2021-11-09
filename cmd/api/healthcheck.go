package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, _ *http.Request) {
	data := map[string]string{
		"status":     "available",
		"environmet": app.config.env,
		"version":    version,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Println(data)
		http.Error(w, "The server encountered a problem and cloud not process yaur request", http.StatusInternalServerError)
	}
}
