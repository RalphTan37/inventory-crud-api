package application

import (
	"os"
	"strconv"
)

// config properties
type Config struct {
	RedisAddress string
	ServerPort   uint16
}

// returns instance of config
func LoadConfig() Config {
	cfg := Config{ //creates instance of config
		RedisAddress: "localhost:6379",
		ServerPort:   3000,
	}

	//checks if env is "REDIS_ADDR"
	if redisAddr, exists := os.LookupEnv("REDIS_ADDR"); exists {
		cfg.RedisAddress = redisAddr
	}

	//checks if env is "SERVER_PORT"
	if serverPort, exists := os.LookupEnv("SERVER_PORT"); exists {
		if port, err := strconv.ParseUint(serverPort, 10, 16); err != nil {
			cfg.ServerPort = uint16(port)
		}
	}

	return cfg //returns config instance
}
