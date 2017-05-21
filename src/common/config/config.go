package config

import (
	"sync"
	"os"
	"strings"
)

var instance *Config

type Config struct {
	DbName string
	DbHost string
	DbPort string
	DbUser string
	DbPass string
	DbPool int
}

func Instance() *Config {
	once := new(sync.Once)
	once.Do(func () {
		instance = new(Config)
	})
	return instance
}

func (c *Config) Load() {
	if strings.EqualFold(os.Getenv(ENV_DBNAME), "") {
		c.DbName = DB_NAME
	}

	if strings.EqualFold(os.Getenv(ENV_DBHOST), "") {
		c.DbHost = DB_HOST
	}

	if strings.EqualFold(os.Getenv(ENV_DBPORT), "") {
		c.DbPort = DB_PORT
	}

	if strings.EqualFold(os.Getenv(ENV_DBUSER), "") {
		c.DbUser = DB_USER
	}

	if strings.EqualFold(os.Getenv(ENV_DBPASS), "") {
		c.DbPass = DB_PASS
	}

	if strings.EqualFold(os.Getenv(ENV_DBHOST), "") {
		c.DbPool = DB_POOL
	}

}
