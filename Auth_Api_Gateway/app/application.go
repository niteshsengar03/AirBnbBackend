package app

import (
	"Auth_Api_Gateway/config"
	dbConfig "Auth_Api_Gateway/config/db"
	"Auth_Api_Gateway/controller"
	"Auth_Api_Gateway/db/repositories"
	"Auth_Api_Gateway/router"
	"Auth_Api_Gateway/service"
	"fmt"
	"net/http"
	"time"
)




type Config struct {
	Addr string
}
// Constructor of Config Class
func NewConfig () *Config{
	return &Config{
		Addr: config.GetString("PORT", ":3002"),
	}
}





type Application struct {
	Config Config
	// Store db.Storage
}
// Constructor of Application Class
func NewApplication(cfg Config)*Application{
	return &Application{
		Config: cfg,
		// Store: *db.NewStorage(), // initialise or make objects of all repository objects 
	}
}







func (app *Application) Run() error {
	DB,err:=dbConfig.SetupDB();
	if err!=nil{
		fmt.Println("Cannot connect to database")
		return err
	}
	ur := db.NewRepository(DB);
	us := service.NewUserService(ur)
	uc := controller.NewUserController(us)
	urouter := router.NewUserRouter(uc)
	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      router.SetupRouter(urouter),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Starting server on", app.Config.Addr)
	return server.ListenAndServe()
	
}
