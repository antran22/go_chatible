package env

import (
	"os"

	"github.com/joho/godotenv"
)

func Load() {
	env := os.Getenv("APP_ENV")
	if "" == env {
		env = "development"
	}

	_ = godotenv.Load("env/.env." + env + ".local")
	if "test" != env {
		_ = godotenv.Load("env/.env.local")
	}
	_ = godotenv.Load("env/.env." + env)
	_ = godotenv.Load("env/.env")
}
