package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Debug         bool
	ServerAddress string
}

func NewConfig() Config {
	debug := getBoolean("DEBUG", true)

	if debug {
		if err := godotenv.Load(); err != nil {
			log.Println(err.Error())
		}
	}

	return Config{
		Debug:         debug,
		ServerAddress: getString("SERVER_ADDR", ":8080"),
	}
}

func getBoolean(key string, def bool) bool {
	v, ex := os.LookupEnv(key)

	if ex {
		if res, err := strconv.ParseBool(v); err == nil {
			return res
		}
	}

	return def
}

func getString(key string, def string) string {
	v, ex := os.LookupEnv(key)

	if ex {
		return v
	}

	return def
}
