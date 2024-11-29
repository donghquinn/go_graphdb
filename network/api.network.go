package network

import (
	"net/http"
	"time"

	"org.donghyuns.com/graph/neo4j/middlewares"
)

func OpenServer() *http.Server {
	server := http.NewServeMux()

	middlewares := middlewares.CorsMiddlewares(server)

	serving := &http.Server{
		Addr:         ":8080",
		Handler:      middlewares,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	return serving
}
