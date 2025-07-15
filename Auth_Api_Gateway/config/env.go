package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Load(){
	err:=godotenv.Load()
	if err != nil{
		fmt.Println("Error in loading .env file")
	}
}

func GetString(key string ,fallback string)string{
	// ok is boolean
	// value is the string returns two values
	value,ok := os.LookupEnv(key)

	if !ok{
		return fallback
	}

	return  value
}

func GetInt(key string,fallback int) int {
	value,ok := os.LookupEnv(key)

	if !ok{
		return  fallback
	}
	// convert string to int
	intValue,err :=strconv.Atoi(value)
	if err!=nil{
		fmt.Printf("Error in converting string: \"%s\" to int: %v\n", value, err)
		return fallback
	}
	return  intValue
}