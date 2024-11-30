package main

import (
	"gopoc/config/logger"
	ApiHandler "gopoc/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// @host localhost:8000
// @BasePath /tasks
func main() {
	logger.InitaiteLogger()
	var chiMux *chi.Mux = chi.NewRouter()
	ApiHandler.ApiHandler(chiMux)
	ApiHandler.CronHandler()
	logger.Log(`Server starting at http://localhost:8000`)
	err := http.ListenAndServe("localhost:8080", chiMux)
	if err != nil {
		logger.Error(err)
	}
}
