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
	ClientAddress string
	Token         tokenConfig
}

type cookieConfig struct {
	Name     string
	MaxAge   int
	Path     string
	Secure   bool
	HttpOnly bool
}

type tokenConfig struct {
	Secret string
	Cookie cookieConfig
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
		Token: tokenConfig{
			Secret: getString("TOKEN_SECRET", "test_secret"),
			Cookie: cookieConfig{
				Name:     getString("TOKEN_COOKIE_NAME", "refresh_token"),
				MaxAge:   int(getInteger("TOKEN_COOKIE_MAX_AGE", 10000)),
				Path:     getString("TOKEN_COOKIE_PATH", "refresh_token"),
				Secure:   getBoolean("TOKEN_COOKIE_SECURE", false),
				HttpOnly: getBoolean("TOKEN_COOKIE_HTTP_ONLY", true),
			},
		},
	}
}

func getBoolean(key string, def bool) bool {
	if v, exists := os.LookupEnv(key); exists {
		if res, err := strconv.ParseBool(v); err == nil {
			return res
		}
	}

	return def
}

func getString(key string, def string) string {
	if v, exists := os.LookupEnv(key); exists {
		return v
	}

	return def
}

func getInteger(key string, def int64) int64 {

	if v, exists := os.LookupEnv(key); exists {
		if i, err := strconv.ParseInt(v, 10, 64); err == nil {
			return i
		}
	}

	return def
}
