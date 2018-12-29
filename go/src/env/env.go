package env

import (
	"os"
)

func GetEnvOrDefault(env string, def string) string {
	v := os.Getenv(env)

	if v != "" {

		return v

	}

	return def

}
