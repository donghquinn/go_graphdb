package main

import (
	"github.com/joho/godotenv"
	"org.donghyuns.com/graph/neo4j/configs"
)

func main() {
	godotenv.Load(".env")

	configs.SetDatabaseConfiguration()
	// server := network.OpenServer()

	// server.ListenAndServe()

	// dbCon, dbConErr := database.InitGraphConnect()

	// dbCon.CheckConnection()

	// 테스트용
	// person.CreateSinglePerson("Kim", "25", []string{"Typescript", "Golang", "Python"}, "Golang")
	// person.GetPerson("Kim")
}
