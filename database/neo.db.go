package database

import (
	"context"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/db"
	"org.donghyuns.com/graph/neo4j/configs"
)

type GraphDb struct {
	neo4j.DriverWithContext
}

func InitGraphConnect() GraphDb {
	// ctx := context.Background()
	databaseConfig := configs.DatabaseConfiguration

	log.Printf("[GRAPH_DB] Configuration: %v", databaseConfig)

	driver, err := neo4j.NewDriverWithContext(
		databaseConfig.GraphUri,
		neo4j.BasicAuth(databaseConfig.GraphUser, databaseConfig.GraphPasswd, ""),
	)

	if err != nil {
		log.Printf("[GRAPH_DB] Creating Instance Error: %v", err)
		return GraphDb{}
	}

	log.Printf("[GRAPH_DB] Instance Created")

	connection := GraphDb{driver}

	return connection
}

func (db *GraphDb) CheckConnection() error {
	ctx := context.Background()

	err := db.VerifyConnectivity(ctx)

	if err != nil {
		log.Printf("[GRAPH_DB] Check Connection Error: %v", err)
		return err
	}

	defer db.Close(ctx)

	log.Printf("[GRAPH_DB] Connection Verified")
	return nil
}

func (db *GraphDb) QueryOne(queryString string, argumentList map[string]any) ([]*db.Record, error) {
	ctx := context.Background()

	result, queryErr := neo4j.ExecuteQuery(ctx, db, queryString, argumentList, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))

	if queryErr != nil {
		log.Printf("[GRAPH_DB] Query Error: %v", queryErr)
		return nil, queryErr
	}

	defer db.Close(ctx)

	for _, record := range result.Records {
		log.Printf("[GRAPH_DB] Query Result: %v", record.Values...)
	}

	log.Printf("The query `%v` returned %v records in %+v.\n",
		result.Summary.Query().Text(), len(result.Records),
		result.Summary.ResultAvailableAfter())

	return result.Records, nil
}

func (db *GraphDb) Insert(queryString string, argumentList map[string]any) ([]*db.Record, error) {
	ctx := context.Background()

	result, queryErr := neo4j.ExecuteQuery(ctx, db, queryString, argumentList, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))

	if queryErr != nil {
		log.Printf("[GRAPH_DB] Query Error: %v", queryErr)
		return nil, queryErr
	}

	defer db.Close(ctx)

	summary := result.Summary
	log.Printf("Created %v nodes in %+v.\n",
		summary.Counters().NodesCreated(),
		summary.ResultAvailableAfter())

	return result.Records, nil
}
