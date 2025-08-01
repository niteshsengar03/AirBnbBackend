package config

import (
	"Reviews_Service/config"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

func SetupDB() (*sql.DB,error){
	cfg := mysql.NewConfig();
	cfg.User = config.GetString("DB_USER","root")
	cfg.Passwd = config.GetString("DB_PASS","xyz")
	cfg.Net = config.GetString("DB_NET","tcp")
	cfg.Addr = config.GetString("DB_ADDR","127.0.0.1:3306")
	cfg.DBName = config.GetString("DB_NAME","review_dev")

	db,err := sql.Open("mysql",cfg.FormatDSN())
	fmt.Println("Connecting to database: ",cfg.DBName, cfg.FormatDSN());
	if err!=nil{
		fmt.Println("Error connecting to database ",err)
		return nil,err
	}
	fmt.Println("Trying to connect to database..");
	pingErr := db.Ping()
	if pingErr !=nil{
		fmt.Println("Error pinging the database ",pingErr);
		return nil,pingErr
	}
	fmt.Println("Connected to database sucessfully: ",cfg.DBName)
	return  db,nil
}
