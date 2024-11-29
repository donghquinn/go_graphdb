package routers

import (
	"net/http"

	"org.donghyuns.com/graph/neo4j/biz/person"
)

func PersonRouter(server *http.ServeMux) {
	server.HandleFunc("/person", person.CreatePersonController)
	server.HandleFunc("/person/get", person.GetPersonController)
}
