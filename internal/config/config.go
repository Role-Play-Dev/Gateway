package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	envRelease             = "RELEASE"
	envServerAddress       = "SERVER_ADDRESS"
	envClientAddress       = "CLIENT_ADDRESS"
	envTokenReadSecret     = "TOKEN_READ_SECRET"
	envTokenCookieName     = "TOKEN_COOKIE_NAME"
	envTokenCookieMaxAge   = "TOKEN_COOKIE_MAX_AGE"
	envTokenCookiePath     = "TOKEN_COOKIE_PATH"
	envTokenCookieSecure   = "TOKEN_COOKIE_SECURE"
	envTokenCookieHttpOnly = "TOKEN_COOKIE_HTTP_ONLY"
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
	ReadSecret string
	Cookie     cookieConfig
}

func NewConfig() Config {
	release := getBoolean(envRelease, false)

	if !release {
		if err := godotenv.Load(); err != nil {
			log.Println(err.Error())
		}
	}

	return Config{
		Release:       release,
		ServerAddress: getString(envServerAddress, ":8080"),
		ClientAddress: getString(envClientAddress, "localhost:80"),
		Token: tokenConfig{
			ReadSecret: getString(envTokenReadSecret, "test_secret"),
			Cookie: cookieConfig{
				Name:     getString(envTokenCookieName, "refresh_token"),
				MaxAge:   int(getInteger(envTokenCookieMaxAge, 10000)),
				Path:     getString(envTokenCookiePath, "refresh_token"),
				Secure:   getBoolean(envTokenCookieSecure, false),
				HttpOnly: getBoolean(envTokenCookieHttpOnly, true),
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
