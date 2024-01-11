package util

import (
	"os"

	"github.com/joho/godotenv"

	"github.com/programkingstar/task-management-go.git/types"
)

func GetEnv() types.Env {
	err := godotenv.Load(".env.local")
	if err != nil {
		panic(err)
	}

	return types.Env{
		Host:    os.Getenv("HOST"),
		Port:    os.Getenv("PORT"),
		DBUrl:   os.Getenv("DB_URL"),
		SSLMode: os.Getenv("SSL_MODE"),
	}
}
