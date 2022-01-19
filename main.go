package main

import (
	"my_kanban_board/database"
	"my_kanban_board/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
