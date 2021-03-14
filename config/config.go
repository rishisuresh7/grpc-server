package config

import (
	"fmt"
	"os"
	"strings"

	"grpc-server/wrapper"
)

type AppConfig struct {
	Port     int
	MongoUri string
	Token    string
	LogFile  *os.File
}

func NewConfig(w wrapper.Wrapper) (*AppConfig, error) {
	var missing []string
	portString := os.Getenv("PORT")
	port, err := w.Atoi(portString)
	if err != nil {
		return nil, fmt.Errorf("NewConfig: invalid value for port: %s", portString)
	}

	token := os.Getenv("TOKEN")
	if token == "" {
		missing = append(missing, "TOKEN")
	}

	mongoUri := os.Getenv("MONGO_URI")
	if mongoUri == "" {
		missing = append(missing, "MONGO_URI")
	}

	mongoUsername := os.Getenv("MONGO_USERNAME")
	if mongoUsername == "" {
		missing = append(missing, "MONGO_USERNAME")
	}

	mongoPassword := os.Getenv("MONGO_USERNAME")
	if mongoPassword == "" {
		missing = append(missing, "MONGO_USERNAME")
	}

	if len(missing) > 0 {
		return nil, fmt.Errorf("NewConfig: unable to init config, missing %s", strings.Join(missing, ", "))
	}

	return &AppConfig{
		Port:     port,
		MongoUri: fmt.Sprintf("mongodb://%s:%s@%s", mongoUsername, mongoPassword, mongoUri),
		Token:    token,
		LogFile:  os.Stdout,
	}, nil
}
