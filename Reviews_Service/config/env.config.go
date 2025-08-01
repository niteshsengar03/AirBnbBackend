package config

import (
	"os"

	"github.com/joho/godotenv"
)


func Load(){
	godotenv.Load()
}

func GetString(key string,fallback string)string{
	val,ok :=os.LookupEnv(key)
	if !ok{
		return fallback
	} 
	return val
}

