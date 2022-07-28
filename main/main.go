package main

import (
	"todo/mapping"

	_ "github.com/go-sql-driver/mysql"
)
//tsting
func main() {
	mapping.CreateUrlMappings()
	// Listen and server on 0.0.0.0:8080
	//shhkjhjk0098933h97897

mapping.Router.Run(":8080")

}
