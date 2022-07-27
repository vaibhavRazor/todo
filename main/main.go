package main

import (
	"todo/mapping"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mapping.CreateUrlMappings()
	// Listen and server on 0.0.0.0:8080
<<<<<<< HEAD
	//ssh3
=======
	//sh97897
>>>>>>> 556605ca885eaddad7423e70db6cc084e47dcd73

	mapping.Router.Run(":8080")

}
