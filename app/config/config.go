package config

import (
	"github.com/vavas/workchan/pkg/env"
	"os"
)

// Server contains the configuration for Gin server
type Server struct {
	TCPPort string
}

// Database contains configuration for the DB.
type Database struct {
	MasterURL  string
	ReaderURL  string
	PoolSize   int
	LogQueries bool
}

type Config struct {
	*Server
	*Database
}

const unknown = "unknown"

var applicationEnv = unknown

func Load(appEnv string) (*Config, error) {
	applicationEnv = appEnv
	config := Config{
		Server: &Server{
			TCPPort: serverPort(),
		},
		Database: &Database{
			MasterURL:  databaseMasterURL(),
			ReaderURL:  databaseReaderURL(),
			PoolSize:   databasePoolSize(),
			LogQueries: databaseLogQueries(),
		},
	}

	return &config, nil
}

func serverPort() string {
	return os.Getenv("PORT")
}

func databaseMasterURL() string {
	return os.Getenv("DATABASE_URL")
}

func databaseReaderURL() string {
	return os.Getenv("DATABASE_READER_URL")
}

func databasePoolSize() int {
	return 100
}

func databaseLogQueries() bool {
	return env.IsEnabled("DATABASE_QUERY_LOGGING_ENABLED")
}
