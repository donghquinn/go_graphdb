package configs

import "os"

type DatabaseConfig struct {
	GraphUri string
	GraphUser string
	GraphPasswd string
}

var DatabaseConfiguration = DatabaseConfig{
	GraphUri: os.Getenv("GRAPH_URL"),
	GraphUser: os.Getenv("GRAPH_USER"),
	GraphPasswd: os.Getenv("GRAPH_PASSWORD"),
}