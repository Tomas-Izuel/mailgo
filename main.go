package main

import (
	"mailgo/lib/db"
	"mailgo/rabbit"
	"mailgo/rest"
)

func main() {
	db.ConnectDatabase()
	defer db.DisconnectDatabase()

	rabbit.Init()
	rest.Init()
}
