package config

import "os"

// Server contains the configuration for Gin server
type Server struct {
	TCPPort string
}

type Config struct {
	*Server
}

const unknown = "unknown"

var applicationEnv = unknown

func Load(appEnv string) (*Config, error) {
	applicationEnv = appEnv
	config := Config{
		Server: &Server{
			TCPPort: serverPort(),
		},
	}

	return &config, nil
}

func serverPort() string {
	return os.Getenv("PORT")
}
