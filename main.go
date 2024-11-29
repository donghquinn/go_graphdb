package main

import (
	"github.com/joho/godotenv"
	"org.donghyuns.com/graph/neo4j/configs"
	"org.donghyuns.com/graph/neo4j/network"
)

func main() {
	godotenv.Load(".env")

	configs.SetDatabaseConfiguration()
	server := network.OpenServer()

	server.ListenAndServe()

	// dbCon, dbConErr := database.InitGraphConnect()

	// dbCon.CheckConnection()
	// dbCon.QueryOne("CREATE (n:Person{name:'Kim'})-[l:LIKES]->(t:Technology{name:'golang'})", nil)
}
