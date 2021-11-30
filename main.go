package main

import (
	database "github.com/andikanugraha11/rest-api-jwt/databases"
	"github.com/andikanugraha11/rest-api-jwt/routers"
	"log"
)

func main()  {
	database.InitDB()
	r := routers.InitApplication()
	log.Fatal(r.Run(":8080"))
}
