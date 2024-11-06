package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	JWTExpirationInSeconds int64
	JWTSecret              string
}

var ENVS Config = initConfig()

func initConfig() Config {
	godotenv.Load()
	var db_url string = fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306"))
	log.Println(db_url)
	return Config{
		PublicHost:             getEnv("PUBLIC_HOST", "http://127.0.0.1"),
		Port:                   getEnv("PORT", "8080"),
		DBUser:                 getEnv("DB_USER", "root"),
		DBPassword:             getEnv("DB_PASSWORD", ""),
		DBAddress:              db_url,
		DBName:                 getEnv("DB_NAME", "ecommerce"),
		JWTExpirationInSeconds: getEnvInt("JWT_EXP", 3600*24*7),
		JWTSecret:              getEnv("JWT_SECRET", "121323"),
	}
}

func getEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int64) int64 {
	val, ok := os.LookupEnv(key)
	if ok {
		i, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}

	return fallback
}
