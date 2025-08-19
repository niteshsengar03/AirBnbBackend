package app

import (
	"Auth_Api_Gateway/config"
	dbConfig "Auth_Api_Gateway/config/db"
	"Auth_Api_Gateway/controller"
	db "Auth_Api_Gateway/db/repositories"
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
func NewConfig() *Config {
	return &Config{
		Addr: config.GetString("PORT", ":3002"),
	}
}

type Application struct {
	Config Config
}

// Constructor of Application Class
func NewApplication(cfg Config) *Application {
	return &Application{
		Config: cfg,
	}
}

func (app *Application) Run() error {
	DB, err := dbConfig.SetupDB()
	if err != nil {
		fmt.Println("Cannot connect to database")
		return err
	}
	rr := db.NewRoleRepository(DB)
	rpr := db.NewRolePermissionRepository(DB)
	ur := db.NewRepository(DB)
	urp := db.NewUserRoleRepository(DB)
	rs := service.NewRoleService(rr, rpr, urp)
	us := service.NewUserService(ur)
	rc := controller.NewRoleController(rs)
	uc := controller.NewUserController(us)
	rrouter := router.NewRoleRouter(rc)
	urouter := router.NewUserRouter(uc)
	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      router.SetupRouter(urouter, rrouter),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Starting server on", app.Config.Addr)
	return server.ListenAndServe()

}
