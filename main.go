package main

import (
	"go_mini-project/config"
	m "go_mini-project/middlewares"
	"go_mini-project/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	m.LogMiddlewares(e)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
