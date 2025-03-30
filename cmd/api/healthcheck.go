package main

import "net/http"

func (app *application) HealthHandler(w http.ResponseWriter, r *http.Request) {
	err := app.writeJSON(w, http.StatusOK, envelope{
		"env":    app.cfg.Env,
		"port":   app.cfg.Port,
		"status": "running",
	}, nil)
	if err != nil {
		app.logger.Error("App is not running correctly")
	}
}
