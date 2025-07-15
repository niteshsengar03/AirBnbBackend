package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

func load(){
	err:=godotenv.Load()
	if err != nil{
		fmt.Println("Error in loading .env file")
	}
}

func GetString(key string ,fallback string)string{
	load()
	// ok is boolean
	// value is the string returns two values
	value,ok := os.LookupEnv(key)

	if !ok{
		return fallback
	}

	return  value
}