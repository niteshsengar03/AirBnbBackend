package app

import (
	ConfigDB "Reviews_Service/config/db"
	"Reviews_Service/controller"
	"Reviews_Service/db"
	"Reviews_Service/service"
	"fmt"
	"net/http"
	"time"
)

type Application struct {
	Addr string
}

func NewApplication(addr string) *Application {
	return &Application{
		Addr: addr,
	}
}

func (app *Application) Run() error {
	DB, err := ConfigDB.SetupDB()
	if err != nil {
		fmt.Println("Cannot connect to database")
		return err
	}
	repoObj := db.NewRepository(DB)
	serviceObj := service.NewReviewService(repoObj)
	controller.NewReviewController(serviceObj)
	// routerObj :=router.NewReviewRouter(controllerObj)
	server := &http.Server{
		Addr:         app.Addr,
		Handler:      nil,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Starting server on", app.Addr)
	return server.ListenAndServe()
}
