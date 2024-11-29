package database

import (
	"context"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"org.donghyuns.com/graph/neo4j/configs"
)

type GraphDb struct {
	neo4j.DriverWithContext
}

func InitGraphConnect() GraphDb {
	ctx := context.Background()
	databaseConfig := configs.DatabaseConfiguration

	driver, err := neo4j.NewDriverWithContext(
		databaseConfig.GraphUri,
		neo4j.BasicAuth(databaseConfig.GraphUser, databaseConfig.GraphPasswd, ""),
	)

	if err != nil {
		log.Printf("[GRAPH_DB] Creating Instance Error: %v", err)
		return GraphDb{}
	}

	defer driver.Close(ctx)

	connection := GraphDb{driver}

	return connection
}