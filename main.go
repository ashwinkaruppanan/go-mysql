package main

import (
	"log"

	"github.com/ashwin/go-mysql/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.RegisterBookStoreRoutes(r)

	log.Fatal(r.Run())
}
