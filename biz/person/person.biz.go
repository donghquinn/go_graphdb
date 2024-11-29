package person

import (
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
	"org.donghyuns.com/graph/neo4j/database"
)

// Insert New Person Node And Tech Node
func CreateSinglePerson(name string, age string, languageList []string, techName string) error {
	dbCon, dbConErr := database.InitGraphConnect()

	if dbConErr != nil {
		return dbConErr
	}

	queryArguments := map[string]interface{}{
		"name":      name,
		"age":       age,
		"languages": languageList,
		"tech":      techName,
	}

	queryErr := dbCon.Insert(CreatePersonQuery, queryArguments)

	if queryErr != nil {
		log.Printf("[SINGLE_PERSON] Create single person: %v", queryErr)
		return queryErr
	}

	return nil
}

// Get Person Node List By Name
func GetPerson(name string) ([]Person, error) {
	dbCon, dbConErr := database.InitGraphConnect()

	if dbConErr != nil {
		return []Person{}, dbConErr
	}

	result, resultErr := dbCon.QueryOne(GetPersonQueryByName, map[string]interface{}{"name": name})

	if resultErr != nil {
		return []Person{}, resultErr
	}

	var personList []Person

	for _, record := range result {
		person, isPersonExist := record.Get("p")

		if !isPersonExist {
			log.Printf("[SINGLE_PERSON] No Person Data Found: %v", isPersonExist)
			continue
		}

		log.Printf("[SINGLE_PERSON] Get single person: %v", person)

		personProperties := person.(dbtype.Node).Props

		personData := Person{
			Name:      personProperties["name"].(string),
			Age:       personProperties["age"].(string),
			Languages: personProperties["languages"].([]string),
		}

		personList = append(personList, personData)
	}

	log.Printf("[SINGLE_PERSON] Get single person: %v", personList)

	return personList, nil
}
