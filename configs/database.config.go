package configs

import "os"

type DatabaseConfig struct {
	GraphUri    string
	GraphUser   string
	GraphPasswd string
}

var DatabaseConfiguration DatabaseConfig

func SetDatabaseConfiguration() {
	DatabaseConfiguration.GraphUri = os.Getenv("GRAPH_URL")
	DatabaseConfiguration.GraphUser = os.Getenv("GRAPH_USER")
	DatabaseConfiguration.GraphPasswd = os.Getenv("GRAPH_PASSWORD")
}
