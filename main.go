package main

import (
	"fmt"
)

func main() {
	database := setupDB()
	fmt.Println("DATABASE CONNECTION ESTABLISHED", database)
	database.getUser()
	r := setupRouter()
	fmt.Println("ROUTER", r)
	// Listen and Server in 0.0.0.0:8080
	r.router.Run(":8080")
}
