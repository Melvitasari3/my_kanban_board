package main

import (
	"my_kanban_board/database"
	"my_kanban_board/router"
	"os"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	port := os.Getenv("PORT")
	//r.Run(":8080")
	r.Run(":"+port)
}
