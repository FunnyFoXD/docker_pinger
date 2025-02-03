package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/funnyfoxd/docker_pinger/db"
	"github.com/funnyfoxd/docker_pinger/routes"
)

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatalf("Fail to connect to DB: %v", err)
	}

	r := gin.Default()
	routes.SetupRouter(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Fail to start server: %v", err)
	}
}
