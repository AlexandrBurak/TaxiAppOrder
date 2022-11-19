package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_HOST         string
	DB_PORT         string
	RPC_PORT_USER   string
	RPC_PORT_DRIVER string
	RPC_HOST_USER   string
	RPC_HOST_DRIVER string
}

func GetCfg() (Config, error) {
	cfg := Config{}
	err := readEnvironmentVariables(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func readEnvironmentVariables(cfg *Config) (err error) {
	err = godotenv.Load(".env")
	if err != nil {
		return err
	}
	vars := []string{
		"DB_HOST",
		"DB_PORT",
		"RPC_PORT",
		"RPC_PORT_USER",
		"RPC_PORT_DRIVER",
		"RPC_HOST_USER",
		"RPC_HOST_DRIVER"}

	for i := range vars {
		if os.Getenv(vars[i]) == "" {
			return errors.New("empty environment variable " + vars[i])
		}
	}
	cfg.DB_PORT = os.Getenv("DB_PORT")
	cfg.DB_HOST = os.Getenv("DB_HOST")
	cfg.RPC_PORT_USER = os.Getenv("RPC_PORT_USER")
	cfg.RPC_PORT_DRIVER = os.Getenv("RPC_PORT_DRIVER")
	cfg.RPC_PORT_USER = os.Getenv("RPC_HOST_USER")
	cfg.RPC_PORT_DRIVER = os.Getenv("RPC_HOST_DRIVER")
	return nil
}
