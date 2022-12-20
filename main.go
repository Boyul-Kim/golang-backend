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
	r.router.Run(":8080")
}
