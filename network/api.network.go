package network

import (
	"fmt"
	"net/http"
	"time"

	"org.donghyuns.com/graph/neo4j/configs"
	"org.donghyuns.com/graph/neo4j/middlewares"
)

func OpenServer() *http.Server {
	server := http.NewServeMux()

	middlewares := middlewares.CorsMiddlewares(server)

	serving := &http.Server{
		Addr:         fmt.Sprintf(":%s", configs.GlobalConfiguration.AppPort),
		Handler:      middlewares,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	return serving
}
