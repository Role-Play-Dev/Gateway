package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Release       bool
	ServerAddress string
}

func NewConfig() Config {
	release := getBoolean("RELEASE", false)

	if !release {
		if err := godotenv.Load(); err != nil {
			log.Println(err.Error())
		}
	}

	return Config{
		Release:       release,
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
