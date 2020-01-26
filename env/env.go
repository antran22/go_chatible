package env

import (
	"os"

	"github.com/joho/godotenv"
)

func Load(prefix string) {
	env := os.Getenv("APP_ENV")
	if "" == env {
		env = "development"
	}

	_ = godotenv.Load(prefix + "env/.env." + env + ".local")
	if "test" != env {
		_ = godotenv.Load(prefix + "env/.env.local")
	}
	_ = godotenv.Load(prefix + "env/.env." + env)
	_ = godotenv.Load(prefix + "env/.env")
}
